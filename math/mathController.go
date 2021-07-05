package math

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	div   = "div"
	minus = "minus"
	multi = "multi"
	sum   = "sum"
)

//division

func Div(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if div, ok := params[div]; ok {
		fmt.Fprintf(w, "div %s", div[0])

		text := strings.Split(div[0], "-")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0]-arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}

// Subtraction

func Minus(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	fmt.Println(params)
	if minus, ok := params[minus]; ok {
		fmt.Fprintf(w, "minus %s", minus[0])

		text := strings.Split(minus[0], "-")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0]-arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}

// multiplication

func Multi(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	// get "name" query
	if multi, ok := params[multi]; ok {
		fmt.Fprintf(w, "multi %s", multi[0])

		text := strings.Split(multi[0], "*")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0]*arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}

// summation

func Sum(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	// get "name" query
	if sum, ok := params[sum]; ok {
		fmt.Fprintf(w, "sum %s", sum[0])

		text := strings.Split(sum[0], " ")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0]+arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}
