package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/go-supertest/supertest"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"

	"github.com/restuwahyu13/gin-rest-api/configs"
	"github.com/restuwahyu13/gin-rest-api/helpers"
)

var routerStudent = configs.NewRouterTesting()
var accessTokenStudent string
var studentId interface{}

func TestCreateStudentHandler(t *testing.T) {

	Convey("Login Success Get accessTokenStudent", func() {
		payload := gin.H{
			"email":    "eduardo.wehner@greenholtadams.net",
			"password": "qwerty12345",
		}

		test := supertest.NewSuperTest(routerStudent, t)

		test.Post("/api/v1/login")
		test.Send(payload)
		test.Set("Content-Type", "application/json")
		test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

			response := helpers.Parse(rr.Body.Bytes())
			t.Log(response)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, "Login successfully", response.Message)

			var token map[string]interface{}
			encoded := helpers.Strigify(response.Data)
			_ = json.Unmarshal(encoded, &token)

			accessTokenStudent = token["accessTokenStudent"].(string)
		})
	})

	Convey("Test Handler Create Student Group", t, func() {

		Convey("Create New Student Is Conflict", func() {

			payload := gin.H{
				"name": "bagus budiawan",
				"npm":  201543502292,
				"fak":  "mipa",
				"bid":  "tehnik informatika",
			}

			test := supertest.NewSuperTest(routerStudent, t)

			test.Post("/api/v1/student")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusConflict, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Npm student already exist", response.Message)
			})
		})

		Convey("Create New Student", func() {

			payload := gin.H{
				"name": faker.Internet().FreeEmail(),
				"npm":  faker.RandomInt(25, 50),
				"fak":  "mipa",
				"bid":  "tehnik informatika",
			}

			test := supertest.NewSuperTest(routerStudent, t)

			test.Post("/api/v1/student")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusCreated, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Create new student accessTokenStudentount successfully", response.Message)
			})
		})
	})
}

func TestResultsStudentHandler(t *testing.T) {

	Convey("Test Handler Results Student By ID Group", t, func() {

		Convey("Results All Student", func() {

			test := supertest.NewSuperTest(routerStudent, t)

			test.Get("/api/v1/student")
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				var objects []map[string]interface{}
				encoded := helpers.Strigify(response.Data)
				_ = json.Unmarshal(encoded, &objects)

				studentId = objects[0]["ID"]

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "Results Students data successfully", response.Message)
			})
		})
	})
}

func TestResultStudentHandler(t *testing.T) {

	Convey("Test Handler Result Student By ID Group", t, func() {

		Convey("Result Specific Student If StudentID Is Not Exist", func() {

			ID := "00f85d71-083b-4089-9d20-bb1054df4575"

			test := supertest.NewSuperTest(routerStudent, t)

			test.Get("/api/v1/student/" + ID)
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusNotFound, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "Student data is not exist or deleted", response.Message)
			})
		})

		Convey("Result Specific Student By ID", func() {

			ID := studentId

			test := supertest.NewSuperTest(routerStudent, t)

			test.Get("/api/v1/student/" + ID.(string))
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)

				mapping := make(map[string]interface{})
				encode := helpers.Strigify(response.Data)
			_:
				json.Unmarshal(encode, &mapping)

				assert.Equal(t, ID, mapping["ID"])
				assert.Equal(t, "Result Student data successfully", response.Message)
			})
		})
	})
}

func TestDeleteStudentHandler(t *testing.T) {

	Convey("Test Handler Delete Student By ID Group", t, func() {

		Convey("Delete Specific Student If StudentID Is Not Exist", func() {

			ID := "00f85d71-083b-4089-9d20-bb1054df4575"

			test := supertest.NewSuperTest(routerStudent, t)

			test.Delete("/api/v1/student/" + ID)
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusForbidden, rr.Code)
				assert.Equal(t, http.MethodDelete, req.Method)
				assert.Equal(t, "Student data is not exist or deleted", response.Message)
			})
		})

		Convey("Delete Specific Student By ID", func() {

			ID := studentId

			test := supertest.NewSuperTest(routerStudent, t)

			test.Delete("/api/v1/student/" + ID.(string))
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.Set("Authorization", "Bearer "+accessTokenStudent)
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := helpers.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodDelete, req.Method)
				assert.Equal(t, "Delete student data successfully", response.Message)
			})
		})
	})
}
