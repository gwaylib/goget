# 本地搭建go get的MetaImport

安装
```text
go go -v -u github.com/gwaylib/goget
```

修改自go1.9的go get源码, 编译依赖于go1.9及以上, [go源码地址](https://github.com/golang/go)

使用goget代替go get执行:
```text
原指令
go get 
替换指令
goget 
```

重定向配置.goget时, 优先在当前目录下查找.goget，若没有，查找$PJ_ROOT/.goget

配置.goget文件
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


