package foobar

//move this file inside ~GOPATH/src/foobar/ (create foobar directory) and then run go install from the directory

//find sum of numbers
func Sum(args ...int) (sum int) {
	for _, val := range args {
		sum += val
	}
	return
}
