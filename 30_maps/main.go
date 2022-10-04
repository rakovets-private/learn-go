package main

import "fmt"

func main() {
	// Инициализация
	var mm map[string]string
	fmt.Println("unitialized map", mm)

	// Полная инициализация
	// var mm2 map[string]string = map[string]string{}
	mm2 := map[string]string{}
	fmt.Println(mm2)

	// Добавление елемента
	mm2["test"] = "ok"
	fmt.Println(mm2)

	// Короткая инициализация
	mm3 := make(map[string]string)
	mm3["firstName"] = "Vasily"
	fmt.Println("mm3", mm3)

	// Получение елемента
	firstName := mm3["firstName"]
	fmt.Println("firstName", firstName, len(firstName))

	// Обращение по не существующему ключу
	lastName := mm3["lastName"]
	fmt.Println("lastName", lastName, len(lastName))

	// Проверка на наличие ключа
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is", lastName, "exist:", ok)

	// Получени ТОЛЬКО признака существования
	_, exist := mm3["firstName"]
	fmt.Println("firstName exist:", exist)

	// Удаление элемента
	key:= "firstName1"
	mm3[key] = "Dmitry"
	delete(mm3, key)
	_, exist = mm3[key]
	fmt.Println(key, "exist:", exist)

	// Присваивание по ссылке
	mm4 := mm3
	mm4[key] = "Dmitry"
	fmt.Println(mm3, mm4)
}
