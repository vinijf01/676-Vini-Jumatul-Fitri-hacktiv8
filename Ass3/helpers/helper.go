package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func RandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

func WaterSatuan(number int) string {
	desc := fmt.Sprint(number, " meter")
	return desc

}

func WindSatuan(number int) string {
	desc := fmt.Sprint(number, " meter per second")
	return desc

}

func WindStatus(number int) string {
	if number < 6 {
		return "Aman"
	} else if number >= 7 && number < 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func WaterStatus(number int) string {
	if number < 5 {
		return "Aman"
	} else if number >= 6 && number < 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func GetData() Data {
	fileBytes, err := ioutil.ReadFile("./jsons/data.json")

	if err != nil {
		panic(err)
	}

	var data Data

	err = json.Unmarshal(fileBytes, &data)

	if err != nil {
		panic(err)
	}

	return data
}
