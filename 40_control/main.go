package main

func main() {
	// Условный оператор
	a := true
	if a {
		println("hello world")
	}

	// Нет неявного преобразования
	b := 1
	if b == 1 {
		println("неявное преобразование для if не работает")
	}

	// Сложный условный оператор
	mp := map[string]string{
		"firstName": "Dmitry",
		"lastName":  "Rakovets",
	}
	if _, ok := mp["firstName"]; ok {
		println("firstName key is exist")
	} else {
		println("firstName key isn't exist")
	}

	// Условный оператор с else if
	if firstName, ok := mp["firstName"]; !ok {
		println("firstName key isn't exist")
	} else if firstName == "Dmitry" {
		println("firstName is Dmitry")
	} else {
		println("firstName isn't Dmitry")
	}

	// Конструкция switch
	mp["flag"] = "Ok"

	switch mp["firstName"] {
	case "Dmitry", "Tom":
		println("switch - name is Dmitry")
		// по-умолчанию не перейдет в другой вариант
	case "Petr":
		if mp["flag"] == "Ok" {
			break // выходим из switch
		}
		println("switch - name is Petr")
		fallthrough // перейти в следующий вариант
	default:
		println("switch - some other name")
	}

	switch {
	case mp["firstName"] == "Dmitry":
		println("switch2 - name is Dmitry")
	case mp["firstName"] == "Petr":
		println("switch2 - name is Petr")
	}

	// Циклы
	for {
		println("бесконечный цикл - аналог while(true)")
		break // использую что бы сразу выйти
	}

	// while-style loop
	sl := []int{3, 4, 5, 6, 7, 8}
	value := 0
	i := 0

	for i < 4 {
		if i < 2 {
			i++
			continue
		}
		value = sl[i]
		i++
		println("index:", i, "value:", value)
	}

	// c-style loop
	for i := 0; i < len(sl); i++ {
		println("index:", i, "value:", sl[i])
	}

	// range slice by index
	for i := range sl {
		println("index:", i)
	}

	// range slice by index and value
	for i, val := range sl {
		println("index:", i, "value:", val)
	}

	// range slice by value
	for _, val := range sl {
		println("value:", val)
	}

	// range map by key
	for key := range mp {
		println("key:", key)
	}

	// range map by key and value
	for key, value := range mp {
		println("key:", key, "value", value)
	}

	// range map by value
	for _, value := range mp {
		println("value", value)
	}

	// Label
MyLoop:
for key, val := range mp {
	println("switch in loop", key, val)
	switch {
	case key == "firstName" && val == "Dmitry":
		println("switch - break loop here")
		break MyLoop
	}
}
}
