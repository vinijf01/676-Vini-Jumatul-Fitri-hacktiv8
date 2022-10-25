package services

import (
	"ass3/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func UpdateJSON(disaster helpers.Data) {
	/*
		update (write) the JSON file
	*/
	const FilePath = "./jsons/data.json"

	disaster.Status.Water = helpers.RandomNumber(1, 100)
	disaster.Status.Wind = helpers.RandomNumber(1, 100)

	disasterByte, err := json.Marshal(disaster)

	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(FilePath, disasterByte, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", disaster)
}
