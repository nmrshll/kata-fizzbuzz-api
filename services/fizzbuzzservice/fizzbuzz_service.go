package fizzbuzzservice

import "fmt"

type POSTFizzBuzzParameters struct {
	limit int
	int1  int
	int2  int
	str1  string
	str2  string
}

type FizzBuzzService struct{}

func (fbservice *FizzBuzzService) FizzBuzz(limit, int1, int2 int, str1, str2 string) string {
	results := ""
	for i := 1; i <= limit; i++ {
		result := ""
		if i%int1 == 0 {
			result += str1
		}
		if i%int2 == 0 {
			result += str2
		}
		if result != "" {
			results += result + "\n"
			continue
		}
		results += fmt.Sprintf("%d\n", i)
	}
	return results
}
