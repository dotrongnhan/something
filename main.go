package main

import "fmt"

func main() {

	//fmt.Println(sort.QuickSort([]int{1,4,2,3}))
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


func hello(name string) string {
	if name == "" {
		return fmt.Sprintf("What is your name ?")
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}