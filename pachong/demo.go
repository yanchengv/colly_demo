package pachong

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func colly7() {
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/8.0; .NET4.0C; .NET4.0E)"
	c.OnRequest(func(r *colly.Request) {
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Host", "xx.xx.40.202:8084")
			r.Headers.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
			r.Headers.Set("Connection", "keep-alive")
			r.Headers.Set("X-Requested-With", "XMLHttpRequest")
			r.Headers.Set("X-MicrosoftAjax", "Delta=true")
			r.Headers.Set("Cache-Control", "no-cache")
			r.Headers.Set("Accept", "*/*")
			//r.Headers.Set("Origin", "http://www.sse.com.cn")
			r.Headers.Set("Referer", "http://xx.xx.40.202:8084/hotelbs254/Main.aspx?ScreenHeight=692") //关键头 如果没有 则返回 错误
			r.Headers.Set("Accept-Encoding", "gzip, deflate")
			r.Headers.Set("Accept-Language", "zh-CN")
			r.Headers.Set("Cookie", "ASP.NET_SessionId=mwokiogxpjpn22syxsqdb3kt")
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		log.Println(strings.NewReader(string(r.Body)))

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.OnHTML("#gvList", func(e *colly.HTMLElement) {
		log.Println("1111")
		log.Println(e.DOM.First())
	})

	c.Visit("xx")

}
