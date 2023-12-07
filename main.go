package main

import (
	"github.com/jacobsa/go-serial/serial"
	"log"
	"time"
)

//var mutex = &sync.Mutex{}

func main() {
	options := serial.OpenOptions{
		PortName:   "COM3",
		BaudRate:   9600,
		DataBits:   8,
		StopBits:   1,
		ParityMode: serial.PARITY_NONE,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("串口打开错误：", err)
	}

	//responseCh := make([]byte, 1024)
	//responseChYX := make(chan []byte)
	dataWithCRC, err := HandleInput()
	dataWithCRCYX, err1 := HandleInputYX()

	if err != nil {
		log.Printf("处理发生错误", err)
	}

	if err1 != nil {
		log.Printf("处理发生错误", err)
	}

	//go Writetoserial(port, dataWithCRC, responseCh)
	//time.Sleep(time.Second * 10)
	//go WritetoserialYX(port, dataWithCRCYX, responseChYX)

	for {
		err, response := Writetoserial(port, dataWithCRC)
		if err != nil {
			log.Printf("函数执行出错：%v", err)
		}
		go Dealwith(response)
		time.Sleep(time.Second)
		err, responseYX := WritetoserialYX(port, dataWithCRCYX)
		if err != nil {
			log.Printf("函数执行出错：%v", err)
		}
		go DealwithYX(responseYX)

		time.Sleep(time.Second)
	}
	//time.Sleep(time.Second * 5)
	//go WritetoserialYX(port, dataWithCRCYX, responseChYX)

	//for {
	//	mutex.Lock() // 获取锁
	//	go func() {
	//		defer mutex.Unlock() // 在goroutine结束后释放锁
	//		err := Writetoserial(port, dataWithCRC, responseCh)
	//		if err != nil {
	//			log.Printf("Writetoserial执行错误：%v", err)
	//		}
	//	}()
	//
	//	mutex.Lock() // 再次获取锁
	//	go func() {
	//		defer mutex.Unlock() // 在goroutine结束后释放锁
	//		err := WritetoserialYX(port, dataWithCRCYX, responseChYX)
	//		if err != nil {
	//			log.Printf("Writetoserial执行错误：%v", err)
	//		}
	//	}()
	//
	//	time.Sleep(time.Second * 4)
	//}
	//
	//for {
	//
	//	select {
	//	case response := <-responseCh:
	//		go Dealwith(response)
	//		//case responseYX := <-responseChYX:
	//		//	go DealwithYX(responseYX)
	//
	//	}
	//}

}
