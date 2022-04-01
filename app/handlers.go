package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Jestinepa/doctor/service"
)

type Timings struct {
	From string `json:"from" xml:"from"`
	To   string `json:"to" xml:"to"`
}

type Doctor struct {
	DoctorId     int    `json:"doctorid" xml:"doctorid"`
	Fullname     string `json:"fullname" xml:"fullname"`
	Availability []Timings
}

type DoctorHandlers struct {
	service service.DoctorService
}

func (dh *DoctorHandlers) getavailability(w http.ResponseWriter, r *http.Request) {
	// doctors := []Doctor{
	// 	{
	// 		DoctorId: 1,
	// 		Fullname: "jackson",
	// 		Availability: []Timings{
	// 			{
	// 				From: "10",
	// 				To:   "12",
	// 			},
	// 			{
	// 				From: "4",
	// 				To:   "6",
	// 			},
	// 		},
	// 	},
	// }

	doctors, _ := dh.service.GetAllDoctor()

	if r.Header.Get("Conten-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(doctors)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doctors)
	}
}
