package main

import (
	"encoding/json"
	"fmt"
)

type Restaurant struct {
	ID int64 `json:"id"`
	Name    string `json:"name"`
	Distance  int `json:"distance"`
	DistanceFromStation  int `json:"distance_from_station"`
	DinnerBudget  interface{} `json:"dinner_budget"`
	LunchBudget  interface{} `json:"lunch_budget"`
	PaymentMethods []string `json:"payment_methods"`
	RecommendersCount int `json:"recommenders_count"`
	Status   int    `json:"status"`
}

func main() {
	var rs []Restaurant

	rs = append(rs, Restaurant{
		ID: 19831008,
		Name: "おにぎり屋さん",
		Distance: 1000,
		DistanceFromStation: 400,
		DinnerBudget: interface{}(nil),
		LunchBudget: 300,
		PaymentMethods: []string {"QR", "Credit", "Cash"},
		RecommendersCount: 420,
		Status: 1,
	})

	rs = append(rs, Restaurant{
		ID: 12109451,
		Name: "焼き鳥屋",
		Distance: 300,
		DistanceFromStation: 400,
		DinnerBudget: 1000,
		LunchBudget: 1000,
		PaymentMethods: []string {"Cash"},
		RecommendersCount: 312,
		Status: 1,
	})

	rs = append(rs, Restaurant{
		ID: 16244132,
		Name: "伝説のうな丼",
		Distance: 150,
		DistanceFromStation: 300,
		DinnerBudget: 6000,
		LunchBudget: interface{}(nil),
		PaymentMethods: []string {"Credit", "Cash"},
		RecommendersCount: 14694,
		Status: 0,
	})

	rs = append(rs, Restaurant{
		ID: 15614324,
		Name: "焼肉一筋",
		Distance: 30,
		DistanceFromStation: 100,
		DinnerBudget: 3000,
		LunchBudget: 2000,
		PaymentMethods: []string {"QR", "Credit", "Cash"},
		RecommendersCount: 821,
		Status: 1,
	})

	rs = append(rs, Restaurant{
		ID: 13455313,
		Name: "ぱんや",
		Distance: 250,
		DistanceFromStation: 450,
		DinnerBudget: interface{}(nil),
		LunchBudget: 300,
		PaymentMethods: []string {"Credit", "Cash"},
		RecommendersCount: 519,
		Status: 1,
	})

	rs = append(rs, Restaurant{
		ID: 18136859,
		Name: "The Pizza",
		Distance: 100,
		DistanceFromStation: 200,
		DinnerBudget: 2000,
		LunchBudget: 1000,
		PaymentMethods: []string {"QR", "Credit", "Cash"},
		RecommendersCount: 123,
		Status: 2,
	})

	j, err := json.Marshal(&rs)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(j))
}

