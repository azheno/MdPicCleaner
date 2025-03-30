package GetUrl

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

// 比对md中image url 和oss image url所有的完整路径
func CompareAndOutputDifferences(mdImageUrls []string, ossImageUrls []string) []string {
	mdImageSet := make(map[string]struct{})
	var diffImage []string
	// 将本地md数据转换为map形式保存
	for _, url := range mdImageUrls {
		mdImageSet[url] = struct{}{}
	}
	// 把map中的数据和oss image url完整路径进行比对并去重保存到diffimage数组中
	for _, ossUrl := range ossImageUrls {
		if _, exists := mdImageSet[ossUrl]; !exists {
			diffImage = append(diffImage, ossUrl)
		}
	}
	return diffImage
}

// 移动oss image短名称到dirty缓冲删除目录中
func MoveDirtyImage(bucket *oss.Bucket, sourceDir, destDir string) error {

	// 复制oss image短名称到dirty缓冲删除目录中
	_, err := bucket.CopyObject(sourceDir, destDir)
	if err != nil {
		return err
	}
	// 删除oss image源路径冗余文件
	err = bucket.DeleteObject(sourceDir)
	if err != nil {
		return err
	}
	return nil

}
