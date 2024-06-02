package main
import(
	"fmt"
	"net"
	"bufio"
	"os"
)

func main(){
	//❗️链接服务器（这里为本机的8888端口）
	conn, err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}
	fmt.Println("conn success=",conn)

	for{
		//❗️创建一个reader，其中os.Stdin表示标准输入（即终端）
		reader := bufio.NewReader(os.Stdin) 
		//❗️从终端读取一行用户输入，当读到第一个'\n'换行符时结束
		line, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("readString err=",err)
		}
		if line == "exit\n"{
			fmt.Println("客户端退出了")
			break
		}

		//❗️将输入发送给服务器，注意要将string转为字符数组，返回的n是输入的字节个数
		n, err := conn.Write([]byte(line))
		if err != nil{
			fmt.Println("conn.Write err=",err)
		}
		fmt.Printf("客户端发送了%d字节的数据\n",n)
		
	}
	
	
}