package mp

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// Accessible Determine whether the given value is array accessible.
func Accessible[T any](value T) bool {
	k := reflect.ValueOf(value).Kind()
	return k == reflect.Map || k == reflect.Slice || k == reflect.Array
}

// Add an element to an array using “dot” notation if it doesn't exist.
func Add(arr map[string]any, key string, value any) (map[string]any, error) {
	if len(key) == 0 {
		return nil, ErrInvalidKey
	}
	v, err := Get(arr, key, nil)
	if err != nil {
		return nil, err
	}
	if v == nil {
		Set(&arr, key, value)
	}
	return arr, nil
}

// Collapse collapses an array of arrays into a single array.
func Collapse(array any) []interface{} {
	res := make([]any, 0)

	recursiveCollapse(array, &res)
	return res
}

func recursiveCollapse(value any, res *[]any) {
	switch v := value.(type) {
	case map[string]map[string]interface{}:
		for _, vv := range v {
			recursiveCollapse(vv, res)
		}
	case map[string]interface{}:
		for _, vv := range v {
			recursiveCollapse(vv, res)
		}
	default:
		*res = append(*res, v)
	}
}

// CrossJoin returns all possible permutations of the given arrays.
// todo: check & test cases
func CrossJoin(arrays ...[]any) [][]any {
	results := [][]any{{}}

	for _, array := range arrays {
		var appendResult [][]any

		for _, product := range results {
			for _, item := range array {
				newProduct := make([]any, len(product))
				copy(newProduct, product)
				newProduct = append(newProduct, item)

				appendResult = append(appendResult, newProduct)
			}
		}

		results = appendResult
	}

	return results
}

// Divide an array into two arrays. One with keys and the other with values.
// todo: check & test cases
func Divide(input map[any]any) ([]any, []any) {
	keys := make([]any, 0, len(input))
	values := make([]any, 0, len(input))

	for key, value := range input {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

// Dot returns a flattened associative array with dot notation
// todo: check & test cases
func Dot(array map[any]any, prepend string) map[string]any {
	results := make(map[string]any)

	for key, value := range array {
		switch v := value.(type) {
		case map[any]any:
			if len(v) != 0 {
				subResults := Dot(v, fmt.Sprintf("%s%v.", prepend, key))
				for k, val := range subResults {
					results[k] = val
				}
			} else {
				results[fmt.Sprintf("%s%v", prepend, key)] = value
			}
		default:
			results[fmt.Sprintf("%s%v", prepend, key)] = value
		}
	}

	return results
}

// Undot returns an expanded array from flattened dot notation array
// todo: check & test cases
func Undot(array map[string]any) map[string]any {
	res := make(map[string]any)

	for key, value := range array {
		Set(&res, key, value)
	}

	return res
}

// Except returns all the given array except for a specified array of keys.
// todo: check & test cases
func Except(array map[any]any, keys []any) map[any]any {
	for _, key := range keys {
		delete(array, key)
	}

	return array
}

// Exists determines if the given key exists in the provided array.
// todo: check & test cases
func Exists(array map[any]any, key any) bool {
	_, exists := array[key]
	return exists
}

// First returns the first element in an array that passes a given truth test.
// todo: check & test cases
func First(array map[any]any, callback func(any, any) bool, defaultValue any) any {
	if callback == nil {
		if len(array) == 0 {
			return defaultValue
		}

		for _, value := range array {
			return value
		}
	}

	for key, value := range array {
		if callback(value, key) {
			return value
		}
	}

	return defaultValue
}

// Last returns the last element in an array passing a given truth test.
// todo: check & test cases
func Last(array map[any]any, callback func(any, any) bool, defaultValue any) any {
	if callback == nil {
		if len(array) == 0 {
			return defaultValue
		}

		var lastValue any
		for _, value := range array {
			lastValue = value
		}
		return lastValue
	}

	reversedArray := make(map[any]any)
	for key, value := range array {
		reversedArray[key] = value
	}

	return First(reversedArray, callback, defaultValue)
}

// Flatten flattens a multi-dimensional array into a single level.
// todo: check & test cases
func Flatten(array map[any]any, depth int) []any {
	result := []any{}

	for _, item := range array {
		subArray, ok := item.(map[any]any)
		if !ok {
			result = append(result, item)
		} else {
			values := make([]any, 0)
			if depth == 1 {
				for _, v := range subArray {
					values = append(values, v)
				}
			} else {
				values = Flatten(subArray, depth-1)
			}

			for _, value := range values {
				result = append(result, value)
			}
		}
	}

	return result
}

// Forget Remove one or many array items from a given array.
// todo: check & test cases
func Forget(array map[any]any, keys []any) {
	for _, key := range keys {
		delete(array, key)
	}
}

// Get an item from a map using "dot" notation.
func Get(array map[string]any, key string, def ...any) (any, error) {
	if len(def) > 1 {
		return nil, ErrInvalidParam
	}

	if array == nil || len(array) == 0 {
		if len(def) == 0 {
			return nil, nil
		}
		return def[0], nil
	}

	if len(key) == 0 {
		return array, nil
	}

	if value, exists := array[key]; exists {
		return value, nil
	}

	if !strings.Contains(key, ".") {
		return def[0], nil
	}

	m := array
	keys := strings.Split(key, ".")
	for _, v := range keys {
		value, exists := m[v]
		if !exists {
			return def[0], nil
		}

		if value == nil {
			return nil, nil
		}

		subArray, ok := value.(map[string]any)
		if !ok {
			return value, nil
		}

		m = subArray
	}

	return m, nil
}

// Has checks if an item or items exist in an array using "dot" notation.
// todo: check & test cases
func Has(array map[any]any, keys []any) bool {
	if len(array) == 0 || len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		subKeyArray := array
		if Exists(array, key) {
			continue
		}

		for _, segment := range strings.Split(fmt.Sprintf("%v", key), ".") {
			if Exists(subKeyArray, segment) {
				subKeyArray = subKeyArray[segment].(map[any]any)
			} else {
				return false
			}
		}
	}

	return true
}

// HasAny Determine if any of the keys exist in an array using int key
// todo: check & test cases
func HasAny(array map[any]any, keys []any) bool {
	if len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		if Has(array, []any{key}) {
			return true
		}
	}

	return false
}

// IsAssoc Determines if an array is associative.
// todo: finish & test cases
//func IsAssoc(array map[any]any) bool {
//	for i, key := range getSortedKeys(array) {
//		if i != key {
//			return true
//		}
//	}
//	return false
//}

// IsList Determines if an array is a list.
// todo: finish & test cases
//func IsList(array map[any]any) bool {
//	return !IsAssoc(array)
//}

// Join concatenates elements of a slice into a string with a specified delimiter and final separator
// todo: finish & test cases
//func Join(array []any, glue string, finalGlue string) string {
//	if finalGlue == "" {
//		return strings.Join(toStringArray(array), glue)
//	}
//
//	if len(array) == 0 {
//		return ""
//	}
//
//	if len(array) == 1 {
//		return fmt.Sprintf("%v", array[0])
//	}
//
//	finalItem := array[len(array)-1]
//	array = array[:len(array)-1]
//	return strings.Join(toStringArray(array), glue) + finalGlue + fmt.Sprintf("%v", finalItem)
//}

// KeyBy Key an associative array by a field or using a callback.
// todo: check & test cases
func KeyBy(array map[any]any, keyBy func(any) any) map[any]any {
	result := make(map[any]any)

	for _, item := range array {
		key := keyBy(item)
		result[key] = item
	}

	return result
}

// PrependKeysWith Prepend the key names of an associative array.
// todo: check & test cases
func PrependKeysWith(array map[any]any, prependWith string) map[any]any {
	result := make(map[any]any)

	for key, value := range array {
		newKey := fmt.Sprintf("%s%v", prependWith, key)
		result[newKey] = value
	}

	return result
}

// Only returns a subset of the items from the given map with specified keys.
// todo: check & test cases
func Only(array map[any]any, keys []any) map[any]any {
	result := make(map[any]any)

	for _, key := range keys {
		if value, ok := array[key]; ok {
			result[key] = value
		}
	}

	return result
}

// Pluck an array of values from an array.
// todo: finish & test cases
//func Pluck[T any](array []T, value T, key ...int) ([]T, error) {
//}

// ExplodePluckParameters Explode the "value" and "key" arguments passed to "pluck".
// todo: finish & test cases
//func ExplodePluckParameters[T any](array []T, key ...int) ([]T, error) {
//}

// Map Run a map over each of the items in the array.
// todo: finish & test cases
//func Map[T, U any](arr []T, fn func(T, int) U) []U {
//}

// Prepend the given value to the beginning of an array or associative array.
// todo: check & test cases
func Prepend[K comparable, V any](arr map[K]V, value V, key K) map[K]V {
	arr[key] = value
	return arr
}

// Pull Get a value from the array, and remove it.
// todo: finish & test cases
//func Pull[T any](arr []T, key int, def T) ([]T, T, error) {
//}

// Query converts a map into a query string.
// todo: check & test cases
func Query(arr map[string]string) string {
	values := url.Values{}
	for key, value := range arr {
		values.Add(key, value)
	}
	return values.Encode()
}

// Random returns one or a specified number of random values from a slice.
// todo: finish & test cases
//func Random[T any](arr []T, number *int) ([]T, error) {
//}

// Set a map item to a given value using int key
func Set(dict *map[string]any, key string, value any) error {
	if len(key) == 0 {
		if isMapStringInterface(value) {
			*dict = value.(map[string]any)
			return nil
		}

		return ErrInvalidValue
	}

	keys := strings.Split(key, ".")
	m := *dict

	for i, v := range keys {
		if i == len(keys)-1 {
			m[v] = value
		} else {
			if !isMapStringInterface(m[v]) {
				m[v] = make(map[string]any)
			} else if _, exists := m[v]; !exists {
				m[v] = make(map[string]any)
			}

			subArray, ok := m[v].(map[string]any)
			if !ok {
				return nil
			}

			m = subArray
		}
	}

	return nil
}

func isMapStringInterface(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Map && rv.Type().Key().Kind() == reflect.String
}

// Shuffle the given array and return the result.
// todo: finish & test cases
//func Shuffle[T any](arr []T, seed *int64) []T {
//}

// Sort the nested array using the given callback.
// todo: finish & test cases
//func Sort(arr []any, fn func(i, j int) bool) []any {
//}

// Sort the nested array in descending order using the given callback.
// todo: finish & test cases
//func SortDesc(arr []any, fn func(i, j int) bool) []any {
//}

// SortRecursive Recursively sort an array by values.
// todo: finish & test cases
//func SortRecursive(arr []any, descending bool) ([]any, error) {
//}

// ToCssClasses Convert a map of string-bool pairs to a string of CSS classes.
func ToCssClasses(mp map[string]bool) string {
	var classList []string

	for i, v := range mp {
		if v {
			classList = append(classList, i)
		}
	}

	return strings.Join(classList, " ")
}

// ToCssStyles Convert a map of string-bool pairs to a string of CSS styles.
func ToCssStyles(mp map[string]bool) string {
	var styleStrings []string

	for i, v := range mp {
		if v {
			styleStrings = append(styleStrings, i)
		}
	}

	return strings.Join(styleStrings, "; ")
}

// Where Filter the array using the given callback.
// todo: finish & test cases
//func Where[T any](arr []T, fn func(T) bool) []T {
//}

// WhereNotNull Filter items where the value is not null.
// todo: finish & test cases
//func WhereNotNull[T any](arr []T) []T {
//}

// Wrap If the given value is not an array and not null, wrap it in one.
// todo: finish & test cases
//func Wrap(value any) []any {
//}
