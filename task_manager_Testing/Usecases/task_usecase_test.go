package usecase_test

import (
	"context"
	"errors" 
	model "task_manager_Testing/Domain/entities"
	usecase_interface "task_manager_Testing/Domain/usecase_interfaces"
	usecase "task_manager_Testing/Usecases"
	"task_manager_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCaseSuite struct {
	suite.Suite
	ctx 		context.Context
	mockRepo 	*mocks.TaskRepo
	uc			usecase_interface.TaskUsecase
}

func(s *TaskUseCaseSuite) SetupTest(){
	s.ctx = context.Background()
	s.mockRepo = new(mocks.TaskRepo)
	s.uc = usecase.NewTaskUseCase(s.mockRepo)
}

func(s *TaskUseCaseSuite) TestAddTask(){
	title := "Testing projects"
	task := model.Task{Title: title, Completed: false}
	s.mockRepo.On("IAddTask", s.ctx, task).Return(&task, nil)

	res,err := s.uc.IAddTask(s.ctx, task)

	s.NoError(err)
	s.Equal(title, res.Title)
	s.False(res.Completed)

	s.mockRepo.AssertExpectations(s.T())
}

func(s *TaskUseCaseSuite) TestGetAllTasks(){
	task1Title := "Starter Project"
	task2Title := "Revise task 7"
	mockTets := []model.Task{
		{Title : task1Title, Completed: false},
		{Title : task2Title, Completed: true},
	}
	s.mockRepo.On("IGetAllTasks", s.ctx).Return(mockTets, nil)
	res,err := s.uc.IGetAllTasks(s.ctx)

	s.NoError(err)
	s.Equal(len(mockTets), len(res))
	s.Equal(task1Title, res[0].Title)
	s.False(res[0].Completed)
	s.Equal(task2Title, res[1].Title)
	s.True(res[1].Completed)

	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUseCaseSuite) TestGetTaskById() {
	id := primitive.NewObjectID()
	expectedTitle := "From MongoDB"
	expected := &model.Task{ID: id, Title: expectedTitle, Completed: false}

	s.mockRepo.On("IGetTaskById", s.ctx, id.Hex()).Return(expected, nil)

	res, err := s.uc.IGetTaskById(s.ctx, id.Hex())

	s.NoError(err)
	s.Equal(expectedTitle, res.Title)
	s.Equal(id, res.ID)
	s.False(res.Completed)

	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUseCaseSuite) TestUpdateTask(){
	id := primitive.NewObjectID()
	newTitle := "updated task"
	updated := &model.Task{
		Title: newTitle, Completed: true,
	}
	s.mockRepo.On("IUpdateTask", s.ctx, id.Hex(), *updated).Return(updated,nil)

	res,err := s.uc.IUpdateTask(s.ctx, id.Hex(),*updated)

	s.NoError(err)
	s.Equal(newTitle, res.Title)
	s.True(res.Completed)

	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUseCaseSuite) TestDeleteTask(){
	id := primitive.NewObjectID()
	s.mockRepo.On("IDeleteTask", s.ctx, id.Hex()).Return(nil)
	err := s.uc.IDeleteTask(s.ctx, id.Hex())

	s.NoError(err)
	s.mockRepo.AssertExpectations(s.T())
}

//Negative Test Cases

func (s *TaskUseCaseSuite) TestGetTaskById_NotFound() {
	id := primitive.NewObjectID().Hex()
	s.mockRepo.On("IGetTaskById", s.ctx, id).Return(nil, errors.New("task not found"))

	res, err := s.uc.IGetTaskById(s.ctx, id)

	s.Error(err)
	s.Nil(res)
	s.EqualError(err, "task not found")

	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUseCaseSuite) TestAddTask_Error() {
	task := model.Task{Title: "Failed Add", Completed: false}
	s.mockRepo.On("IAddTask", s.ctx, task).Return(nil, errors.New("db error")) 

	res, err := s.uc.IAddTask(s.ctx, task)

	s.Error(err)
	s.Nil(res)
	s.EqualError(err, "db error")

	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUseCaseSuite) TestUpdateTask_Error() {
	id := primitive.NewObjectID().Hex()
	task := model.Task{Title: "bad update", Completed: true}
	s.mockRepo.On("IUpdateTask", s.ctx, id, task).Return(nil, errors.New("update failed")) 
	res, err := s.uc.IUpdateTask(s.ctx, id, task)

	s.Error(err)
	s.Nil(res)
	s.EqualError(err, "update failed")

	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUseCaseSuite) TestDeleteTask_Error() {
	id := primitive.NewObjectID().Hex()
	s.mockRepo.On("IDeleteTask", s.ctx, id).Return(errors.New("delete error")) 
	err := s.uc.IDeleteTask(s.ctx, id)

	s.Error(err)
	s.EqualError(err, "delete error")

	s.mockRepo.AssertExpectations(s.T())
}

func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseSuite))
}
