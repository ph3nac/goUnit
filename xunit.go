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

func (t *TestCase) run(instance interface{ setUp() }) {
	instance.setUp()
	var method = reflect.ValueOf(instance).MethodByName(t.name)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	} else {
		panic("Method not found")
	}
}

// ================== WasRun ==================

type WasRun struct {
	TestCase
	wasRun   bool
	wasSetUp bool
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}

func (w *WasRun) setUp() {
	w.wasRun = false
	w.wasSetUp = true
}

// ================== TestCaseTest ==================

type TestCaseTest struct {
	TestCase
	test *WasRun
}

func (t *TestCaseTest) setUp() {
	t.test = &WasRun{TestCase: TestCase{name: "TestMethod"}}
}

func (t *TestCaseTest) TestRunning() {
	t.test.run(t.test)
	assert(t.test.wasRun)
}

func (t *TestCaseTest) TestSetUp() {
	t.test.run(t.test)
	assert(t.test.wasSetUp)
}

// ================== main ==================

func main() {
	var testRunning = &TestCaseTest{TestCase: TestCase{name: "TestRunning"}}
	testRunning.run(testRunning)
	var testSetUp = &TestCaseTest{TestCase: TestCase{name: "TestSetUp"}}
	testSetUp.run(testSetUp)
}
