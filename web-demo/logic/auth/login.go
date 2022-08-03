package auth

import (
	"gf-example/web-demo/cf"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/skill/jsonx"
)

func BeforeLogin(c *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37,"tok":"t:QnBQTHNDT3RIS2V2aFJyUk1o.rEnZy6QeaS/fDtG3Kj/eBBwKDRbfJs8/nAqIxtmzdM"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/login?account=admin\&pass=abc
func LoginByAccPass(c *fst.Context) {
	// 模拟验证登录，写入 user_id
	account := c.GetString("account")
	password := c.GetString("password")

	var cus_parent_id string
	if len(account) == 11 {
		//手机号登录模式
		cus_parent_id, _ = cf.R61501.Cli.HGet(c, `cus_parent_mobile`, account).Result()
	} else if len(account) == 18 {
		//身份证号登录模式
		cus_parent_id, _ = cf.R61501.Cli.HGet(c, `cus_parent_card`, account).Result()
	} else {
		c.Fai(0, `账户名错误`, nil)
		return
	}

	userInfo, _ := cf.R61501.Cli.HGetAll(c, `cus_parent_`+cus_parent_id).Result()
	pwd := userInfo[`hashed_password`]
	if pwd != password {
		c.Fai(0, `账户名密码错误`, nil)
		return
	}

	wechat, _ := jsonx.GetMapFromString(userInfo[`wechat`])
	fund, _ := jsonx.GetMapFromString(userInfo[`fund`])
	base, _ := jsonx.GetMapFromString(userInfo[`base`])

	println(base)

	sdx.DestroySession(c)
	sdx.NewSession(c)

	pms := fst.KV{
		sdx.MySessDB.GuidField: cus_parent_id,
		`nickName`:             base[`nickname`],
		`avatarUrl`:            base[`avatar_url`],
		`wechatBind`:           wechat[`is_bind`],
		`open_gm`:              fund[`identity_card`] != nil,
	}

	c.Sess.SetKV(pms)
	c.Sess.Save()
	c.SucKV(pms)
	return
}
