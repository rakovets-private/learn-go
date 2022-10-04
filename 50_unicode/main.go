package main

import "fmt"

func main() {
	var symbol rune = 'a'
	var autoSymbol = 'a'
	unicodeSumbol := '©'
	unicodeSymbolByNumber := '\u2318'
	println(symbol, autoSymbol, unicodeSumbol, unicodeSymbolByNumber)

	str1 := "Привет, Мир!"
	fmt.Println("ru: ", str1, len(str1))
	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	str2 := "你好世界"
	fmt.Println("cn: ", str2, len(str2))
	for index, runeValue := range str2 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	bin := []byte(str2)
	fmt.Println("binary cn: ", bin, len(bin))
	for index, value := range str2 {
		fmt.Printf("raw binary index: %v, oct: %v, hex: %x\n", index, value, value)
	}
}
