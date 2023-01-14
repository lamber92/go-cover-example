# go-cover-example

[go-cover](https://github.com/lamber92/go-cover)项目的简单使用示例。


#### 跳转到以下分支查看具体示例与效果

- [CASE1](https://github.com/lamber92/go-cover-example/tree/feature/test_2_api) - 全量/增量覆盖率报告及中间文件展示


#### 执行步骤

- 安装[goc](https://github.com/qiniu/goc)
- 运行goc server
```shell
goc server
```
- 编译本分支的go-cover-example并运行
```shell
goc build .
./go-cover-example
```
- 请求go-cover-example的/test_2接口
```shell
curl 127.0.0.1:8080/test_2
```
> {"code":0,"data":{"text":"something"},"msg":"success"}
- 请求goc，获取go cover-profile
```shell
curl 127.0.0.1:7777/v1/cover/profile > profile
```
- 未完待续...