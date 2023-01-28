# go-cover-example

[go-cover](https://github.com/lamber92/go-cover)项目的简单使用示例。


#### 跳转到以下分支查看具体示例与效果

- [CASE1](https://github.com/lamber92/go-cover-example/tree/feature/test_2_api) - 全量/增量覆盖率报告及中间文件展示

#### 执行步骤

1. 安装[goc](https://github.com/qiniu/goc), [go-cover](https://github.com/lamber92/go-cover)
2. 运行goc server
    ```shell
    goc server
    ```
3. 编译本分支的go-cover-example并运行
    ```shell
    git checkout feature/test_2_api
    goc build .
    ./go-cover-example
    ```
4. 请求go-cover-example的/test_2接口
    ```shell
    curl 127.0.0.1:8080/test_2
    ```
    > 输出内容：{"code":0,"data":{"text":"something"},"msg":"success"}
5. 请求goc，获取基于 feature/test_2_api 分支的全量覆盖率基础数据文件：profile
    ```shell
    curl 127.0.0.1:7777/v1/cover/profile > profile
    ```
6. 在go-cover-example根目录下执行以下语句
    ```shell
    go-cover convert ./profile -t main
    ```  
   > 自动获取当前分支相对于main分支的代码差异，并生成全量/增量覆盖率报告html文件。  
   > 全量报告：[full.html](http://htmlpreview.github.io/?https://github.com/lamber92/go-cover-example/blob/feature/test_2_api/full.html)  
   > 增量报告：[diff.html](http://htmlpreview.github.io/?https://github.com/lamber92/go-cover-example/blob/feature/test_2_api/diff.html)    
   > 
   > 从 diff.html 中可以得到 curl 127.0.0.1:8080/test_2 请求后代码的执行痕迹；  
   > 其中 D:\GitHub\go-cover-example\internal\logic\test.go 第17行没有被执行过；
   > 此时执行 curl 127.0.0.1:8080/test_2?flag=1 再重复上面步骤5和步骤6, 即可得到叠加的覆盖结果。