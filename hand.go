package main

type byHand struct {
	Companies
}

func (o *byHand) getCompanyNames() []string {
	var names = make([]string, len(o.Companies))
	for i, v := range o.Companies {
		names[i] = v.Name
	}
	return names
}

func (o *byHand) findAvgEmployeesUSA() float64 {
	var count int
	var sum float64
	for _, v := range o.Companies {
		if v.Country == "United States" {
			count++
			sum += v.Employees
		}
	}
	return sum / float64(count)
}
