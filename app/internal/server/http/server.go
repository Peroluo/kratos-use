package http

import (
	"net/http"

	pb "app/api"
	"app/internal/model"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var svc pb.DemoServer

// New new a bm server.
func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/app")
	{
		g.GET("/start", howToStart)
		// 路径参数有两个特殊符号":"和"*"
		// ":" 跟在"/"后面为参数的key，匹配两个/中间的值 或 一个/到结尾(其中不再包含/)的值
		// "*" 跟在"/"后面为参数的key，匹配从 /*开始到结尾的所有值，所有*必须写在最后且无法多个
		// NOTE：这是不被允许的，会和 /start 冲突
		// g.GET("/:xxx")
		// NOTE: 可以拿到一个key为name的参数。注意只能匹配到/param1/felix
		g.GET("/param1/:name", pathParam)
		// NOTE: 可以拿到多个key参数。注意只能匹配到/param2/felix/hao/love
		g.GET("/param2/:name/:value/:felid", pathParam)
		// NOTE: 如/params3/felix/hello/hi/，action的值为"/hello/hi/"
		g.GET("/param3/:name/*action", pathParam)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func pathParam(c *bm.Context) {
	name, _ := c.Params.Get("name")
	value, _ := c.Params.Get("value")
	felid, _ := c.Params.Get("felid")
	action, _ := c.Params.Get("action")
	path := c.RoutePath // NOTE: 获取注册的路由原始地址，如: /kratos-demo/param1/:name
	c.JSONMap(map[string]interface{}{
		"name":   name,
		"value":  value,
		"felid":  felid,
		"action": action,
		"path":   path,
	}, nil)
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
