package message

//确定一些消息类型
const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
)

 
//定义消息结构
type Message struct{
	Type string `json:"type"` //消息类型
	Data string `json:"data"`//消息内容 ❗️这里用string类型是因为，序列化后的数据其实使用string类型来存的
}

//定义具体消息
//登陆信息
type LoginMes struct{
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}
//登陆结果消息
type LoginResMes struct {
	Code int `json:"code"`//状态码 500表示还没有注册；200表示登陆成功
	Error string `json:"error"`//返回错误信息
}