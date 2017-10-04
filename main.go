package main

import "fmt"

type String string

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func dec_to_62(symbol string, n int) string {
	var url string
	for int(n/62) > 0 {
		url += string(symbol[n%62])
		n = int(n / 62)
	}
	url += string(symbol[n])
	return reverse(url)
}

func main() {

	fmt.Println(dec_to_62("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 3843))
}
