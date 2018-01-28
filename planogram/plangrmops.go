package planogram

import (
	"encoding/json"
	"fmt"
)

func checkPlanogram(){

}

type SlotsQty struct {
	Sku string `json:"sku"`
	Plan int `json:"plan"`
	Fact int `json:"fact"`
	Status bool `json:"status"`
}
type AbsentSKU struct {
	Sku string `json:"sku"`
	Plan int `json:"plan"`
	Fact int `json:"fact"`
	Status bool `json:"status"`
}
type PlanogramReport struct {
	SlotsQtyBySKU []SlotsQty `json:"SlotsQtyBySKU"`
	AbsentSKU []AbsentSKU `json:"AbsentSKU"`
	EmptySells int `json:"EmptyCells"`
}

type Planogram struct {

}

func BuildCheckPlanogramReport() []byte{

	slotsArr, absentArr := buildTestData()
	report := new(PlanogramReport)
	report.SlotsQtyBySKU = slotsArr
	report.AbsentSKU = absentArr
	report.EmptySells = 12

	jreport, _ := json.Marshal(report)
	fmt.Print(jreport)
	return jreport
}

func buildTestData() ([]SlotsQty, []AbsentSKU){
	s1 := new(SlotsQty)
	s1.Sku = "cola-67868"
	s1.Plan = 20
	s1.Fact = 15
	s1.Status = false

	s2 := new(SlotsQty)
	s2.Sku = "cola-112"
	s2.Plan = 6
	s2.Fact = 6
	s2.Status = true

	s3 := new(SlotsQty)
	s3.Sku = "cola-434"
	s3.Plan = 10
	s3.Fact = 10
	s3.Status = true

	a1 := new(AbsentSKU)
	a1.Sku = "cola-67868"
	a1.Plan = 20
	a1.Fact = 15
	a1.Status = false

	a2 := new(AbsentSKU)
	a2.Sku = "cola-112"
	a2.Plan = 6
	a2.Fact = 6
	a2.Status = true

	a3 := new(AbsentSKU)
	a3.Sku = "cola-434"
	a3.Plan = 10
	a3.Fact = 10
	a3.Status = true

	var slotsArr []SlotsQty
	slotsArr = append(slotsArr, *s1)
	slotsArr = append(slotsArr, *s2)
	slotsArr = append(slotsArr, *s3)

	var absentArr []AbsentSKU
	absentArr = append(absentArr, *a1)
	absentArr = append(absentArr, *a2)
	absentArr = append(absentArr, *a3)

	return slotsArr, absentArr
}
