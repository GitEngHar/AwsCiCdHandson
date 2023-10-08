package main

import (
	"fmt"
	"os"
	"strconv"
)

type HumanStatus struct {
	Name          string
	BirthdayMonth int
	BirthdayDay   int
}

func personalColor(personalData HumanStatus) (personalColorData string) {
	nameLength := len(personalData.Name)
	redElement := int((nameLength * 255) / 10)
	greenElement := int((personalData.BirthdayMonth * 255) / 12)
	blueElement := int((personalData.BirthdayDay * 255) / 31)
	personalColorData = fmt.Sprintf("%d,%d,%d", redElement, greenElement, blueElement)
	return personalColorData
}

func main() {
	month, _ := strconv.Atoi(os.Getenv("MONTH"))
	day, _ := strconv.Atoi(os.Getenv("DAY"))
	personalData := HumanStatus{os.Getenv("MyName"), month, day}
	myLiteral := fmt.Sprintf("YourColorCode is %s", personalColor(personalData))
	fmt.Printf(myLiteral)
}
