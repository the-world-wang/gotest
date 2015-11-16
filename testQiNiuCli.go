package main

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"time"
)

func main() {
	ACCESS_KEY := "YZebTjnnoLzr0fGR7wD31WqcGE2KM1EKhs006S7L"
	SECRET_KEY := "V3fZaFSgFOuX3GjNL0uouTyz9QWyKvxlA0zymCce"
	kodo.SetMac(ACCESS_KEY, SECRET_KEY)

	zone := 0 // 您空间(Bucket)所在的区域

	client := kodo.New(zone, nil) // 用默认配置创建 Client
	deadline_sec := time.Now().Add(3000 * time.Second).Unix()
	reg := "timetable/img/"
	filename := uuid.New()
	reg = reg + filename + ".png"
	fmt.Println(reg)
	put_policy := kodo.PutPolicy{
		Scope:   "wanghao:" + reg,
		Expires: uint32(deadline_sec),
	}
	utoken := client.MakeUptoken(&put_policy)

	uploader := kodocli.NewUploader(zone, nil)
	ctx := context.Background()
	local_image_file := "D:/java/eclipse/project/TestLoad/WebContent/img/ic_game.png" //本地json文件地址
	ret := kodocli.PutRet{}
	err := uploader.PutFile(ctx, ret, utoken, reg, local_image_file, nil)

	//http://7xl0tq.com2.z0.glb.qiniucdn.com/timetable/data/
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(ret.Hash))
}
