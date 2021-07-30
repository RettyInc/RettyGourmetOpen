package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
)

type Restaurant struct {
	ID                  int64       `json:"id"`
	Name                string      `json:"name"`
	Distance            int64       `json:"distance"`
	DistanceFromStation int         `json:"distance_from_station"`
	DinnerBudget        interface{} `json:"dinner_budget"`
	LunchBudget         interface{} `json:"lunch_budget"`
	PaymentMethods      []string    `json:"payment_methods"`
	RecommendersCount   int64       `json:"recommenders_count"`
	Status              int         `json:"status"`
}

func main() {
	b, err := ioutil.ReadFile("./restaurants.json")
	if err != nil {
		panic(err)
	}

	var rs []Restaurant
	if err := json.Unmarshal(b, &rs); err != nil {
		panic(err)
	}
	idToScore := map[int64]int64{}
	for _, r := range rs {
		idToScore[r.ID] = 0
	}

	// 「店舗の代表予算」の安い順に、1番目に3pt, 2番目に2pt, 3番目に1pt与える
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].calcTypicalPrice() < rs[j].calcTypicalPrice()
	})
	idToScore[rs[0].ID] += 3
	idToScore[rs[1].ID] += 2
	idToScore[rs[2].ID] += 1

	// 「現在地からの距離」の近い順に、1番目に3pt, 2番目に2pt, 3番目に1pt与える
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].Distance < rs[j].Distance
	})
	idToScore[rs[0].ID] += 3
	idToScore[rs[1].ID] += 2
	idToScore[rs[2].ID] += 1

	// 「QRコード決済対応」と判定される店舗に2pt与える
	for _, r := range rs {
		for _, pm := range r.PaymentMethods {
			if pm == "QR" {
				idToScore[r.ID] += 2
			}
		}
	}

	// 「店舗のおすすめ人数」の多い順に、1番目に3pt, 2番目に2pt, 3番目に1pt与える
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].RecommendersCount > rs[j].RecommendersCount
	})
	idToScore[rs[0].ID] += 3
	idToScore[rs[1].ID] += 2
	idToScore[rs[2].ID] += 1

	// 「店舗ステータス」が「閉店」の店舗を0ptとする
	for _, r := range rs {
		if r.Status == 0 {
			idToScore[r.ID] = 0
		}
	}

	// 算出したスコアの降順にソート
	sort.Slice(rs, func(i, j int) bool {
		return idToScore[rs[i].ID] > idToScore[rs[j].ID]
	})

	/*
		以下のforループでこのような出力になる

		焼肉一筋:15614324:9
		The Pizza:19831008:4
		おにぎり屋さん:18136859:2
		ぱんや:13455313:1
		伝説のうな丼:16244132:0
		焼き鳥屋:12109451:0
	*/
	for _, r := range rs {
		fmt.Println(r.Name + ":" + strconv.FormatInt(r.ID, 10) + ":" + strconv.FormatInt(idToScore[r.ID], 10))
	}

	// 先頭の次のID -> 19831008
	fmt.Println(strconv.FormatInt(rs[1].ID, 10))
}

// 代表価格の算出
func (p Restaurant) calcTypicalPrice() float64 {
	if p.DinnerBudget == nil && p.LunchBudget == nil {
		return 0
	}
	if p.DinnerBudget == nil && p.LunchBudget != nil {
		return p.LunchBudget.(float64)
	}
	if p.DinnerBudget != nil && p.LunchBudget == nil {
		return p.DinnerBudget.(float64)
	}
	if p.DinnerBudget != nil && p.LunchBudget != nil {
		return math.Round(float64(p.LunchBudget.(float64))*0.4 + float64(p.DinnerBudget.(float64))*0.6)
	}

	return 0
}
