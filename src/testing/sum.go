package sum


//find sum of numbers
func Sum(args []int) (sum int) {
	for _,val := range args {
		sum += val
	}
	return
}

