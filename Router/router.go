package Router

import (
	controller "Exam/Controller"
	middleware "Exam/Middleware"
	utils "Exam/Utils"
	"github.com/gin-gonic/gin"
)

// InitRouter 加入路由访问路径
func InitRouter(e *gin.Engine) {
	e.Use(utils.CORS(utils.Options{Origin: "*"}))
	example := e.Group("/example")
	{
		example.GET("/search", controller.Search)
		example.POST("/add", controller.Add)
		example.PUT("/update", controller.Update)
		example.DELETE("/delete", controller.Delete)
	}
	user := e.Group("/user")
	{
		user.GET("/findAll", controller.FindAllUser)
		user.GET("/findByN", controller.FindAllUserByName)
		user.POST("/findUserByToken",middleware.JWTAuthMiddleware(),controller.GetUserByToken)
		user.POST("/login", controller.Login)
		user.POST("/insert", controller.UserInsert)
		user.PUT("/update", controller.UserUpdate)
		user.DELETE("/delete", controller.UserDelete)
	}
	student := e.Group("/student")
	{
		student.GET("/findAllSubject", controller.GetExamSubject)   //获取考试科目
		student.GET("/findAllExamSource", controller.GetExamSource) //获取考试科目内容
		student.POST("/SaveAnswer", controller.SaveAnswer)          //保存答案
		student.PUT("/InExam", controller.InExam)                   //进入考试
		student.PUT("/OutExam", controller.OutExam)                 //离开考试
	}
	teacher := e.Group("/teacher")
	{
		teacher.POST("/insertSubject", controller.AddSubject)       //添加科目
		teacher.POST("/insertExamSource", controller.AddExamSource) //添加题库
		teacher.POST("/startExam", controller.StartExam)            //发布考试
		teacher.GET("/findStatus", controller.FindStatus)           //获取考试状态
		teacher.GET("/getAllGrade", controller.GetAllGrade)         //成绩汇总
	}
	// 中间件拦截示例
	e.GET("/cookie", utils.SetCK)
	e.GET("/home", middleware.AuthMiddleWare(), controller.Home)
	e.GET("/home2", middleware.JWTAuthMiddleware(), middleware.TokenCheck(), controller.Home2)
}

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}
