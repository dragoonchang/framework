package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessible(t *testing.T) {
}

func TestAdd(t *testing.T) {
}

func TestCollapse(t *testing.T) {
}

func TestCrossJoin(t *testing.T) {
}

func TestDivide(t *testing.T) {
}

func TestDot(t *testing.T) {
}

func TestUndot(t *testing.T) {
}

func TestExcept(t *testing.T) {
}

func TestExists(t *testing.T) {
}

func TestFirst(t *testing.T) {
}

func TestLast(t *testing.T) {
}

func TestFlatten(t *testing.T) {
}

func TestForget(t *testing.T) {
}

func TestGet(t *testing.T) {
}

func TestHas(t *testing.T) {
}

func TestHasAny(t *testing.T) {
}

func TestJoin(t *testing.T) {
}

func TestKeyBy(t *testing.T) {
}

func TestPrependKeysWith(t *testing.T) {
}

func TestOnly(t *testing.T) {
}

func TestMap(t *testing.T) {
}

func TestPrepend(t *testing.T) {
}

func TestPull(t *testing.T) {
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

func TestRandom(t *testing.T) {
}

func TestSet(t *testing.T) {
}

func TestShuffle(t *testing.T) {
}

func TestSort(t *testing.T) {
}

func TestSortDesc(t *testing.T) {
}

func TestSortRecursive(t *testing.T) {
}

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

func TestWhere(t *testing.T) {
}

func TestWhereNotNull(t *testing.T) {
}

func TestWrap(t *testing.T) {
}
