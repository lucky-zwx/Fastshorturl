package main

import (
	"awesomeProject/zwxurl/credis"
	"awesomeProject/zwxurl/uid"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/go-redis/redis"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"time"
)

var Credis *redis.Client = credis.RedisInit()
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
	myurl, _ := credis.Cfg.GetValue("domain", "url")
	ctx.SetBodyString(myurl + key)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/:shorturl", Shorturl)
	router.POST("/", Postset)
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
