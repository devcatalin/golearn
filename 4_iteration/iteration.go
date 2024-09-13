package iteration

func Repeat(character string, times int) (repeated string) {
	if times <= 0 {
		return character
	}
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return
}
