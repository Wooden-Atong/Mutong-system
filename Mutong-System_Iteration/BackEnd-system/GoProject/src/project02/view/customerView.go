package main
import (
	"fmt";
	"project02/service";
	"project02/model";
)

//就在本包内用所以不用大写
type customerView struct{
	//接受用户输入
	key string
	//判断是否退出
	loop bool
	//增删查改各种操作服务
	customerService *service.CustomerService
}
func (cv *customerView) customerServiceList() {
	customers := cv.customerService.List()
	fmt.Println("---------------------客户列表---------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0; i <len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------------客户列表完成--------------------")
}
func (cv *customerView) customerServiceAdd(){
	fmt.Println("添 加 客 户")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)

	newCustomer := model.NewCustomer2(name,gender,age,phone,email)
	cv.customerService.Add(newCustomer)
}

func (cv *customerView)findCustomerById(id int) int {
	for i:=0;i<cv.customerService.CustomerNum;i++{
		if cv.customerService.Customers[i].Id==id{
			return i
		}
	}
	return -1
}

func (cv *customerView) customerServiceDelete(){
	fmt.Println("删 除 客 户")
	id := 0
	fmt.Println("请输入要删除的客户的id号：")
	fmt.Scanln(&id)
	key:=""
	index := cv.findCustomerById(id)
	if index!=-1 {
			for{
				fmt.Println("确认要删除吗？y/n")
				fmt.Scanln(&key)
				if key == "y" || key =="n"{
					if key =="y"{
						cv.customerService.Delete(index)
						fmt.Printf("id为%v的客户已被删除！\n",id)
					}
					break
				}
			}
		}else{
		fmt.Printf("没有id为%v的客户\n",id)
	}
}
func (cv *customerView)customerServiceModify(){
	fmt.Println("修 改 客 户")
	fmt.Println("请输入要修改的客户id：")
	id:=0
	fmt.Scanln(&id)
	index := cv.findCustomerById(id)
	if index!= -1{
		fmt.Printf("姓名：%v\n" , cv.customerService.Customers[index].Name)
		fmt.Println("请输入修改后的姓名：")
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("性别：%v\n" , cv.customerService.Customers[index].Gender)
		fmt.Println("修改后的性别：")
		gender := ""
		fmt.Scanln(&gender)
		fmt.Printf("年龄：%v\n" , cv.customerService.Customers[index].Age)
		fmt.Println("修改后的年龄：")
		age := 0
		fmt.Scanln(&age)
		fmt.Printf("电话：%v\n" , cv.customerService.Customers[index].Phone)
		fmt.Println("修改后的电话：")
		phone := ""
		fmt.Scanln(&phone)
		fmt.Printf("电邮：%v\n" , cv.customerService.Customers[index].Email)
		fmt.Println("修改后的电邮：")
		email := ""
		fmt.Scanln(&email)

		customer := model.NewCustomer2(name,gender,age,phone,email)
		cv.customerService.Modify(index,customer)
	}else{
		fmt.Printf("没有ID为%v的客户",id)
	}

}

//❗️掉了个坑，结构体是值类型，默认传递，最初我写cv customerView，里面修改cv.loop根本不顶用，因为实际值并没有修改到值无法退出程序
//❗️所以应该还是要用指针类型传递。
func (cv *customerView) exist(){
	cv.loop = false
}

func (cv *customerView) mainMenu(){
	for{
		fmt.Println("-------------------客户信息管理软件-------------------")
		fmt.Println("                   1 添 加 客 户")
		fmt.Println("                   2 修 改 客 户")
		fmt.Println("                   3 删 除 客 户")
		fmt.Println("                   4 客 户 列 表")
		fmt.Println("                   5 退       出")
		fmt.Println("请选择（1 - 5）")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.customerServiceAdd()
		case "2":
			cv.customerServiceModify()
		case "3":
			cv.customerServiceDelete()
		case "4":
			cv.customerServiceList()
		case "5":
			cv.exist()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !cv.loop{
			break
		}
	}
}

func main(){

	customerView := &customerView{
		key:"",
		loop:true,
		customerService:service.NewCustomerService(),
	}
	

	customerView.mainMenu()

}