package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type LogEntry struct {
	CalledFunction  string
	CallingFunction string
	CallingPosition string
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
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024) // 将缓冲区大小设置为 64 MB
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
	re := regexp.MustCompile(`Calling (.+) from (.+) at (.+), arguments:(.+), returns:(.+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 6 {
		return nil
	}

	logEntry := &LogEntry{
		CalledFunction:  matches[1],
		CallingFunction: matches[2],
		CallingPosition: matches[3],
		Arguments:       matches[4],
		Results:         matches[5],
	}

	return logEntry
}
