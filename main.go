package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	listen := flag.String("listen", "0.0.0.0:6610", "本地监听地址 IP:PORT")
	upstream := flag.String("upstream", "xx.xxx.xx.xx:6610", "上游 IPTV 源地址 IP:PORT")
	flag.Parse()

	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = *upstream // 保留原始 path/query
		},
		Transport:     roundTripper(http.DefaultClient.Do),
		FlushInterval: -1, // 直播流：每次写立即 flush，不缓冲
	}

	log.Printf("listening on %s -> %s", *listen, *upstream)
	log.Fatal(http.ListenAndServe(*listen, proxy))
}

type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.RequestURI = "" // Client.Do 不允许 RequestURI 非空，必须清掉
	return f(r)
}