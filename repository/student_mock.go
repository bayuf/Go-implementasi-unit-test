package repository

import (
	"session-9/model"

	"github.com/stretchr/testify/mock"
)

// type MockStudentRepository struct {
// 	Students []model.Student
// 	ErrGet   error
// 	ErrSave  error
// }

// func (mockStudentRepository *MockStudentRepository) GetAll() ([]model.Student, error) {
// 	return mockStudentRepository.Students, mockStudentRepository.ErrGet
// }

// func (mockStudentRepository *MockStudentRepository) SaveAll(students []model.Student) error {
// 	mockStudentRepository.Students = students
// 	return mockStudentRepository.ErrSave
// }

type MockStudentRepository struct {
	mock.Mock
}

func (mockStudentRepository *MockStudentRepository) GetAll() ([]model.Student, error) {
	args := mockStudentRepository.Called()
	return args.Get(0).([]model.Student), args.Error(1)
}

func (mockStudentRepository *MockStudentRepository) SaveAll(students []model.Student) error {
	args := mockStudentRepository.Called()
	return args.Error(0)
}
