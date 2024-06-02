# 0 零碎

## 0.1 有两个变量 a和b，要求将其进行交换，但是不允许使用中间变量。

```go
func main(){
  var int a = 10
  var int b = 20
  
  a = a + b
  
  b = a - b // b = a + b - b 
  a = a - b // a = a + b - a
}
```



## 0.2 斐波那契数

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

## 0.3 猴子吃桃子问题

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



# 1 排序

## 1.1 排序的分类

**1）内部排序**

将需要处理的所有数据都加载到内部存储器中进行排序。

（**交换式排序法**、**选择式排序法**和**插入式排序法**）

**2）外部排序**

数据量过大，无法全部加载到内存中，需要借助外部存储排序。（**合并排序法**和**直接排序法**）

## 1.2 交换式排序

交换式排序属于内部排序法，运用数据值比较后，依判断规则对数据位置进行交换，以达到排序目的。

### 1.2.1 冒泡排序Bubble sort

```go

func BubbleSort(arr *[5]int){
  for i := 0 ; i < len(arr)-1 ; i++ {
    for j := 0; j < len(arr) - 1 - i; j++ {
      if (*arr)[j] > (*arr)[j+1]{
        temp = (*arr)[j]
        (*arr)[j+1] = temp
        (*arr)[j] = (*arr)[j+1]
      }
    }
  }
}
```

### 1.2.2 快速排序Quick sort

# 2 查找

## 2.1 顺序查找

```go
for i:=0; i < len(names); i++ {
  if hereName == names[i]{
    fmt.Printf("找到了！！")
    break
  }
  fmt.Printf("没有找到。")  
}
```



## 2.2 二分查找（🚩面试）

```go
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int){
  //判断leftIndex 是否大于 rightIndex
  if leftIndex > rightIndex{
    fmt.Println("找不到！")
    return
  }
  
  //获取中间下标
  middle = (leftIndex + rightIndex) /2
  
  if (*arr)[middle] > findVal{
    //要查找的数在middle左边
    BinaryFind(arr, leftIndex, middle - 1, findVal)
  }else if (*arr[middle] < findVal){
    //要查找的数在middle右边
    BinaryFind(arr, middle + 1, rightIndex, findVal)
  }else{
    fmt.Println("找到了！下标为：",middle)
  }
}
```

