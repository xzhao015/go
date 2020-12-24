package ch1

import (
	"testing"
)

func TestRange(t *testing.T) {
	//Range is use in for
	//slice
	a := []int{1, 2, 3, 4, 5}
	for index, value := range a {
		t.Log(index, value)
	}
	//map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		t.Log("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		t.Log("Capital of", country, "is", capital)
	}

}
