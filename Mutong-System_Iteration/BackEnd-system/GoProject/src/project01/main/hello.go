// 表示hello.go所在的包是main go中每个文件都必须归属一个包
package main
//引入一个包
import "fmt"
// func是关键字，表示后面是一个函数，main是一个主函数（程序入口）


func eatPeach(days int) int{
	if days == 10 {
		return 1
	}else{
		return  (eatPeach(days+1)+1)*2
	}
	
	
  }
func main(){
	// 调用fmt包中的Println函数，输出字符串
	fmt.Println("hello, world!")

	//声明i是一个int变量
	var i int
	//给变量i赋值
	i = 10
	//打印使用变量
	fmt.Println(i)

	var n6,n7 int64 = 3,4
	fmt.Println(n6,n7)

	var n_3w int = 2

	fmt.Println(n_3w)

	if age:= 20; age>18 {
		fmt.Println("hhhh")
		fmt.Println(age)
	}

	// fmt.Println(age)



	var peachNum int


	peachNum = eatPeach(1)// 输入1-10之间的某一天
	fmt.Println("第一天的桃子有： ",peachNum)
	

}