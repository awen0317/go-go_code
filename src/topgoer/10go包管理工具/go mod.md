#go Mod

###Golang 无法下载依赖解决方案 unrecognized import path 
Go 1.11 版本开始，官方支持了 go module 包依赖管理工具。
1.环境变量中加入下列命令：  

    # Enable the go modules feature
    export GO111MODULE=on
    # Set the GOPROXY environment variable
    export GOPROXY=https://goproxy.io
