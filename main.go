package main

import (
	"./credis"
	"./uid"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/go-redis/redis"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"net/url"
)

var Credis *redis.Client = credis.RedisInit()
var Indexfile, _ = ioutil.ReadFile("templates/index.html")

func Shorturl(ctx *fasthttp.RequestCtx) {
	shorturl := ctx.UserValue("shorturl").(string)
	Dshorturl, _ := url.QueryUnescape(shorturl)
	ret, err := Credis.HGet("url", Dshorturl).Result()
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
	if message == "" {
		ctx.SetBodyString("error")
		return
	}
	_, err := Credis.HSet("url", key, message).Result()
	if err != nil {
		fmt.Println(err)
	}
	myurl, _ := credis.Cfg.GetValue("domain", "url")
	if strings.Index(myurl, "http") == -1 {
		ctx.SetBodyString("请输入正确的地址")
	}
	ctx.SetBodyString(myurl + key)
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
