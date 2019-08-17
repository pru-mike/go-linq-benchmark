package main

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

type Company struct {
	Country      string  `csv:"country"`
	Name         string  `csv:"company"`
	Headquarters string  `csv:"headquarters"`
	Founded      int     `csv:"founded"`
	Employees    float64 `csv:"employees"`
	Revenue      float64 `csv:"revenue"`
	MarketCap    float64 `csv:"marketcap"`
}

type Companies []*Company

func (c *Companies) Load() {
	companyFile, err := os.OpenFile("company.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	if err := gocsv.UnmarshalFile(companyFile, c); err != nil {
		panic(err)
	}

	err = companyFile.Close()
	if err != nil {
		panic(err)
	}
}
