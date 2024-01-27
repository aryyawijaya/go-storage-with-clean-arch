package filedelivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
	"github.com/gin-gonic/gin"
)

func (fh *FileHandler) Create(ctx *gin.Context) {
	var req fileusecase.CreateDto
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	fileReq, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	// Get file extension
	fileExt := fh.getFileExt(fileReq)

	// Get file content
	fileContent, err := fh.getFileContent(fileReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fh.wrapper.ErrResp(err))
		return
	}

	// New file dto
	fileDto := fh.newFileDto(fileContent, fileExt)

	file, err := fh.fileUseCase.Create(ctx, &req, fileDto)
	if err != nil {
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}

	ctx.JSON(http.StatusCreated, file)
}

func (fh *FileHandler) GetPublicByName(ctx *gin.Context) {
	var req fileusecase.GetPublicByNameDto
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	filepath, err := fh.fileUseCase.GetPublicByName(ctx, &req)
	if err != nil {
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}
	fullpath := fmt.Sprintf("%s%s", fileusecase.PrefixPath, filepath)

	ctx.File(fullpath)
}

func (fh *FileHandler) GetPrivateByName(ctx *gin.Context) {
	var req fileusecase.GetPrivateByNameDto
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	filepath, err := fh.fileUseCase.GetPrivateByName(ctx, &req)
	if err != nil {
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}
	fullpath := fmt.Sprintf("%s%s", fileusecase.PrefixPath, filepath)

	ctx.File(fullpath)
}

func (fh *FileHandler) CreateBulk(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	// Get fields from form-data
	filesEntity := form.Value["filesEntity"]
	files := form.File["files"]

	// Length filesEntity & files should be same
	if len(filesEntity) != len(files) {
		err = entity.ErrNotSameLenSlice
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}

	// Lenght > 0
	if len(filesEntity) == 0 {
		err = entity.ErrEmptyPayload
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}

	// Validate each fileEntity & create slice of CreateDto
	var dto []*fileusecase.CreateDto
	for _, fe := range filesEntity {
		// Parse json string
		var createDto *fileusecase.CreateDto
		err = json.Unmarshal([]byte(fe), &createDto)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
			return
		}

		// Validate struct with tag
		err = fh.validate.Struct(createDto)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
			return
		}

		dto = append(dto, createDto)
	}

	// Create slice of fileDto
	var filesDto []*fileusecase.FileDto
	for _, file := range files {
		// Get file extension
		fileExt := fh.getFileExt(file)

		// Get file content
		fileContent, err := fh.getFileContent(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fh.wrapper.ErrResp(err))
			return
		}

		// New file dto
		fileDto := fh.newFileDto(fileContent, fileExt)

		filesDto = append(filesDto, fileDto)
	}

	err = fh.fileUseCase.CreateBulk(ctx, dto, filesDto)
	if err != nil {
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (fh *FileHandler) Update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindUri(&req.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	updateDto := &fileusecase.UpdateDto{
		Name:   req.Uri.Name,
		Access: req.Body.Access,
	}
	file, err := fh.fileUseCase.Update(ctx, updateDto)
	if err != nil {
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}

	ctx.JSON(http.StatusOK, file)
}

func (fh *FileHandler) DeleteBulk(ctx *gin.Context) {
	var req fileusecase.DeleteBulkDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, fh.wrapper.ErrResp(err))
		return
	}

	err := fh.fileUseCase.DeleteBulk(ctx, &req)
	if err != nil {
		ctx.JSON(fh.wrapper.GetStatusCode(err), fh.wrapper.ErrResp(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
