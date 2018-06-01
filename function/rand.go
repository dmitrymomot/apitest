package function

import (
	"fmt"

	"github.com/jmcvetta/randutil"
	"github.com/robertkrimen/otto"
)

// Rand registers rand function
// Name: Rand.
// Arguments: string.
// Return: string.
func Rand(vm *otto.Otto) {
	vm.Set("rand", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		r, err := randutil.AlphaString(10)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		rs := fmt.Sprintf("%s+%s", r, a0)
		v, err := vm.ToValue(rs)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		return v
	})
}
