package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		// auth.PUT("profile/:id", v1.UpdateProfile)
		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)
		//日志系统
		auth.POST("log", v1.AddLogAPI)
		auth.GET("getlogs", v1.GetAllLogsAPI)
		auth.DELETE("deletelogs", v1.DeleteAllLogsAPI)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		router.GET("articlefront", v1.GetArt)
		router.GET("articlefrontbyauthor", v1.GetArtByAuthor)
		router.PUT("articlefront/:id", v1.EditArt)
		router.DELETE("articlefront/:id", v1.DeleteArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)

		//文章点赞数量记录
		router.PUT("articlefront/like/:id", v1.AddLikeCountAPI)
		router.PUT("articlefront/cancellike/:id", v1.CancelLikeCountAPI)
		//点赞详情信息
		router.POST("like", v1.AddLikeRecordAPI)
		router.DELETE("deletelike/:userid/:aid", v1.DeleteLikeRecordAPI)
		router.GET("getlike/:userid/:aid", v1.GetLikeRecordsAPI)
		//获取点赞文章列表
		router.GET("getlikelist/:userid", v1.GetLikedArticlesAPI)

		//文章收藏数量记录
		router.PUT("articlefront/star/:id", v1.AddStarCountAPI)
		router.PUT("articlefront/cancelstar/:id", v1.CancelStarCountAPI)
		//收藏详情信息
		router.POST("star", v1.AddStarRecordAPI)
		router.DELETE("deletestar/:userid/:aid", v1.DeleteStarRecordAPI)
		router.GET("getstar/:userid/:aid", v1.GetStarRecordsAPI)
		//获取收藏文章列表
		router.GET("getstarlist/:userid", v1.GetStarArticlesAPI)

		// 文章列表功能迁移 测试--------------------------------------------------------

		// router.POST("category/add", v1.AddCategory)
		// router.PUT("category/:id", v1.EditCate)

		// router.GET("article", v1.GetArt)
		// router.PUT("article/:id", v1.EditArt)
		// router.DELETE("article/:id", v1.DeleteArt)

		//-----------------------------------------------------------------------------

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 个人信息设置模块
		router.GET("profile/:id", v1.GetProfile)
		router.PUT("profile/:id", v1.UpdateProfile)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)

		//上传文件接口
		router.POST("uploadfront", v1.UpLoad)
	}

	_ = r.Run(utils.HttpPort)

}
