package main

import (
	//"code.google.com/p/go-uuid/uuid"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"strings"
	"time"
)

func main() {
	access_key := "epbKZnxFUtJ9bTWufWtvXkAwtsseutpa8xRpJ3KI"
	secret_key := "TEWIuwOCs9KTeRqrOuQEDmUHDDd6RYBkoZ32m2Is"
	kodo.SetMac(access_key, secret_key)

	zone := 0 // 您空间(Bucket)所在的区域

	client := kodo.New(zone, nil) // 用默认配置创建 Client
	deadline_sec := time.Now().Add(3000 * time.Second).Unix()

	//遍历文件夹...上传文件
	path := "D:/temp/img"
	files, _ := ListDir(path, "png")
	for _, v := range files {
		reg := "timetable/img/"
		// filename := uuid.New()
		reg = reg + strings.Replace(v, path+"\\", "", -1)
		fmt.Println(reg)
		put_policy := kodo.PutPolicy{
			Scope:   "otaku-resource:" + reg,
			Expires: uint32(deadline_sec),
		}
		utoken := client.MakeUptoken(&put_policy)

		local_image_file := v //本地img文件地址
		uploader := kodocli.NewUploader(zone, nil)
		ret := kodocli.PutRet{}
		ctx := context.Background()
		err := uploader.PutFile(ctx, &ret, utoken, reg, local_image_file, nil)
		//http://7xl0tq.com2.z0.glb.qiniucdn.com
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ret.Key)
	}

}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	pathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+pathSep+fi.Name())
		}
	}
	return files, nil
}
