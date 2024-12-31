package main

import (
	"fmt"
	"strings"
)

// Struct to represent climate data for a city
type CityData struct {
	CityName       string
	AvgTemperature float64 // Average temperature in °C
	Rainfall       float64 // Rainfall in mm
}

// Function to find the city with the highest temperature
func highestTemperature(cities []CityData) (string, float64) {
	var highestCity string
	var highestTemp float64
	for _, city := range cities {
		if city.AvgTemperature > highestTemp {
			highestTemp = city.AvgTemperature
			highestCity = city.CityName
		}
	}
	return highestCity, highestTemp
}

// Function to find the city with the lowest temperature
func lowestTemperature(cities []CityData) (string, float64) {
	var lowestCity string
	var lowestTemp float64 = cities[0].AvgTemperature
	for _, city := range cities {
		if city.AvgTemperature < lowestTemp {
			lowestTemp = city.AvgTemperature
			lowestCity = city.CityName
		}
	}
	return lowestCity, lowestTemp
}

// Function to calculate the average rainfall across all cities
func averageRainfall(cities []CityData) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter and display cities with rainfall above a given threshold
func filterCitiesByRainfall(cities []CityData, threshold float64) {
	fmt.Println("\nCities with rainfall above", threshold, "mm:")
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("%s: %.2f mm\n", city.CityName, city.Rainfall)
		}
	}
}

// Function to search for a city by name and display its data
func searchByCityName(cities []CityData, cityName string) {
	cityName = strings.ToLower(cityName)
	found := false
	for _, city := range cities {
		if strings.ToLower(city.CityName) == cityName {
			fmt.Printf("\nCity: %s\nAverage Temperature: %.2f°C\nRainfall: %.2f mm\n", city.CityName, city.AvgTemperature, city.Rainfall)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("City not found.")
	}
}

func main() {
	// Hardcoded data for multiple cities
	cities := []CityData{
		{"New York", 15.2, 120.4},
		{"London", 10.8, 114.5},
		{"Tokyo", 16.3, 120.8},
		{"Paris", 12.5, 90.1},
		{"Sydney", 18.4, 102.3},
	}

	// Display highest and lowest temperature
	highestCity, highestTemp := highestTemperature(cities)
	lowestCity, lowestTemp := lowestTemperature(cities)
	fmt.Printf("\nCity with the highest temperature: %s (%.2f°C)\n", highestCity, highestTemp)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowestCity, lowestTemp)

	// Calculate and display average rainfall
	avgRainfall := averageRainfall(cities)
	fmt.Printf("\nAverage rainfall across all cities: %.2f mm\n", avgRainfall)

	// Get rainfall threshold from user input
	var threshold float64
	fmt.Print("\nEnter the rainfall threshold (in mm) to filter cities: ")
	_, err := fmt.Scan(&threshold)
	if err != nil || threshold < 0 {
		fmt.Println("Invalid input. Please enter a positive number for rainfall threshold.")
		return
	}
	filterCitiesByRainfall(cities, threshold)

	// Search by city name
	var cityName string
	fmt.Print("\nEnter the city name to search: ")
	fmt.Scan(&cityName)
	searchByCityName(cities, cityName)
}
