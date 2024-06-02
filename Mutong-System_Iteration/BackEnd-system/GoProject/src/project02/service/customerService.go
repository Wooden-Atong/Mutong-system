package service
import(
	"project02/model"
)

//完成对Customer的操作，包括增删改查
type CustomerService struct{
	Customers []model.Customer
	//表示当前切片有多少客户
	CustomerNum int
}

func NewCustomerService() *CustomerService{
	//初始化一个客户并返回
	return &CustomerService{
		//❗️踩坑了，model.NewCustomer返回的是单个Customer类型，而Customers是customer切片类型，所以要使用[]model.Customer{}进行强转
		Customers:[]model.Customer{model.NewCustomer(1,"张三","男",20,"112","zs@163.com")},
		CustomerNum:1,
	}
}

func (cs *CustomerService) List() []model.Customer{
	return cs.Customers
}

func (cs *CustomerService) Add(customer model.Customer) bool { 
	cs.CustomerNum++
	customer.Id = cs.CustomerNum
	cs.Customers = append(cs.Customers,customer)
	return true
}

func (cs *CustomerService) Delete(index int) {
	
		//❗️从切片中删除一个元素的方法
		//❗️cs.customers[index+1:]...的意思是将切片从索引'index+1'开始的部分展开为单独的参数列表，用于可变参数函数调用
		cs.Customers = append(cs.Customers[:index],cs.Customers[index+1:]...)
		cs.CustomerNum--
}
func (cs *CustomerService) Modify(index int, customer model.Customer){
	cs.Customers[index] = customer
}