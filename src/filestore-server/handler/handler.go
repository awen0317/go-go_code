package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/**
   http.ResponseWriter,，用户返回给用户的responsewrite对象
   r *http.Request 用于接受对象请求的request对象指针
 */
func UploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method =="GET"{
		//返回html请求
		data,err :=ioutil.ReadFile("./static/view/index.html")

		if err !=nil{
			io.WriteString(w,"interel server error")
			return
		}
		io.WriteString(w,string(data))

	}else if r.Method =="POST"{
		/**
			1.存储到本地目录
			2.接受到本地文件及存储到本地目录
			3.函数惨数1文件句柄，header头新
		 */


		file,head,err := r.FormFile("file")
		if err !=nil{
			fmt.Printf("Failed to get data,err %s\n",err.Error())
			return
		}
		defer file.Close()

		newFile,err := os.Create("/tmp/"+head.Filename)
		if err !=nil{
			fmt.Printf("Failed to create file,err:%s\n",err.Error())
			return
		}
		defer newFile.Close()

		//
 		_,err =io.Copy(newFile,file)
 		if err !=nil{
 			fmt.Printf("Failed to save data into ,err :%s\n",err.Error())
		}
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Upload finished!")
}
