package main


func Sum(array []int) (sum int) {
	for _, i := range array {
		sum += i
	}
	return 
}

func SumAll(array ...[]int) (sum []int) {
	sum = make([]int, len(array))
	for i, arr := range array {
		sum[i] = Sum(arr)
	}
	return 
}

func main() {

}