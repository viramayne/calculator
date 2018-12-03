package main

import (
	"fmt"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func zesarEncript(key int, str string) (result string) {
	if key > 25 {
		return ""
	}
	arr := []byte(strings.ToLower(str))
	var newIndex int
	for i := 0; i < len(arr); i++ {
		n := strings.Index(alphabet, string(arr[i]))
		if n != -1 {
			newIndex = n + key
			if newIndex > 25 {
				arr[i] = alphabet[newIndex-25]
			} else {
				arr[i] = alphabet[newIndex]
			}
		}

	}
	return string(arr)
}

func zesarDecript(key int, str string) (result string) {
	if key > 25 {
		return ""
	}
	arr := []byte(strings.ToLower(str))
	var newIndex int
	for i := 0; i < len(arr); i++ {
		if n := strings.Index(alphabet, string(arr[i])); n != -1 {
			newIndex = n - key
			if newIndex < 0 {
				arr[i] = alphabet[25-(key-n)]
			} else {
				arr[i] = alphabet[newIndex]
			}
		}

	}
	return string(arr)
}

func encriptDecodeZesar() {
	strin := "today is monday"
	key := 5
	encriptedStr := zesarEncript(key, strin)
	fmt.Println(encriptedStr)
	decriptedStr := zesarDecript(key, encriptedStr)
	fmt.Println(decriptedStr)
}
