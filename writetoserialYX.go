package main

import (
	"io"
	"log"
	"time"
)

func WritetoserialYX(port io.ReadWriteCloser, dataWithCRC []byte) (error, responseCh1 []byte) {
	log.Printf("准备开始写入遥信数据")
	_, err := port.Write(dataWithCRC)
	if err != nil {
		log.Printf("数据写入错误： %v", err)
	}

	log.Printf("数据写入完成")
	time.Sleep(time.Second)

	buf := make([]byte, 1024)

	n, err := port.Read(buf)
	if err != nil {
		log.Printf("读取数据错误：%v", err)
	}

	responseCh1 = make([]byte, n)
	copy(responseCh1, buf[:n])
	log.Printf("遥信数据获取完成")
	time.Sleep(time.Second)

	return nil, responseCh1

}
