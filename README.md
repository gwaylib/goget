# 本地搭建go get的MetaImport

在中国的某墙的作用下，若不翻墙，无法go get到某些包，于是对go get的源码进行了些修改，就可以本地直接配置go get的MetaImport值，在联网前直接本地重定向。

本代码修改自go1.9的go get源码, 编译依赖于go1.9及以上, [go源码地址](https://github.com/golang/go)

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

# 配置.goget文件

重定向配置.goget时, 优先在当前目录下查找.goget，若没有，查找$PJ_ROOT/.goget

以下是配置文件格式，若未配置，goget源码中已内置以下的内容，可以直接使用goget访问golang.org/x/的包
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


