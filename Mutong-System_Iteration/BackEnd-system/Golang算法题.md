# 1 零碎

## 1.1 有两个变量 a和b，要求将其进行交换，但是不允许使用中间变量。

```go
func main(){
  var int a = 10
  var int b = 20
  
  a = a + b
  
  b = a - b // b = a + b - b 
  a = a - b // a = a + b - a
}
```



## 1.2 斐波那契数

**题目：**斐波那契数为 [1,1,2,3,5,8,13...]，随便给一个正整数n，求出它对应的斐波那契数。

**核心：递归**（很标准的递归，**向前找规律**）

**题解：**

```go
import "fmt"

func feibona(n int) int{
  if (n == 1 || n ==2 ){
    return 1
  }
  else{//n>2 ❗️f(n) = f(n-1) +f(n-2)
    return feibona(n-1) + feibona(n-2)
  }
  
}

func main(){
  feibonaNum = feibona(3)
  fmt.Println("斐波那契数： ",feibonaNum)
  
}
```

## 1.3 猴子吃桃子问题

**题目**：有一堆桃子，猴子第一天吃了其中的一半，

**核心：递归**（非标准递归，麻烦一丢丢，**向后找规律**）

**题解：**

```go
import "fmt"


func eatPeach(days int) int{
	if days == 10 {
		return 1
	}else{
    //❗️要向后推，因为已知的条件是最后一天的结果。
    //❗️所以不要找这一天和前一天桃子的关系，要找输入当天和后一天的关系，然后递归
		return  (eatPeach(days+1)+1)*2
	}
}

func main(){
  peachNum = eatPeach(1)
  fmt.Println("第1天桃子有： ",peachaNum)//1534
  
}
```

