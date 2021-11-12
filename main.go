package main

import (
	"fmt"
	"math"
)

func printPrimeNumbers(num1, num2 int) []int {
	var array []int
	if num1 < 2 || num2 < 2 {
		fmt.Println("Numbers must be greater than 2.")
		return nil
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			array = append(array, num1)
		}
		num1++
	}

	return array
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func arrayGrouping(arr []string) [][]string {
	var ret [][]string
	var group []string
	for idx, i := range arr {
		if idx > 0 && i != group[len(group)-1] {
			ret = append(ret, group)
			group = make([]string, 0)
		}
		group = append(group, i)
	}
	if len(group) > 0 {
		ret = append(ret, group)
	}
	return ret
}

func modifCountElement(arr [][]string) []interface{} {

	var output []interface{}

	for i := 0; i < len(arr); i++ {

		var a = []interface{}{
			len(arr[i]), arr[i][0],
		}

		output = append(output, a)
	}

	return output
}

func main() {
	h := isPalindrome("abcba")

	pn := printPrimeNumbers(11, 40)

	arr := []string{
		"a", "a", "a", "b", "c", "c", "b", "b", "b", "d", "d", "e", "e", "e",
	}
	sa := arrayGrouping(arr)

	ce := modifCountElement(arrayGrouping(arr))

	fmt.Println("isPalindrome=", h)
	fmt.Println("primeNumber=", pn)
	fmt.Println("grouping=", sa)
	fmt.Println("countSomeElement=", ce)

}
