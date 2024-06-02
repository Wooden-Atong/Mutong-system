# 1 Redis基本介绍

## 1.1 Redis基本了解

1）Redis是**NoSQL**数据库，不是传统的关系型数据库。

2）Redis：REmote Dictionary Server（远程字典服务器），Redis性能非常高，单机能够达到15w qps（读取的速度非常快），通常适合做缓存，也可以持久化。

3）完全免费开源，高性能的（key/value）**分布式**内存数据库，基于内存运行并支持持久化的NoSQL数据库，也称为**数据结构**服务器。

## 1.2 Redis原理示意图

![31981716944731_.pic_hd](/Users/mac/Desktop/markdown笔记/Mutong-System_Iteration/BackEnd-system/RedisNoteImage/31981716944731_.pic_hd.jpg)

## 1.3 Redis的五大数据类型

### 1.3.1 String（字符串）

1）string是redis最基本的类型，**一个key对应一个value**

2）string类型是二进制安全的，除了普通的字符串文本外，也可以存放图片等数据（不过一般没人把图片存放进去）。

3）redis中一个字符串value最大是512M。

### 1.3.2 Hash（哈希）

1）Redis中的hash是一个**键值对集合**（**key不能重复但值可以**），类似go中的map

2）Redis中的hash是一个string类型的field和value的映射表，特别适合存储对象。

### 1.3.3 List（列表）

1）redis中的list就是简单的string列表，按照插入顺序排序（⚠️**比如从左边插入，那么最先插入的就在最右边，但是它的下标为0**），可以从头部或者尾部插入。

2）redis中的list本质是**链表**，元素是有序的，元素的值可以重复。

3）redis中的list本质**存储的顺序和下标并不对应**（最左边的元素下标一定为0，如果从左向右插入的话，它是最后插入的）



<img src="/Users/mac/Desktop/markdown笔记/Mutong-System_Iteration/BackEnd-system/RedisNoteImage/32031716988502_.pic.jpg" alt="32031716988502_.pic" style="zoom:70%;" />

### 1.3.4 Set（集合）

1）redis中的set是**string类型的集合**，字符串元素是**无序的**，而且**元素的value都不能重复**

2）底层是HashTable数据结构，

### 1.3.5 zset（有序集合）

1）类似set只不过元素是**有序**的。

# 2 Redis基本使用

## 2.1 Redis启动及连接

**Redis命令一览：http://doc.redisfans.com

**启动：**终端输入redis-server

**连接：**新开终端输入redis-cli

**说明：**Redis安装后，默认有16个数据库，初始默认使用0号数据库，编号是0...15

## 2.2 Redis基本操作

```bash
#⚠️这里是在redis-cli中的指令操作⚠️

#1⃣️切换数据库 [select]
select 1
#2⃣️查看当前数据库的key-value数量 [dbsize]
dbsize 1 #不写index默认查看当前数据库的key-value数量
#3⃣️清空当前数据库的key-value [flushdb]
flushdb
#4⃣️清空16个数据库所有key-value [flushall]
flushall
```



# 3 Redis对string的操作

## 3.1 Redis对string的CRUD操作

### 3.1.1 RU：添加（存在则相当于修改）key-value

```bash
# [set]
set key1 value1
# [setex] 指定时间内，在内存中存在，超过这个时间则销毁
setex key1 10 hello
# [mset] 一次性设置一个或多个key-value(如果原先key已经存在，则会覆盖)
mset key1 value1 key2 value2
```

### 3.1.2 C：获取key对应的value

```bash
# [get]
get key1
# [mget] 一次性获取一个或多个key-value
mget key1 key2
```

### 3.1.3 D：删除key-value

```bash
# [del]
del key1
```

# 4 Redis对Hash的操作

## 4.1 Redis对hash的CRUD操作

### 4.1.1 RU：添加（存在则相当于修改）key对应的field-value

```bash
# [hset]
hset key1 name "smith"
hset key1 age 20
hset key1 job "golang coder"
# [hmset] 一次性存入多个field-value
hmset key1 name "smith" age 20 job "golang coder"
```

### 4.1.2 C：获取key中某个field对应的value

```bash
# [hget]
hget key1 name
hget key1 age
# [hmget] 一次性取出多个field-value
hmget key1 name age job
# [hgetall] 一次性将key中所有field-value都取出来
hgetall key1

```

### 4.1.3 D：删除key中某个field对应的value

```bash
# [hdel]
hdel key1 name
```

## 4.2 hash使用细节和注意事项

1）统计hash有多少个field- value：hlen key1

2）判断hash表key中是否有某个field字段：hexists key1 field



# 5 Redis对List的操作

## 5.1 Redis对list的CRUD操作

### 5.1.1 RU：添加元素到list中

```bash
# [lpush] / [rpush] 从 左边 或 右边 插入元素
lpush key1 hello
rpush key1 world
```

### 5.1.2 C：返回列表中指定区间内的元素

```bash
# [lrange] 只取数据不拿走
lrange key1 0 6 #返回列表key1中下标为0到下标为6的元素 ⚠️左闭右闭
lrange key1 0 -1 #返回列表key1中下标为0到-1的元素
# [lpop] [rpop] 从最左边 或 最右边 取出并拿走数据
lpop key1
rpop key1
```

### 5.1.3 D：删除list

```bash
# [del] ⚠️删掉后key1还在，只不过其中放了一个nil值，再删一次则list变空列表完全不存在了
del key1
```

## 5.2 list使用细节和注意事项

1）按照索引获取元素（从左到右，编号从0开始）：lindex key1 index

2）返回list的长度，如果list不存在，返回0：llen key1

3）如果list中元素没有了，那么这个list对应的key（list名字）也就消失了



# 6 Redis对Set的操作

## 6.1 Redis对Set的CRUD操作

### 6.1.1 RU：添加元素到list中

```bash
# [sadd]
sadd key1 hello world
```

### 6.1.2 C：取出所有值

```bash
# [smenmers]
smenmers key1
```

### 6.1.3 D：删除set中指定值

```bash
# [srem]
srem key1 hello
```

## 6.2 set使用细节和注意事项

1）判断set中有没有某个值：sismenber key1 hello （是 就返回 1；反之返回 0）

# 7 Redis对zset的操作

-暂无-有需要再回来补充

# 8 Golang操作redis

## 8.1 go连接redis

### 8.1.1 go安装第三方库redis

1）获取GOROOT路径并切换进去（⚠️我这里本机配置是GOROOT下放包等，GOPATH下就是我自己的go项目）

```bash
go env GOROOT
```

⚠️要确保GO111MODULE是on的状态，否则执行：export GO111MODULE=on

2）在GOPATH路径下执行命令

```bash
go install github.com/garyburd/redigo/redis@latest
```

3）还是不行就手动去github指定路径下载然后放在GOROOT位置下面

（我这里把redigo-master改名为redigo）

<img src="/Users/mac/Desktop/markdown笔记/Mutong-System_Iteration/BackEnd-system/RedisNoteImage/32041716992966_.pic_hd.jpg" alt="32041716992966_.pic_hd" style="zoom:50%;" />

### 8.1.2 go连接redis

```go
//❗️go连接到redis
	conn, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("连接失败！",err)
		return
	}

	defer conn.Close()
```

## 8.2 go操作redis

### 8.2.1 go操作redis-单个数据

（给的例子是set字符串，其他单个数据都类似，只不过Do当中的指令变一下）

```go
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
```

### 8.2.2 go操作redis-批量数据

（给的例子是hmset哈希，其他批量数据类似，Do当中指令变一下）

```go
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
```

**注意⚠️：**

1）批量返回将元素用redis.String()转成字符串，得到的是一个字符串切片。

# 9 Redis连接池

## 9.1 Redis连接池流程

1）事先初始化一定数量链接，放入链接池

2）当go需要操作redis时，**直接从redis链接池中取出链接**即可。

3）这样可以**节省临时获取redis链接的时间**，从而提高效率

## 9.2 Redis链接池核心代码

```go
var pool *redis.Pool
pool = &redis.Pool{
  MaxIdle : 8, //最大空闲连接数
  MaxActive : 0, //表示和数据库的最大连接数，0表示没有限制
  IdleTimeout : 100, //最大空闲时间
  Dial: func() (redis.Conn, error){//初始化链接，指定协议 及 链接哪个ip监听哪个端口的redis
    return redis.Dial("tcp","localhost:6379")
  },
}
c := pool.Get()//从链接池取出一个链接
pool.Close() //关闭链接池，再也不能从链接池中取出链接了
```

