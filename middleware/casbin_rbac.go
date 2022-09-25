package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"kuikui/models"
	"log"
	"net/http"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//获取请求用户,从请求头中获取用户信息
		//sub := ctx.Request.Header.Get("USERNAME")
		//获取请求用户可以从通过验证的token中获取用户信息
		claims, ok := ctx.Get("claims")
		var sub string
		if ok {
			sub = claims.(*CustomClaims).Username
		} else {
			ctx.JSON(http.StatusUnauthorized, "用户信息丢失,请重新登录")
			return
		}

		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method

		conf := models.MysqlConf{}
		_, dn, dsn := conf.GetConfStr()

		a, _ := gormadapter.NewAdapter(dn, dsn, true) // Your driver and data source.
		e, _ := casbin.NewEnforcer("./configs/model.conf", a)
		//e.AddPolicy(sub, obj, act)
		//_ = e.LoadPolicy()
		fmt.Println(obj, act, sub)
		// 判断策略中是否存在
		success, err := e.Enforce(sub, obj, act)
		if err != nil {
			// 处理错误
			fmt.Printf("%s", err)
		}
		if success {
			log.Println("恭喜您,权限验证通过")
			ctx.Next()
		} else {
			log.Printf("e.Enforce err: %s", "很遗憾,权限验证没有通过")
			ctx.JSON(http.StatusUnauthorized, "很遗憾,权限验证没有通过")
			ctx.Abort()
			return
		}

	}
}

func CasbinTest() {
	//e, err := casbin.NewEnforcer("./configs/model.conf", "./configs/policy.csv")
	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(192.168.2.157:3306)/kuikui", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./configs/model.conf", a)
	sub := "alice"
	obj := "data1"
	act := "read"
	added, err := e.AddPolicy("alice", "data1", "read")
	fmt.Println(added)
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// 处理错误
		fmt.Printf("%s", err)
	}

	if ok == true {
		fmt.Println("通过")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("未通过")
	}
}
