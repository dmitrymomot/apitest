package function

import (
	"fmt"

	"github.com/jmcvetta/randutil"
	"github.com/robertkrimen/otto"
)

// Randomizer registers randomizer function
// Name: Randomizer.
// Arguments: string.
// Return: string.
func Randomizer(vm *otto.Otto) {
	vm.Set("randomizer", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		if !a0.IsString() {
			fmt.Println("ERROR", "randomizer(string)")
			return otto.Value{}
		}
		s, err := a0.ToString()
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		r, err := randutil.AlphaString(n)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		rs = fmt.Sprintf("%s+%s", r, s)
		v, err := vm.ToValue(rs)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		return v
	})
}
