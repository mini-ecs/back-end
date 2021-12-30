# Mini-ECS 后端代码

产生的日志位于`logs`文件夹下，（可在config.toml文件中更改）。

```shell
\[\d*m
```  

用该正则表达式可以去除日志的乱码字符

## 项目骨架安排

[参考文件](https://github.com/golang-standards/project-layout) 。

## [swagger接口](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

在项目根目录运行`sh build/build.sh`即可生成接口文档。