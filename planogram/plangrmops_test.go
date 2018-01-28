package planogram

import (
	"testing"
)

func Test_BuildReport(t *testing.T){
	checkfunc := BuildCheckPlanogramReport()
	if checkfunc != true {
		t.Errorf("BuildCheckPlanogramReport: %v", checkfunc)
	}
}
