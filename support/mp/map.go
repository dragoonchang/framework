package mp

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Accessible(input any) bool {
	value := reflect.ValueOf(input)

	switch value.Kind() {
	case reflect.Map:
		return true
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			if field.PkgPath == "" {
				return true
			}
		}
	}

	return false
}

func Add(input map[any]any, key any, value any) map[any]any {
	if Get(input, key) == nil {
		Set(input, key, value)
	}
	return input
}

func Collapse(input []any) []any {
	results := make([]any, 0)

	for _, values := range input {
		switch v := values.(type) {
		case []any:
			results = append(results, v...)
		}
	}

	return results
}

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

func Divide(input map[any]any) ([]any, []any) {
	keys := make([]any, 0, len(input))
	values := make([]any, 0, len(input))

	for key, value := range input {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

func Get(array map[any]any, key any, defaultValue ...any) any {
	if key == nil {
		return array
	}

	if value, exists := array[key]; exists {
		return value
	}

	keyStr, ok := key.(string)
	if !ok || !strings.Contains(keyStr, ".") {
		return defaultValue
	}

	segments := strings.Split(keyStr, ".")
	for _, segment := range segments {
		value, exists := array[segment]
		if !exists {
			return defaultValue
		}

		subArray, ok := value.(map[any]any)
		if !ok {
			return defaultValue
		}

		array = subArray
	}

	return array
}

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

func Undot(array map[string]any) map[any]any {
	results := make(map[any]any)

	for key, value := range array {
		Set(results, key, value)
	}

	return results
}

func Except(array map[any]any, keys []any) map[any]any {
	for _, key := range keys {
		delete(array, key)
	}

	return array
}

func Exists(array map[any]any, key any) bool {
	_, exists := array[key]
	return exists
}

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

func Forget(array map[any]any, keys []any) {
	for _, key := range keys {
		delete(array, key)
	}
}

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

//func IsAssoc(array map[any]any) bool {
//	for i, key := range getSortedKeys(array) {
//		if i != key {
//			return true
//		}
//	}
//	return false
//}

//func IsList(array map[any]any) bool {
//	return !IsAssoc(array)
//}

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

func KeyBy(array map[any]any, keyBy func(any) any) map[any]any {
	result := make(map[any]any)

	for _, item := range array {
		key := keyBy(item)
		result[key] = item
	}

	return result
}

func PrependKeysWith(array map[any]any, prependWith string) map[any]any {
	result := make(map[any]any)

	for key, value := range array {
		newKey := fmt.Sprintf("%s%v", prependWith, key)
		result[newKey] = value
	}

	return result
}

func Only(array map[any]any, keys []any) map[any]any {
	result := make(map[any]any)

	for _, key := range keys {
		if value, ok := array[key]; ok {
			result[key] = value
		}
	}

	return result
}

// Prepend the given value to the beginning of an array or associative array.
func Prepend[K comparable, V any](arr map[K]V, value V, key K) map[K]V {
	arr[key] = value
	return arr
}

func Set(array map[any]any, key any, value any) map[any]any {
	if key == nil {
		return map[any]any{
			nil: value,
		}
	}

	keys := strings.Split(fmt.Sprintf("%v", key), ".")
	current := array

	for i, key := range keys {
		if i == len(keys)-1 {
			current[key] = value
		} else {
			if _, exists := current[key]; !exists {
				current[key] = make(map[any]any)
			}

			subArray, ok := current[key].(map[any]any)
			if !ok {
				return array
			}

			current = subArray
		}
	}

	return array
}

// ToCssClasses Convert a map of string-bool pairs to a string of CSS classes.
func ToCssClasses(mp map[any]bool) string {
	var classList []string

	for k, v := range mp {
		if v {
			classList = append(classList, fmt.Sprint(k))
		}
	}

	return strings.Join(classList, " ")
}

// ToCssStyles Convert a map of string-bool pairs to a string of CSS styles.
func ToCssStyles[T comparable](mp map[T]bool) string {
	//func ToCssStyles(mp map[any]bool) string {
	var styleStrings []string

	for k, v := range mp {
		if v {
			styleStrings = append(styleStrings, fmt.Sprint(k))
		}
	}

	return strings.Join(styleStrings, "; ")
}

// Query converts an array (map) into a query string.
func Query(arr map[string]string) string {
	values := url.Values{}
	for key, value := range arr {
		values.Add(key, value)
	}
	return values.Encode()
}
