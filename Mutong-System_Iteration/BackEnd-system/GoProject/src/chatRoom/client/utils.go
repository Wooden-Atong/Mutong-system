package main
import (
	"net"
	"encoding/binary"
	"encoding/json"
	"chatRoom/common/message"
	"fmt"
)


func readPkg(conn net.Conn)(mes message.Message, err error){
	buf := make([]byte,8096)
	//发了4个读4个字节
	_, err= conn.Read(buf[:4])
	if err !=nil {
		//❗️自定义一个错误
		// err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n!=int(pkgLen) || err != nil{
		//❗️自定义一个错误
		// err = errors.New("read pkg body error")
		return
	}
	//把读取到的buf[:pkgLen]反序列化
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error){
	//❗️获取data的长度，并将其转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	//❗️这里也有坑，conn.Write()传入的是byte切片，如果直接传入byte是数组，会报错
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	//向服务器发送真正的消息内容
	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen){ 
		fmt.Println("conn.Write([]byte(data)) fail",err)
		return
	}
	return
}