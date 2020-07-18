package iteration

const repeatCount = 5

// Repeat a character and repeats it
func Repeat(character string, repeatCount int) string {
	result := ""
	for i := 0; i < repeatCount; i++ {
		result += character
	}
	return result
}
