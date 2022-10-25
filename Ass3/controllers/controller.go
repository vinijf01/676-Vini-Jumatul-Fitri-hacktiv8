package controllers

import (
	"ass3/helpers"
	"html/template"
	"net/http"
	"path"
)

func DataController(w http.ResponseWriter, r *http.Request) {

	var FilePath = path.Join("views", "web.html")

	var tmpl, err = template.ParseFiles(FilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	getdata := helpers.GetData()
	numWater := getdata.Status.Water
	numWind := getdata.Status.Wind

	waterSatuan := helpers.WaterSatuan(numWater)
	windSatuan := helpers.WindSatuan(numWind)
	waterStatus := helpers.WaterStatus(numWater)
	windStatus := helpers.WindStatus(numWind)

	var data = map[string]interface{}{
		"waterSatuan": waterSatuan,
		"windSatuan":  windSatuan,
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
