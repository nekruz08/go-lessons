package main

import "fmt"

type Животное interface {
    ИздатьЗвук() string
}

type Собака struct{}
func (s Собака) ИздатьЗвук() string {
    return "Гав!"
}

type Кошка struct{}
func (k Кошка) ИздатьЗвук() string {
    return "Мяу!"
}

func ПоказатьЗвук(j Животное) {
    fmt.Println(j.ИздатьЗвук())
}

func main() {
    var животное Животное

    животное = Собака{}
    ПоказатьЗвук(животное) // Вывод: Гав!

    животное = Кошка{}
    ПоказатьЗвук(животное) // Вывод: Мяу!
}
