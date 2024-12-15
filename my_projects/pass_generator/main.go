package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	fmt.Println("\n!!!ВАС ПРИВЕТСТВУЕТ ГЕНЕРАТОР ПАРОЛЕЙ!!!")

	for {

		fmt.Print("\n0. ВЫЙТИ\n1. СГЕНЕРИРОВАТЬ ПАРОЛЬ\n\nВЫБЕРИТЕ ПУНКТ: ")
		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Ошибка ввода: ", err, ". Введите число 0 или 1!")
			continue
		}
		if a == 0 {
			break
		} else if a != 1 {
			fmt.Println("Ошибка ввода. Введите число 0 или 1!")
			continue
		}

		var passLen int
		for {
			fmt.Println("Укажите длинну пароля от 6 до 40: ")
			_, err = fmt.Scan(&passLen)
			if err != nil {
				fmt.Println("Ошибка ввода: ", err, ". Введите число от 6 до 40!")
				continue
			} else {
				break
			}
		}

		password := ""

		// password allow symbols in ASCII
		// 65-90 A-Z   97-122 a-z   48-57 0-9   96 `  94 ^  64 @  33 !  35-44 #-'
		// 26 26 10 1 1 1 10 1  = 76

		for i := 0; i < passLen; i++ {
			var passElem int
			for {
				passElem = rand.Intn(122)
				if passElem == 33 || (passElem >= 35 && passElem <= 44) || passElem == 64 || passElem == 94 || passElem == 96 || (passElem >= 48 && passElem <= 57) ||
					(passElem >= 65 && passElem <= 90) || (passElem >= 97 && passElem <= 122) {
					break
				} else {
					continue
				}
			}
			password += string(passElem)
		}

		fmt.Println(password)

		file, err := os.OpenFile("password.txt", os.O_APPEND, 02)
		if err != nil {
			file, err = os.Create("password.txt")
		}
		file.WriteString("Сгенерированный пароль: " + password + "\n")
		file.Close()

	}

	fmt.Println("Завершение работы ПО\n")

}
