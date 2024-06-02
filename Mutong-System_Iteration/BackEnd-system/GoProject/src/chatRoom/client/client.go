package main
import (
	"fmt"
)

var userId int
var userPwd string

func main(){
	//接收用户的选择
	var key int
	//判断是否还继续显示界面
	var loop = true
	for loop{
		fmt.Println("------------------欢迎登陆多人聊天系统----------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 4 请选择(1-3)")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("\t\t\t 1 登陆聊天室")
			loop = false
		case 2:
			fmt.Println("\t\t\t 2 注册用户")
			loop = false
		case 3:
			fmt.Println("\t\t\t 3 退出系统")
			loop = false
		default:
			fmt.Println("\t\t\t 输入有误，请重新输入")
		}
	}
	if key ==1 {
		fmt.Println("请输入用户id")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&userPwd)
		err := login(userId,userPwd)
		if err != nil{
			fmt.Println("登陆失败")
		}
	}
}
