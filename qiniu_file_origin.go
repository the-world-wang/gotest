package main

import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"time"
)

func main() {
	fmt.Println("qiu niu file upload")
	access_key := "epbKZnxFUtJ9bTWufWtvXkAwtsseutpa8xRpJ3KI"
	secret_key := "TEWIuwOCs9KTeRqrOuQEDmUHDDd6RYBkoZ32m2Is"
	kodo.SetMac(access_key, secret_key)

	client := kodo.New(0, nil)
	deadline_sec := time.Now().Add(3000 * time.Second).Unix()
	reg := "timetable/data/"
	reg = reg + "filename.json" // filename = uuid.New()
	fmt.Println(reg)
	put_policy := kodo.PutPolicy{
		Scope:   "otaku-resource:" + reg,
		Expires: uint32(deadline_sec),
	}
	utoken := client.MakeUptoken(&put_policy)
	fmt.Println(utoken)

	uploader := kodocli.NewUploader(0, nil)
	ctx := context.Background()
	local_image_file := "..." //本地json文件地址
	ret := kodocli.PutRet{}
	err := uploader.PutFile(ctx, &ret, utoken, reg, local_image_file, nil)
	fmt.Println(err)
	//http://7xl0tq.com2.z0.glb.qiniucdn.com/timetable/data/
	fmt.Println(ret)

}
