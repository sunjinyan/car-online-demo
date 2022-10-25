package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	parse, err := url.Parse("http://nginx-service")
	if err != nil {
		panic(err)
	}
	fmt.Println(parse)
	p := httputil.NewSingleHostReverseProxy(parse)

	p.Transport = &http.Transport{
		DisableKeepAlives:      false,
	}
	http.Handle("/lb-keepalive",p)

	//ParseHttp("http://localhost:8080/go?a=123&b=456")

	p1 := httputil.NewSingleHostReverseProxy(parse)

	p1.Transport = &http.Transport{
		DisableKeepAlives:      true,
	}
	http.Handle("/lb-nokeepalive",p1)
	//grpc是采用的HTTP2协议，默认情况下使用多路复用io，保持长链接，也就是采用了keepalive保持链接方式，所以k8s不能进行keepalive的负载均衡
	//所以需要istio来进行负载分发
	err = http.ListenAndServe(":80",nil)
}


func ParseHttp(clientUrl string) {

	u, _ := url.Parse(clientUrl)            //将string解析成*URL格式
	fmt.Println(u)                          // go?a=123&b=456
	fmt.Println(u.RawQuery)                 //编码后的查询字符串，没有'?' a=123&b=456
	values, _ := url.ParseQuery(u.RawQuery) //返回Values类型的字典
	fmt.Println(values)                     // map[a:[123] b:[456]]
	fmt.Printf(" %v \n",values["a"])                //[123]
	fmt.Println(values.Get("b"))            //456
}