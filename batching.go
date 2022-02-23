package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Prices struct {
	timestamp            string  `json:"date"`
	last_traded_price    float64 `json:"last_traded_price"`
	total_buy_qty        float64 `json:"total_buy_qty"`
	total_sell_qty       float64 `json:"total_sell_qty"`
	average_traded_price float64 `json:"average_traded_price"`
	pct_change           float64 `json:"pct_change"`
}
type Depth struct {
	timestamp  string  `json:"date"`
	asset_type string  `json:"type"`
	price      float64 `json:"price"`
	num_orders float64 `json:"number_of_orders_at_pricepoint"`
	total_qty  float64 `json:"total_qty"Z`
}

func createPricesList(data [][]string) []Prices {
	var pricesList []Prices
	for i, line := range data {
		if i > 0 {
			var rec Prices
			for j, field := range line {
				if j == 0 {
					rec.timestamp = field
					fmt.Println(rec.timestamp)
				} else if j == 1 {
					rec.last_traded_price, _ = strconv.ParseFloat(field, 64)

				} else if j == 2 {
					rec.total_buy_qty, _ = strconv.ParseFloat(field, 64)
				} else if j == 3 {
					rec.total_sell_qty, _ = strconv.ParseFloat(field, 64)
				} else if j == 4 {
					rec.average_traded_price, _ = strconv.ParseFloat(field, 64)
				} else if j == 5 {
					rec.pct_change, _ = strconv.ParseFloat(field, 64)
				}

			}
			pricesList = append(pricesList, rec)
		}
	}
	return pricesList

}
func createDepthList(data [][]string) []Depth {
	var depthList []Depth
	for i, line := range data {
		if i > 0 {
			var rec Depth
			for j, field := range line {
				if j == 0 {
					rec.timestamp = field
				} else if j == 1 {
					rec.asset_type = field

				} else if j == 2 {
					rec.price, _ = strconv.ParseFloat(field, 64)
				} else if j == 3 {
					rec.num_orders, _ = strconv.ParseFloat(field, 64)
				} else if j == 4 {
					rec.total_qty, _ = strconv.ParseFloat(field, 64)
				}
			}
			depthList = append(depthList, rec)
		}
	}
	return depthList
}

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("File Name: %s\n", info.Name())
		if filepath.Ext(info.Name()) == ".csv" {
			readCSV(info.Name())
		}

		return nil
	})
}

func readCSV(name string) string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	pricesList := createPricesList(data)
	jsonData, err := json.MarshalIndent(pricesList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
	return string(jsonData)

}
func main() {
	absPath, _ := filepath.Abs("../data/prices")
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		readCSV(absPath + "/" + file.Name())
	}
	//absPath, _ = filepath.Abs("../data/depth")
	/*files, err = ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		readCSV(file.Name())
	}*/

}
