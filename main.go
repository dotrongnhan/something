package main

import (
	"example.com/hello/sort"
	"fmt"
)

func main() {

	fmt.Println(sort.QuickSort([]int{1, 2, 3, 4, 5}))
	//db, err := database.Connection()
	//if err != nil {
	//	//Catch error trong quá trình thực thi
	//	log.Printf("Error %s when getting db connection", err)
	//	return
	//}
	//defer db.Close()
	//log.Printf("Successfully connected to database")
	//
	//crawler.Crawler(db)
}
