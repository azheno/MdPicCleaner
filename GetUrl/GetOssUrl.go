package GetUrl

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"strings"
)

// 获取oss image url完整路径
func GetOssImageUrl(bucket *oss.Bucket) []string {
	var ossStr []string
	url := viper.GetString("url")
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			handleError(err)
		}
		// 获取文件，存储在ossStr数组变量中
		for _, object := range lsRes.Objects {
			if strings.HasPrefix(object.Key, "dirty/") {
				continue
			}
			if !strings.HasSuffix(object.Key, "/") {
				ossStr = append(ossStr, url+object.Key)
			}
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return ossStr

}
