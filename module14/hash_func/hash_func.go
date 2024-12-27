package main

import "fmt"

func hashint64(val int64) uint64 {
	return uint64(val) % 1000
}

func hashstr(val string) uint64 {
	fmt.Println(val)

	runes := []rune(val)
	var splitted uint64
	for i, r := range runes {
		splitted += uint64((i + 1) * int(r))
	}

	return splitted % 1000
}

func main() {
	sourceNumber := int64(12345)
	hashNumber := hashint64(sourceNumber)

	fmt.Printf("Source number: %d\nHash number: %d\n", sourceNumber, hashNumber)

	sourceString := "АБВГД"
	hashString := hashstr(sourceString)

	fmt.Printf("Source string: %s\nHash string: %d\n", sourceString, hashString)

	sourceString = "ДГВБА"
	hashString = hashstr(sourceString)

	fmt.Printf("Source string: %s\nHash string: %d\n", sourceString, hashString)
}
