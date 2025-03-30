package Error

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Println("Error:", err) // 打印报错信息
	os.Exit(1)                 // 退出码为1 表示错误
}
