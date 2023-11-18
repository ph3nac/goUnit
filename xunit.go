package main

import (
	"reflect"
)

func assert(b bool) {
	if !b {
		var errorMessage = ("Assertion failed")
		panic(errorMessage)
	}
}

type TestCase struct {
	name string
}

func (t *TestCase) run(i interface{}) {
	var method = reflect.ValueOf(i).MethodByName(t.name)
	method.Call([]reflect.Value{})
}

type WasRun struct {
	TestCase
	WasRun bool
}

func (w *WasRun) TestMethod() {
	w.WasRun = true
}

type TestCaseTest struct {
	TestCase
}

func (t *TestCaseTest) TestRunning() {
	var test = &WasRun{TestCase: TestCase{name: "TestMethod"}}
	assert(!test.WasRun)
	test.run(test)
	assert(test.WasRun)
}

func main() {
	var test = &TestCaseTest{TestCase: TestCase{name: "TestRunning"}}
	test.run(test)
}
