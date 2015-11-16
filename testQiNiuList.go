package main

import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	// "qiniupkg.com/api.v7/kodocli"
)

func main() {
	access_key := "epbKZnxFUtJ9bTWufWtvXkAwtsseutpa8xRpJ3KI"
	secret_key := "TEWIuwOCs9KTeRqrOuQEDmUHDDd6RYBkoZ32m2Is"
	kodo.SetMac(access_key, secret_key)

	zone := 0                // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil) // 用默认配置创建 Client

	bucket := c.Bucket("otaku-resource")
	ctx := context.Background()

	list, marker_outer, _ := bucket.List(ctx, "timetable/data", "", 4000)
	for _, v := range list {
		bucket.Delete(ctx, v.Key)
	}
	fmt.Println(len(list), marker_outer)

	// err := bucket.Delete(ctx, "timetable/data")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("delte success")
	// list2, _, _ := bucket.List(ctx, "timetable/data", "", 2000)
	// fmt.Println(len(list2))
}
