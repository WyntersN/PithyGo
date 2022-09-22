/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-13 18:46:28
 */
package admin

import (
	"PithyGo/common"
	"PithyGo/models/user"
	"PithyGo/templates/admin"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

const Version = "0.0.1"
const WebVersion = "0.0.1"

type ReturnJson struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AdminController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *AdminController) MvcError(code int, msg string) mvc.View {
	//println(c.Ctx.GetHeader("user-agent"))
	return mvc.View{
		Code: code,
		Name: "errors/mobile.html",
	}
}

func ExecuteTemplate(ctx iris.Context, tmpl admin.Partial) {
	//ctx.GzipResponseWriter().WriteString("Hello World!")
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	// if isMobile(ctx.GetHeader("User-Agent")) {
	// 	tmplError := &errors.Mobile{Vars: map[string]interface{}{"title": "出错啦"}}
	// 	templates.WriteTemplate(ctx.ResponseWriter(), tmplError)
	// 	return
	// }
	admin.WriteTemplate(ctx.ResponseWriter(), tmpl)
}
func isMobile(userAgent string) bool {
	if strings.Contains(userAgent, "Mobile") {
		return true
	}
	return false
}

/**
 * @summary: 路由 验证授权
 * @description:
 * @param {iris.Context} c
 * @router:
 */

const secretKey = "YCOAOS-p#M*JW1bl2!ON#IHunorILFjVX*u8CZa$IPAnvw6@&unHc&nkMs8BYM9yXuslC5Vlxv6rbA74Cb3AaG7k6pFjyM0*7GmFmjSY1Z"

// 生成Member token
func GenerateToken(user_id uint, username string) (tokenString string, err error) {

	claims := jwt.MapClaims{
		"user_id":  user_id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 720).Unix(), //2小时有效期，过期需要重新登录获取token
	}
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	tokenString, err = token.SignedString([]byte("YCOAOS-p#M*JW1bl2!ON#IHunorILFjVX*u8CZa$IPAnvw6@&unHc&nkMs8BYM9yXuslC5Vlxv6rbA74Cb3AaG7k6pFjyM0*7GmFmjSY1Z"))
	return
}

func AuthApp(token string) (uint, string, error) {

	claim, err := common.ParseToken(secretKey, "Bearer "+token)
	if err != nil {
		return 0, "", err
	} else {
		if data := user.GetLoginLogByAuthorization(token); time.Now().Unix() < data.ExpireAt.Unix() {
			if claim["user_id"] != nil && claim["username"] != nil {
				return uint(claim["user_id"].(float64)), claim["username"].(string), nil
			}
		}
	}
	return 0, "", errors.New("authorization expired")
}

func SessionLoginSignAuthApp(c iris.Context) {

	// c.Header("Access-Control-Allow-Origin", "*")
	// if c.Request().Method == "OPTIONS" {
	// 	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Api, Accept, Authorization, Version, Token")
	// 	c.StatusCode(204)
	// 	return
	// }
	// if appAuthorization := SessManager.Start(c).Get("appAuthorization"); appAuthorization != nil {
	// 	println(appAuthorization)
	// }
	var retData ReturnJson
	{
		retData.Code = -1
		retData.Message = "FAIL"
	}
	appAuthorization := c.GetHeader("Authorization")
	if appAuthorization == "" {
		appAuthorization = "Bearer " + c.GetCookie("NR_Authorization")
	}

	if appAuthorization != "" {

		claim, err := common.ParseToken(secretKey, appAuthorization)
		if err != nil {
			retData.Message = err.Error()
			c.RemoveCookie("NR_Authorization")
			c.StatusCode(401)
			c.JSON(retData)
		} else {

			appAuthorizationStr := strings.Replace(appAuthorization, "Bearer ", "", -1)
			if data := user.GetLoginLogByAuthorization(appAuthorizationStr); time.Now().Unix() < data.ExpireAt.Unix() {
				common.SessStart = common.SessManager.Start(c)
				if claim["user_id"] != nil && claim["username"] != nil {
					common.SessStart.Set("USER_ID", uint(claim["user_id"].(float64)))
					common.SessStart.Set("USERNAME", claim["username"].(string))
					c.Next()
				}

			} else {
				c.StatusCode(401)
				retData.Message = "authorization expired"
				c.JSON(retData)
			}
		}
	} else {
		c.StatusCode(403)
		c.JSON(retData)
	}

}
