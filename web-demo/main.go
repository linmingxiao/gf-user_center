package main

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/route"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

// 应用启动前设置全局参数
func init() {
	// log.SetPrefix("[GoFast]")    // 前置字符串加上特定标记
	// log.SetFlags(log.Lmsgprefix) // 取消前置字符串
	// log.SetFlags(log.LstdFlags)  // 设置成日期+时间 格式
}

func main() {
	// TODO: 1. 初始化配置，连接数据库，创建Server
	cf.InitEnvConfig()
	app := fst.CreateServer(&cf.AppCnf.WebServerCnf)

	// TODO：2. 加载中间件、路由
	route.LoadRoutes(app)

	// TODO: 3. 启动Server Listen, 等待请求
	logx.Infof("Listening and serving HTTP on %s", app.Addr)
	if lisErr := app.Listen(); lisErr != nil {
		logx.Error(lisErr)
	}
}
