package main

import (
  "fmt"
  "math"
)

type Converter struct {}

type Centimeter float64
type Feet float64

type Minutes float64
type Seconds float64

type Celsius float64
type Fahrenheit float64

type Kilograms float64
type Pounds float64

type Radian float64
type Degree float64


func (cvr Converter) CentimeterToFeet(C Centimeter) Feet {
  inFeet := C / 30.48
  fmt.Printf("%v cm  = %v feet\nApprox ~ ", C, inFeet)
  // Returns the rounded Feet value
  return Feet(math.Round(float64(inFeet)))
}

func (cvr Converter) FeetToCentimeter(F Feet) Centimeter {
  inCentimeter := F * 30.48
  fmt.Printf("\n%v feet = %v cm\nApprox ~ ", F, inCentimeter)
  // Returns the rounded Centimeter value
  return Centimeter(math.Round(float64(inCentimeter)))
}

func (cvr Converter) MinutesToSeconds(M Minutes) Seconds {
  inSeconds := M * 60
  fmt.Printf("\n%v minute(s) = %v second(s)\nApprox ~ ", M, inSeconds)
  // Returns the rounded Seconds value
  return Seconds(math.Round(float64(inSeconds)))
}

func (cvr Converter) SecondsToMinutes(S Seconds) Minutes {
  inMinutes := S / 60
  fmt.Printf("\n%v second(s) = %v minute(s)\nApprox ~ ", S, inMinutes)
  // Returns the rounded Minutes value
  return Minutes(math.Round(float64(inMinutes)))
}

func (cvr Converter) CelsiusToFahrenheit(Cel Celsius) Fahrenheit {
  inFahrenheit := (Cel * 1.8) + 32
  fmt.Printf("\n%v degree(s) celsius = %v degree(s) Fahrenheit\nApprox ~ ", Cel, inFahrenheit)
  // Returns the rounded Fahrenheit value
  return Fahrenheit(math.Round(float64(inFahrenheit)))
}

func (cvr Converter) FahrenheitToCelsius(Fah Fahrenheit) Celsius {
  inCelsius := (Fah - 32) / 1.8
  fmt.Printf("\n%v degree(s) fahrenheit = %v degree(s) celsius\nApprox ~ ", Fah, inCelsius)
  // Returns the rounded Celsius value
  return Celsius(math.Round(float64(inCelsius)))
}

func (cvr Converter) KilogramsToPounds(K Kilograms) Pounds {
  inPounds := K * 2.2046
  fmt.Printf("\n%v kg = %v ibs\nApprox ~ ", K, inPounds)
  // Returns the rounded value in pounds
  return Pounds(math.Round(float64(inPounds)))
}

func (cvr Converter) PoundsToKilograms(P Pounds) Kilograms {
  inKilograms := P * 0.4536
  fmt.Printf("\n%v ibs = %v kg\nApprox ~ ", P, inKilograms)
  // Returns the rounded Kilograms value
  return Kilograms(math.Round(float64(inKilograms)))
}

func (cvr Converter) RadianToDegree(R Radian) Degree {
  inDegree := R * (180 / math.Pi)
  fmt.Printf("\n%v radian = %v degree(s)\nApprox ~ ", R, inDegree)
  // Returns the rounded Degree value
  return Degree(math.Round(float64(inDegree)))
}

func (cvr Converter) DegreeToRadian(D Degree) Radian {
  inRadian := D * (math.Pi / 180)
  fmt.Printf("\n%v degree(s) = %v radian\nApprox ~ ", D, inRadian)
  // Returns the rounded Radian value
  return Radian(math.Round(float64(inRadian))) 
}


func main() {
  cvr := Converter{}
  
  fmt.Println("Welcome to our Converter")
  // Line 102 to 111 make use of the printf string formatting
  // To print out the results of each method, which are of different types
  fmt.Printf("%v", cvr.CentimeterToFeet(1))
  fmt.Printf("%v", cvr.FeetToCentimeter(1))
  fmt.Printf("%v", cvr.MinutesToSeconds(1))
  fmt.Printf("%v", cvr.SecondsToMinutes(1))
  fmt.Printf("%v", cvr.CelsiusToFahrenheit(1))
  fmt.Printf("%v", cvr.FahrenheitToCelsius(1))
  fmt.Printf("%v", cvr.KilogramsToPounds(1))
  fmt.Printf("%v", cvr.PoundsToKilograms(1))
  fmt.Printf("%v", cvr.RadianToDegree(1))
  fmt.Printf("%v", cvr.DegreeToRadian(1))

}
