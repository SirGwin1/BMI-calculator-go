package functions

import (
	"net/http"
	"strconv"
)

type Result struct {
	Name   string
	BMI    float64
	Status string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl.Execute(w, nil)
}

func CalculateBMI(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	name := r.FormValue("name")

	weight, err := strconv.ParseFloat(
		r.FormValue("weight"),
		64,
	)
	if err != nil {
		http.Error(w, "Invalid Weight", http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseFloat(
		r.FormValue("height"),
		64,
	)
	if err != nil {
		http.Error(w, "Invalid Height", http.StatusBadRequest)
		return
	}

	bmi := ComputeBMI(weight, height)

	data := Result{
		Name:   name,
		BMI:    bmi,
		Status: BMIStatus(bmi),
	}

	tmpl.Execute(w, data)
}
