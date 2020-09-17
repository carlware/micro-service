// +build unit

package helpers

import (
	"testing"

	e "github.com/carlware/go-common/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestNewInternalErr(t *testing.T) {
	type args struct {
		op  string
		err error
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{

			name: "Test error message opcode",
			args: func(t *testing.T) args {
				return args{
					op:  "path.to.func.operation",
					err: xerrors.New("error message"),
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: <Internal> error message")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := NewInternalErr(tArgs.op, tArgs.err)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewInternalErr error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestNewNilErr(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{

			name: "Test nil err",
			args: func(t *testing.T) args {
				return args{
					op: "path.to.func.operation",
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: <InvalidArgument> nil request error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := NewNilErr(tArgs.op)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewNilErr error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestNewEmptyIDErr(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{

			name: "Test empty string",
			args: func(t *testing.T) args {
				return args{
					op: "path.to.func.operation",
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: <InvalidArgument> id cannot be blank")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := NewEmptyIDErr(tArgs.op)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewEmptyIDErr error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestNewNilRepositoryErr(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{

			name: "Test nil err",
			args: func(t *testing.T) args {
				return args{
					op: "path.to.func.operation",
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: <Internal> nil repository")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := NewNilRepositoryErr(tArgs.op)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewNilRepositoryErr error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestNewArgsErr(t *testing.T) {
	type args struct {
		op  string
		err error
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{

			name: "Test nil err",
			args: func(t *testing.T) args {
				return args{
					op:  "path.to.func.operation",
					err: xerrors.New("error message"),
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: error message")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := NewArgsErr(tArgs.op, tArgs.err)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewArgsErr error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestNewDatabaseError(t *testing.T) {
	type args struct {
		op  string
		err error
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{

			name: "Test default",
			args: func(t *testing.T) args {
				return args{
					op:  "path.to.func.operation",
					err: xerrors.New("error message"),
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: error message")
			},
		},
		{
			name: "Test SQL no rows",
			args: func(t *testing.T) args {
				return args{
					op:  "path.to.func.operation",
					err: xerrors.New("pg: no rows in result set"),
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "path.to.func.operation: pg: no rows in result set")
				assert.Equal(t, e.ErrorMessage(err), "A resource with this ID does not exists")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := NewDatabaseError(tArgs.op, tArgs.err)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewDatabaseError error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
