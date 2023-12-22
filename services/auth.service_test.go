package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/IrvanWijaya/dealls-tinder/models"
	"github.com/IrvanWijaya/dealls-tinder/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestAuthService_SignUpUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBItf := NewMockDBItf(ctrl)

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		mock    func(args)
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			args: args{ctx: gin.CreateTestContextOnly(httptest.NewRecorder(), &gin.Engine{})},
			mock: func(args args) {
				fnUHashPassword = func(password string) (string, error) {
					defer func() {
						fnUHashPassword = utils.HashPassword
					}()

					return "password123", nil
				}

				fnTimeNow = func() time.Time {
					defer func() {
						fnTimeNow = time.Now
					}()

					return time.Unix(0, 0)
				}

				req := &models.SignUpInput{
					Name:            "Irv",
					Email:           "Irv@gmail.com",
					Password:        "password123",
					PasswordConfirm: "password123",
				}
				reqJSON, _ := json.Marshal(req)

				args.ctx.Request = httptest.NewRequest(
					"POST", "/", io.NopCloser(bytes.NewReader(reqJSON)),
				)

				mockDBItf.EXPECT().Create(&models.User{
					Name:      "Irv",
					Email:     "irv@gmail.com",
					Password:  "password123",
					IsPremium: false,
					CreatedAt: time.Unix(0, 0),
					UpdatedAt: time.Unix(0, 0),
				}).Return(nil)
			},
			wantErr: false,
			want: &models.User{
				Name:      "Irv",
				Email:     "irv@gmail.com",
				Password:  "password123",
				IsPremium: false,
				CreatedAt: time.Unix(0, 0),
				UpdatedAt: time.Unix(0, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(tt.args)
			}

			ac := &AuthService{
				DB: mockDBItf,
			}

			got, err := ac.SignUpUser(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.SignUpUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.SignUpUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
