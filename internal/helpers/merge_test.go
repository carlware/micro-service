// +build unit

package helpers

import (
	"testing"
)

func TestUpdateStruct(t *testing.T) {
	type args struct {
		dst interface{}
		tmp interface{}
		src interface{}
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		// TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := UpdateStruct(tArgs.dst, tArgs.tmp, tArgs.src)

			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateStruct error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
