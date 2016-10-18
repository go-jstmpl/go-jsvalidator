package slices

import (
	"fmt"
	"reflect"
)

func toSlice(input interface{}) ([]interface{}, error) {
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(input)
		l := s.Len()
		slice := make([]interface{}, l)
		for i := 0; i < l; i++ {
			slice[i] = s.Index(i)
		}
		return slice, nil
	default:
		return nil, TypeError{fmt.Sprintf("%T should be slice", input)}
	}
}
