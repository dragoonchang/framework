package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToCssClasses(t *testing.T) {

	classes := ToCssClasses(map[interface{}]bool{"font-bold": true, "mt-4": true, "ml-2": true, "mr-2": false})
	expected := "font-bold mt-4 ml-2"
	assert.Equal(t, expected, classes)
}

func TestToCssStyles(t *testing.T) {
	styles := ToCssStyles(map[string]bool{
		"font-weight: bold;": true,
		"margin-top: 4px;":   true,
		"margin-left: 2px;":  true,
		"margin-right: 2px":  false,
	})

	expected := "font-weight: bold; margin-top: 4px; margin-left: 2px;"
	if styles != expected {
		t.Errorf("ToCssStyles() = %q, expected %q", styles, expected)
	}
}

func TestPrepend(t *testing.T) {
	//arr1 := []string{"one", "two", "three", "four"}
	//result1 := Prepend(arr1, "zero")
	//expected1 := []string{"zero", "one", "two", "three", "four"}
	//assert.Equal(t, expected1, result1)

	arr2 := map[string]int{"one": 1, "two": 2}
	result2 := Prepend(arr2, 0, "zero")
	expected2 := map[string]int{"zero": 0, "one": 1, "two": 2}
	assert.Equal(t, expected2, result2)

	//arr3 := map[string]int{"one": 1, "two": 2}
	//result3 := Prepend(arr3, 0, nil)
	//expected3 := map[interface{}]int{nil: 0, "one": 1, "two": 2}
	//assert.Equal(t, expected3, result3)
}

func TestQuery(t *testing.T) {
	// Test case 1: Normal input
	arr := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	expected := "key1=value1&key2=value2"
	result := Query(arr)
	assert.Equal(t, expected, result)

	// Test case 2: Empty input
	arr = map[string]string{}
	expected = ""
	result = Query(arr)
	assert.Equal(t, expected, result)
}
