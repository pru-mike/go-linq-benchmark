package main

import (
	"gopkg.in/ahmetb/go-linq.v3"
)

type byLinq struct {
	Companies
}

func (o *byLinq) getCompanyNames() []string {
	var names = make([]string, len(o.Companies))
	linq.From(o.Companies).
		Select(func(c interface{}) interface{} { return c.(*Company).Name }).
		ToSlice(&names)
	return names
}

func (o *byLinq) getCompanyNamesT() []string {
	var names = make([]string, len(o.Companies))
	linq.From(o.Companies).SelectT(func(c *Company) string { return c.Name }).ToSlice(&names)
	return names
}

func (o *byLinq) findTAvgEmployUSA() float64 {
	return linq.From(o.Companies).
		WhereT(func(c *Company) bool { return c.Country == "United States" }).
		SelectT(func(c *Company) float64 { return c.Employees }).Average()
}
