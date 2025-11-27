package service

import (
	"errors"
	"session-9/model"
	"session-9/repository"
	"session-9/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newTestService() (*StudentService, *repository.MockStudentRepository) {
	mokeRepo := new(repository.MockStudentRepository)
	service := NewStudentService(mokeRepo)
	return service, mokeRepo
}

func TestStudentService_GetAll_Success(t *testing.T) {
	expected := []model.Student{
		{ID: 1, Name: "Bayu", Age: 25},
	}

	svc, repo := newTestService()

	repo.On("GetAll").Return(expected, nil).Once()

	result, err := svc.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}

func TestStudentService_GetAll_Error(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, errors.New("no file")).Once()

	// menjalankan svc getall
	result, err := svc.GetAll()

	assert.Error(t, err)
	assert.EqualError(t, err, "no file")
	assert.Empty(t, result)
	repo.AssertExpectations(t)
}

func TestStudentService_GetByID_Error_GetAll(t *testing.T) {
	svc, repo := newTestService()

	// memaksa error di getAll
	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	result, err := svc.GetByID(1)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)

	repo.AssertExpectations(t)
}

func TestStudentService_GetByID_Found(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	svc, repo := newTestService()
	repo.On("GetAll").Return(initial, nil).Once()

	result, err := svc.GetByID(2)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "Siti", result.Name)
	assert.Equal(t, 22, result.Age)

	repo.AssertExpectations(t)
}

func TestStudentService_GetByID_NotFound(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	svc, repo := newTestService()
	repo.On("GetAll").Return(initial, nil).Once()

	result, err := svc.GetByID(999)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, utils.ErrNotFound, err)

	repo.AssertExpectations(t)
}

func TestStudentService_GetByID_FileError(t *testing.T) {
	svc, repo := newTestService()
	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	_, err := svc.GetByID(1)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != utils.ErrFile {
		t.Fatalf("expected error file, got %v", err)
	}
}

func TestStudentService_Create_Success(t *testing.T) {
	existing := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	input := model.Student{Name: "Bayu", Age: 20}

	expected := model.Student{
		ID:   3,
		Name: "Bayu",
		Age:  20,
	}

	svc, repo := newTestService()
	// return existing dan nil
	repo.On("GetAll").Return(existing, nil).Once()

	// return nil karena berhasil tambah
	repo.On("SaveAll", append(existing, expected)).Return(nil).Once()

	// jalankan service asli
	result, err := svc.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	repo.AssertExpectations(t)

}

func TestStudentService_Create_Error_GetAll(t *testing.T) {
	svc, repo := newTestService()

	input := model.Student{Name: "Bayu", Age: 25}

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	result, err := svc.Create(input)

	assert.Error(t, err)
	assert.EqualError(t, err, utils.ErrFile.Error())
	assert.Equal(t, model.Student{}, result)

	repo.AssertExpectations(t)

}

func TestStudentService_Create_Error_SaveAll(t *testing.T) {
	svc, repo := newTestService()

	existing := []model.Student{{Name: "Andi", Age: 10}}
	input := model.Student{Name: "Bayu", Age: 25}

	repo.On("GetAll").Return(existing, nil).Once()

	repo.On("SaveAll", mock.Anything).Return(utils.ErrFile)

	result, err := svc.Create(input)

	assert.Error(t, err)
	assert.EqualError(t, err, utils.ErrFile.Error())
	assert.Equal(t, model.Student{}, result)

	repo.AssertExpectations(t)
}

func TestStudentService_Update_Success(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	updatedInput := model.Student{Name: "Siti Baru", Age: 23}
	expected := model.Student{ID: 2, Name: "Siti Baru", Age: 23}

	// mengembalikan initial error nil
	repo.On("GetAll").Return(initial, nil).Once()

	// mengembalikan data lama dengan expected update
	repo.On("SaveAll", []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		expected,
	}).Return(nil).Once()

	// menjalankan service
	result, err := svc.Update(2, updatedInput)

	assert.NoError(t, err)            //tidak ada error
	assert.Equal(t, expected, result) // expected sama dengan result

	repo.AssertExpectations(t)
}

func TestStudentService_Update_Error_IdNotFound(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()

	// id 999 tidak ada
	result, err := svc.Update(999, model.Student{Name: "X", Age: 30})

	assert.Error(t, err)
	assert.Equal(t, utils.ErrNotFound, err)  //apakah error sama
	assert.Equal(t, model.Student{}, result) //apakah result objek kosong

	repo.AssertExpectations(t)
}

func TestStudentService_Update_Error_GetAll(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	result, err := svc.Update(1, model.Student{Name: "Bayu", Age: 25})

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
	assert.Equal(t, model.Student{}, result)
}

func TestStudentService_Update_Error_SaveAll(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	input := model.Student{Name: "Andi Updated", Age: 22}
	expectedStudents := []model.Student{
		{ID: 1, Name: "Andi Updated", Age: 22},
	}

	repo.On("GetAll").Return(initial, nil).Once() //getAll success
	repo.On("SaveAll", expectedStudents).Return(utils.ErrFile).Once()

	result, err := svc.Update(1, input)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
	assert.Equal(t, model.Student{}, result)

	repo.AssertExpectations(t)
}

func TestStudentService_Delete_Success(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	expectedAfterDelete := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()       //GetAllSuccess
	repo.On("SaveAll", expectedAfterDelete).Return(nil) //SaveAll Success

	err := svc.Delete(2)

	assert.NoError(t, err)

	repo.AssertExpectations(t)

}

func TestStudentService_Delete_Error_IdNotFound(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()

	err := svc.Delete(999) //id 999 tidak ada di initial

	assert.Error(t, err)                    //apa ada error
	assert.Equal(t, utils.ErrNotFound, err) //apa error notFoundError

	repo.AssertExpectations(t)
}

func TestStudentService_Delete_Error_GetAll(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	err := svc.Delete(3)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)

	repo.AssertExpectations(t)
}

func TestStudentService_Delete_Error_SaveAll(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	expectedAfterDelete := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()                        //GetAll Success
	repo.On("SaveAll", expectedAfterDelete).Return(utils.ErrFile).Once() // SaveAll Gagal

	err := svc.Delete(2)

	assert.Error(t, err)                // apakah error tidak nil
	assert.Equal(t, utils.ErrFile, err) //apakah errorFile

	repo.AssertExpectations(t)
}
