# Go MetaImport本地重定向

对go get的源码进行了些修改，使得本地可以直接配置go get的MetaImport值，在连网之前直接本地重定向以便跳过获取go远程包被强的问题。

本代码修改自go1.9的go get源码, 编译依赖于go1.9及以上, [go源码地址](https://github.com/golang/go)。

安装
```text
go go -v -u github.com/gwaylib/goget
```

使用goget代替go get执行:
```text
原指令
go get 
替换指令
goget 
```
用例：
```text
goget -v golang.org/x/net/websocket
```

# 配置.goget文件

查找.goget配置文件时, 优先在当前目录下查找.goget，若没有，查找全局的$PJ_ROOT/.goget

以下是配置文件格式，在goget源码中已内置以下的内容的配置文件，未配置.goget文件可以直接使用goget访问golang.org/x/的包
```text
#
# TODO: Implements goget configuration.
# Library version default by go1 tag
# others version work like gopkg.in version control
#
# this is configuration depend on $PJ_ROOT enviorement for looking up the root of project.
#


# Regexp				MetaImport
golang.org/x/net*:			{"Prefix":"golang.org/x/net", "VCS":"git", "RepoRoot":"https://github.com/golang/net.git"} 

```


