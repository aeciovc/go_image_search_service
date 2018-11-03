package rabbitmq

import (
	"testing"
	test "github.com/aeciovc/go-image-search/test"
	"reflect"
)

func TestInvokerWithNoParamsStringResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func() string{
		return "hello"
	}

	result, err := invoker.Invoke(function)

	test.Equals(t, err, nil)
	test.Equals(t, result, "hello")
}

func TestInvokerWithParamsSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func(param1 string) string{
		return param1
	}

	result, err := invoker.Invoke(function, "this one")

	test.Equals(t, err, nil)
	test.Equals(t, result, "this one")
}

func TestInvokerWithArrayResultSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func(param1 string, param2 string) []string{
		return []string{param1, param2}
	}

	result, err := invoker.Invoke(function, "one", "two")

	test.Equals(t, err, nil)
	test.Equals(t, result, []string{"one", "two"})
}

func TestInvokerWithNoParamsBooleanResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func() bool{
		return true
	}

	result, err := invoker.Invoke(function)

	test.Equals(t, err, nil)
	test.Equals(t, result, true)
}

func TestInvokerWithNoParamsIntegerResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func() int{
		return 20
	}

	result, err := invoker.Invoke(function)
	test.Equals(t, err, nil)
	test.Equals(t, reflect.TypeOf(result).Kind(), reflect.Int)
	test.Equals(t, result, 20)
}

func TestInvokerWithNoParamsFloatResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func() float64{
		return 5.6
	}

	result, err := invoker.Invoke(function)

	test.Equals(t, err, nil)
	test.Equals(t, result, 5.6)
}

func TestInvokerWithNoParamsStructResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	type myStruct struct{
		A string
		B bool
		C float64
		D int
	}

	function := func() myStruct{
		return myStruct{A:"text", B:true, C:2.6, D:50}
	}

	result, err := invoker.Invoke(function)

	expectedResult := myStruct{A:"text", B:true, C:2.6, D:50}
	
	test.Equals(t, err, nil)
	test.Equals(t, reflect.TypeOf(result).Kind(), reflect.TypeOf(expectedResult).Kind())
	test.Equals(t, result.(myStruct).A, expectedResult.A)
	test.Equals(t, result.(myStruct).B, expectedResult.B)
	test.Equals(t, result.(myStruct).C, expectedResult.C)
	test.Equals(t, result.(myStruct).D, expectedResult.D)
}

func TestInvokerWithNoParamsArrayOfStructResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	type myStruct struct{
		A string
		B bool
		C float64
		D int
	}

	function := func() []myStruct{
		var array []myStruct
		array = append(array, myStruct{A:"text", B:true, C:2.6, D:50})
		return array
	}

	result, err := invoker.Invoke(function)

	arrayResult := result.([]myStruct)

	expectedResult := myStruct{A:"text", B:true, C:2.6, D:50}
	
	test.Equals(t, err, nil)
	test.Equals(t, len(arrayResult), 1)
	test.Equals(t, arrayResult[0].A, expectedResult.A)
	test.Equals(t, arrayResult[0].B, expectedResult.B)
	test.Equals(t, arrayResult[0].C, expectedResult.C)
	test.Equals(t, arrayResult[0].D, expectedResult.D)
}