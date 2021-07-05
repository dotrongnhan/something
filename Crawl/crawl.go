package crawler

import (
	"database/sql"
	"example.com/hello/model"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

func Crawler(db *sql.DB) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		//Đang gửi request get HTML
		fmt.Printf("Visiting: %sn", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		//Handle error trong quá trình craw html
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		//Sau khi đã lấy được HTML
		fmt.Printf("Visited: %sn", r.Request.URL)
	})

	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {
		//Bóc tách dữ liệu từ HTML lấy được
		data := model.Film{}
		data.Title = e.ChildText(".titleColumn a")
		yearElement := e.ChildText(".titleColumn span")
		year, _ := strconv.Atoi(yearElement[1 : len(yearElement)-1])
		//Tìm đến thẻ con h2 và lấy nội dung
		data.Year = year
		//tìm đến thẻ con cite và lấy nội dung
		data.Rating, _ = strconv.ParseFloat(e.ChildText(".ratingColumn strong"), 32)
		//Tìm đến thẻ con p và lấy nội dung
		data.Link = e.ChildAttr("a", "href")
		fmt.Printf("- Title: %s- Year: %d- Rating: %f- Link: %s", data.Title, data.Year, data.Rating, data.Link)
		//In ra màn hình giá trị đã lấy được
		stmt, err := db.Prepare("INSERT film SET title=?,year=?,Rating=?")
		//Prepare SQL cho việc insert
		checkErr(err)
		//Handle error
		res, err := stmt.Exec(data.Title, data.Year, data.Rating)
		//Binding data vào câu query
		checkErr(err)
		//Handle error

		lastId, err := res.LastInsertId()
		//Lấy ra ID vừa được insert

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("=&gt;Insert ID: %dnn", lastId)
	})

	c.OnScraped(func(r *colly.Response) {
		//Hoàn thành job craw
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
	//Trình thu thập truy cập URL đó
}

func checkErr(err error) { //Thêm function để handle error
	if err != nil {
		panic(err)
	}
}
