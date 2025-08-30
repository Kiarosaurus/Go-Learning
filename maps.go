package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Uso de la función string.Fields(string) + maps.
func WordCount(s string) map[string]int {
	str := strings.Fields(s)	
	m := make(map[string]int)
		
	for _,value := range str {
		m[value]++	// Si existe, aumenta un valor. Si no existe, se crea con 1.
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
