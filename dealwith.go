package main

import (
	"fmt"
	"log"
	"strconv"
)

func Dealwith(responseCh []byte) {

	//for _, value := range responseCh {
	//
	//	log.Printf("遥测数据：%02x", value)
	//}
	//responselength := len(responseCh)
	//
	//log.Println("接收到的数据长度是：", responselength)

	subResponse := responseCh[3 : len(responseCh)-2]

	log.Printf("开始解析遥测数据")

	var result []string
	for i := 0; i < len(subResponse); i += 2 {
		if i+1 < len(subResponse) {
			group := fmt.Sprintf("%x%x", subResponse[i], subResponse[i+1])
			result = append(result, group)
		}
	}
	//resultlength := len(result)
	//log.Println("解析后的数据长度是：", resultlength)

	log.Printf("遥测数据解析完成开始输出")

	for index, value := range result {
		decimal, err := strconv.ParseInt(value, 16, 64)
		if err != nil {
			fmt.Printf("解析错误：%v\n", err)
		} else {
			log.Printf("采集模块第 %d 个数据为 %d\n", index+1, decimal)
		}
	}

	log.Printf("开始准备下一次的数据采集")

}
