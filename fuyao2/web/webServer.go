package web

import (
	"fmt"
	"net/http"

	controller "github.com/fuyao/web/appController"
)

// 启动Web服务并指定路由信息
func WebStart(app controller.Application) {

	// 指定路由信息(匹配请求)
	http.HandleFunc("/query", app.FindByID)
	http.HandleFunc("/addInfo", app.AddInfo)
	http.HandleFunc("/delInfo", app.DeleteInfo)
	fmt.Println("启动Web服务, 监听端口号为: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}

}
