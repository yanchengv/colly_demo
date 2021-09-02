//爬虫balawo网站并且数据保存在数据库中
package pachong

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func Colly4() {


	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.github.*", Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(".post", func(e *colly.HTMLElement) {
		var artcles []Article
		oa := Article{Title: strings.TrimSpace(e.DOM.Find("h1").Eq(0).Text())}
		DB.Find(&artcles,"title = ?",oa.Title)
		if len(artcles) == 0  {
			result := DB.Create(&oa)
			if result == nil {
				log.Println("创建失败")
			}
		}
		log.Println(
			//strings.TrimSpace(e.DOM.Find("h1").Eq(0).Text())
			strings.TrimSpace(e.DOM.Find("h1").Eq(0).Text()))
	})

	c.OnHTML(".pagination a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://yanchengv.github.io/")
	c.Wait()

}
