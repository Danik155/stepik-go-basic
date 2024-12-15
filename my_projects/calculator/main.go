package main

import (
	"fmt"
)

func main() {

	fmt.Print("\n\n\n\nВАС ПРИВЕТСТВУЕТ ПО 'КАЛЬКУЛЯТОР'\nВВЕДИТЕ 1, ЕСЛИ ХОТИТЕ ПРИСТУПИТЬ К ВЫПОЛНЕНИЮ ОПЕРАЦИИ\nВВЕДИТЕ 0, ЕСЛИ ХОТИТЕ ЗАВЕРШИТЬ РАБОТУ С ПО\nВАШ ВВОД: ")
	var a int
	breakFlag := false

	for {
		_, err := fmt.Scan(&a)

		if err != nil || (a != 0 && a != 1) {
			fmt.Println("\nНеобходимо ввести цифру 0 или 1! Повторите попытку ввода.\nВАШ ВВОД: ")
			continue
		}

		if a == 1 {
			fmt.Println("\n\n\nПО 'КАЛЬКУЛЯТОР' ЗАПУЩЕНО")
		} else {
			break
		}

		for {
			var firstDigit float64
			for {
				fmt.Println("\n\n\nВВЕДИТЕ ПЕРВОЕ ЧИСЛО ДЛЯ ОПЕРАЦИИ: ")
				_, err = fmt.Scan(&firstDigit)

				if err != nil {
					fmt.Println("ОШИБКА: " + err.Error() + "\nВЫПОЛНИТЕ ВВОД ПЕРВОГО ЧИСЛА ЕЩЁ РАЗ\n\n\n")
					continue
				} else {
					break
				}
			}

			var secondDigit float64
			for {
				fmt.Println("\nВВЕДИТЕ ВТОРОЕ ЧИСЛО ДЛЯ ОПЕРАЦИИ: ")
				_, err = fmt.Scan(&secondDigit)

				if err != nil {
					fmt.Println("ОШИБКА: " + err.Error() + "\nВЫПОЛНИТЕ ВВОД ВТОРОГО ЧИСЛА ЕЩЁ РАЗ\n\n\n")
					continue
				} else {
					break
				}
			}

			var operationSymbol string
			for {
				fmt.Println("ДОСТУПНЫЕ ОПЕРАЦИИ: сумма(+), разность(-), умножение(*), деление(/)\nВВЕДИТЕ СИМВОЛ НЕОБХОДИМОЙ ОПЕРАЦИИ (+, -, *, /, ^, %): ")
				_, err = fmt.Scan(&operationSymbol)

				if err != nil {
					fmt.Println("ОШИБКА: " + err.Error() + "\nВЫПОЛНИТЕ ВВОД СИМВОЛА ОПЕРАЦИИ ЕЩЁ РАЗ\n\n\n")
					continue
				} else if operationSymbol != "+" && operationSymbol != "-" && operationSymbol != "*" && operationSymbol != "/" {
					fmt.Println("ОШИБКА: " + "ВВЕДЁН НЕКОРРЕКТНЫЙ СИМВОЛ - " + operationSymbol + "\nВЫПОЛНИТЕ ВВОД СИМВОЛА ОПЕРАЦИИ ЕЩЁ РАЗ\n\n\n")
					continue
				} else {
					break
				}
			}

			if operationSymbol == "/" {
				if secondDigit == 0 {
					fmt.Println("ОШИБКА: ДЕЛЕНИЕ НА 0" + "\nВЫПОЛНИТЕ ВВОД ДАННЫХ ДЛЯ ОПЕРАЦИИ ПОВТОРНО")
					continue
				}
			}

			fmt.Print("РЕЗУЛЬТАТ ВАШЕЙ ОПЕРАЦИИ:\n", firstDigit, " "+operationSymbol+" ", secondDigit, " = ")
			switch operationSymbol {
			case "+":
				fmt.Print(firstDigit + secondDigit)
			case "-":
				fmt.Print(firstDigit - secondDigit)
			case "*":
				fmt.Print(firstDigit * secondDigit)
			case "/":
				fmt.Print(firstDigit / secondDigit)
			}

			fmt.Print("\n\n\nЕСЛИ ХОТИТЕ ПРОДОЛЖИТЬ ВЫПОЛНЕНИЕ ОПЕРАЦИЙ, ВВЕДИТЕ 1. ДЛЯ ЗАВЕРШЕНИЯ РАБОТЫ ПО ВВЕДИТЕ 0.\nВАШ ВВОД: ")
			for {
				_, err = fmt.Scan(&a)

				if err != nil || (a != 0 && a != 1) {
					fmt.Println("\nНеобходимо ввести цифру 0 или 1! Повторите попытку ввода.\nВАШ ВВОД: ")
					continue
				}

				if a == 1 {
					fmt.Println("ЗАПУЩЕНО ВЫПОЛНЕНИЕ НОВОЙ ОПЕРАЦИИ")
				} else {
					breakFlag = true
				}
				break
			}

			if breakFlag {
				break
			}

		}

		if breakFlag {
			break
		}

	}

	fmt.Println("Завершение работы ПО\n")

}
