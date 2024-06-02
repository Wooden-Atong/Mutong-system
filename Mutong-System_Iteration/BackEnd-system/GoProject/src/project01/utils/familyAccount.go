package utils
import(
	"fmt"
)

type FamilyAccount struct{
	//接受用户输入选项
	key string
	//用于判断是否退出程序
	loop bool
	//定义账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的一个说明
	note string
	//收支详情表
	details string
	//定义一个变量判断是否有收支行为
	flag bool
	//用于判断是否退出程序
	choice string
}

//编写一个工厂模式的构造方法，返回一个*Family Account实例
func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		//接受用户输入选项
		key : "",
		//用于判断是否退出程序
		loop : true,
		//定义账户余额
		balance : 10000.0,
		//每次收支的金额
		money : 0.0,
		//每次收支的一个说明
		note : "",
		//收支详情表
		details : "收支\t账户金额\t收支金额\t说明",
		//定义一个变量判断是否有收支行为
		flag : false,
		//用于判断是否退出程序
		choice : "",
	}
}


func (familyAccount *FamilyAccount) ShowDetails(){
	fmt.Println("--------------当前收支明细表--------------")
	if familyAccount.flag{
		fmt.Println(familyAccount.details)
	}else{
		fmt.Println("当前没有收支信息，来一笔吧！")
	}
}

func (familyAccount *FamilyAccount) RegisterIncomes(){
	fmt.Println("-----------------登记收入----------------")
	fmt.Println("本次收入金额：")
	fmt.Scanln(&familyAccount.money)
	fmt.Println("本次收入说明：")
	fmt.Scanln(&familyAccount.note)
	//有了一笔收支行为
	familyAccount.flag = true
	//余额上涨
	familyAccount.balance += familyAccount.money
	//details表进行拼接
	familyAccount.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",familyAccount.balance,familyAccount.money,familyAccount.note)//❗️fmt.Sprintf()以格式化拼接字符串
}
func (familyAccount *FamilyAccount) RegisterOutcomes(){
	fmt.Println("--------------登记支出-------------")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&familyAccount.money)
	//判断有没有足够的金额支出
	if familyAccount.money>familyAccount.balance{
		fmt.Println("余额不足")
		return
	}
	//有了一笔收支行为
	familyAccount.flag = true
	//余额减少
	familyAccount.balance -= familyAccount.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&familyAccount.note)
	//details表进行拼接
	familyAccount.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v",familyAccount.balance,familyAccount.money,familyAccount.note)//❗️fmt.Sprintf()以格式化拼接字符串
}

func (familyAccount *FamilyAccount) Exist(){
	fmt.Println("你确定要退出吗？ y/n")
	
	for {
		fmt.Scanln(&familyAccount.choice)
		if familyAccount.choice == "y" || familyAccount.choice == "n"{
			break
		}
		fmt.Println("输入有误，请重新输入！")
	}
	if familyAccount.choice == "y"{
		familyAccount.loop = false
	}
	
}

func (familyAccount *FamilyAccount) MainMenu(){
	for {
		fmt.Println("--------------家庭收支记账软件--------------")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Scanln(&familyAccount.key)
		switch (familyAccount.key){
		case "1":
			familyAccount.ShowDetails()
		case "2":
			familyAccount.RegisterIncomes()
		case "3":
			familyAccount.RegisterOutcomes()
		case "4":
			familyAccount.Exist()
		default:
			fmt.Println("请输入正确选项...")
	}
	if !familyAccount.loop{
		break
	}
}
fmt.Println("已退出当前家庭记账收支软件...")
	}


