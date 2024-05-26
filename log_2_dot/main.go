package main

import (
	"bufio"
	"fmt"
	"os"
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
