package process

import(
	"fmt"
	"net"
	"encoding/json"
	"encoding"
)

type UserProcess struct{
	Conn net.Conn
}


func (this *UserProcess)serverProcessLogin( mes *message.Message)(err error){
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
	tf := &utils.Transfer{
		Conn : this.Conn
	}
	err = tf.WritePkg(data)
	return
}