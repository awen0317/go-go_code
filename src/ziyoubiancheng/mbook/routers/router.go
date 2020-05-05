package routers

import (
	"ziyoubiancheng/mbook/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//test
	//beego.Router("/initdata", &controllers.MainController{}, "get:InitData")

	//首页&分类&详情
	beego.Router("/", &controllers.HomeController{}, "get:Index")	//首页
	beego.Router("/2", &controllers.HomeController{}, "get:Index2")
	beego.Router("/explore", &controllers.ExploreController{}, "*:Index") //点击分类之后出现有哪些书
	beego.Router("/books/:key", &controllers.DocumentController{}, "*:Index") //点击数据，显示分级信息，并显示章节

	//读书
	beego.Router("/read/:key/:id", &controllers.DocumentController{}, "*:Read") //点击阅读
	beego.Router("/read/:key/search", &controllers.DocumentController{}, "post:Search") //点击所搜搜搜

	//搜索
	//beego.Router("/search", &controllers.SearchController{}, "get:Search")
	//beego.Router("/search/result", &controllers.SearchController{}, "get:Result")
	beego.Router("/search", &controllers.ElasticsearchController{}, "get:Search")
	beego.Router("/search/result", &controllers.ElasticsearchController{}, "get:Result")

	//login
	beego.Router("/login", &controllers.AccountController{}, "*:Login")  //登陆
	beego.Router("/regist", &controllers.AccountController{}, "*:Regist") //注册
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout") //登出
	beego.Router("/doregist", &controllers.AccountController{}, "post:DoRegist") //注册

	//编辑
	beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit") //文档编辑显示
	beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content") //文档保存
	beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload") //图片上环
	beego.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create") //文档发送
	beego.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete") //删除文档

	//用户图书管理
	beego.Router("/book", &controllers.BookController{}, "*:Index")                         //我的图书
	beego.Router("/book/create", &controllers.BookController{}, "post:Create")              //创建图书
	beego.Router("/book/:key/setting", &controllers.BookController{}, "*:Setting")          //图书设置
	beego.Router("/book/setting/upload", &controllers.BookController{}, "post:UploadCover") //图书封面
	beego.Router("/book/star/:id", &controllers.BookController{}, "*:Collection")           //收藏图书
	beego.Router("/book/setting/save", &controllers.BookController{}, "post:SaveBook")      //保存
	beego.Router("/book/:key/release", &controllers.BookController{}, "post:Release")       //发布
	beego.Router("/book/setting/token", &controllers.BookController{}, "post:CreateToken")  //创建Token

	//个人中心
	beego.Router("/user/:username", &controllers.UserController{}, "get:Index")                 //分享
	beego.Router("/user/:username/collection", &controllers.UserController{}, "get:Collection") //收藏
	beego.Router("/user/:username/follow", &controllers.UserController{}, "get:Follow")         //关注
	beego.Router("/user/:username/fans", &controllers.UserController{}, "get:Fans")             //粉丝
	beego.Router("/follow/:uid", &controllers.BaseController{}, "get:SetFollow")                //关注或取消关注
	beego.Router("/book/score/:id", &controllers.BookController{}, "*:Score")                   //评分
	beego.Router("/book/comment/:id", &controllers.BookController{}, "post:Comment")            //评论

	//个人设置
	beego.Router("/setting", &controllers.SettingController{}, "*:Index")
	beego.Router("/setting/upload", &controllers.SettingController{}, "*:Upload")

	//管理后台
	beego.Router("/manager/category", &controllers.ManagerController{}, "post,get:Category")
	beego.Router("/manager/update-cate", &controllers.ManagerController{}, "get:UpdateCate")
	beego.Router("/manager/del-cate", &controllers.ManagerController{}, "get:DelCate")
	beego.Router("/manager/icon-cate", &controllers.ManagerController{}, "post:UpdateCateIcon")
}
