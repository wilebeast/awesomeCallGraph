package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
    <style>
        #tooltip {
            position: absolute;
            background-color: #333;
            color: #fff;
            padding: 10px;
            border-radius: 5px;
            font-size: 14px;
            display: none;
        }
    </style>
    <script>
        function copyToClipboard(text) {
            navigator.clipboard.writeText(text).then(function() {
            }, function(err) {
                console.error('Async: Could not copy text: ', err);
            });
        }

        document.addEventListener('DOMContentLoaded', function() {
            var svg = document.querySelector('svg');
      		// 遍历 SVG 文档中的所有 <a> 元素(包含 tooltip 的链接)
      		var links = svg.getElementsByTagName('a');
    		var tooltip = document.getElementById('tooltip');
    		var tooltipContent = document.getElementById('tooltipContent');
    		//var isTooltipHovered = false;

      		for (var i = 0; i < links.length; i++) {
      		  var link = links[i];

              // 为每个 <a> 元素添加鼠标悬浮事件监听器
              link.addEventListener('mouseover', function() {
                  // 从 'xlink:title' 属性中获取 tooltip 文本内容
                  var tooltipText = this.getAttribute('xlink:title');
                  if (tooltipText) {
                      showTooltip(tooltipText);
                  }
              });

              //link.addEventListener('mouseout', function() {
            	  // 只有当鼠标离开 tooltip 区域时,才隐藏 tooltip
            	  //if (!isTooltipHovered) {
                	//  hideTooltip();
            	  //}
			  //});

      		  // 为每个 <a> 元素添加双击事件监听器
      		  link.addEventListener('dblclick', function() {
      		    // 从 'xlink:title' 属性中获取 tooltip 文本内容
      		    var tooltipText = this.getAttribute('xlink:title');
      		    if (tooltipText) {
      		      // 将内容复制到剪贴板
      		      navigator.clipboard.writeText(tooltipText);
      		    }
      		  });
      		}
    		//tooltip.addEventListener('mouseover', function() {
    		    //isTooltipHovered = true;
                //hideTooltip();
    		//});

    		tooltip.addEventListener('mouseout', function() {
    		    //isTooltipHovered = false;
    		    //hideTooltip();
    		});
        });

        function showTooltip(text) {
            var tooltip = document.getElementById('tooltip');
            var tooltipContent = document.getElementById('tooltipContent');
            tooltipContent.textContent = text;
            tooltip.style.display = 'block';
            tooltip.style.left = event.pageX + 10 + 'px';
            tooltip.style.top = event.pageY + 10 + 'px';
        }

        function hideTooltip() {
            var tooltip = document.getElementById('tooltip');
            tooltip.style.display = 'none';
        }
    </script>
</head>
<body>
    <div id="tooltip" class="tooltip" style="display: none;">
        <pre id="tooltipContent"></pre>
    </div>
    %s
</body>
</html>
`
	fmt.Fprintf(w, html, svgContent)
}
