package main

import (
	"awesomeProject/zwxurl/credis"
	"awesomeProject/zwxurl/uid"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"time"
)

var Credis, _ = credis.RedisInit("127.0.0.1:6379")
var Indexfile, _ = ioutil.ReadFile("templates/index.html")

func Shorturl(ctx *fasthttp.RequestCtx) {
	shorturl := ctx.UserValue("shorturl").(string)
	ret, err := Credis.HGet("h1", shorturl).Result()
	if err != nil {
		fmt.Println(err)
	}
	ctx.Redirect(ret, fasthttp.StatusMovedPermanently)

}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.Response.SetBody(Indexfile)
}

func Postset(ctx *fasthttp.RequestCtx) {
	key := uid.DecimalToAny(int(time.Now().UnixNano()), 76)
	message := string(ctx.FormValue("url"))
	_, err := Credis.HSet("h1", key, message).Result()
	if err != nil {
		fmt.Println(err)
	}
	ctx.SetBodyString("http://zwxurl.top/"+key)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/:shorturl", Shorturl)
	router.POST("/", Postset)
	fasthttp.ListenAndServe(":80", router.Handler)
	//log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
