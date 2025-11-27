package service

import (
	"session-9/model"
	"session-9/repository"
	"session-9/utils"
	"testing"
)

// func newTestService(initial []model.Student) (*StudentService, *repository.MockStudentRepository) {
// 	repo := &repository.MockStudentRepository{Students: initial}
// 	svc := NewStudentService(repo)
// 	return svc, repo
// }

// func TestStudentService_Create(t *testing.T) {
// 	svc, repo := newTestService([]model.Student{})

// 	created, err := svc.Create(model.Student{
// 		Name: "Budi",
// 		Age:  20,
// 	})
// 	if err != nil {
// 		t.Fatalf("Create returned error: %v", err)
// 	}

// 	if created.ID != 1 {
// 		t.Errorf("expected ID 1, got %d", created.ID)
// 	}
// 	if created.Name != "Budi" {
// 		t.Errorf("expected Name Budi, got %s", created.Name)
// 	}

// 	if len(repo.Students) != 1 {
// 		t.Fatalf("expected repo to have 1 student, got %d", len(repo.Students))
// 	}
// }

// func TestStudentService_GetByID_Found(t *testing.T) {
// 	initial := []model.Student{
// 		{ID: 1, Name: "Andi", Age: 21},
// 		{ID: 2, Name: "Siti", Age: 22},
// 	}
// 	svc, _ := newTestService(initial)

// 	st, err := svc.GetByID(2)
// 	if err != nil {
// 		t.Fatalf("GetByID returned error: %v", err)
// 	}

// 	if st.Name != "Siti" {
// 		t.Errorf("expected Name Siti, got %s", st.Name)
// 	}
// }

// func TestStudentService_GetByID_NotFound(t *testing.T) {
// 	initial := []model.Student{
// 		{ID: 1, Name: "Andi", Age: 21},
// 		{ID: 2, Name: "Siti", Age: 22},
// 	}
// 	svc, _ := newTestService(initial)

// 	_, err := svc.GetByID(999)
// 	if err == nil {
// 		t.Fatalf("expected error, got nil")
// 	}

// 	if err != utils.ErrNotFound {
// 		t.Fatalf("expected ErrNotFound, got %v", err)
// 	}
// }

// func TestStudentService_GetByID_fileError(t *testing.T) {
// 	svc, _ := newTestService([]model.Student{})

// 	_, err := svc.GetByID(1)
// 	if err == nil {
// 		t.Fatalf("expected error, got nil")
// 	}

// 	if err != utils.ErrFile {
// 		t.Fatalf("expected error file, got %v", err)
// 	}
// }

func newTestService() (*StudentService, *repository.MockStudentRepository) {
	mokeRepo := new(repository.MockStudentRepository)
	service := NewStudentService(mokeRepo)
	return service, mokeRepo
}

// func TestStudent_Create(t *testing.T) {
// 	service, repo := newTestService([]model.Student{})

// 	created, err := service.Create(model.Student{
// 		Name: "Rudi",
// 		Age:  20,
// 	})

// 	if err != nil {
// 		t.Fatalf("Created returned error: %v", err)
// 	}

// 	if created.ID != 1 {
// 		t.Errorf("expected ID 1, got %d", created.ID)
// 	}

// 	if created.Name != "Rudi" {
// 		t.Errorf("expected Name Budi, got %s", created.Name)
// 	}

// 	if len(repo.Students) != 1 {
// 		t.Fatalf("expected repo to have 1 student, got %d", len(repo.Students))
// 	}
// }

func TestStudentService_GetByID_Found(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}
	svc, repo := newTestService()
	repo.On("GetAll").Return(initial, nil).Once()

	st, err := svc.GetByID(2)
	if err != nil {
		t.Fatalf("GetByID returned error: %v", err)
	}

	if st.Name != "Siti" {
		t.Errorf("expected Name Siti, got %s", st.Name)
	}
}

func TestStudentService_GetByID_NotFound(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}
	svc, repo := newTestService()
	repo.On("GetAll").Return(initial, nil).Once()

	_, err := svc.GetByID(999)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != utils.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
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
