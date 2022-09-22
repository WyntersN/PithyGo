/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-22 14:57:59
 */
package controllers

import (
	"PithyGo/application/app"
)

type AppController struct {
	app.AppController
}

func (c *AppController) Get() {

	// tmpl := &app.Index{
	// 	Vars: map[string]interface{}{"title": ""},
	// }
	// app.ExecuteTemplate(c.Ctx, tmpl)
	//model.CreatUser(&c.Ctx)
	c.Ctx.HTML("ok")
}

func (c *AppController) GetAi() {
	//	c.GetImgs()
	c.Ctx.HTML(`<html lang="en"><head><meta charset="utf-8"><title>This Person Does Not Exist</title><link rel="icon" type="image/png" href="/favicon.png"><meta name="viewport" content="width=device-width,initial-scale=1"><meta name="description" content="This Person Does Not Exist"><style> body { -webkit-transition: 3s -webkit-filter linear; -moz-transition: 3s -moz-filter linear; -moz-transition: 3s filter linear; -ms-transition: 3s -ms-filter linear; -o-transition: 3s -o-filter linear; transition: 3s filter linear,3s -webkit-filter linear; filter: grayscale(0)}body,html { display: flex; justify-content: center; align-items: center; background-color: #000; margin: 0; padding: 0; width: 100%; height: 100%; position: relative; font-family: Arial; font-size: .9em; overflow-x: hidden; transition: background-color .5s ease}body.show #face { max-width: 97%; max-height: 97%; margin: 1.5%; transition: all .5s ease}body.show #face:focus,body.show #face:hover { max-width: 100%; max-height: 100%; margin: 0; box-shadow: none}#face { max-width: 100%; max-height: 100%; text-align: center; box-shadow: 0 0 5px rgba(10,10,10,.8)}</style></head><body><img id="face" style="width:300px;height: 300px;" src=""></body><script type="text/javascript"> var time = new Date().getTime(); document.getElementById('face').src='https://owr.yanchu.co/imgs?t=_' + time;</script></html>`)
}

// func (c *AppController) GetImgs() {
// 	str := Import.GetImg()
// 	c.Ctx.Write([]byte(str))
// }

// func (c *AppController) GetReptiley() {
// 	Import.ReptileYy()

// }

// func (c *AppController) GetBind() {
// 	//Import.Bind()
// }

// func (c *AppController) GetLocationIn() {
// 	Import.LocationIn()
// }
