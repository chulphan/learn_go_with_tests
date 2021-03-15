package iteration

func Repeat() string {
	repeated := ""

	for i := 0; i < 5; i++ {
		repeated += "a"
	}

	return repeated
}
