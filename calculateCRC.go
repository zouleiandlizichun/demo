package main

func CalculateCRC(data []byte) []byte {
	crc := 0xFFFF
	for _, b := range data {
		crc ^= int(b)

		for i := 0; i < 8; i++ {
			if (crc & 0x0001) != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				crc >>= 1
			}
		}
	}

	return []byte{byte(crc & 0xFF), byte(crc >> 8)}
}
