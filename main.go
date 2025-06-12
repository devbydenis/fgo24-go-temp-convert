package main

import (
	"fmt"
	"os"
	"temp_convert/calculate"
	"temp_convert/models"
)

func MainMenu() {
	fmt.Println("-----------------------------------------------")	
	fmt.Println("-------Convert Temperature From Celcius-------")	
	fmt.Println("Masukan temperature celcius:")
}

func ConvertMenu() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("Silakan pilih menu berikut:")
	fmt.Println("1. Convert to Reamur")
	fmt.Println("2. Convert to Fahrenheit")
	fmt.Println("3. Convert to Kelvin")
	fmt.Println("4. History Calculation")
	fmt.Println("-----------------------------------------------")
}

func main() {
	var temperatureInput float64
	var choiceDestination int
	var histories []models.History
	
	if temperatureInput == 0 {
		MainMenu()
		_, err := fmt.Scan(&temperatureInput)
		if err != nil {
			fmt.Println("Invalid harus angka")
			return
		}
	}
	if temperatureInput != 0 {
		ConvertMenu()
		fmt.Print("your choice: ")
		_, err := fmt.Scan(&choiceDestination)
		if err != nil {
			fmt.Println("Invalid harus angka")
			return
		}
	}
	
	switch choiceDestination {
	case 1:
		fmt.Printf("%.2f Celcius = %.2f Reamur\n", temperatureInput, calculate.FromCelciusToReamur(temperatureInput))
		histories = append(histories, models.History{
			Input: temperatureInput,
			Output: calculate.FromCelciusToReamur(temperatureInput),
			Type: "Celcius to Reamur",
		})
	case 2:
		fmt.Printf("%.2f Celcius = %.2f Fahrenheit\n", temperatureInput, calculate.FromCelciusToFahrenheit(temperatureInput))
		histories = append(histories, models.History{
			Input: temperatureInput,
			Output: calculate.FromCelciusToFahrenheit(temperatureInput),
			Type: "Celcius to Fahrenheit",
		})
	case 3:
		fmt.Printf("%.2f Celcius = %.2f Kelvin\n", temperatureInput, calculate.FromCelciusToKelvin(temperatureInput))
		histories = append(histories, models.History{
			Input: temperatureInput,
			Output: calculate.FromCelciusToKelvin(temperatureInput),
			Type: "Celcius to Kelvin",
		})
	case 4:
		os.Exit(0)
	case 5:
		// if len(histories) == 0 {
		// 	fmt.Println("No history available.")
		// } else {
		// 	fmt.Println("History:")
		// 	for _, history := range histories {
		// 		fmt.Printf("%s\n", history)
		// 	}
		// }
		// return
	default:
		fmt.Println("Invalid choice. Make sure you enter a number between 1 and 5.")
	}
	
		
		
		
	
}