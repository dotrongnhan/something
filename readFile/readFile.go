package readFile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadTextFile() (resultMax, resultMin, average int) {
	data, err := ioutil.ReadFile("text.txt")
	content := string(data)
	newArray := strings.Split(content, " ")
	if err != nil {
		fmt.Println(err)
	}
	var newArrayInt = []int{}
	for _, i := range newArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		newArrayInt = append(newArrayInt, j)
	}
	var sum int
	resultMax = newArrayInt[0]
	resultMin = newArrayInt[0]
	for i := 1; i < len(newArrayInt); i++ {
		if newArrayInt[i] > resultMax {
			resultMax = newArrayInt[i]
		}
		if newArrayInt[i] < resultMin {
			resultMin = newArrayInt[i]
		}

		sum += newArrayInt[i]
	}
	average = sum/len(newArrayInt)
	return resultMax, resultMin, average
}

func ReadFile()  {
	content, err := ioutil.ReadFile("./db.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Read file success")

	var config map[string]string

	error := json.Unmarshal([]byte(content), &config)
	if error != nil {
		panic(error)
	}
	fmt.Println(config)

}