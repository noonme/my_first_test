# one_qyc_case and mxiaoxu_case
## 示例图
![mark](http://images.opsblogs.cn/blog/20241025/VLJh7priTlnW.png?imageslim)
## 接口定义
Go 接口通常属于使用接口类型值的包，而不是实现这些值的包。实现包应返回具体（通常是指针或结构）类型：这样，可以将新方法添加到实现中，而无需进行大量重构。

不要在 API 的实现器端定义“用于模拟”的接口;相反，请设计 API，以便可以使用实际实现的公共 API 对其进行测试。

在使用接口之前不要定义接口：如果没有实际的使用示例，就很难看到接口是否必要，更不用说它应该包含什么方法了。

## 参考:
https://github.com/golang/go/wiki/CodeReviewComments#interfaces