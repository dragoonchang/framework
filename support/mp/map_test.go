package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessible(t *testing.T) {
	assert.True(t, Accessible([]any{}))
	assert.True(t, Accessible([]any{1, 2}))
	assert.True(t, Accessible([5]any{1, 2, 3, 4, 5}))
	assert.True(t, Accessible(map[int]any{1: "a", 2: "b"}))
	assert.True(t, Accessible(map[string]any{"a": 1, "b": 2}))

	assert.False(t, Accessible("abc"))
	assert.False(t, Accessible(new(struct{})))
}

func TestAdd(t *testing.T) {
	array := map[string]any{
		"name": "Desk",
	}
	expected := map[string]any{
		"name":  "Desk",
		"price": 100,
	}
	result, err := Add(array, "price", 100)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	expected = map[string]any{
		"surname": "Mövsümov",
	}
	result, err = Add(map[string]any{}, "surname", "Mövsümov")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	expected = map[string]any{
		"developer": map[string]any{
			"name": "Ferid",
		},
	}
	result, err = Add(map[string]any{}, "developer.name", "Ferid")
	assert.NoError(t, err)

	assert.Equal(t, expected, result)

	expected = map[string]any{
		"1": "hAz",
	}
	result, err = Add(map[string]any{}, "1", "hAz")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	expected = map[string]any{
		"1": map[string]any{
			"1": "hAz",
		},
	}
	result, err = Add(map[string]any{}, "1.1", "hAz")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestCollapse(t *testing.T) {
	array := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := []any{"value1", "value2", "value3"}
	result := Collapse(array)
	assert.ElementsMatch(t, expected, result)

	array = map[string]any{
		"outer1": map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
		"outer2": map[string]any{
			"key3": "value3",
		},
	}
	expected = []any{"value1", "value2", "value3"}
	result = Collapse(array)
	assert.ElementsMatch(t, expected, result)

	input := map[string]any{
		"outer": map[string]any{
			"inner1": map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			"inner2": map[string]any{
				"key3": "value3",
				"key4": "value4",
			},
			"inner3": map[string]map[string]any{
				"nested": {
					"key5": "value5",
					"key6": "value6",
				},
			},
		},
		"outer2": "value7",
	}

	expected = []any{"value1", "value2", "value3", "value4", "value5", "value6", "value7"}
	result = Collapse(input)
	assert.ElementsMatch(t, expected, result)

	input = map[string]any{
		"outer": map[string]any{
			"inner1": map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			"inner2": []any{"value3", "value4"},
		},
		"outer2": "value5",
	}

	expected = []any{"value1", "value2", "value3", "value4", "value5"}
	result = Collapse(input)
	assert.ElementsMatch(t, expected, result)

	input = map[string]any{
		"outer": []any{
			map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			"value3",
			"value4",
		},
		"outer2": map[string]any{
			"inner1": []any{"value5", "value6"},
		},
	}

	expected = []any{"value1", "value2", "value3", "value4", "value5", "value6"}
	result = Collapse(input)
	assert.ElementsMatch(t, expected, result)
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
	array := map[string]any{
		"products.desk": map[string]any{
			"price": 100,
		},
	}
	expected := map[string]any{"price": 100}
	value, err := Get(array, "products.desk")
	assert.NoError(t, err)
	assert.Equal(t, expected, value)

	// Test null array values
	array = map[string]any{
		"foo": nil,
		"bar": map[string]any{
			"baz": nil,
		},
	}
	value, err = Get(array, "foo", "default")
	assert.NoError(t, err)
	assert.Nil(t, value)

	value, err = Get(array, "bar.baz", "default")
	assert.NoError(t, err)
	assert.Nil(t, value)

	// Test null key returns the whole array
	array = map[string]any{
		"foo": "bar",
	}
	value, err = Get(array, "")
	assert.NoError(t, err)
	assert.Equal(t, array, value)

	// Test array not an array
	value, err = Get(nil, "foo", "default")
	expectedStr := "default"
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, value)

	// Test array not an array and key is null
	value, err = Get(nil, "", "default")
	expectedStr = "default"
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, value)

	// Test array is empty and key is null
	value, err = Get(map[string]any{}, "")
	assert.NoError(t, err)
	assert.Empty(t, value)

	value, err = Get(map[string]any{}, "", "default")
	expectedStr = "default"
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, value)

	// Test numeric keys
	array = map[string]any{
		"products": map[string]any{
			"0": map[string]any{"name": "desk"},
			"1": map[string]any{"name": "chair"},
		},
	}
	expectedStr = "desk"
	value, err = Get(array, "products.0.name")
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, value)

	value, err = Get(array, "products.1.name")
	expectedStr = "chair"
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, value)

	// Test return default value for non-existing key.
	array = map[string]any{
		"names": map[string]any{
			"developer": "taylor",
		},
	}
	expectedStr = "dayle"
	value, err = Get(array, "names.otherDeveloper", "dayle")
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, value)
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
	// dot notation
	array := map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 100,
			},
		},
	}
	expected := map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 200,
			},
		},
	}
	err := Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// No key is given
	array = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 100,
			},
		},
	}
	expected = map[string]any{"price": 300}
	err = Set(&array, "", map[string]any{"price": 300})
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// The key doesn't exist at the depth
	array = map[string]any{
		"products": "desk",
	}
	expected = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 200,
			},
		},
	}
	err = Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// No corresponding key exists
	array = map[string]any{
		"": "products",
	}
	expected = map[string]any{
		"": "products",
		"products": map[string]any{
			"desk": map[string]any{
				"price": 200,
			},
		},
	}
	err = Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 100,
			},
		},
	}
	expected = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 100,
			},
		},
		"table": 500,
	}
	err = Set(&array, "table", 500)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 100,
			},
		},
	}
	expected = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 100,
			},
		},
		"table": map[string]any{
			"price": 350,
		},
	}
	err = Set(&array, "table.price", 350)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]any{}
	expected = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 200,
			},
		},
	}
	err = Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// Override
	array = map[string]any{
		"products": "table",
	}
	expected = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": 300,
			},
		}}
	err = Set(&array, "products.desk.price", 300)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]any{
		"1": "test",
	}
	expected = map[string]any{
		"1": "hAz",
	}
	err = Set(&array, "1", "hAz")
	assert.NoError(t, err)
	assert.Equal(t, expected, array)
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
	array := map[string]bool{
		"font-bold": true,
		"mt-4":      true,
		"ml-2":      true,
		"mr-2":      false,
	}
	result := ToCssClasses(array)
	expected := [...]string{"font-bold", "mt-4", "ml-2"}
	for _, v := range expected {
		assert.Contains(t, result, v)
	}
}

func TestToCssStyles(t *testing.T) {
	array := map[string]bool{
		"font-weight: bold;": true,
		"margin-top: 4px;":   true,
		"margin-left: 2px;":  true,
		"margin-right: 2px":  false,
	}
	result := ToCssStyles(array)
	expected := [...]string{"font-weight: bold;", "margin-top: 4px;", "margin-left: 2px;"}
	for _, v := range expected {
		assert.Contains(t, result, v)
	}
}

func TestWhere(t *testing.T) {
}

func TestWhereNotNull(t *testing.T) {
}

func TestWrap(t *testing.T) {
}
