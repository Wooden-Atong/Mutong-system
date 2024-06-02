package main
import (
	"redigo/redis"
	"fmt"
)

func main(){
	//❗️go连接到redis
	conn, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("连接失败！",err)
		return
	}

	defer conn.Close()

	//❗️通过go向redis写入单个数据
	_, err := conn.Do("set","name","tom")
	if err != nil{
		fmt.Println("写入出错！",err)
		return
	}
	//❗️通过go从redis读取单个数据，并转成string类型
	r, err := redis.String(conn.Do("get","name"))
	if err != nil{
		fmt.Println("读取出错！",err)
		return
	}
	


	//❗️通过go向redis批量写入数据
	_, err := conn.Do("HMset","user","name","tom","age",18)
	if err != nil{
		fmt.Println("写入出错！",err)
		return
	}
	//❗️通过go从redis批量读取数据，并将单个数据均转成string类型，最后获得的 r 是切片类型
	r, err := redis.Strings(conn.Do("HMget","user","name","age"))
	if err != nil{
		fmt.Println("读取出错！",err)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s",i,v)
	}

	fmt.Println("连接成功1",conn)
}