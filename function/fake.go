package function

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jmcvetta/randutil"
	"github.com/manveru/faker"
	"github.com/robertkrimen/otto"
)

// Fake registers fake generator function
// Name: Fake.
// Arguments: string.
// Return: string.
func Fake(vm *otto.Otto) {
	vm.Set("fake", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		if !a0.IsString() {
			fmt.Println("ERROR", "fake(string)")
			return otto.Value{}
		}
		s, err := a0.ToString()
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		s = strings.ToLower(s)

		var res string

		switch s {
		case "email":
			res = getEmail()
		case "name":
			res = getName()
		case "first_name":
			res = getFirstName()
		case "last_name":
			res = getLastName()
		case "phone":
			res = getPhoneNumber()
		}

		v, err := vm.ToValue(res)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		return v
	})
}

func getEmail() string {
	rInt, _ := randutil.IntRange(0, 64)
	rStr, _ := randutil.AlphaStringRange(5, 10)
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(int64(rInt)))
	return fmt.Sprintf("%s-%s@autotest.dev", fake.UserName(), rStr)
}

func getName() string {
	r, _ := randutil.IntRange(0, 64)
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(int64(r)))
	return fake.Name()
}

func getFirstName() string {
	r, _ := randutil.IntRange(0, 64)
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(int64(r)))
	return fake.FirstName()
}

func getLastName() string {
	r, _ := randutil.IntRange(0, 64)
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(int64(r)))
	return fake.LastName()
}

func getPhoneNumber() string {
	r, _ := randutil.IntRange(0, 64)
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(int64(r)))
	return fake.PhoneNumber()
}
