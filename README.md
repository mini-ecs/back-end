# Mini-ECS 后端代码

产生的日志位于`logs`文件夹下，（可在config.toml文件中更改）。

```shell
\[\d*m
```  

用该正则表达式可以去除日志的乱码字符

## 项目骨架安排

[参考文件](https://github.com/golang-standards/project-layout) 。

`api`文件夹中的代码用于处理前端发送过来的restful请求，请求的路由文件位于`internal/router`目录下。原理是在项目初始化时，将这些api处理函数注册到gin的路由内，
路由管理器在接收到对应的请求时，会调用已注册的函数来响应该请求。


## 备注

因项目开发时间有限，部分测试代码仅用于实验对应的接口是否可行。
