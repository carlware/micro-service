package helpers

import (
	"reflect"
	"time"

	"github.com/imdario/mergo"
)

type timeTransformer struct {
}

func (t timeTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(time.Time{}) {
		return func(dst, src reflect.Value) error {
			if dst.CanSet() {
				isZero := dst.MethodByName("IsZero")
				result := isZero.Call([]reflect.Value{})
				if result[0].Bool() {
					dst.Set(src)
				}
			}
			return nil
		}
	}
	return nil
}

func UpdateStruct(dst, tmp, src interface{}) error {
	_ = Copy(tmp, src)
	return mergo.Merge(dst, tmp, mergo.WithTransformers(timeTransformer{}), mergo.WithOverride)
}
