package CollectionUtils

import "reflect"

// extract string field from list
func MapToStringSlice(length int, action func(index int) string) []string {
	if length <= 0 || action == nil {
		return []string{}
	}
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, action(i))
	}
	return result
}

// extract int field from list
func MapToIntSlice(list interface{}, action func(index int) int) []int {
	if list == nil || action == nil {
		return []int{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []int{}
	}
	result := make([]int, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract int64 field from list
func MapToInt64Slice(list interface{}, action func(index int) int64) []int64 {
	if list == nil || action == nil {
		return []int64{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []int64{}
	}
	result := make([]int64, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract float64 field from list
func MapToFloat64Slice(list interface{}, action func(index int) float64) []float64 {
	if list == nil || action == nil {
		return []float64{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []float64{}
	}
	result := make([]float64, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract float32 field from list
func MapToFloat32Slice(list interface{}, action func(index int) float32) []float32 {
	if list == nil || action == nil {
		return []float32{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []float32{}
	}
	result := make([]float32, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}
