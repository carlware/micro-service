// +build unit

package account

import (
	"carlware/accounts/internal/helpers"
	"carlware/accounts/internal/interfaces/mocks"
	"carlware/accounts/internal/models"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

var fa = &models.Account{
	ID:            "1",
	BankName:      "BBVA",
	AccountNumber: "123456789012345678",
	Holder:        "Me",
}

var de = xerrors.New("database error")

func TestCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *Opts
		req  *CreateRequest
	}
	tests := []struct {
		name string
		args func(t *gomock.Controller) args

		want1      *models.Account
		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation
	}{
		{
			name: "new account was added successfully",
			args: func(t *gomock.Controller) args {
				r := &CreateRequest{}
				_ = helpers.Copy(r, fa)
				p := mocks.NewMockAccount(t)
				// TODO: test the reciver match
				p.EXPECT().Add(gomock.Any(), gomock.Any()).Return(fa, nil).Times(1)
				return args{
					ctx: context.TODO(),
					opts: &Opts{
						Repository: p,
					},
					req: r,
				}
			},
			want1: fa,
		},
		{
			name: "if req is nill return an error",
			args: func(t *gomock.Controller) args {
				return args{}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Create: <InvalidArgument> nil request error")
			},
		},
		{
			name: "if req is not nil but validation fail return an error",
			args: func(t *gomock.Controller) args {
				return args{
					req: &CreateRequest{},
				}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Create: bank: cannot be blank; holder: cannot be blank; number: cannot be blank.")
			},
		},
		{
			name: "if req is valid but database throw an error, returns an error",
			args: func(t *gomock.Controller) args {
				r := &CreateRequest{}
				_ = helpers.Copy(r, fa)
				p := mocks.NewMockAccount(t)
				p.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil, helpers.NewDatabaseError("path", de)).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					req: r,
				}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Create: path: database error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			tArgs := tt.args(mc)

			got1, err := Create(tArgs.ctx, tArgs.opts, tArgs.req)

			assert.Equal(t, tt.want1, got1, "Create returned unexpected result")

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRetrieve(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *Opts
		id   string
	}
	tests := []struct {
		name string
		args func(t *gomock.Controller) args

		want1      *models.Account
		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation
	}{
		{
			name: "new account get successfully",
			args: func(t *gomock.Controller) args {
				p := mocks.NewMockAccount(t)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fa, nil).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					id: "1",
				}
			},
			want1: fa,
		},
		{
			name: "if req is nill return an error",
			args: func(t *gomock.Controller) args {
				return args{}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Retrieve: <InvalidArgument> id cannot be blank")
			},
		},
		{
			name: "if req is valid but database throw an error, returns an error",
			args: func(t *gomock.Controller) args {
				r := &CreateRequest{}
				_ = helpers.Copy(r, fa)
				p := mocks.NewMockAccount(t)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, helpers.NewDatabaseError("path", de)).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					id: "1",
				}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Retrieve: path: database error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			tArgs := tt.args(mc)

			got1, err := Retrieve(tArgs.ctx, tArgs.opts, tArgs.id)

			assert.Equal(t, tt.want1, got1, "Retrieve returned unexpected result")

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	updatedA := &models.Account{}
	_ = helpers.Copy(updatedA, fa)
	updatedA.BankName = "BBVA"
	type args struct {
		ctx  context.Context
		opts *Opts
		id   string
		req  *UpdateRequest
	}
	tests := []struct {
		name string
		args func(t *gomock.Controller) args

		want1      *models.Account
		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation
	}{
		{
			name: "new account was added successfully",
			args: func(t *gomock.Controller) args {
				r := &UpdateRequest{}
				_ = helpers.Copy(r, fa)
				c := &models.Account{}
				_ = helpers.Copy(r, fa)
				p := mocks.NewMockAccount(t)
				// TODO: test the reciver match
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(c, nil).Times(1)
				p.EXPECT().Update(gomock.Any(), gomock.Any()).Return(updatedA, nil).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					req: r,
				}
			},
			want1: fa,
		},
		{
			name: "if req is nil return an error",
			args: func(t *gomock.Controller) args {
				return args{}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Update: <InvalidArgument> nil request error")
			},
		},
		{
			name: "if req is not nil but validation fail return an error",
			args: func(t *gomock.Controller) args {
				p := mocks.NewMockAccount(t)
				c := &models.Account{}
				_ = helpers.Copy(c, fa)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(c, nil).Times(1)
				return args{
					req: &UpdateRequest{AccountNumber: "a"},
					opts: &Opts{
						Repository: p,
					},
				}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Update: number: the length must be exactly 18.")
			},
		},
		{
			name: "if req is valid but database throw an error, returns an error",
			args: func(t *gomock.Controller) args {
				r := &UpdateRequest{}
				_ = helpers.Copy(r, fa)
				c := &models.Account{}
				_ = helpers.Copy(r, fa)
				p := mocks.NewMockAccount(t)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(c, nil).Times(1)
				p.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, helpers.NewDatabaseError("path", de)).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					req: r,
				}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Update: path: database error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			tArgs := tt.args(mc)

			got1, err := Update(tArgs.ctx, tArgs.opts, tArgs.id, tArgs.req)

			assert.Equal(t, tt.want1, got1, "Update returned unexpected result")

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *Opts
		id   string
	}
	tests := []struct {
		name string
		args func(t *gomock.Controller) args

		want1      *models.Account
		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation
	}{
		{
			name: "new account deleted successfully",
			args: func(t *gomock.Controller) args {
				p := mocks.NewMockAccount(t)
				p.EXPECT().Remove(gomock.Any(), &models.Account{ID: "1"}).Return(fa, nil).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					id: "1",
				}
			},
			want1: fa,
		},
		{
			name: "if req is nill return an error",
			args: func(t *gomock.Controller) args {
				return args{}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Delete: <InvalidArgument> id cannot be blank")
			},
		},
		{
			name: "if req is valid but database throw an error, returns an error",
			args: func(t *gomock.Controller) args {
				r := &CreateRequest{}
				_ = helpers.Copy(r, fa)
				p := mocks.NewMockAccount(t)
				p.EXPECT().Remove(gomock.Any(), gomock.Any()).Return(nil, xerrors.New("database error")).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
					id: "1",
				}
			},
			want1:   nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "cases.account.Delete: database error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			tArgs := tt.args(mc)

			got1, err := Delete(tArgs.ctx, tArgs.opts, tArgs.id)

			assert.Equal(t, tt.want1, got1, "Delete returned unexpected result")

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		cxt  context.Context
		opts *Opts
	}
	tests := []struct {
		name string
		args func(t *gomock.Controller) args

		want1      []*models.Account
		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation
	}{
		{
			name: "Get a list of accounts",
			args: func(t *gomock.Controller) args {
				var list []*models.Account
				list = append(list, fa, fa, fa)
				p := mocks.NewMockAccount(t)
				p.EXPECT().List(gomock.Any()).Return(list, nil).Times(1)
				return args{
					opts: &Opts{
						Repository: p,
					},
				}
			},
			want1: []*models.Account{fa, fa, fa},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			tArgs := tt.args(mc)

			got1, err := List(tArgs.cxt, tArgs.opts)

			assert.Equal(t, tt.want1, got1, "List returned unexpected result")

			if tt.wantErr {
				if assert.Error(t, err) && tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
