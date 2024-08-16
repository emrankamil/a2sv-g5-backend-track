// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// TaskController is an autogenerated mock type for the TaskController type
type TaskController struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *TaskController) Create(c *gin.Context) {
	_m.Called(c)
}

// Delete provides a mock function with given fields: c
func (_m *TaskController) Delete(c *gin.Context) {
	_m.Called(c)
}

// FetchAll provides a mock function with given fields: c
func (_m *TaskController) FetchAll(c *gin.Context) {
	_m.Called(c)
}

// FetchByTaskID provides a mock function with given fields: c
func (_m *TaskController) FetchByTaskID(c *gin.Context) {
	_m.Called(c)
}

// Update provides a mock function with given fields: c
func (_m *TaskController) Update(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewTaskController interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskController creates a new instance of TaskController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskController(t mockConstructorTestingTNewTaskController) *TaskController {
	mock := &TaskController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
