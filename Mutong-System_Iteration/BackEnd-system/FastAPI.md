# 0.并发async/await

## 0.1 并发（异步）和并行的区别

1）对于**I/O密集型**操作，倾向于使用**并发**。

​      因为I/O密集型操作，很多时候需要等待，并发允许在等待时去处理其他事情。

2）对于**CPU密集型**操作，倾向于使用**并行**。

​     因为需要处理的任务很多，基本上一直需要工作，很少有等待时间，所以用多线程都专注于自己的事情，并行处理效率更高。

## 0.2 使用异步编程

```python
async def get_burgers(number:int):
  return burgers

#告诉python必需等待get_burgers(2)完成它的工作，此时python就会知道它可以去做其他事情。
burgers=await get_burgers(2)
```

 **注意⚠️：**

1）只能对async def 创建的函数使用await

2）带async def的函数也只能在async def定义的函数内部调用。

# 1.快速上手

# 2.路径参数

## 2.1 直接声明路径参数

```python
from fastapi import FastAPI

app = FastAPI()

#路径中的item_id作为参数传入下方函数
@app.get("/items/{item_id}")
async def read_item(item_id):
    return {"item_id": item_id}

```



## 2.2 路径参数类型转换

```python
from fastapi import FastAPI

app = FastAPI()

@app.get("/items/{item_id}")
#从路径中传入的item_id参数在进入函数后转换为int型
async def read_item(item_id: int):
  #此时返回的item_id就是一个整数
    return {"item_id": item_id}

```

**参数类型**：int、str、float等

**注意⚠️：**如果已经声明为int，输入item_id是str或float，则会报错



## 2.3 声明路径顺序的重要性

```python
from fastapi import FastAPI

app = FastAPI()

#先声明则当成一个单独的me路径，如果后声明，传入则会先被当成user_id，执行read_user操作。
@app.get("/users/me")
async def read_user_me():
    return {"user_id": "the current user"}


@app.get("/users/{user_id}")
async def read_user(user_id: str):
    return {"user_id": user_id}

```



## 2.4 路径参数预设值（枚举型）

```python
from enum import Enum

from fastapi import FastAPI

#继承自str和Enum的自定义子类
class ModelName(str, Enum):
    alexnet = "alexnet"
    resnet = "resnet"
    lenet = "lenet"


app = FastAPI()


@app.get("/models/{model_name}")
#拿到model_name将其转换为ModelName枚举类型
async def get_model(model_name: ModelName):
  #这里是枚举类型名称比较，所以不是==
    if model_name is ModelName.alexnet:
        return {"model_name": model_name, "message": "Deep Learning FTW!"}
#这里是值进行比较可以直接==。注意⚠️："lenet"等价于ModelName.lenet.value。
    if model_name.value == "lenet":
        return {"model_name": model_name, "message": "LeCNN all the images"}
#返回枚举成员
    return {"model_name": model_name, "message": "Have some residuals"}

```



## 2.5 路径参数本身就是路径

**注意⚠️：不允许在不声明路径参数的情况下直接包含路径。**比如@app.get("/files/{file_path}")，此时file_path是一个路径，但此时像普通传参一样直接传进去了，而这是不被允许的。

```python
from fastapi import FastAPI

app = FastAPI()

#这里在get路径时声明了一个包含路径的参数
@app.get("/files/{file_path:path}")
#在获得这个包含路径参数供下面函数使用时再转换为str
async def read_file(file_path: str):
    return {"file_path": file_path}

```

# 3.查询参数

```python
from typing import Union

from fastapi import FastAPI

app = FastAPI()


@app.get("/users/{user_id}/items/{item_id}")
async def read_user_item(
    user_id: int, item_id: str, q: Union[str, None] = None, short: bool = False, needy:str
):
    item = {"item_id": item_id, "owner_id": user_id}
    if q:
        item.update({"q": q})
    if not short:
        item.update(
            {"description": "This is an amazing item that has a long description"}
        )
    return item

```



## 3.1 基础使用

### 3.1.1 查询参数默认值（设置不为None的默认值）

​		  short : bool = False

​		  将short查询参数转为bool值，并且默认为False。

​		  **注意⚠️：**则http://127.0.0.1:8000/items/和http://127.0.0.1:8000/items/short=bool等价。

### 3.1.2 查询参数可选值（默认值为None）

​		  q: Union[str, None] = None

​		  将q查询参数指定为自定义的Union型，并且默认为None。

### 3.1.3 查询参数必需值（不设置默认值）

​		  needy:str

​		  指定needy为str类型，类似路径参数，但无法在路径当中查找到这个参数所以会被识别为查询参数。

​		  **注意⚠️：**此时不指定默认值如果在路径中没有在 ?后接needy=值 则会报错。（所以说必需）

​		  **注意⚠️：**直接needy不指定类型则没有进行类型转换。

### 3.1.4 多路径参数和查询参数

​		  如上代码所示，在函数传参处，⚠️**不需要按照什么特殊顺序去传入路径参数和查询参数**。路径参数会自动根据参数名被识别（如user_id, item_id），剩下的没有对应到路径参数的则是查询参数。

## 3.2 字符串校验



# 4.请求体

## 4.1 基础定义

**请求体**是客户端发送给API的数据，而**响应体**是API发送给客户端的数据。

我们定义的API几乎总是要发送**响应体**，但是客户端却不一定总是发送**请求体**。

## 4.2 请求体的定义和使用（Basic）

### 4.2.1 代码

```python
from typing import Union

from fastapi import FastAPI
#要先从pydantic中引入BaseModel
from pydantic import BaseModel

#自定义请求体，继承BaseModel
class Item(BaseModel):
  #名称：指定类型
  #和前面的查询参数类似，不指定默认值（如name和price）则他们是必需的；如果指定默认值（如description和tax）则他们是可选（请求时写或不写）。
    name: str
    description: Union[str, None] = None
    price: float
    tax: Union[float, None] = None


app = FastAPI()


@app.post("/items/")
#将item转化为Item请求体
async def create_item(item: Item):
  #将item转换为标准的python字典形式，便于后续读取操作
    item_dict = item.dict()
    if item.tax:
        price_with_tax = item.price + item.tax
        item_dict.update({"price_with_tax": price_with_tax})
    return item_dict
```

### 4.2.3 item格式辨析

**1）**当Fast API去读取客户端的请求时候是正常的按照**json格式**读取的；

**2）**如上述代码，当进入creat_item函数之后会将读入的json转化为自定义的**Pydantic模型对象**（会检验是否正确）；

**3）**接着在item.dict()这一步又将其转换成了标准的**python字典**。



### 4.2.2 Union[str,None]

**1）使用注意**

​	 似乎**python3.6-python3.10**用它；python3.10+似乎换了一种使用方法。

​	 用前要记得引入：**from typing import Union**

**2）作用**

​	 表示一个变量可以具有**两种可能的类型**，如str和None。

## 4.3 请求体+路径参数+查询参数



```python
from typing import Union

from fastapi import FastAPI
from pydantic import BaseModel


class Item(BaseModel):
    name: str
    description: Union[str, None] = None
    price: float
    tax: Union[float, None] = None

app = FastAPI()

@app.put("/items/{item_id}")
#如下一行，其中包含了请求体、路径参数、查询参数
async def create_item(item_id: int, item: Item, q: Union[str, None] = None):
    result = {"item_id": item_id, **item.dict()}
    if q:
        result.update({"q": q})
    return result
```

### 4.3.1 FastAPI识别请求体、路径参数、查询参数

1）如果**路径中有**，则被识别为**路径参数**

2）如果被声明为一个**Pydantic模型**，则被视作**请求体**

3）如果没有上述两种特殊情况，则被视作**查询参数**



# 5.查询参数和字符串校验



🍦🍦🍦🍦🍦5这一整块其实都是在教Query的各种用法满足各种功能。





## 5.1 长度校验

```python
from fastapi import FastAPI,Query
from typing import Union
app = FastAPI()


@app.get("/items/")
# 校验字符串最小长度为2，最大长度为5，并且符合pattern的正则表达式。
async def read_items(q: Union[str, None] = Query(default=None, min_length=3,max_length=50,pattern="^fixedquery$"):
    ...
```



## 5.2 是否必需校验

```python
# q是必须参数（省略default值）
  
# q是必须参数（将default值设为... 或required）
  
# q是必需的 但也可以是None
q: Union[str, None] = Query(default=...)

```

## 5.3 查询参数列表（多值）

q: Union[List[str], None] = Query(default=None)

上述List注意：必须显式用query声明，要不然会被解释为请求体



List[str]相当于加了一个校验全部为str的条件。可以直接List

## 5.4高级玩法

### 5.4.1 声明更多元数据

在Query中加title，加description。

就是说会声明一些额外信息，不过似乎是在文档界面或用其他工具使用，不是直接呈现在url中的。

### 5.4.2 别名参数

在Query的alias=“别名”，就是说假如url中这个参数名python解析会错误（等到进到python中来可能参数名变了，但你不想变），那我用这个别名查询一下（感觉用处不大。。。也没特别理解）

### 5.4.3 弃用参数

不影响客户端，但想让大家在文档当中知道被弃用。Query中加deprecated=True。

# 6.路径参数和数值校验

## 6.1 路径参数的字符串校验

类似上述Query，这里就是Path，可以声明上述Query可以声明的所有参数，作用也相同（也就是字符串校验）。

唯一不同的是无论如何声明，参数都是必须的不存在可选。

## 6.2 参数按需排序

⚠️如果你将带有「默认值」的参数放在没有「默认值」的参数之前，Python 将会报错。

排序做法：

1）不带默认值的参数放在最前面。

2）最前面为*。解释为函数第一个参数，但没有任何实际作用（kwargs），并且将它之后的都视为关键字参数（键值对），也就不会因为有无参数值的顺序报错了。



## 6.3 路径参数的数值校验

整型和浮点型都适用：（就是在Path当中国加的参数ge、gt、le、lt）

大于等于 ge 

大于gt（grater than）

小于等于le

小于lt （less than）

# 7.请求体-多参数

🍦就是解决多参数请求混合使用的时候这个请求体格式问题怎么处理。

## 7.1 Annotated[]使用

```python
from typing import Annotated
函数传参处比如：
item_id:Annotated[int,Path(title="")]
```

Union也是typing中引进来，使用也差不多这个格式，但是Annotated跟高级准许int、Path等杂糅在一起。

## 7.2 多个请求体参数结果

原本多个请求体（且请求体参数名也在）会再被一个{}包起来，也就是还是一个请求体形式。

## 7.3 Body()使用

作用：单一值的请求体防止被识别为查询参数。

使用Body()，和Query、Path类似，该有的属性都有，只不过这里为了强调importence是一个请求体而不是查询参数。

```python
async def update_item(
    item_id: int, item: Item, user: User, importance: Annotated[int, Body()]
):
```



# 8.请求体-字段

## 8.1 Field()使用

```python
#用法类似Query、Path和Body，浩瀚的参数也相同
#但要注意import是从pydantic引入，而不是从fastapi引入
from pydantic import Field,BaseModel

# Field使用是在Pydantic模型内部声明校验和元数据
#（比如下面继承BaseModel的Pydantic模型Item）
class Item(BaseModel):
    name: str
    description: Union[str, None] = Field(
        default=None, title="The description of the item", max_length=300
    )
    price: float = Field(gt=0, description="The price must be greater than zero")
    tax: Union[float, None] = None
```

# 9.请求体-嵌套模型

## 9.1 List和Set

### 9.1.1 基本使用

```python
# python3.8+版本
from typing import List,Set
from pydantic import BaseModel

class Item(BaseModel):
  #声明子类型都为str
  tags1:List[str]=[]
  tags2:Set[str]=()

# python3.9+,3.10+
# 可以不引入List、Set，直接使用list和set
class Item(BaseModel):
  # 不声明子类型
  tags3:list=[]
  # 声明子类型都为str
  tags4:list[str]=[]
  tags5:set[str]=()
```

### 9.1.2 纯列表请求体

```python
#类似python转类型
#由于一般继承了BaseModel定义的请求体都是json格式，可以在穿参数的时候使用list[]，将它转成json array（python list）
async def create_multiple_images(images: list[Image]):
```



## 9.2 嵌套模型

```python
from pydantic import BaseModel

class Item(BaseModel):
  name:str
  price:int
#Student嵌套Item，Item作为其中的子模型
class Student(BaseModel):
  #单个子模型
  imfo:Item
  #一组子模型
  imfo_li:list[Item]
  description:str
```

不止上面只嵌套一次，它完全可以深度嵌套。

## 9.3 补充知识

### 9.3.1 特殊类型和校验

详见：https://docs.pydantic.dev/latest/api/base_model/

比如HttpUrl

```python
from pydantic import BaseModel, HttpUrl
class Image(BaseModel):
    #将被检查是否为有效url而不仅仅是str
    url: HttpUrl
    name: str
```



### 9.3.2 任意dict构成请求体

```python
async def get_item(dict[int,str]):
```

1）虽然一般请求体都是pedantic模型构建，但是直接用dict构建也可以。

2）dict的键和值可以定义任意类型，可以进行数据类型转换（比如str->int）。

3）JSON仅支持str作为键。

# 10.模式的额外信息-例子

⚠️：不会有额外的作用，只是添加了一个注释，帮助文档

## 10.1 pydantic的schema_extra添加额外信息

```python
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()


class Item(BaseModel):
    name: str
    description: str | None = None
    price: float
    tax: float | None = None
  #在model_config下的json_schema_extra下加examples
    model_config = {
        "json_schema_extra": {
            "examples": [
                {
                    "name": "Foo",
                    "description": "A very nice Item",
                    "price": 35.4,
                    "tax": 3.2,
                }
            ]
        }
    }
```



## 10.2 Field函数添加额外信息

```python
from pydantic import BaseModel,Field

class Item(BaseModel):
  description:float=Field[exsamples=[...]]
```

## 10.3 Body添加额外信息(搭配Annotated)

```python
from fastapi import Body
from pydantic import BaseModel
from typing import Annotated

class Item(BaseModel):
  name:str
  price:int
    
async def use_item(item:Annotated[Item,Body(examples={...})]):
  
```

11.
