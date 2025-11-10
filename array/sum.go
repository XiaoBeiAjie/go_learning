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

func SumAllTails(array ...[]int) (sum []int) {
	for _, arr := range array {
		if len(arr) == 0 {
			sum = append(sum, 0)
		} else { 
			sum = append(sum, Sum(arr[1:]))
		}
	}
	return 
}

func main() {

}