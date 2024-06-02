package utils

import(
	"fmt"
	"encoding/binary"
	"encoding/json"
	"net"
	"chatRoom/common/message"
)

type Transfer struct{
	Conn net.Conn
	Buf [8096]byte //传输时使用的缓冲
}


func (this *Transfer)ReadPkg()(mes message.Message, err error){
	//发了4个读4个字节
	_, err= conn.Read(this.Buf[:4])
	if err !=nil {
		//❗️自定义一个错误
		// err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n!=int(pkgLen) || err != nil{
		//❗️自定义一个错误
		// err = errors.New("read pkg body error")
		return
	}
	//把读取到的buf[:pkgLen]反序列化
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

func WritePkg(data []byte) (err error){
	//❗️获取data的长度，并将其转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	//发送长度
	//❗️这里也有坑，conn.Write()传入的是byte切片，如果直接传入byte是数组，会报错
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	//向服务器发送真正的消息内容
	n, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen){ 
		fmt.Println("conn.Write([]byte(data)) fail",err)
		return
	}
	return
}