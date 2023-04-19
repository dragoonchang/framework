package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessible(t *testing.T) {
	assert.True(t, Accessible([]interface{}{}))
	assert.True(t, Accessible([]interface{}{1, 2}))
	assert.True(t, Accessible([5]interface{}{1, 2, 3, 4, 5}))
	assert.True(t, Accessible(map[int]interface{}{1: "a", 2: "b"}))
	assert.True(t, Accessible(map[string]interface{}{"a": 1, "b": 2}))

	assert.False(t, Accessible("abc"))
	assert.False(t, Accessible(new(struct{})))
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
	// dot notation
	array := map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 100,
			},
		},
	}
	expected := map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 200,
			},
		},
	}
	err := Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// No key is given
	array = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 100,
			},
		},
	}
	expected = map[string]interface{}{"price": 300}
	err = Set(&array, "", map[string]interface{}{"price": 300})
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// The key doesn't exist at the depth
	array = map[string]interface{}{
		"products": "desk",
	}
	expected = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 200,
			},
		},
	}
	err = Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// No corresponding key exists
	array = map[string]interface{}{
		"": "products",
	}
	expected = map[string]interface{}{
		"": "products",
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 200,
			},
		},
	}
	err = Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 100,
			},
		},
	}
	expected = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 100,
			},
		},
		"table": 500,
	}
	err = Set(&array, "table", 500)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 100,
			},
		},
	}
	expected = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 100,
			},
		},
		"table": map[string]interface{}{
			"price": 350,
		},
	}
	err = Set(&array, "table.price", 350)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]interface{}{}
	expected = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 200,
			},
		},
	}
	err = Set(&array, "products.desk.price", 200)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	// Override
	array = map[string]interface{}{
		"products": "table",
	}
	expected = map[string]interface{}{
		"products": map[string]interface{}{
			"desk": map[string]interface{}{
				"price": 300,
			},
		}}
	err = Set(&array, "products.desk.price", 300)
	assert.NoError(t, err)
	assert.Equal(t, expected, array)

	array = map[string]interface{}{
		"1": "test",
	}
	expected = map[string]interface{}{
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
