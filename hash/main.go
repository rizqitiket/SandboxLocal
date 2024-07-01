package main

import (
	"crypto/md5"
	"fmt"

	"github.com/samber/lo"
)

func main() {
	data := []map[string]string{
		{
			"en": "1 Double or Twin",
		},
		{
			"en": "1 Double and 1 King",
		},
		{
			"en": "1 Single bed and 2 Double bed",
		},
	}
	var hash []string
	for _, d := range data {
		hash = append(hash, ComputeHash(d))
	}

	fmt.Println(hash)
	testString := []string{}
	curr := map[string]any{"a": nil, "b": nil}
	_, r := lo.Find[string](testString, func(item string) bool {
		if _, ok := curr[item]; ok {
			return false
		}
		return true
	})

	fmt.Println(r)

}

func ComputeHash(b map[string]string) string {
	var hash string

	for key, value := range b {
		hash = fmt.Sprintf("%s%s%s", hash, key, value)
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(hash)))
}
