package api

import (
	"fmt"
	"net/http"

	"github.com/ev-chargingpoint/backend-evchargingpoint/charge"
	"github.com/ev-chargingpoint/backend-evchargingpoint/charging_station"
	"github.com/ev-chargingpoint/backend-evchargingpoint/login"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"github.com/ev-chargingpoint/backend-evchargingpoint/signup"

	"github.com/rs/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			if r.Method == "GET" {
				fmt.Fprintf(w, "API BISA DIPAKE")
				return
			}

		case "/login":
			if r.Method == "POST" {
				fmt.Fprint(w, login.Post("PASETOPRIVATEKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
		case "/register":
			if r.Method == "POST" {
				fmt.Fprint(w, signup.Post("MONGOSTRING", "db_evchargingpoint", r))
				return
			}
		case "/chargingstation":
			if r.Method == http.MethodPost {
				fmt.Fprint(w, charging_station.Post("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
			if r.Method == http.MethodPut {
				fmt.Fprint(w, charging_station.Put("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
			if r.Method == http.MethodDelete {
				fmt.Fprint(w, charging_station.HapusChargingStationHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
			fmt.Fprint(w, charging_station.GetChargingStationHandler("MONGOSTRING", "db_evchargingpoint", r))
			return
		case "/profile":
			if r.Method == http.MethodPut {
				fmt.Fprint(w, profile.Put("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
			fmt.Fprint(w, profile.Get("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
			return
		case "/chargecar":
			if r.Method == http.MethodPost {
				fmt.Fprint(w, charge.ChargeHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
			if r.Method == http.MethodPut {
				fmt.Fprint(w, charge.PaymentAndStatusChargeHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
				return
			}
			fmt.Fprint(w, charge.GetChargeHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_evchargingpoint", r))
		}

	}))

	handler.ServeHTTP(w, r)
}
