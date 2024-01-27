package fileusecase_test

import (
	"context"
	"fmt"
	"testing"

	mockdb "github.com/aryyawijaya/go-storage-with-clean-arch/db/mock"
	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
	filevalidator "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/validator"
	utilrandom "github.com/aryyawijaya/go-storage-with-clean-arch/util/random"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	// create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	// create use case
	fValidator := filevalidator.NewFileValidator(repo)
	fUseCase := fileusecase.NewFileUseCase(repo, fValidator)

	t.Run("success", func(t *testing.T) {
		mockDto := &fileusecase.CreateDto{
			Name:      utilrandom.RandomString(5),
			Access:    utilrandom.RandomEnum[entity.Access](entity.AllAccessValues()),
			Tribe:     utilrandom.RandomPath(),
			Service:   utilrandom.RandomPath(),
			Module:    utilrandom.RandomPath(),
			SubFolder: utilrandom.RandomPath(),
		}
		mockFileDto := &fileusecase.FileDto{
			Content: utilrandom.RandomByteSlices(),
			Ext:     utilrandom.RandomFileExt(),
		}

		filepath := fmt.Sprintf("%s%s%s%s", mockDto.Tribe, mockDto.Service, mockDto.Module, mockDto.SubFolder)

		mockFile := &entity.File{
			Name:   mockDto.Name,
			Access: mockDto.Access,
			Path:   filepath,
			Ext:    mockFileDto.Ext,
		}

		mockCreatedFile := &entity.File{
			ID:        utilrandom.RandomInt(1, 100),
			Name:      mockDto.Name,
			Access:    mockDto.Access,
			Path:      filepath,
			Ext:       mockFileDto.Ext,
			CreatedAt: utilrandom.RandomTime(),
			UpdatedAt: utilrandom.RandomTime(),
		}

		// stubs
		repo.EXPECT().
			GetByName(gomock.Any(), gomock.Eq(mockDto.Name)).
			Times(1).
			Return(nil, entity.ErrNotFound)
		repo.EXPECT().
			CreateWithTx(gomock.Any(), gomock.Eq(mockFile), gomock.Any()).
			Times(1).
			Return(mockCreatedFile, nil)

		resp, err := fUseCase.Create(context.TODO(), mockDto, mockFileDto)
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		require.Equal(t, mockCreatedFile, resp)
	})

	t.Run("duplicate name", func(t *testing.T) {
		mockDto := &fileusecase.CreateDto{
			Name:      utilrandom.RandomString(5),
			Access:    utilrandom.RandomEnum[entity.Access](entity.AllAccessValues()),
			Tribe:     utilrandom.RandomPath(),
			Service:   utilrandom.RandomPath(),
			Module:    utilrandom.RandomPath(),
			SubFolder: utilrandom.RandomPath(),
		}
		mockFileDto := &fileusecase.FileDto{
			Content: utilrandom.RandomByteSlices(),
			Ext:     utilrandom.RandomFileExt(),
		}

		// stubs
		repo.EXPECT().
			GetByName(gomock.Any(), gomock.Eq(mockDto.Name)).
			Times(1).
			Return(&entity.File{ID: -1}, nil)
		repo.EXPECT().
			CreateWithTx(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)

		resp, err := fUseCase.Create(context.TODO(), mockDto, mockFileDto)
		require.Error(t, err)
		require.Empty(t, resp)

		require.EqualError(t, err, entity.ErrUnique.Error())
	})
}

func TestGetPublicByName(t *testing.T) {
	// create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	// create use case
	fValidator := filevalidator.NewFileValidator(repo)
	fUseCase := fileusecase.NewFileUseCase(repo, fValidator)

	t.Run("success", func(t *testing.T) {
		mockDto := &fileusecase.GetPublicByNameDto{
			Name: utilrandom.RandomString(5),
		}
		mockCurrentFile := &entity.File{
			ID:     utilrandom.RandomInt(1, 100),
			Name:   utilrandom.RandomString(5),
			Access: entity.AccessPUBLIC,
			Path:   utilrandom.RandomPath(),
			Ext:    utilrandom.RandomFileExt(),
		}

		// stubs
		repo.EXPECT().
			GetByName(gomock.Any(), gomock.Eq(mockDto.Name)).
			Times(1).
			Return(mockCurrentFile, nil)

		resp, err := fUseCase.GetPublicByName(context.TODO(), mockDto)
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expectedResp := fmt.Sprintf("%s%s%s", mockCurrentFile.Path, mockCurrentFile.Name, mockCurrentFile.Ext)

		require.Equal(t, expectedResp, resp)
	})

	t.Run("private file", func(t *testing.T) {
		mockDto := &fileusecase.GetPublicByNameDto{
			Name: utilrandom.RandomString(5),
		}
		mockCurrentFile := &entity.File{
			ID:     utilrandom.RandomInt(1, 100),
			Name:   utilrandom.RandomString(5),
			Access: entity.AccessPRIVATE,
			Path:   utilrandom.RandomPath(),
			Ext:    utilrandom.RandomFileExt(),
		}

		// stubs
		repo.EXPECT().
			GetByName(gomock.Any(), gomock.Eq(mockDto.Name)).
			Times(1).
			Return(mockCurrentFile, nil)

		resp, err := fUseCase.GetPublicByName(context.TODO(), mockDto)
		require.Error(t, err)
		require.Empty(t, resp)

		require.EqualError(t, err, entity.ErrForbiddenAccess.Error())
	})
}

func TestGetPrivateByName(t *testing.T) {
	// create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	// create use case
	fValidator := filevalidator.NewFileValidator(repo)
	fUseCase := fileusecase.NewFileUseCase(repo, fValidator)

	t.Run("success", func(t *testing.T) {
		mockDto := &fileusecase.GetPrivateByNameDto{
			Name: utilrandom.RandomString(5),
		}
		mockCurrentFile := &entity.File{
			ID:     utilrandom.RandomInt(1, 100),
			Name:   utilrandom.RandomString(5),
			Access: entity.AccessPRIVATE,
			Path:   utilrandom.RandomPath(),
			Ext:    utilrandom.RandomFileExt(),
		}

		// stubs
		repo.EXPECT().
			GetByName(gomock.Any(), gomock.Eq(mockDto.Name)).
			Times(1).
			Return(mockCurrentFile, nil)

		resp, err := fUseCase.GetPrivateByName(context.TODO(), mockDto)
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expectedResp := fmt.Sprintf("%s%s%s", mockCurrentFile.Path, mockCurrentFile.Name, mockCurrentFile.Ext)

		require.Equal(t, expectedResp, resp)
	})

	t.Run("public file", func(t *testing.T) {
		mockDto := &fileusecase.GetPrivateByNameDto{
			Name: utilrandom.RandomString(5),
		}
		mockCurrentFile := &entity.File{
			ID:     utilrandom.RandomInt(1, 100),
			Name:   utilrandom.RandomString(5),
			Access: entity.AccessPUBLIC,
			Path:   utilrandom.RandomPath(),
			Ext:    utilrandom.RandomFileExt(),
		}

		// stubs
		repo.EXPECT().
			GetByName(gomock.Any(), gomock.Eq(mockDto.Name)).
			Times(1).
			Return(mockCurrentFile, nil)

		resp, err := fUseCase.GetPrivateByName(context.TODO(), mockDto)
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expectedResp := fmt.Sprintf("%s%s%s", mockCurrentFile.Path, mockCurrentFile.Name, mockCurrentFile.Ext)

		require.Equal(t, expectedResp, resp)
	})
}

func TestCreateBulk(t *testing.T) {
	// create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	// create use case
	fValidator := filevalidator.NewFileValidator(repo)
	fUseCase := fileusecase.NewFileUseCase(repo, fValidator)

	t.Run("success", func(t *testing.T) {
		var mockDto []*fileusecase.CreateDto
		var mockFileDto []*fileusecase.FileDto
		var names []string
		n := 2

		for i := 0; i < n; i++ {
			currFile := &fileusecase.CreateDto{
				Name:      utilrandom.RandomString(5),
				Access:    utilrandom.RandomEnum[entity.Access](entity.AllAccessValues()),
				Tribe:     utilrandom.RandomPath(),
				Service:   utilrandom.RandomPath(),
				Module:    utilrandom.RandomPath(),
				SubFolder: utilrandom.RandomPath(),
			}
			currFileDto := &fileusecase.FileDto{
				Content: utilrandom.RandomByteSlices(),
				Ext:     utilrandom.RandomFileExt(),
			}

			mockDto = append(mockDto, currFile)
			mockFileDto = append(mockFileDto, currFileDto)
			names = append(names, currFile.Name)
		}

		var mockCreateFiles []*entity.File

		for i := 0; i < len(mockDto); i++ {
			currDto := mockDto[i]
			currFileDto := mockFileDto[i]

			filepath := fmt.Sprintf("%s%s%s%s", currDto.Tribe, currDto.Service, currDto.Module, currDto.SubFolder)

			createFile := &entity.File{
				Name:   currDto.Name,
				Access: currDto.Access,
				Path:   filepath,
				Ext:    currFileDto.Ext,
			}

			mockCreateFiles = append(mockCreateFiles, createFile)
		}

		// stubs
		repo.EXPECT().
			GetByNames(gomock.Any(), gomock.Eq(names)).
			Times(1).
			Return([]*entity.File{}, nil)
		repo.EXPECT().
			CreateBulkWithTx(gomock.Any(), gomock.Eq(mockCreateFiles), gomock.Any()).
			Times(1).
			Return(nil)

		err := fUseCase.CreateBulk(context.TODO(), mockDto, mockFileDto)
		require.NoError(t, err)
	})

	t.Run("duplicate name", func(t *testing.T) {
		var mockDto []*fileusecase.CreateDto
		var mockFileDto []*fileusecase.FileDto
		var names []string
		n := 2

		for i := 0; i < n; i++ {
			currFile := &fileusecase.CreateDto{
				Name:      utilrandom.RandomString(5),
				Access:    utilrandom.RandomEnum[entity.Access](entity.AllAccessValues()),
				Tribe:     utilrandom.RandomPath(),
				Service:   utilrandom.RandomPath(),
				Module:    utilrandom.RandomPath(),
				SubFolder: utilrandom.RandomPath(),
			}
			currFileDto := &fileusecase.FileDto{
				Content: utilrandom.RandomByteSlices(),
				Ext:     utilrandom.RandomFileExt(),
			}

			mockDto = append(mockDto, currFile)
			mockFileDto = append(mockFileDto, currFileDto)
			names = append(names, currFile.Name)
		}

		// stubs
		repo.EXPECT().
			GetByNames(gomock.Any(), gomock.Eq(names)).
			Times(1).
			Return([]*entity.File{{ID: -1}}, nil)
		repo.EXPECT().
			CreateWithTx(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)

		err := fUseCase.CreateBulk(context.TODO(), mockDto, mockFileDto)
		require.Error(t, err)

		require.EqualError(t, err, entity.ErrUnique.Error())
	})
}

func TestUpdate(t *testing.T) {
	// create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	// create use case
	fValidator := filevalidator.NewFileValidator(repo)
	fUseCase := fileusecase.NewFileUseCase(repo, fValidator)

	t.Run("success", func(t *testing.T) {
		currFile := utilrandom.RandomFile()

		dto := &fileusecase.UpdateDto{
			Name:   currFile.Name,
			Access: utilrandom.RandomEnum[entity.Access](entity.AllAccessValues()),
		}

		expectedResult := currFile
		expectedResult.Access = dto.Access

		// stubs
		repo.EXPECT().
			Update(gomock.Any(), gomock.Eq(&entity.File{Name: currFile.Name, Access: dto.Access})).
			Times(1).
			Return(expectedResult, nil)

		result, err := fUseCase.Update(context.TODO(), dto)
		require.NoError(t, err)
		require.NotEmpty(t, result)

		require.Equal(t, expectedResult, result)
	})
}

func TestDeleteBulk(t *testing.T) {
	// create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	// create use case
	fValidator := filevalidator.NewFileValidator(repo)
	fUseCase := fileusecase.NewFileUseCase(repo, fValidator)

	t.Run("success", func(t *testing.T) {
		var names []string
		n := 2
		for i := 0; i < n; i++ {
			createdFile := utilrandom.RandomFile()
			names = append(names, createdFile.Name)
		}
		dto := &fileusecase.DeleteBulkDto{
			Names: names,
		}

		// stubs
		repo.EXPECT().
			DeleteBulkWithTx(gomock.Any(), gomock.Eq(dto.Names)).
			Times(1).
			Return(nil)

		err := fUseCase.DeleteBulk(context.TODO(), dto)
		require.NoError(t, err)
	})
}
