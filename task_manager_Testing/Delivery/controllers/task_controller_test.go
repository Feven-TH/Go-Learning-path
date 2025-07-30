package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager_Testing/Delivery/controllers"
	model "task_manager_Testing/Domain/entities"
	"task_manager_Testing/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	mockUC     *mocks.TaskUsecase
	controller *controllers.TaskController
	router     *gin.Engine
}

func (s *TaskControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockUC = new(mocks.TaskUsecase)
	s.controller = controllers.NewTaskController(s.mockUC)

	router := gin.Default()
	router.GET("/tasks", s.controller.GetAllTasks)
	router.GET("/tasks/:id", s.controller.GetTaskById)
	router.POST("/tasks", s.controller.AddTask)
	router.PUT("/tasks/:id", s.controller.UpdateTask)
	router.DELETE("/tasks/:id", s.controller.DeleteTask)

	s.router = router
}

func (s *TaskControllerSuite) TestGetAllTasks_Success() {
	expected := []model.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Completed: false},
		{ID: primitive.NewObjectID(), Title: "Task 2", Completed: true},
	}

	s.mockUC.On("IGetAllTasks", mock.Anything).Return(expected, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)

	var actual []model.Task
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	s.NoError(err)
	s.Len(actual, len(expected))
	s.Equal(expected[0].Title, actual[0].Title)

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestGetAllTasks_Failure() {
	s.mockUC.On("IGetAllTasks", mock.Anything).Return(nil, errors.New("some error"))

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusInternalServerError, w.Code)
	s.Contains(w.Body.String(), "Failed to retrieve tasks")

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestGetTaskById_Success() {
	id := primitive.NewObjectID()
	expected := model.Task{ID: id, Title: "Task 1", Completed: false}

	s.mockUC.On("IGetTaskById", mock.Anything, id.Hex()).Return(&expected, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+id.Hex(), nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)

	var actual model.Task
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	s.NoError(err)
	s.Equal(expected.Title, actual.Title)
	s.Equal(expected.ID, actual.ID)

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestGetTaskById_NotFound() {
	id := primitive.NewObjectID()
	s.mockUC.On("IGetTaskById", mock.Anything, id.Hex()).Return(nil, errors.New("not found"))

	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+id.Hex(), nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusNotFound, w.Code)
	s.Contains(w.Body.String(), "Task not found")

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestAddTask_Success() {
	newTask := model.Task{
		Title:     "New Task",
		Completed: false,
	}
	body, _ := json.Marshal(newTask)

	s.mockUC.On("IAddTask", mock.Anything, newTask).Return(&newTask, nil)


	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusCreated, w.Code)
	
	var actual model.Task
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	s.NoError(err)
	s.Equal(newTask.Title, actual.Title)
	s.Equal(newTask.Completed, actual.Completed)

	s.mockUC.AssertExpectations(s.T())
}


func (s *TaskControllerSuite) TestAddTask_BadRequest() {
	body := []byte(`{invalid json}`)

	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Contains(w.Body.String(), "Invalid input")
}

func (s *TaskControllerSuite) TestAddTask_Failure() {
	task := model.Task{
		Title:     "Test Failure",
		Completed: false,
	}
	body, _ := json.Marshal(task)

	s.mockUC.On("IAddTask", mock.Anything, mock.Anything).Return((*model.Task)(nil), errors.New("insert failed"))


	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusInternalServerError, w.Code)
	s.Contains(w.Body.String(), "Failed to add task")

	s.mockUC.AssertExpectations(s.T())
}


func (s *TaskControllerSuite) TestUpdateTask_Success() {
	id := primitive.NewObjectID().Hex()
	input := model.Task{Title: "Updated", Completed: true}
	output := input
	output.ID, _ = primitive.ObjectIDFromHex(id)

	s.mockUC.On("IUpdateTask", mock.Anything, id, input).Return(&output, nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Updated")

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestUpdateTask_BadRequest() {
	id := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+id, bytes.NewBuffer([]byte(`{invalid json}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Contains(w.Body.String(), "Invalid Input")
}

func (s *TaskControllerSuite) TestUpdateTask_Failure() {
	id := primitive.NewObjectID().Hex()
	task := model.Task{Title: "New", Completed: false}

	s.mockUC.On("IUpdateTask", mock.Anything, id, task).Return(nil, errors.New("update error"))

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusInternalServerError, w.Code)
	s.Contains(w.Body.String(), "update error")

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestDeleteTask_Success() {
	id := primitive.NewObjectID().Hex()
	s.mockUC.On("IDeleteTask", mock.Anything, id).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+id, nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Task deleted successfully")

	s.mockUC.AssertExpectations(s.T())
}

func (s *TaskControllerSuite) TestDeleteTask_NotFound() {
	id := primitive.NewObjectID().Hex()
	s.mockUC.On("IDeleteTask", mock.Anything, id).Return(errors.New("not found"))

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+id, nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusNotFound, w.Code)
	s.Contains(w.Body.String(), "Task not found")

	s.mockUC.AssertExpectations(s.T())
}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
