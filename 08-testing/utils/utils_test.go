package utils

import (
	"fmt"
	"testing"
)

/*
func Test_IsPrime(t *testing.T) {
	// arrange
	no := 17
	expectedResult := true

	//act
	actualResult := IsPrime(no)

	//assert
	if actualResult != expectedResult {
		t.Errorf("IsPrime(17) : expected [%v] but got [%v]\n", expectedResult, actualResult)
	}
}
*/

func Test_IsPrime(t *testing.T) {
	testData := []struct {
		no             int
		expectedResult bool
	}{
		{no: 13, expectedResult: true},
		{no: 14, expectedResult: true},
		{no: 15, expectedResult: false},
		{no: 17, expectedResult: true},
	}
	for _, td := range testData {
		testName := fmt.Sprintf("Test_IsPrime(%d)\n", td.no)
		t.Run(testName, func(t *testing.T) {

			//act
			actualResult := IsPrime(td.no)

			//assert
			if actualResult != td.expectedResult {
				t.Errorf("IsPrime(%v) : expected [%v] but got [%v]\n", td.no, td.expectedResult, actualResult)
			}
		})
	}
}
