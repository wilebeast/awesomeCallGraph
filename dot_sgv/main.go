package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// 将 dot 文件转换为 svg 文件
	convertDotToSvg()

	// 启动 Web 服务
	http.HandleFunc("/", handleRequest)
	fmt.Println("Starting web server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
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
            svg.addEventListener('dblclick', function(event) {
                var target = event.target;
                if (target.tagName.toLowerCase() === 'path') {
                    var tooltip = target.getAttribute('tooltip');
                    copyToClipboard(tooltip);
                }
            });
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
