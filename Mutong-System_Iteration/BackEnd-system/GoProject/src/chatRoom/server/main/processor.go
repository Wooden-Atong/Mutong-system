package main

import(
	"fmt"
	"net"
	"chatRoom/common/message"
	"chatRoom/process/process"
)


type Processor struct {
	Conn net.Conn
}

//根绝客户端发送的消息种类不同，决定调用哪个函数来处理
func (this *Processor)serverProcessMes(mes *message.Message)(err error){
	switch mes.Type{
		case message.LoginMesType:
			up := &UserProcess{
				Conn : this.Conn,
			}
			err =  up.ServerProcessLogin(mes)
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}