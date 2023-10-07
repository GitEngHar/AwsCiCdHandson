package main

import "fmt"

type HumanStatus struct{
  Name string
  BirthdayMonth int
  BirthdayDay int
}

func personalColor(personalData HumanStatus) (personalColorData string){
  nameLength := len(personalData.Name)
  redElement := int((nameLength*255)/10)
  greenElement := int((personalData.BirthdayMonth*255)/12)
  blueElement := int((personalData.BirthdayDay*255)/31)
  personalColorData =fmt.Sprintf("%d,%d,%d",redElement,greenElement,blueElement)
  return personalColorData
}

func main() {
  personalData := HumanStatus{"satoshi",12,10}
  myLiteral := fmt.Sprintf("YourColorCode is %s",personalColor(personalData))
  fmt.Printf(myLiteral)
}