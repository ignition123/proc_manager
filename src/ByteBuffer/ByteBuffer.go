package ByteBuffer

import (
	"errors"
	"bytes"
	"encoding/binary"
	"math"
	"log"
	"lib"
)

type Buffer struct{

	packetBuffer bytes.Buffer

	Endian string

}

func (obj *Buffer) Wrap(data []byte) {

	defer lib.Handlepanic()

	obj.packetBuffer.Write(data)

}

func (obj *Buffer) PutShort(value int) {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return

	}

	buff := make([]byte, 2)

	if obj.Endian == "big"{

		binary.BigEndian.PutUint16(buff, uint16(value))

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint16(buff, uint16(value))

	}


	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetShort() []byte{
	
	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	shortValue := tempBuff[:2]

	restValue := tempBuff[2:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return shortValue

}

func (obj *Buffer) PutInt(value int) {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return
		
	}

	buff := make([]byte, 4)

	if obj.Endian == "big"{

		binary.BigEndian.PutUint32(buff, uint32(value))

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint32(buff, uint32(value))

	}
	
	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetInt() []byte{
	
	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	intValue := tempBuff[:4]

	restValue := tempBuff[4:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return intValue

}


func (obj *Buffer) PutLong(value int) {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return
		
	}

	buff := make([]byte, 8)

	if obj.Endian == "big"{

		binary.BigEndian.PutUint64(buff, uint64(value))

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint64(buff, uint64(value))

	}

	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetLong() []byte{
	
	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	longValue := tempBuff[:8]

	restValue := tempBuff[8:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return longValue

}

func (obj *Buffer) PutFloat(value float32) {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return
		
	}

	bits := math.Float32bits(value)

	buff := make([]byte, 4)
	
	if obj.Endian == "big"{

		binary.BigEndian.PutUint32(buff, bits)

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint32(buff, bits)

	}
	
	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetFloat() []byte{
	
	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	floatValue := tempBuff[:4]

	restValue := tempBuff[4:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return floatValue

}

func (obj *Buffer) PutDouble(value float64) {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return
		
	}

	bits := math.Float64bits(value)

	buff := make([]byte, 8)
	
	if obj.Endian == "big"{

		binary.BigEndian.PutUint64(buff, bits)

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint64(buff, bits)

	}
	
	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetDouble() []byte{
	
	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	doubleValue := tempBuff[:8]

	restValue := tempBuff[8:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return doubleValue

}

func (obj *Buffer) Put(value []byte) {

	defer lib.Handlepanic()

	obj.packetBuffer.Write(value)

}

func (obj *Buffer) PutByte(value byte) {

	defer lib.Handlepanic()

	var tempByte []byte

	tempByte = append(tempByte, value)

	obj.packetBuffer.Write(tempByte)

}

func (obj *Buffer) Get(size int) []byte{

	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	value := tempBuff[:size]

	restValue := tempBuff[size:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return value
}

func (obj *Buffer) GetByte() []byte{

	defer lib.Handlepanic()

	tempBuff := obj.packetBuffer.Bytes()

	value := tempBuff[:1]

	restValue := tempBuff[1:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return value

}

func (obj *Buffer) Array() []byte{

	defer lib.Handlepanic()

	return obj.packetBuffer.Bytes()

}

func (obj *Buffer) Size() int{

	defer lib.Handlepanic()

	return len(obj.packetBuffer.Bytes())

}

func (obj *Buffer) Flip(){

	defer lib.Handlepanic()

	bytesArr := obj.packetBuffer.Bytes()

	for i, j := 0, len(bytesArr)-1; i < j; i, j = i+1, j-1 {

		bytesArr[i], bytesArr[j] = bytesArr[j], bytesArr[i]

	}

	var byteBuffer bytes.Buffer

	byteBuffer.Write(bytesArr) 

	obj.packetBuffer = byteBuffer;
}

func (obj *Buffer) Clear(){

	defer lib.Handlepanic()

	var byteBuffer bytes.Buffer

	obj.packetBuffer = byteBuffer;

}

func (obj *Buffer) Slice(start int, end int) error{

	defer lib.Handlepanic()

	bytesArr := obj.packetBuffer.Bytes()

	if len(bytesArr) < (start + end){
		return errors.New("Buffer does not contain that much of limit")
	}

	bytesArr = bytesArr[start:end]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(bytesArr) 

	obj.packetBuffer = byteBuffer;

	return nil

}

func (obj *Buffer) Bytes2Str(data []byte) string{

	defer lib.Handlepanic()

	return string(data)

}

func (obj *Buffer) Str2Bytes(data string) []byte{

	defer lib.Handlepanic()

	return []byte(data)

}

func (obj *Buffer) Bytes2Short(data []byte) uint16{

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return 0
		
	}

	if obj.Endian == "big"{

		return binary.BigEndian.Uint16(data)

	}else if obj.Endian == "little"{

		return binary.LittleEndian.Uint16(data)

	}

	return 0
}

func (obj *Buffer) Bytes2Int(data []byte) uint32{

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return 0
		
	}

	if obj.Endian == "big"{

		return binary.BigEndian.Uint32(data)

	}else if obj.Endian == "little"{

		return binary.LittleEndian.Uint32(data)

	}

	return 0
}

func (obj *Buffer) Bytes2Long(data []byte) uint64{

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return 0
		
	}

	if obj.Endian == "big"{

		return binary.BigEndian.Uint64(data)

	}else if obj.Endian == "little"{

		return binary.LittleEndian.Uint64(data)

	}

	return 0
}

func (obj *Buffer) Short2Bytes(data uint16) []byte{

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return nil
		
	}

	bs := make([]byte, 2)

	if obj.Endian == "big"{

		binary.BigEndian.PutUint16(bs, data)

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint16(bs, data)

	}

	return bs
}

func (obj *Buffer) Int2Bytes(data uint32) []byte{

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return nil
		
	}

	bs := make([]byte, 4)

	if obj.Endian == "big"{

		binary.BigEndian.PutUint32(bs, data)

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint32(bs, data)

	}

	return bs

}

func (obj *Buffer) Long2Bytes(data uint64) []byte{

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return nil
		
	}

	bs := make([]byte, 8)

	if obj.Endian == "big"{

		binary.BigEndian.PutUint64(bs, data)

	}else if obj.Endian == "little"{

		binary.LittleEndian.PutUint64(bs, data)

	}

	return bs

}

func (obj *Buffer) Bytes2Float(bytes []byte) float32 {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return 0.0
		
	}

    if obj.Endian == "big"{

    	bits := binary.BigEndian.Uint32(bytes)

	    float := math.Float32frombits(bits)

	    return float

    }else if obj.Endian == "little"{

    	bits := binary.LittleEndian.Uint32(bytes)

	    float := math.Float32frombits(bits)

	    return float

    }

    return 0.0
}

func (obj *Buffer) Float2Bytes(float float32) []byte {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return nil
		
	}

	if obj.Endian == "big"{

		bits := math.Float32bits(float)

	    bytes := make([]byte, 4)

	    binary.BigEndian.PutUint32(bytes, bits)

	    return bytes

	}else if obj.Endian == "little"{

		bits := math.Float32bits(float)

	    bytes := make([]byte, 4)

	    binary.LittleEndian.PutUint32(bytes, bits)

	    return bytes

	}

	return nil
}

func (obj *Buffer) Bytes2Double(bytes []byte) float64 {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return 0
		
	}

	if obj.Endian == "big"{

		bits := binary.BigEndian.Uint64(bytes)

	    float := math.Float64frombits(bits)

	    return float

	}else if obj.Endian == "little"{

		bits := binary.LittleEndian.Uint64(bytes)

	    float := math.Float64frombits(bits)

	    return float

	}

	return 0
}

func (obj *Buffer) Double2Bytes(float float64) []byte {

	defer lib.Handlepanic()

	if obj.Endian != "big" && obj.Endian != "little"{

		log.Println("Invalid endianness, must be big or little")

		return nil
		
	}

	if obj.Endian == "big"{

		bits := math.Float64bits(float)

	    bytes := make([]byte, 8)

	    binary.BigEndian.PutUint64(bytes, bits)

	    return bytes

	}else if obj.Endian == "little"{

		bits := math.Float64bits(float)

	    bytes := make([]byte, 8)

	    binary.LittleEndian.PutUint64(bytes, bits)

	    return bytes

	}

	return nil

}