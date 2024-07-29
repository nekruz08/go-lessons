// Задания для 6 – го урока
//
//Учет накопительных счетов с ежемесячным пополнением
// Начальный баланс накопительного счета равен 100000 рублей.
// Реализуйте функции для пополнения счета каждый месяц на фиксированную сумму.
// Выводите баланс после каждого пополнения.
// Если баланс становится больше 500000 рублей, выведите сообщение "Достигнут лимит накоплений".
package main

import (
	"fmt"
)

func main() {
	var balance float64 = 100000.0
	monthlyDeposit := 1500.0 // Фиксированная сумма пополнения каждый месяц

	fmt.Printf("Начальный баланс накопительного счета: %.2f рублей\n", balance)

	for month := 1; month <= 12; month++ {
		depositMonthly(&balance, monthlyDeposit)
	}
}

// Функция для пополнения счета на фиксированную сумму каждый месяц
func depositMonthly(balance *float64, monthlyDeposit float64) {
	*balance += monthlyDeposit
	fmt.Printf("Баланс после пополнения: %.2f рублей\n", *balance)
	if *balance > 500000 {
		fmt.Println("Достигнут лимит накоплений")
	}
}
//-------------------------------------------------------------------------------------------------
// Рассчет выплат по кредиту с фиксированной ставкой
// Начальная сумма кредита равна 3_000_000 рублей.
// Реализуйте функции для ежемесячного расчета выплат по кредиту с фиксированной процентной ставкой 10%.
// Выводите остаток по кредиту после каждой выплаты.
// Если остаток по кредиту становится меньше 500000 рублей, выведите сообщение "Почти погашен кредит".
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	principal := 3000000.0 // Начальная сумма кредита
// 	annualRate := 10.0     // Годовая процентная ставка
// 	months := 36           // Срок кредита в месяцах

// 	calculateLoanPayments(principal, annualRate, months)
// }

// func calculateLoanPayments(principal float64, annualRate float64, months int) {
// 	monthlyRate := annualRate / 12 / 100
// 	monthlyPayment := principal * (monthlyRate * math.Pow(1+monthlyRate, float64(months))) / (math.Pow(1+monthlyRate, float64(months)) - 1)

// 	balance := principal
// 	for month := 1; month <= months; month++ {
// 		interestPayment := balance * monthlyRate
// 		principalPayment := monthlyPayment - interestPayment
// 		balance -= principalPayment
// 		fmt.Printf("Месяц %d: Остаток по кредиту: %.2f рублей\n", month, balance)

// 		if balance < 500000 {
// 			fmt.Println("Почти погашен кредит")
// 			break
// 		}
// 	}
// 	fmt.Printf("Ежемесячный платеж: %.2f рублей\n", monthlyPayment)
// }
