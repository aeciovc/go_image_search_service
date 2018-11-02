package rabbitmq

import (
	"testing"
	test "github.com/aeciovc/go-image-search/test"
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

//TODO Support boolean returns
func TestInvokerWithNoParamsBooleanResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func() bool{
		return true
	}

	result, err := invoker.Invoke(function)

	test.Equals(t, err, nil)
	test.Equals(t, result, true)
}

//TODO Support integer returns
func TestInvokerWithNoParamsIntegerResultMethodSuccess(t *testing.T) {

	invoker := GetInvoker()

	function := func() int{
		return 20
	}

	result, err := invoker.Invoke(function)

	test.Equals(t, err, nil)
	test.Equals(t, result, 20)
}