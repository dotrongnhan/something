package checkExist

func checkNumber(search []int, arr []int) []int {

	var numberExist []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(search); j++ {
			if arr[i] == search[j] {
				numberExist = append(numberExist, arr[i])
			}
		}
	}

	return numberExist
}
