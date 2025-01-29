package app

import "fmt"

func CreatePlaceholders(cnt int) []string {
	pl := []string{}
	for i := 0; i < cnt; i++ {
		pl = append(pl, fmt.Sprintf("$%d", i+1))
	}
	return pl
}

func GenUpdateSets(from []string) []string {
	result := []string{}
	for i := 0; i < len(from); i++ {
		indexOf := i + 2 // because $1 is a first parameter (starts with 1, not 0), and also $1 is reserved for ID
		result = append(result, from[i]+" = $"+fmt.Sprintf("%d", indexOf))
	}
	return result
}
