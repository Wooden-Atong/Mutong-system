package main
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"chatRoom/common/message"
)

//登陆函数
//❗️最好是一个error返回，真假TF返回描述能力太弱了
func login(userId int, userPwd string)(err error){
	// fmt.Println(userId,userPwd)
	// return nil

	//连接到服务器
	conn, err := net.Dial("tcp","localhost:8889")
	if err!= nil {
		fmt.Println("net.Dial error=",err)
		return
	}

	defer conn.Close()

	//通过conn发消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err!=nil{
		fmt.Println("json Marshal error=",err)
		return
	}
	//序列化后的data是一个字符串序列，所以要先string()强转
	//1） 现将mes.Data序列化
	mes.Data = string(data)
	//2） 再把整个mess序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json Marshal error=",err)
		return
	}

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
	if err != nil{
		fmt.Println("conn.Write([]byte(data)) fail",err)
		return
	}

	resMes,err := readPkg(conn)
	if err != nil{
		fmt.Println("readPkg(conn) fail",err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data),&loginResMes)
	
	if loginResMes.Code == 200{
		fmt.Println("登陆成功")
	}else if loginResMes.Code == 500{
		fmt.Println(loginResMes.Error)
	}


	// fmt.Println("客户端发送的消息长度内容切片",buf[:4])
	return
}