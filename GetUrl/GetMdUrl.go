package GetUrl

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func handleError(err error) {
	fmt.Println("Error:", err) // 打印报错信息
	os.Exit(1)                 // 退出码为1 表示错误

}

// 获取md image url完整路径
func GetMdImageUrl(note string) []string {
	var mdStr []string
	// 指定路径
	re := regexp.MustCompile(`!\[.*?\]\(([^)]+)\)`)
	err := filepath.Walk(note, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 过滤出所有md文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			// 读取文件内容
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("读取文件内容错误", err)
				return err
			}
			// 传所有md文件路径到matches变量中
			matches := re.FindAllStringSubmatch(string(content), -1)
			if matches != nil {
				for _, match := range matches {
					if len(match) > 1 {
						mdStr = append(mdStr, match[1])
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("未找到目录，请核对路径: \n", err)
	}

	return mdStr
}
