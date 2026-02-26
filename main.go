package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type APIResponse struct {
	BaseCode string             `json:"base_code"`
	Rates    map[string]float64 `json:"conversion_rates"`
}

func main() {
	exchangeRates := getRates()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите название валюты которую хотите перевести:")
	sliceRates := make([]string, 0, len(exchangeRates))
	for i := range exchangeRates {
		sliceRates = append(sliceRates, i)
	}
	sort.Strings(sliceRates)
	for i := range sliceRates {
		fmt.Print(sliceRates[i], " ")
	}
	fmt.Println()
	fmt.Println()
	scanner.Scan()
	currencyFrom := factoring(scanner.Text())
	_, ok := exchangeRates[currencyFrom]
	if !ok {
		log.Fatal("Такой валюты нет в списке")
	}
	fmt.Println("Введите название валюты в какую хотите перевести:")
	scanner.Scan()
	currencyTo := factoring(scanner.Text())
	_, ok = exchangeRates[currencyTo]
	if !ok {
		log.Fatal("Такой валюты нет в списке")
	}
	fmt.Printf("Введите количество денег из %s в %s\n", currencyFrom, currencyTo)
	scanner.Scan()
	valueStr := scanner.Text()
	valueStr = strings.ReplaceAll(valueStr, ",", ".")
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatal("Неправильно введено количество")
	}
	result := exchangeRates[currencyTo] / exchangeRates[currencyFrom] * value
	fmt.Printf("%.2f %s в %s будет %.2f\n", value, currencyFrom, currencyTo, result)
}

func factoring(input string) string {
	return (strings.ToUpper(strings.TrimSpace(input)))
}

func getRates() map[string]float64 {
	apiKey := os.Getenv("EXCHANGE_API_KEY")
	if apiKey == "" {
		log.Fatal("API ключ не найден в переменных окружения")
	}
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/RUB", apiKey)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Api не отвечает")
	}
	defer response.Body.Close()

	var data APIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal("Ошибка декодирования JSON")
	}
	return data.Rates
}
