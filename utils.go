package main

import (
	"encoding/json"
	"math"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func SumOfDigits(n int) int {
	numStr := strconv.Itoa(n)
	sum := 0

	for _, digit := range numStr {
		digitInt, _ := strconv.Atoi(string(digit))

		sum += digitInt
	}

	return sum
}

type FFResponseData struct {
	Found  bool   `json:"found"`
	Number int    `json:"number"`
	Text   string `json:"text"`
	Type   string `json:"type"`
}

func GetFunFact(n int) (string, error) {
	requestURL := fmt.Sprintf("http://numbersapi.com/%d/math", n)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var data FFResponseData
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return "", err
	}

	return data.Text, nil
}

func IsPerfectFn(n int) bool {
	// A perfect number is a positive integer
	// equal to the sum of its proper divisors (excluding itself).

	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func IsPrimeFn(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsArmStrongFn(n int) []string {
	order := len(strconv.Itoa(n))
	status := "odd"
	if n%2 == 0 {
		status = "even"
	}

	var temp, remainder, result int
	// initialize the variables
	temp = n

	// Use of For Loop
	for temp > 0 {
		remainder = temp % 10
		result += int(math.Pow(float64(remainder), float64(order)))
		temp /= 10
	}

	if result == n {
		return []string{"armstrong", status}
	}

	return []string{status}
}
