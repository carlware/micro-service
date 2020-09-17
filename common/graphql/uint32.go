package graphql

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUInt32(i uint32) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(int64(i), 10))
	})
}

func UnmarshalUInt32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(iv), nil
	case int:
		return uint32(v), nil
	case int64:
		return uint32(v), nil
	case json.Number:
		iv, err := strconv.ParseInt(string(v), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(iv), nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
