package iteration

func Repeat(char string, repeatCount int) string {
	var repeated string

	if char == "" {
		char = "a"
	}

	for i := 0; i < repeatCount; i++ {
		repeated += char
	}

	return repeated
}
