package cf

import (
	"flag"
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/connx/gfrds"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/skill/conf"
)

type AppConfigEntity struct {
	WebServerCnf   fst.GfConfig     `v:"required"`
	RedisSessCnf   sdx.RedisSessCnf `v:"required"`
	MysqlGoZeroCnf gform.ConnCnf    `v:"required"`

	RedisDBListCnf gfrds.RedisDBListCnf
}

var AppCnf AppConfigEntity
var cnfFile = flag.String("f", "/Users/lmx/go/pkg/mod/github.com/linmx/gf-user_center/web-demo/cf/env.yaml", "-f env.[yaml|yml|json]")

func MustAppConfig() {
	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	logx.MustSetup(&AppCnf.WebServerCnf.LogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.AppName + ", config all ready.")

	initRedisForSession()
	initGoRedis()
	initMysql()
}
