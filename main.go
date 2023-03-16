package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/zwx19981207/Fastshorturl/credis"
	"github.com/zwx19981207/Fastshorturl/uid"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
	"time"
)

var Credis = credis.RedisInit()
var Indexfile, _ = ioutil.ReadFile("templates/index.html")

func Shorturl(ctx *fasthttp.RequestCtx) {
	shorturl := ctx.UserValue("shorturl").(string)
	Dshorturl, _ := url.QueryUnescape(shorturl)
	ret, err := Credis.HGet("url", Dshorturl).Result()
	if err != nil {
		fmt.Println(err)
		ctx.SetBodyString("error")
	} else {
		ctx.Redirect(ret, fasthttp.StatusMovedPermanently)
	}
}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.Response.SetBody(Indexfile)
}

func Postset(ctx *fasthttp.RequestCtx) {
	key := uid.DecimalToAny(int(time.Now().UnixNano()), 76)
	message := string(ctx.FormValue("url"))
	if message == "" || strings.Index(message, "http") == -1 {
		ctx.SetBodyString("error，请添加http://或https://")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	_, err := Credis.HSet("url", key, message).Result()
	if err != nil {
		fmt.Println(err)
	}
	myurl, _ := credis.Cfg.GetValue("domain", "url")
	if strings.Index(myurl, "http") == -1 {
		ctx.SetBodyString("请输入正确的地址")
	} else {
		ctx.SetBodyString(myurl + key)
	}
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/:shorturl", Shorturl)
	router.POST("/", Postset)
	if Credis != nil {
		log.Fatal(fasthttp.ListenAndServe(":88", router.Handler))
	}
}
