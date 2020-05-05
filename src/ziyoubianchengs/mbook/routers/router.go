package routers

import (
	"github.com/astaxie/beego"
	"ziyoubianchengs/mbook/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //首页&分类&详情
    beego.Router("/", &controllers.HomeController{},"get:Index")
    //beego.Router("/explore", &controllers.ExploreController{},"get:Index")
    //beego.Router("/books/:key", &controllers.DocumentController{},"get:Index")
    ////读书
	//beego.Router("/read/:key/:id", &controllers.DocumentController{},"*:Read")
	//beego.Router("/read/:key/search", &controllers.DocumentController{},"post:Search")
    ////图书编辑
	//beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")
	//beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content")
	//beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")
	//beego.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create")
	//beego.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete")
	//
	////搜索
	//beego.Router("/search", &controllers.SearchController{},"get:Search")
	//beego.Router("/search/result", &controllers.SearchController{},"get:Result")
    ////登陆
	//beego.Router("/login", &controllers.AccountController{},"*:Login")
	//beego.Router("/register", &controllers.AccountController{},"*:Register")
	//beego.Router("/logout", &controllers.AccountController{},"*:Logout")
	//beego.Router("/doregister", &controllers.AccountController{},"post:DoRegister")
    ////用户图书馆里
	//beego.Router("/book", &controllers.BookController{},"*:Index") //我的图书展示
	//beego.Router("/book/create", &controllers.BookController{},"post:Create") //创建图书
	//beego.Router("/book/:key/setting", &controllers.BookController{},"*:Setting") //图书设置
	//beego.Router("/book/setting/upload", &controllers.BookController{},"post:UploadCover") //图书封面
	//beego.Router("/book/star/:id", &controllers.BookController{},"*:Collection") //收藏图书
	//beego.Router("/book/setting/save", &controllers.BookController{},"post:SaveBook") //报错
	//beego.Router("/book/:key/release", &controllers.BookController{},"post:Release") //发布图书
	//beego.Router("/book/setting/token", &controllers.BookController{},"post:CreateToken") //创建token
	//
	////个人中心
	//beego.Router("/user/:username", &controllers.UserController{},"get:Index") //分享
	//beego.Router("/user/:username/collection", &controllers.UserController{},"get:Collection") //收藏
	//beego.Router("/user/:username/follow", &controllers.UserController{},"get:Follow") //关注
	//beego.Router("/user/:username/fans", &controllers.UserController{},"get:Fans") //粉丝
	//beego.Router("/follow/:uid", &controllers.BaseController{},"get:SetFollow") //关注或取消关注
	//beego.Router("/book/score/:id", &controllers.BookController{},"*:Score") //评论
	//beego.Router("/book/comment/:id", &controllers.BookController{},"post:Comment") //评分
	//
	////个人设置
	//beego.Router("/setting", &controllers.SettingController{},"*:Index") //个人设置页面
	//beego.Router("/setting/upload", &controllers.SettingController{},"*:Upload") //上传图片
	//
	////管理员后台
	//beego.Router("/manager/category", &controllers.ManagerController{},"post,get:Category") //分类展示
	//beego.Router("/manager/update-cate", &controllers.ManagerController{},"get:UploadCate") //上传分类信息
	//beego.Router("/manager/del-cate", &controllers.ManagerController{},"get:DelCate") //删除分类
	//beego.Router("/manager/icon-cate", &controllers.ManagerController{},"post:UploadCateIcon") //上传分类icon
	//
	//








}
