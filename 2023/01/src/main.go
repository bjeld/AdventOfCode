package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	calibrationValues := readCallibrationValueFromFile("data.txt")
	fmt.Println(calculatesSumOfAllCalibrationValues(calibrationValues))
}

func readCallibrationValueFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func calculatesSumOfAllCalibrationValues(calibrationValues []string) int {
	sumOfAllCalibrationValues := 0
	for _, callibrationValue := range calibrationValues {
		sumOfAllCalibrationValues += combineFirstAndLastDigitsOrZero(callibrationValue)
	}
	return sumOfAllCalibrationValues
}

func combineFirstAndLastDigitsOrZero(input string) int {
	digits := extractDigits(input)

	if len(digits) == 0 {
		return 0
	}

	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	return combineFirstAndLastDigit(digits)
}

func extractDigits(input string) []int {
	var digits []int
	for _, char := range input {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			digits = append(digits, num)
		}
	}
	return digits
}

func combineFirstAndLastDigit(digits []int) int {
	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	return firstDigit*10 + lastDigit
}
