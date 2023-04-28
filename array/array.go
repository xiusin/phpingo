package array

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

type Numerable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type BasicType interface {
	~bool | ~string | Numerable
}

func In[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func Find[T any](haystack []T, cb func(item T, index int) bool) (val T, ok bool) {
	if cb == nil {
		cb = func(item T, index int) bool {
			return true
		}
	}
	for index, t := range haystack {
		if cb(t, index) {
			return t, true
		}
	}
	return
}

func Some[T any](array []T, cb func(item T, index int) bool) bool {
	if cb == nil {
		cb = func(item T, index int) bool {
			return true
		}
	}
	var someResult = false
	for index, t := range array {
		someResult = cb(t, index)
	}
	return someResult
}

func Every[T any](array []T, cb func(item T, index int) bool) bool {
	if cb == nil {
		cb = func(item T, index int) bool {
			return true
		}
	}
	for index, t := range array {
		if !cb(t, index) {
			return false
		}
	}
	return true
}

func Filter[T any](array []T, cb func(item T, index int) bool) []T {
	var ret []T
	if cb == nil {
		cb = func(item T, index int) bool {
			return true
		}
	}

	for index, t := range array {
		if cb(t, index) {
			ret = append(ret, t)
		}
	}
	return ret
}

func Chunk[T any](array []T, length int) [][]T {
	var chunks [][]T
	if length > 0 {
		count, pos := len(array), 0
		for pos < count {
			if pos+length < count {
				chunks = append(chunks, array[pos:pos+length])
			} else {
				chunks = append(chunks, array[pos:])
			}
			pos += length
		}
	}
	return chunks
}

func columnParse[K BasicType, V any](array any, columnKey string, indexKey string, isMap bool) any {
	valueRef := reflect.ValueOf(array)
	if valueRef.Kind() != reflect.Ptr {
		return nil
	}
	valueRef = valueRef.Elem()
	var mapValues = map[K]V{}
	var sliceValues []V
	if length := valueRef.Len(); length > 0 {
		kind := valueRef.Index(0).Kind()
		for i := 0; i < length; i++ {
			var key K
			var value V
			item := valueRef.Index(i)
			switch kind {
			case reflect.Map:
				for _, e := range item.MapKeys() {
					if e.Interface().(string) == indexKey && isMap {
						key, _ = item.MapIndex(e).Interface().(K)
					}
					if e.Interface().(string) == columnKey {
						value, _ = item.MapIndex(e).Interface().(V)
					}
				}
			case reflect.Struct:
				if isMap {
					key, _ = item.FieldByName(indexKey).Interface().(K)
				}
				value, _ = item.FieldByName(columnKey).Interface().(V)

			default:
				panic(fmt.Errorf("暂不支持的类型"))
			}
			if isMap {
				mapValues[key] = value
			} else {
				sliceValues = append(sliceValues, value)
			}
		}
	}
	if isMap {
		return mapValues
	}
	return sliceValues
}

func ColumnWithIndex[K BasicType, V any](array any, columnKey string, indexKey string) map[K]V {
	return columnParse[K, V](array, columnKey, indexKey, true).(map[K]V)
}

func Column[T any](array any, columnKey string) []T {
	return columnParse[int, T](array, columnKey, "", false).([]T)
}

func Sum[T Numerable](array []T) T {
	var sum T
	for _, v := range array {
		sum += v
	}
	return sum
}

func Keys[K ~string | Numerable, T any](arr map[K]T) []K {
	var keys = make([]K, 0, len(arr))
	for k := range arr {
		keys = append(keys, k)
	}
	return keys
}

func Values[K ~string | Numerable, T any](array map[K]T) []T {
	var values = make([]T, 0, len(array))
	for _, v := range array {
		values = append(values, v)
	}
	return values
}

func Push[T any](array *[]T, values ...T) {
	*array = append(*array, values...)
}

func Pop[T any](array *[]T) (T, bool) {
	var ele T
	index := len(*array) - 1
	if index < 0 {
		return ele, false
	}
	ele = (*array)[index]
	*array = (*array)[:index]
	return ele, true
}

func Shift[T any](array *[]T) (T, bool) {
	var firstEle T
	if len(*array) == 0 {
		return firstEle, false
	}
	firstEle = (*array)[0]
	*array = (*array)[1:]
	return firstEle, true
}

func UnShift[T any](array *[]T, values ...T) {
	*array = append(values, *array...)
}

// Flip
func Flip[K BasicType, V BasicType](array map[K]V) map[V]K {
	var values = map[V]K{}
	for k, v := range array {
		values[v] = k
	}
	return values
}

func SearchMap[K BasicType, V BasicType](needle V, haystack map[K]V) (K, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	var k K
	return k, false
}

func SearchSlice[V BasicType](needle V, haystack []V) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return -1, false
}

func KeyExists[K BasicType, V BasicType](k K, array map[K]V) bool {
	_, exists := array[k]
	return exists
}

func Reverse[T BasicType](array []T) []T {
	length := len(array)
	var clone = make([]T, length)
	for k, v := range array {
		clone[length-1-k] = v
	}
	return clone
}

// Unique 仅支持基本类型 Warning: Array to string conversion
func Unique[T BasicType](array []T) []T {
	var values []T
	var filter = map[T]struct{}{}
	for _, v := range array {
		if _, ok := filter[v]; !ok {
			filter[v] = struct{}{}
			values = append(values, v)
		}
	}
	return values
}

func Intersect[T BasicType](array1, array2 []T, arrays ...[]T) []T {
	var intersect []T
	for _, arr := range arrays {
		array2 = append(array2, arr...)
	}
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v2 == v1 {
				intersect = append(intersect, v2)
				break
			}
		}
	}
	return intersect
}

func Combine[K BasicType, V any](keys []K, values []V) map[K]V {
	if len(keys) != len(values) {
		return nil
	}
	var combine = map[K]V{}
	for index, key := range keys {
		combine[key] = values[index]
	}
	return combine
}

func Product[T Numerable](array []T) T {
	var product T
	for _, v := range array {
		product *= v
	}
	return product
}

func ColumnMap[K BasicType, V any](array *[]V, columnKey string) map[K]V {
	return Combine[K, V](Column[K](array, columnKey), *array)
}

func Sort(array any, less func(i, j int) bool) {
	sort.SliceStable(array, less)
}

func Slice[T any](array []T, start int, length ...int) []T {
	arrLen := len(array)
	if start >= arrLen {
		return nil
	}
	if len(length) == 0 || start+length[0] > arrLen {
		return array[start:]
	}
	return array[start : start+length[0]]
}

func Splice[T any](array *[]T, offset int, length int, replacement []T) []T {
	arrLen := len(*array)
	var removed = make([]T, length)
	var new = make([]T, offset, arrLen-length+len(replacement))
	if offset < 0 {
		offset = -offset
	}
	copy(new, (*array)[:offset])
	if length == 0 || offset+length > len(*array) {
		removed = (*array)[offset:]
		new = append(new, replacement...)
	} else {
		removed = (*array)[offset : offset+length]
		new = append(new, replacement...)
		new = append(new, (*array)[offset+length:]...)
	}
	*array = new
	return removed
}

func Diff[T BasicType](array1, array2 []T, arrays ...[]T) []T {
	var single []T
	for _, arr := range arrays {
		array2 = append(array2, arr...)
	}
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v2 == v1 {
				goto CONTINUE_OUT
			}
		}
		single = append(single, v1)
	CONTINUE_OUT:
	}
	return single
}

func Only[T any](array map[string]T, keys []string) map[string]T {
	var only = map[string]T{}
	for _, key := range keys {
		if v, ok := array[key]; ok {
			only[key] = v
		}
	}
	return only
}

func Except[T any](array map[string]T, keys []string) map[string]T {
	var except = map[string]T{}
	for key, value := range array {
		if In[string](key, keys) {
			except[key] = value
		}
	}
	return except
}

func Shuffle[T any](array []T) {
	rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i >= 0; i-- {
		index := rand.Int() % (i + 1)
		array[index], array[i] = array[i], array[index]
	}
}
