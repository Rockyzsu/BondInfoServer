package router

import (
	"bondinfoserver/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	// 配置加载模板路径

	// //  第一个参数是url的路径,第二个是本地的路径
	// router.Static("/upload", "upload")

	// // 将文件映射为url
	// router.StaticFile("/check", "./upload/img.jpg")

	// 基本信息
	router.POST("/info", controllers.BondBaseInfo) //上传

	// router.GET("/list", controllers.ListHistory) //显示url
	// router.GET("/l", controllers.ListHistory)    //快捷方式

	// //router.GET("/", controllers.ListHistory) //显示url
	// router.GET("/ll", controllers.ListHistorys) //快捷方式
	// // 文本
	// router.GET("/copy", controllers.CopyTextIndex)
	// router.GET("/c", controllers.CopyTextIndex)
	// router.POST("/copycontent", controllers.HandlerCopyContent)
	// router.GET("/paste", controllers.GetHistoryText)
	// router.GET("/p", controllers.GetHistoryText)

	// //删除 需要权限
	// router.GET("/d", controllers.DeleteImage) // 删除
	// router.POST("/delimage", controllers.DeleteImage)
	// router.POST("/image/remove", controllers.DeleteImageByBT)

	// // 获取所有图片
	// router.GET("/w", controllers.WalkImages)
	// router.GET("/walk", controllers.WalkImages)
	// router.POST("/image/next", controllers.WalkImages)

	// //router.GET("/cloud", controllers.TencentUpdateCDN)

	// router.GET("/cloud", controllers.ServerManager)
	// router.POST("/update-site1", controllers.BaiduSite1)
	// router.POST("/update-site2", controllers.BaiduSite2)
	// router.POST("/update-cdn", controllers.TencentUpdateCDN)

	// // 测试用
	// router.GET("/node", controllers.DynamicAddNode)

	return router
}
