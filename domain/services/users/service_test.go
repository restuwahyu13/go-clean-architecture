package users_service

import (
	"net/http"

	"testing"

	config "github.com/restuwahyu13/go-clean-architecture/configs"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	"github.com/stretchr/testify/assert"
)

func TestPing(r *testing.T) {
	con := config.NewTest()
	service := NewUsersService(dto.ServiceOptions{ENV: con.ENV, DB: con.DB, RDS: con.RDS})

	r.Run("Test Ping", func(t *testing.T) {
		res := service.Ping(con.CTX)

		if !assert.Equal(t, int(res.StatCode), http.StatusOK) {
			t.FailNow()

		} else if !assert.Equal(t, res.Message, "Ping!") {
			t.FailNow()
		}

		t.Log(res.Message)
	})

	r.Run("Database connected", func(t *testing.T) {
		if err := con.DB.Ping(); err != nil {
			t.FailNow()
		}
	})
}
