# Fastshorturl
基于FastHttp的短连接服务

shell使用示例：
```shell
curl 'http://zwxurl.top/' -H 'Connection: keep-alive' -H 'Accept: */*' -H 'Origin: http://zwxurl.top' -H 'X-Requested-With: XMLHttpRequest' -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36 Edg/79.0.309.47' -H 'DNT: 1' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Referer: http://zwxurl.top/' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' -H 'Cookie: BT_PANEL_6=21c21e28-c767-4756-873d-90f094404517.MPe8sABP_odjjpGOW4lHd_Sbqpo' --data 'url=https://www.baidu.com' --compressed --insecure
```
