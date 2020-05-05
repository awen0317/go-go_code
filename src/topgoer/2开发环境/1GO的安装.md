#1. Go的安装

##1.1. 下载地址

Go官网下载地址：https://golang.org/dl/ (打开有点慢)

##1.2. Windows安装

Windows安装包

![](.1GO的安装_images/26f2b78e.png)

双击文件

![](.1GO的安装_images/28afeecb.png)


定要记住这个文件的位置后面还有用

![](.1GO的安装_images/0ee74c10.png)

##1.3. Linux下安装

1.SSH远程登录你的linux服务器

2.安装 mercurial包

        [root@localhost ~]# yum install mercurial
        
3.安装git包

        [root@localhost ~]# yum install git
        
4.安装gcc

        [root@localhost ~]# yum install gcc
        
5.下载Go的压缩包:(可选择最新的Go版本)

        [root@localhost ~]# cd /usr/local/
        
        [root@localhost local]# wget https://go.googlecode.com/files/go1.13.linux-amd64.tar.gz
        
注意：如果不能翻墙，去go语言资源站 下载相应的包。然后通过ftp上传到此目录。


6.下载完成 or ftp上传完成，用tar 命令来解压压缩包。

    [root@localhost local]# tar -zxvf go1.13.linux-amd64.tar.gz
    
7.建立Go的工作空间（workspace，也就是GOPATH环境变量指向的目录）

GO代码必须在工作空间内。工作空间是一个目录，其中包含三个子目录：

src ---- 里面每一个子目录，就是一个包。包内是Go的源码文件

pkg ---- 编译后生成的，包的目标文件

bin ---- 生成的可执行文件

这里，我们在/home目录下, 建立一个名为go(可以不是go, 任意名字都可以)的文件夹， 然后再建立三个子文件夹(子文件夹名必须为src、pkg、bin)。


    [root@localhost local]# cd /home/
    [root@localhost home]# mkdir go
    [root@localhost home]# cd go/
    [root@localhost go]# mkdir bin
    [root@localhost go]# mkdir src
    [root@localhost go]# mkdir pkg
    
8.添加PATH环境变量and设置GOPATH环境变量


    [root@localhost go]# vi /etc/profile
    

加入下面这三行:

        export GOROOT=/usr/local/go        ##Golang安装目录
        export PATH=$GOROOT/bin:$PATH
        export GOPATH=/home/go  ##Golang项目目录
        


保存后，执行以下命令，使环境变量立即生效:

    [root@localhost go]# source /etc/profile ##刷新环境变量

至此，Go语言的环境已经安装完毕。

9.验证一下是否安装成功，如果出现下面的信息说明安装成功了

    [root@localhost go]# go version        ##查看go版本
    go version go1.13 linux/amd64
    
