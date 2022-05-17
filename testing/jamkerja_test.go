package testing

import (
	"errors"
	"mini-project-go/constants"
	a "mini-project-go/controller/admin"
	"mini-project-go/domain/mocks"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateJamKerjaController(t *testing.T) {
	svc := mocks.SvcAdminMock{}
	svc.On("CreateJamKerja", mock.Anything).Return(errors.New("new")).Once()
	svc.On("CreateJamKerja", mock.Anything).Return(nil).Once()

	adminController := a.JamKerjaEchoController{
		SvcAdmin: &svc,
	}

	e := echo.New()
	t.Run("bad request", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/jamkerja", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)
		adminController.CreateController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("http ok", func(t *testing.T) {
		data := url.Values{}
		data.Set("nama", "test")
		data.Set("harilibur", "['sabtu','minggu']")
		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzU0MDYzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.9_omXwoS8NX4yF-U3us6HIr39sZnfIeVufwpauWkKQU"

		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/jamkerja", strings.NewReader(data.Encode()))
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)
		echoContext := e.NewContext(r, w)
		adminController.CreateController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllJamKerjaController(t *testing.T) {
	svc := mocks.SvcAdminMock{}
	svc.On("GetAllJamKerja", mock.Anything).Return(errors.New("new")).Once()
	svc.On("GetAllJamKerja", mock.Anything).Return(nil).Once()

	adminController := a.JamKerjaEchoController{
		SvcAdmin: &svc,
	}

	t.Run("http ok", func(t *testing.T) {

		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzU0MDYzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.9_omXwoS8NX4yF-U3us6HIr39sZnfIeVufwpauWkKQU"
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/jamkerja", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)
		echoContext := e.NewContext(r, w)
		adminController.GetAllController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetByIdJamKerjaController(t *testing.T) {
	svc := mocks.SvcAdminMock{}
	svc.On("GetByIdJamKerja", mock.Anything).Return(errors.New("new")).Once()
	svc.On("GetByIdJamKerja", mock.Anything).Return(nil).Once()

	adminController := a.JamKerjaEchoController{
		SvcAdmin: &svc,
	}

	t.Run("bad request", func(t *testing.T) {

		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzU0MDYzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.9_omXwoS8NX4yF-U3us6HIr39sZnfIeVufwpauWkKQU"
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/jamkerja/", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)
		echoContext := e.NewContext(r, w)
		adminController.GetByIdController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("http ok", func(t *testing.T) {

		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzU0MDYzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.9_omXwoS8NX4yF-U3us6HIr39sZnfIeVufwpauWkKQU"
		e := echo.New()
		r := httptest.NewRequest("GET", "http://localhost:8888/admin/jamkerja/1", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)
		echoContext := e.NewContext(r, w)
		adminController.GetByIdController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})
}

func TestDeleteJamKerjaController(t *testing.T) {
	svc := mocks.SvcAdminMock{}
	svc.On("DeleteJamKerja", mock.Anything).Return(errors.New("new")).Once()
	svc.On("DeleteJamKerja", mock.Anything).Return(nil).Once()

	adminController := a.JamKerjaEchoController{
		SvcAdmin: &svc,
	}

	t.Run("bad request", func(t *testing.T) {

		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzU0MDYzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.9_omXwoS8NX4yF-U3us6HIr39sZnfIeVufwpauWkKQU"
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/admin/jamkerja/", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)
		echoContext := e.NewContext(r, w)
		adminController.DeleteController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("http ok", func(t *testing.T) {

		TokenInValid := constants.TOKEN_JWT_TYPE + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImFyZGhpa2F5b3ZpeWFudG9AZ21haWwuY29tIiwiZXhwIjoxNjUyNzU0MDYzLCJuYW1hIjoiQXJkaGlrYSBZb3ZpeWFudG8sIFMua29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.9_omXwoS8NX4yF-U3us6HIr39sZnfIeVufwpauWkKQU"
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/admin/jamkerja/1", nil)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", TokenInValid)

		echoContext := e.NewContext(r, w)
		adminController.DeleteController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})
}
