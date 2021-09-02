//爬取外卖订单列表
package pachong

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"time"
)

func CollyTakeoutOrders2() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	//是否停止爬虫
	stop := false

	err := c.Post("https://xxx/sessions", map[string]string{"phone": "xxx", "password": "xxx"})
	if err != nil {
		log.Fatal(err)
	}

	c.Limit(&colly.LimitRule{
		DomainGlob: "*.aranya.*",
		Delay:      1 * time.Second, //间隔一秒请求一次
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, e error) {
		log.Println("Something went wrong:", e)
	})

	c.OnHTML(".table.table-hover > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, row *colly.HTMLElement) {
			//获取每列的数据值
			t1 := row.DOM.Find("td").Eq(0).Text() //订单号
			t2 := row.DOM.Find("td").Eq(2).Text() //价格
			t3 := row.DOM.Find("td").Eq(4).Text() //餐厅
			t4 := row.DOM.Find("td").Eq(5).Text() //创建时间
			createTime, _ := time.Parse("2006-01-02 15:04:05", t4)
			//保存数据到数据库
			var artcles []Article
			oa := Article{Title: strings.TrimSpace(t1), Subtitle: t2, Content: t3, CreatedAt: createTime}
			DB.Find(&artcles, "title = ?", oa.Title)
			if len(artcles) == 0 {
				result := DB.Create(&oa)
				if result == nil {
					log.Println("创建失败")
				}
			} else {
				//如果数据库已存在抓取过的数据，则停止继续爬虫
				stop = true
			}

			fmt.Println(t1)

		})
	})

	//根据下一页的标签获取每一页的内容
	c.OnHTML(".pagination > span.next", func(e *colly.HTMLElement) {
		href, found := e.DOM.Find("a").Attr("href")
		// 如果有下一页，则继续访问
		if found && !stop {
			e.Request.Visit(e.Request.AbsoluteURL(href))
		}
	})

	c.Visit("https://xxx/crms/takeouts/orders")
}
