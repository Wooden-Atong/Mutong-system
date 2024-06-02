package model

import(
	"fmt"
)

//声明一个Customer结构体，表示一个客户信息
type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个工厂模式，返回customer实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer{
	return Customer{
		Id:id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer{
	return Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}
func (c Customer) GetInfo()string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",c.Id,c.Name,c.Gender,c.Age,c.Phone,c.Email)
	return info
}