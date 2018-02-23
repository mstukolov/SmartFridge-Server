package ibmiotf
import (
	"github.com/robertkrimen/otto"
	"fmt"
	"math"
)

func Transform(formula string, value float64) float64{
	fmt.Printf("New value: %f\n", value)
	vm := otto.New()
	vm.Run(formula)
	vm.Set("x", math.Abs(value)/1000)
	vm.Run(`
				result = fullness(x)
				console.log('node fullness(x):' + result)
			`)
	result, _ := vm.Get("result")
	fullness, _ := result.ToFloat()
	return -1*fullness/100
}
