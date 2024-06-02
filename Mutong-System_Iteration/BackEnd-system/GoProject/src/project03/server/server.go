package main
import (
	"fmt"
	"net" //网络socket开发，net包含我们需要的方法和函数
	"io"
)

//❗️conn是一个net.Conn类型接口
func process(conn net.Conn){
	
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte,1024)
		fmt.Printf("服务器在等待客户%s发送信息",conn.RemoteAddr().String())
		//❗️从conn读取，如果客户端一直没有Write，那么就一直阻塞在这里
		//❗️这个n是读到的信息字节数
		n , err := conn.Read(buf)
		//❗️判断客户端是否退出
		if err == io.EOF{
			fmt.Println("客户端退出")
			return
		}
		//显示客户端发送的内容到服务器终端 ❗️注意这个buf要用长度n截取一下，否则可能有额外问题
		fmt.Print(string(buf[:n]))
	}

}

func main(){
	fmt.Println("服务器开始监听....")
	//❗️意思是：基于tcp协议在本地监听8888端口（这里0.0.0.0不仅支持ipv4也支持ipv6）
	listen,err := net.Listen("tcp","0.0.0.0:8888")//❗️"127.0.0.1:8888"也可以，专属ipv4使用
	if err != nil{
		fmt.Println("listen err=",err)
		return
	}

	defer listen.Close()//延时关闭listen

	//循环链接等待客户端来链接（多个客户端）
	for {
		fmt.Println("等待客户端链接...")
		//❗️等待客户端链接
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept() err=",err)
		}else{
			//❗️打印客户端地址和端口（客户端的端口是随机分配的）
			fmt.Printf("Accept() success conn=%v 客户端ip=%v\n",conn,conn.RemoteAddr().String())
		}
		go process(conn)
		
	}

	// fmt.Println("listen success = ",listen)
}

