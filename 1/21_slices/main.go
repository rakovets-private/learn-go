package main

import "fmt"

// slice
func main() {
	// Объявление
	var sl []int
	fmt.Println("пустой слайс ", sl)

	// Добавление елементов
	sl = append(sl, 100)
	fmt.Println("уже не пустой слайс", sl)

	// Длина
	fmt.Println("длина слайса", len(sl))

	// О размерности и расширении
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))
	sl = append(sl, 102)
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))
	sl = append(sl, 103)
	sl = append(sl, 104)
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))
	sl = append(sl, 105)
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))

	// Объявление и инициализация
	sl2 := []int{
		10,
		20,
		30,
	}
	fmt.Println(sl2, len(sl2), cap(sl2))

	// Добавление элементов из одного slice в другой
	sl = append(sl, sl2...)
	fmt.Println(sl)

	// slice of slice
	var slm [][]int
	slm = append(slm, sl)
	fmt.Println(slm)

	// Объявление с нужной длиной
	slice3 := make([]int, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// Объявление с нужной длиной и размерностью
	slice4 := make([]int, 10, 10)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// Передача ссылки при присваивании
	slice5 := slice4
	slice5[1] = 100500
	fmt.Println(slice4, slice5)

	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	slice4[1] = 999
	fmt.Println(slice4, slice5)

	// Неправильное копирование, так размерность меньше чем требуется
	var slice6 []int
	copy(slice6, slice5)
	fmt.Println(slice6)

	// Правильное копирование
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)

	// Получение части
	fmt.Println("часть слайса", slice7[1:5], slice7[:2], slice7[10:])
	slice8 := append(slice7[:2], slice7[10:]...)
	fmt.Println("из кусков слайса", slice8)

	// Создание из массива
	a := [...]int{5, 6, 7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println(sl8)

	// Копирование в конкретные позиции
	sl9 := make([]int, 10)
	copy(sl9[3:6], []int{1, 2, 3})
	fmt.Println(sl9)
	sl9 = append(sl9[3:6], []int{1, 2, 3}...)

	// Вывод
	fmt.Println(sl9)
	fmt.Printf("slice elements is %v\nnewline\n", sl9)
	println(sl9)
}
