package main

import (
	"fmt"
	"reflect"
)

type WasRun struct {
	TestCase
	WasRun bool
}

func (w *WasRun) TestMethod() {
	w.WasRun = true
}

type TestCase struct {
	name string
}

func (t *TestCase) Run(i interface{}) error {
	var method = reflect.ValueOf(i).MethodByName(t.name)
	if !method.IsValid() {
		return fmt.Errorf("method not found: %s", t.name)
	}
	method.Call([]reflect.Value{})
	return nil
}

type TestCaseTest struct {
}

func assert(b bool) {
	if !b {
		var errorMessage = ("Assertion failed")
		panic(errorMessage)
	}
}

func (tct *TestCaseTest) TestRunning() {
	var test = &WasRun{TestCase: TestCase{name: "TestMethod"}}
	assert(test.WasRun)
	var err = test.Run(test)
	if err != nil {
		println(err.Error())
	}
	assert(test.WasRun)
}

func main() {
	tcc := &TestCaseTest{}
	tcc.TestRunning()
}
