package main

import (
	"frank/function"

	"github.com/robertkrimen/otto"
)

// VM
var VM *otto.Otto

// InitVM
func InitVM() {
	VM = otto.New()
	RegisterFunctions()
}

// Register builtin functions
func RegisterFunctions() {
	function.Fake(VM)
	function.Rand(VM)
	function.MD5(VM)
	function.Must(VM)
	function.Exit(VM)
	function.Base64Decode(VM)
	function.Base64Encode(VM)
}
