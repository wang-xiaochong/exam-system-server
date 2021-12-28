package main

import (
	database "Exam/Database"
	redis "Exam/Redis"
	utils "Exam/Utils"
	router "Exam/router"
	"fmt"
)


func init(){
	redis.GetRedisPool()
	utils.InitValidator() //验证
	database.GetMongoDB()
}
func main() {
	r := router.SetupRouter()
	//r.Use(middleware.MiddleWare())    //全局中间件
	router.InitRouter(r)
	if err := r.Run(":7001"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
