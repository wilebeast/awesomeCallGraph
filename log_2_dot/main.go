package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

type LogEntry struct {
	CalledFunction  string
	CallingFunction string
	Arguments       string
	Results         string
}

func main() {
	// 读取日志文件
	file, err := os.Open("./log_2_dot/logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 初始化 map 用于存储调用关系
	callGraph := make(map[string]map[string][]LogEntry)

	// 逐行读取日志文件并解析
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		logEntry := parseLogEntry(line)
		if logEntry != nil {
			// 将调用关系存储到 map 中
			if _, ok := callGraph[logEntry.CallingFunction]; !ok {
				callGraph[logEntry.CallingFunction] = make(map[string][]LogEntry)
			}
			callGraph[logEntry.CallingFunction][logEntry.CalledFunction] = append(callGraph[logEntry.CallingFunction][logEntry.CalledFunction], *logEntry)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// 输出 dot 文件
	outputDotFile(callGraph)
	// 将 dot 文件转换为 svg 文件
	convertDotToSvg()

	// 启动 Web 服务
	http.HandleFunc("/", handleRequest)
	fmt.Println("Starting web server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func parseLogEntry(line string) *LogEntry {
	// 定义正则表达式匹配日志格式
	re := regexp.MustCompile(`Calling (.+) from (.+) at .+, arguments:(.+), result:(.+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 5 {
		return nil
	}

	logEntry := &LogEntry{
		CalledFunction:  matches[1],
		CallingFunction: matches[2],
		Arguments:       matches[3],
		Results:         matches[4],
	}

	return logEntry
}

func outputDotFile(callGraph map[string]map[string][]LogEntry) {
	// 打开输出文件
	file, err := os.Create("call_graph.dot")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 输出 dot 文件头
	fmt.Fprintln(file, "digraph G {")
	fmt.Fprintln(file, "  edge [fontsize=10, penwidth=2];")

	// 遍历调用关系并输出
	for callingFunction, calledFunctions := range callGraph {
		for calledFunction, logEntries := range calledFunctions {
			label := fmt.Sprintf("call%d", len(logEntries))
			for i, logEntry := range logEntries {
				tooltip := marshal2String(map[string]interface{}{"arguments": logEntry.Arguments, "result": logEntry.Results})
				//tooltip := fmt.Sprintf("%s", json.Marshal(map[string]interface{}{"arguments": logEntry.Arguments, "result": logEntry.Results}))
				fmt.Fprintf(file, `"%s" -> "%s" [label="%s[%d]", tooltip=%s]`+"\n",
					callingFunction, calledFunction, label, i+1, marshal2String(tooltip))
			}
		}
	}

	// 输出 dot 文件尾
	fmt.Fprintln(file, "}")
}

func convertDotToSvg() {
	// 检查是否安装了 Graphviz
	_, err := exec.LookPath("dot")
	if err != nil {
		fmt.Println("Graphviz is not installed. Please install Graphviz to generate the SVG file.")
		os.Exit(1)
	}

	// 执行 dot 命令将 dot 文件转换为 svg 文件
	cmd := exec.Command("dot", "-Tsvg", "call_graph.dot", "-o", "call_graph.svg")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to convert DOT file to SVG:", err)
		os.Exit(1)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 读取 svg 文件内容
	svgContent, err := ioutil.ReadFile("call_graph.svg")
	if err != nil {
		http.Error(w, "Failed to read SVG file", http.StatusInternalServerError)
		return
	}

	// 在 HTML 模板中插入 SVG 内容并输出
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Call Graph</title>
    <script>
        function copyToClipboard(text) {
            navigator.clipboard.writeText(text).then(function() {
                console.log('Async: Copying to clipboard was successful!');
            }, function(err) {
                console.error('Async: Could not copy text: ', err);
            });
        }

        document.addEventListener('DOMContentLoaded', function() {
            var svg = document.querySelector('svg');
      		// 遍历 SVG 文档中的所有 <a> 元素(包含 tooltip 的链接)
      		var links = svg.getElementsByTagName('a');
      		for (var i = 0; i < links.length; i++) {
      		  var link = links[i];

      		  // 为每个 <a> 元素添加双击事件监听器
      		  link.addEventListener('dblclick', function() {
      		    // 从 'xlink:title' 属性中获取 tooltip 文本内容
      		    var tooltipText = this.getAttribute('xlink:title');
      		    if (tooltipText) {
      		      // 将内容复制到剪贴板
      		      navigator.clipboard.writeText(tooltipText);
      		      alert('Tooltip copied to clipboard!');
      		    }
      		  });
      		}

        });
    </script>
</head>
<body>
    %s
</body>
</html>
`
	fmt.Fprintf(w, html, svgContent)
}
