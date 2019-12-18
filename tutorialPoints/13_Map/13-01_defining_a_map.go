package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	// create a map
	countryCapitalMap = make(map[string]string)

	// insert key-value paisrs in the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["China"] = "Beijing"

	// pritn map using keys
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	// test if entry is present in the map or not
	capital, ok := countryCapitalMap["United States"]
	fmt.Println("\ntest if entry is present in the map or not")

	// if ok is true, entry is present otherwise entry is absent
	if ok {
		fmt.Println("Capital of United State is ", capital)
	} else {
		fmt.Println("Capital of United State is not present")
	}
}
