package tasktest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/karlosdaniel451/go-rest-api-template/api/middleware"
	"github.com/karlosdaniel451/go-rest-api-template/api/router"
	"github.com/karlosdaniel451/go-rest-api-template/cmd/setup"
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/stretchr/testify/assert"
)

var app *fiber.App

func setupApp() {
	setup.Setup()

	app = fiber.New(fiber.Config{
		AppName:           "TEST | Simple Go RESTful API with Fiber and GORM",
		EnablePrintRoutes: false,
	})
	middleware.Setup(app)
	router.Setup(app, &setup.TaskController)
}

func InsertPredefinedFakeTasks() {
	// predefinedFakeTasks := []model.Task{
	// 	model.Task{Name: "Fake task 123", Description: "Task for testing 123"},
	// 	model.Task{Name: "Fake task 124", Description: "Task for testing 123"},
	// 	model.Task{Name: "Fake task 125", Description: "Task for testing 123"},
	// }
}

func TestMain(m *testing.M) {
	setupApp()
	InsertPredefinedFakeTasks()
	code := m.Run()
	os.Exit(code)
}

func requestTaskRoute(
	t *testing.T,
	method,
	url string,
	taskInRequestBody *model.Task,
) *http.Response {

	var requestBody io.Reader
	if taskInRequestBody == nil {
		requestBody = nil
	} else {
		requestBody = bytes.NewReader(serializeTask(t, taskInRequestBody))
	}

	request := httptest.NewRequest(
		method,
		url,
		requestBody,
	)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	if err != nil {
		t.Errorf("error when processing request: %s", err)
	}

	return response
}

func serializeTask(t *testing.T, task *model.Task) []byte {
	encodedTask, err := json.Marshal(task)
	if err != nil {
		t.Errorf("error when serializing Task: %s", err)
		return nil
	}

	return encodedTask
}

func deserializeResponseBody(t *testing.T, response *http.Response) *model.Task {
	var task model.Task
	if err := json.NewDecoder(response.Body).Decode(&task); err != nil {
		t.Errorf("error when deserializing Task in response body: %s", err)
		return nil
	}

	defer response.Body.Close()

	return &task
}

func TestTaskRouteDeletionAfterCreation(t *testing.T) {
	fakeTask := model.Task{
		Name:        "Fake task",
		Description: "Task for testing",
	}

	creationResponse := requestTaskRoute(
		t, "POST", "http://localhost:8001/tasks/", &fakeTask)

	assert.Equal(t, http.StatusCreated, creationResponse.StatusCode)

	taskInCreationResponseBody := deserializeResponseBody(t, creationResponse)

	deletionResponse := requestTaskRoute(
		t, "DELETE",
		fmt.Sprintf("http://localhost:8001/tasks/%d",
			taskInCreationResponseBody.ID), nil,
	)

	assert.Equal(t, http.StatusNoContent, deletionResponse.StatusCode)
}

func TestTaskRouteCreation(t *testing.T) {
	fakeTask := model.Task{
		Name:        "Fake task",
		Description: "Task for testing",
	}

	testCases := []struct {
		desc                     string
		httpMethod               string
		url                      string
		requestBody              *model.Task
		expectedCode             int
		expectedTaskResponseBody *model.Task
	}{
		{
			desc:                     "get_not_found",
			httpMethod:               "GET",
			url:                      fmt.Sprint("http://localhost:8001/tasks/111"),
			requestBody:              nil,
			expectedCode:             http.StatusNotFound,
			expectedTaskResponseBody: nil,
		},
		{
			desc:                     "creation",
			httpMethod:               "POST",
			url:                      fmt.Sprint("http://localhost:8001/tasks/"),
			requestBody:              &fakeTask,
			expectedCode:             http.StatusCreated,
			expectedTaskResponseBody: &fakeTask,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			response := requestTaskRoute(
				t,
				testCase.httpMethod,
				testCase.url,
				testCase.requestBody,
			)

			assert.Equal(t, testCase.expectedCode, response.StatusCode)

			// If expected body is nil, there is no need to test the response body.
			if testCase.expectedTaskResponseBody == nil {
				return
			}

			taskInResponseBody := deserializeResponseBody(t, response)
			defer response.Body.Close()

			comparer := cmp.Comparer(func(task1, task2 model.Task) bool {
				return task1.Name == task2.Name && task1.Description == task2.Description
			})
			diff := cmp.Diff(
				testCase.expectedTaskResponseBody,
				taskInResponseBody,
				comparer,
			)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
