package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	exchangeRates := map[string]float64{
		"RUB": 1,
		"USD": 80.5,
		"CNY": 12.2,
		"UAH": 2.5,
		"KZT": 5.0,
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите название валюты для перевода в рубли:")
	fmt.Println("USD	CNY	UAH	KZT")
	scanner.Scan()
	currency := factoring(scanner.Text())
	_, ok := exchangeRates[currency]
	if !ok {
		log.Fatal("Такой валюты нет в списке")
	}
	fmt.Println("Введите количество денег")
	scanner.Scan()
	valueStr := scanner.Text()
	valueStr = strings.ReplaceAll(valueStr, ",", ".")
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatal("Неправильно введено количество")
	}
	result := exchangeRates[currency] * value
	fmt.Printf("%.1f %s в %s будет %.1f\n", value, currency, "RUB", result)
}

func factoring(input string) string {
	return (strings.ToUpper(strings.TrimSpace(input)))
}
