package main

import (
	"fmt"
	"log"
	"time"
)

func DealwithYX(responseYX []byte) {
	//for _, value := range responseYX {
	//
	//	log.Printf("%02x", value)
	//}
	//responselength := len(responseYX)
	//log.Println("数据原始长度是：", responselength)
	log.Printf("开始解析遥信数据")

	subResponse := responseYX[3 : len(responseYX)-2]
	//subResponselength := len(subResponse)
	//
	//log.Println("数据解析后长度是：", subResponselength)

	log.Printf("遥信数据解析完成开始输出")

	for _, b := range subResponse {
		binStr := fmt.Sprintf("%08b", b)
		for i := len(binStr) - 1; i >= 0; i-- {
			bit := binStr[i]
			if bit == '1' {
				log.Printf("第%d位状态是: 开启", len(binStr)-i)
			} else {
				log.Printf("第%d位状态是: 关闭", len(binStr)-i)
			}
		}
	}
	time.Sleep(time.Second)
	log.Printf("开始准备下一次的数据采集")
	//var result []string
	//for i := 0; i < len(subResponse); i += 2 {
	//	if i+1 < len(subResponse) {
	//		group := fmt.Sprintf("%x%x", subResponse[i], subResponse[i+1])
	//		result = append(result, group)
	//	}
	//}

	//for index, value := range result {
	//	decimal, err := strconv.ParseInt(value, 16, 64)
	//	if err != nil {
	//		fmt.Printf("解析错误：%v\n", err)
	//	} else {
	//		log.Printf("采集模块第 %d 个数据为 %d\n", index+1, decimal)
	//	}
	//}

}
