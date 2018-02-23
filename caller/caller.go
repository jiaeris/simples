package caller

import (
	"runtime"
	"fmt"
	"strings"
)

//打印日志行
func Line() {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("获取路径失败")
	}
	filePaths := strings.Split(file, "/")
	file = filePaths[len(filePaths)-1]
	fmt.Printf("%s:%d\n", file, line)
}
