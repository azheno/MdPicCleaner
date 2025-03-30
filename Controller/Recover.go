package Controller

import (
	"MdPicCleaner/Error"
	. "MdPicCleaner/GetUrl"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func Recover() {

	// 传入登录参数和操作文件变量
	LoadConfig()
	endpoint := viper.GetString("endpoint")
	accessKeyID := viper.GetString("accessKeyID")
	accessKeySecret := viper.GetString("accessKeySecret")
	note := viper.GetString("note")

	// 登录
	bucket, err := LoginAliyunOss(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		Error.HandleError(err)
	}

	//获取oss image url
	var ossStr []string = GetOssImageUrl(bucket)
	//fmt.Println("oss image url is ", ossStr)

	if err != nil {
		Error.HandleError(err)
	}

	// 获取md image url
	var mdStr []string = GetMdImageUrl(note)
	if err != nil {
		Error.HandleError(err)
	}
	//fmt.Println("--------------------------------------------------------------------------------")
	//fmt.Println("md image url is :", mdStr)

	// 将image url中冗余部分另存
	var diffImage []string = CompareAndOutputDifferences(mdStr, ossStr)

	if err != nil {
		Error.HandleError(err)
	}
	// 将diffimage冗余数据完整路径转为可被oss移动的路径格式，并调用move移动函数进行文件移动
	for i := range diffImage {
		if len(diffImage) >= 1 {
			parts := strings.Split(diffImage[i], "/")
			if len(parts) > 0 {
				//fmt.Println(parts[len(parts)-1]) // 打印文件扩展名
				//fmt.Println("dirty/" + parts[len(parts)-1])
				fmt.Println("--------------------------------------------------------------------------------")
				fmt.Printf("source-->dest:\n%s --> dirty/%s \n", parts[len(parts)-1], parts[len(parts)-1])
				err = MoveDirtyImage(bucket, parts[len(parts)-1], "dirty/"+parts[len(parts)-1])
				if err != nil {
					Error.HandleError(err)
				}
			}
		} else {
			fmt.Println("diffImage 中没有足够的元素")
		}
	}

}
