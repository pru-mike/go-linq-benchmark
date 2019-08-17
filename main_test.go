package main

import (
	"testing"
)

func TestFindAvgEmployeesUSA(t *testing.T) {
	H := byH.findAvgEmployeesUSA()
	LI := byLi.findTAvgEmployUSA()
	if H != LI {
		t.Errorf("different average: %f, %f", H, LI)
	}
}

func TestGetCompanyNames(t *testing.T) {
	H := byH.getCompanyNames()
	LI := byLi.getCompanyNames()
	if len(H) == 0 || len(H) != len(LI) {
		t.Errorf("different lenght of company list: %d <> %d", len(H), len(LI))
	}
	for i := range H {
		if H[i] != LI[i] {
			t.Errorf("different company list: [%v] [%v]", H, LI)
			break
		}
	}
}
