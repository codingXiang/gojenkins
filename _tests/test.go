package main

import (
	"fmt"
	"github.com/codingXiang/gojenkins"
	"strconv"
	"time"
)

func main() {
	client := gojenkins.CreateJenkins(nil, "https://dc.digiwincloud.com.cn", "xiang", "d51001")
	for {
		if ce, err := client.GetCurrentExecute(); err == nil {
			resp := ce.Resp
			fmt.Println("total = " + strconv.Itoa(resp.TotalTaskNum))
			for _, task := range resp.Tasks {
				fmt.Print(task)
				fmt.Println()
			}
			fmt.Println("--------------")
		} else {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
