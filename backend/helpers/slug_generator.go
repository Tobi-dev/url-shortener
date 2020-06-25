package helpers

import(
	"fmt"
)

func GenerateRandomSlug(url string){
	runes := []rune(url)

	for i := len(runes) - 1; i > 1; i-- {
		fmt.Printf("%c", runes[i])
	}
}
