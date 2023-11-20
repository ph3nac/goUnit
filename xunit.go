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

// ================== TestCase ==================
type TestCase struct {
	name string
}

func (t *TestCase) setUp() {}

func (t *TestCase) tearDown() {}

func (t *TestCase) run(instance interface {
	setUp()
	tearDown()
}) {
	instance.setUp()
	var method = reflect.ValueOf(instance).MethodByName(t.name)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	} else {
		panic("Method not found")
	}
	instance.tearDown()
}

// ================== WasRun ==================

type WasRun struct {
	TestCase
	log string
}

func (w *WasRun) TestMethod() {
	w.log += "TestMethod "
}

func (w *WasRun) setUp() {
	w.log = "setUp "
}

func (w *WasRun) tearDown() {
	w.log += "tearDown "
}

// ================== TestCaseTest ==================

type TestCaseTest struct {
	TestCase
}

func (t *TestCaseTest) TestTemplateMethod() {
	test := &WasRun{TestCase: TestCase{name: "TestMethod"}}
	test.run(test)
	assert("setUp TestMethod tearDown " == test.log)
}

// ================== main ==================

func main() {
	var testTemplateMethod = &TestCaseTest{TestCase: TestCase{name: "TestTemplateMethod"}}
	testTemplateMethod.run(testTemplateMethod)
}
