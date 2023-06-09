// Code generated by MockGen. DO NOT EDIT.
// Source: ninth-learn/repository (interfaces: BookRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	model "ninth-learn/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookRepo is a mock of BookRepo interface.
type MockBookRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepoMockRecorder
}

// MockBookRepoMockRecorder is the mock recorder for MockBookRepo.
type MockBookRepoMockRecorder struct {
	mock *MockBookRepo
}

// NewMockBookRepo creates a new mock instance.
func NewMockBookRepo(ctrl *gomock.Controller) *MockBookRepo {
	mock := &MockBookRepo{ctrl: ctrl}
	mock.recorder = &MockBookRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepo) EXPECT() *MockBookRepoMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookRepo) CreateBook(arg0 model.Book) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookRepoMockRecorder) CreateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookRepo)(nil).CreateBook), arg0)
}

// DeleteBook mocks base method.
func (m *MockBookRepo) DeleteBook(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookRepoMockRecorder) DeleteBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookRepo)(nil).DeleteBook), arg0)
}

// GetBookById mocks base method.
func (m *MockBookRepo) GetBookById(arg0 int64) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", arg0)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockBookRepoMockRecorder) GetBookById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockBookRepo)(nil).GetBookById), arg0)
}

// GetBooks mocks base method.
func (m *MockBookRepo) GetBooks() ([]model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooks")
	ret0, _ := ret[0].([]model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks.
func (mr *MockBookRepoMockRecorder) GetBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockBookRepo)(nil).GetBooks))
}

// UpdateBook mocks base method.
func (m *MockBookRepo) UpdateBook(arg0 model.Book) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookRepoMockRecorder) UpdateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookRepo)(nil).UpdateBook), arg0)
}
