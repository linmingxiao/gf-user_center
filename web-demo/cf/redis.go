package cf

import (
	"github.com/qinchende/gofast/connx/gfrds"
	"github.com/qinchende/gofast/sdx"
)

var R61501 *gfrds.GfRedis
var R61508 *gfrds.GfRedis

//
func initGoRedis() {
	if &AppCnf.RedisDBListCnf.R61501 != nil {
		R61501 = gfrds.NewGoRedis(&AppCnf.RedisDBListCnf.R61501)
	}
	if &AppCnf.RedisDBListCnf.R61508 != nil {
		R61508 = gfrds.NewGoRedis(&AppCnf.RedisDBListCnf.R61508)
	}
}

//
//func tryGoRedis() {
//	pong, err := RedisA.Ping()
//
//	if err != nil {
//		fmt.Println("Ping failed", err)
//	} else {
//		fmt.Printf("Ping val is %s", pong)
//	}
//}

// init sdx session with redis store
func initRedisForSession() {
	sdx.SetupSession(&sdx.SessionDB{
		RedisSessCnf: AppCnf.RedisSessCnf,
	})
}
