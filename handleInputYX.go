package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func HandleInputYX() ([]byte, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("现在功能是遥信,请输入装置地址（例如 56）：")
	scanner.Scan()
	deviceAddressstr := scanner.Text()

	fmt.Println("现在功能是遥信,请输入功能码（例如  03）：")
	scanner.Scan()
	functionCodestr := scanner.Text()

	fmt.Println("现在功能是遥信,请输入起始地址（例如 46）：")
	scanner.Scan()
	startingAddressStr := scanner.Text()

	fmt.Println("现在功能是遥信,请输入寄存器数量（例如 5）：")
	scanner.Scan()
	registerQuantityStr := scanner.Text()

	deviceAddress, err1 := strconv.ParseUint(deviceAddressstr, 10, 8)
	functionCode, err2 := strconv.ParseUint(functionCodestr, 16, 8)
	startingAddress, err3 := strconv.ParseUint(startingAddressStr, 10, 16)
	registerQuantity, err4 := strconv.ParseUint(registerQuantityStr, 10, 16)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Printf("输入转换错误：%v %v %v %v", err1, err2, err3, err4)

	}
	startingAddressHigh := byte(startingAddress >> 8)
	startingAddressLow := byte(startingAddress & 0xFF)
	registerQuantityHigh := byte(registerQuantity >> 8)
	registerQuantityLow := byte(registerQuantity & 0xFF)

	data := []byte{byte(deviceAddress), byte(functionCode), startingAddressHigh, startingAddressLow, registerQuantityHigh, registerQuantityLow}
	crc := CalculateCRC(data)
	dataWithCRCYX := append(data, crc...)

	//for _, value := range dataWithCRCYX {
	//
	//	fmt.Printf("%02x", value)
	//}

	return dataWithCRCYX, nil
}
