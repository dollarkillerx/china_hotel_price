package models

import (
	"fmt"
	"sort"
)

type JSModel struct {
	Total     int `json:"total"`
	Parameter struct {
		CityName string `json:"cityname"`
		CityID   string `json:"cityid"`
	} `json:"parameter"`

	Hotels []struct {
		Eid              string `json:"eid"`
		Pic153           string `json:"pic153"`
		Hotelname        string `json:"hotelname"`
		BusinessZoneName string `json:"BusinessZoneName"` // 位置
		DistrictName     string `json:"DistrictName"`     // 区域
		Address          string `json:"address"`          // 具体地址
		RenovationDate   string `json:"renovation_date"`  // 装修时间
		ShowPrice        int    `json:"show_price"`       // 价格
		ID               int    `json:"id"`
		Url              string `json:"url"`
	} `json:"hotels"`

	AveragePrice float64 `json:"average_price"` // 均价
	Median       int     `json:"median"`        // 中位数
}

func (j *JSModel) AvgPrice() *JSModel {
	var total int
	for i, v := range j.Hotels {
		total += v.ShowPrice
		j.Hotels[i].Url = fmt.Sprintf("http://www.zhuna.cn/hotel-%d.html", v.ID)
	}

	j.AveragePrice = float64(total) / float64(len(j.Hotels))

	return j
}

func (j *JSModel) MedianPrice() *JSModel {
	var p []int
	for _, v := range j.Hotels {
		p = append(p, v.ShowPrice)
	}

	sort.Ints(p)
	j.Median = p[len(p)/2]

	return j
}
