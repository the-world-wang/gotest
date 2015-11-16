package main

import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

func main() {
	ACCESS_KEY := "YZebTjnnoLzr0fGR7wD31WqcGE2KM1EKhs006S7L"
	SECRET_KEY := "V3fZaFSgFOuX3GjNL0uouTyz9QWyKvxlA0zymCce"
	kodo.SetMac(ACCESS_KEY, SECRET_KEY)

	zone := 0                // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil) // 用默认配置创建 Client

	bucket := c.Bucket("wanghao")
	ctx := context.Background()

	localFile := "test.txt"
	err := bucket.PutFile(ctx, nil, "foo/test.txt", localFile, nil)
	if err != nil {
		// 上传文件失败处理
		fmt.Println(err)
		return
	}
	// 上传文件成功
	// 这时登录七牛Portal，在 your-bucket-name 空间就可以看到一个 foo/bar.jpg 的文件了
}
