package Controller

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"log"
)

// 导入配置文件
func LoadConfig() {
	viper.SetConfigName("server") // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
}

// 登录接口
func LoginAliyunOss(endpoint, accessKeyID, accessKeySecret string) (*oss.Bucket, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret) // 创建客户端
	if err != nil {
		return nil, err
	}
	bucketName := viper.GetString("bucketName")
	bucket, err := client.Bucket(bucketName) // 连接bucket
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
