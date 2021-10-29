package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestArMap(t *testing.T) {
	fmt.Println(spList(AClassArea, "甲类地区"))
	fmt.Println(spList(BClassArea, "乙类地区"))
}

func TestReptile(t *testing.T) {
	pMap := queryID()

	model, err := reptile("北京", pMap["北京"], "5")
	if err != nil {
		log.Fatalln(err)
	}

	marshal, err := json.Marshal(model)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(marshal))
}

func TestPmAP(t *testing.T) {
	queryID()
}
