package iteration

func Repeat(char string) string {
	var repeated string

	if char == "" {
		char = "a"
	}

	for i := 0; i < 5; i++ {
		repeated += char
	}

	return repeated
}
