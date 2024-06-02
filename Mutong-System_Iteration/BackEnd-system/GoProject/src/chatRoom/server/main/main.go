package main
import (
	"net"
	"fmt"
	"chatRoom/common/message"
	"encoding/binary"
	_ "errors"
	"io"
	"encoding/json"
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

func serverProcessLogin(conn net.Conn, mes *message.Message)(err error){
	//反序列化
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	}else{
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	resMes.Data = string(data)
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}

	//发送res data
	err = writePkg(conn,data)
	return
}

//根绝客户端发送的消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message)(err error){
	switch mes.Type{
		case message.LoginMesType:
			go serverProcessLogin(conn,mes)
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func process(conn net.Conn){
	defer conn.Close()

	//读客户端发送的信息
	for {
		mes, err := readPkg(conn)
		if err!=  nil{
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常退出..")
				return 
			}else{
				fmt.Println("readPkg err=",err)
				return
			}
		}
		fmt.Println("读取的消息：",mes)
		err = serverProcessMes(conn, &mes)
		if err!=nil{
			return
		}
	}

	
}

func main(){
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp","0.0.0.0:8889")

	defer listen.Close()

	if err != nil{
		fmt.Println("监听失败！")
		return
	}
	for{
		fmt.Println("等待客户端来链接服务器...")
		conn,err := listen.Accept()
		if err!=nil{
			fmt.Println("listen.Accept error",err)
		}
		go process(conn)
	}
}