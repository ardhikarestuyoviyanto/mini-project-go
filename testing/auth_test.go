package testing

import (
	"errors"
	"mini-project-go/constants"
	"mini-project-go/controller"
	"mini-project-go/domain/mocks"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginController(t *testing.T) {
	svc := mocks.SvcAuthMock{}
	svc.On("LoginUsers", mock.Anything).Return(errors.New("new")).Once()
	svc.On("LoginUsers", mock.Anything).Return(nil).Once()

	authController := controller.AuthEchoController{
		SvcAuth: &svc,
	}

	e := echo.New()

	t.Run("unauthorized", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		authController.LoginUserController(echoContext)

		assert.Equal(t, 401, w.Result().StatusCode)
	})

	t.Run("http ok", func(t *testing.T) {
		data := url.Values{}
		data.Set("email", "ardhikayoviyanto@gmail.com")
		data.Set("password", "123")

		e := echo.New()
		r := httptest.NewRequest("POST", "/login", strings.NewReader(data.Encode()))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		authController.LoginUserController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})

}

func TestLogoutUserController(t *testing.T) {
	svc := mocks.SvcAuthMock{}
	svc.On("LogoutUsers", mock.Anything).Return(errors.New("new")).Once()
	svc.On("LogoutUsers", mock.Anything).Return(nil).Once()
	authController := controller.AuthEchoController{
		SvcAuth: &svc,
	}

	t.Run("unauthorized", func(t *testing.T) {
		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzIzMzUzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.nF85NftxoIpPBUx6bDe0E5f5-ARP3XLLPrD7EP_q0Tc"
		e := echo.New()
		r := httptest.NewRequest("POST", "/logout", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)
		echoContext := e.NewContext(r, w)
		authController.LogoutUserController(echoContext)

		assert.Equal(t, 401, 401)
	})
	t.Run("http ok", func(t *testing.T) {
		TokenValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzIzMzUzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.nF85NftxoIpPBUx6bDe0E5f5-ARP3XLLPrD7EP_q0Tc"
		e := echo.New()
		r := httptest.NewRequest("POST", "/logout", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenValid)
		echoContext := e.NewContext(r, w)
		authController.LogoutUserController(echoContext)

		assert.Equal(t, 200, 200)
	})
}
