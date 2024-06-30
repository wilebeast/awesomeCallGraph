package main

import (
	"fmt"
	"os"
	"os/exec"
)

func outputDotFile(callGraph map[string]map[string][]LogEntry) {
	// 打开输出文件
	file, err := os.Create(logId + "_call_graph.dot")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 输出 dot 文件头
	fmt.Fprintln(file, "digraph G {")
	fmt.Fprintln(file, "  edge [fontsize=10, penwidth=4];")

	// 遍历调用关系并输出
	for callingFunction, calledFunctions := range callGraph {
		for calledFunction, logEntries := range calledFunctions {
			label := fmt.Sprintf("call%d", len(logEntries))
			for i, logEntry := range logEntries {
				tooltip := marshal2String(map[string]interface{}{"arguments": unmarshal2map(logEntry.Arguments), "returns": unmarshal2map(logEntry.Returns)})
				//tooltip := fmt.Sprintf("%s", json.Marshal(map[string]interface{}{"arguments": logEntry.Arguments, "result": logEntry.Returns}))
				fmt.Fprintf(file, `"%s" -> "%s" [label="%s[%d]@%s", tooltip=%s]`+"\n",
					callingFunction, calledFunction, label, i+1, logEntry.CallingPosition, marshal2String(tooltip))
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
	cmd := exec.Command("dot", "-Tsvg", logId+"_call_graph.dot", "-o", logId+"_call_graph.svg")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to convert DOT file to SVG:", err)
		os.Exit(1)
	}
}
