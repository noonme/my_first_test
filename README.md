# one_qyc_case and mxiaoxu_case
## 序
单侧只是测试代码功能逻辑，如第三方调用,数据库操作等都是非必须的开销，所以可以一定的方法模拟第三方调用，数据库操作等，减少测试代码的开销。



## 示例图
![mark](http://images.opsblogs.cn/blog/20241025/VLJh7priTlnW.png?imageslim)
## 接口定义
面向对象/面向对象提高代码复用性

Go 接口通常属于使用接口类型值的包，而不是实现这些值的包。实现包应返回具体（通常是指针或结构）类型：这样，可以将新方法添加到实现中，而无需进行大量重构。

不要在 API 的实现器端定义“用于模拟”的接口;相反，请设计 API，以便可以使用实际实现的公共 API 对其进行测试。

在使用接口之前不要定义接口：如果没有实际的使用示例，就很难看到接口是否必要，更不用说它应该包含什么方法了。

## 参考:
https://github.com/golang/go/wiki/CodeReviewComments#interfaces

## 第三方工具测试库
### gomock
核心: 
自动构建桩代码，在环境声明好已有的interface接口的基础上

通过EXPECT()方法自动定义桩代码，覆盖源接口的方法，避免原方法带来的开销，如数据库操作等

使用:
```go
依赖包
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen

#在环境声明好或者已有的interface接口的基础上自动构建桩代码

构造的命令
go:generate mockgen -destination=./mock/mock_LearnInf.go -package=mock -source=./interface.go
```


默认只能把对应的方法调用一次，调用多次会报错，或者手动指定调用次数
![mark](http://images.opsblogs.cn/blog/20241026/4IdknuJlnW0X.png?imageslim)
解决
需要设置对应的调用的次数**\.Times\(2\)**，调用次数必须要相等，否则还是会报错
![mark](http://images.opsblogs.cn/blog/20241026/VMcYoEJETYDm.png?imageslim)

可用的方法:

1、设置指定的调用次数
```go
.Times(N)

mockLearnInf.EXPECT().LearnGo(name).Return("Learn Go is ok", nil).Times(2)
```
2、设置宽泛的调用次数，一次或者多次
```go
.AnyTimes()

mockLearnInf.EXPECT().LearnGo(name).DoAndReturn(func(string) (string, error) {
		return "Learn Go is ok", nil
	})

	mockLearnInf.EXPECT().LearnGo(name).Return("Learn Go is ok-case2", nil).AnyTimes()
```
3、写多个方法,实现方法的多次调用，如调用两次就写两个方法的实现
```go
mockLearnInf.EXPECT().LearnGo(name).DoAndReturn(func(string) (string, error) {
		return "Learn Go is ok", nil
	})

	mockLearnInf.EXPECT().LearnGo(name).Return("Learn Go is ok-case2", nil)

	output, err := mockLearnInf.LearnGo(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("output: %s\n", output)

	output2, err := mockLearnInf.LearnGo(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("output: %s\n", output2)
```


### goonkey
相较于gomock更加轻便，功能更加单一
争对具体的实体类和函数进行打桩
没有限制可以多次调用

使用:
```go
go get github.com/agiledragon/gomonkey
```
测试结果有歧义，意思是go版本和gomonkey兼容有问题
![mark](http://images.opsblogs.cn/blog/20241026/3o9PfX0IQc7b.png?imageslim)

![mark](http://images.opsblogs.cn/blog/20241026/eL8tePf9LNwx.png?imageslim)