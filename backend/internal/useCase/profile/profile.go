package profile

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/profile"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type CreateProfileRequest struct {
	UserID string `json:"userId"`
	Image  []byte `json:"image"`
}

type UseCaseProfile struct {
	db *profile.PGProfileDB
}

func NewUseCaseProfile(db *profile.PGProfileDB) *UseCaseProfile {
	return &UseCaseProfile{db: db}
}

func (uc *UseCaseProfile) CreateProfile(ctx *fiber.Ctx, req CreateProfileRequest) (*profileEntity.ResponseProfile, error) {
	filePath := "./static/uploads/profile/images/defaultImage.jpg"
	directoryPath := fmt.Sprintf("./static/uploads/profile/images")
	form, err := ctx.MultipartForm()
	if err != nil {
		logger.Log.Debug(
			"error func CreateProfile, method MultipartForm by path internal/useCase/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	imageFiles := form.File["image"]
	imagesFilePath := make([]string, 0, len(imageFiles))
	imagesCatalog := make([]*profileEntity.ImageProfile, 0, len(imagesFilePath))
	for _, file := range imageFiles {
		filePath = fmt.Sprintf("%s/%s", directoryPath, file.Filename)
		if err := ctx.SaveFile(file, filePath); err != nil {
			logger.Log.Debug(
				"error func CreateProfile, method SaveFile by path internal/useCase/profile/profile.go",
				zap.Error(err))
			return nil, err
		}
		image := profileEntity.ImageProfile{
			UUID:      uuid.New(),
			Name:      file.Filename,
			Url:       filePath,
			Size:      file.Size,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
			IsEnabled: true,
		}
		imagesFilePath = append(imagesFilePath, filePath)
		imagesCatalog = append(imagesCatalog, &image)
	}
	profileRequest := &profileEntity.Profile{
		UUID:      uuid.New(),
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Images:    imagesCatalog,
	}
	newProfile, err := uc.db.Create(ctx, profileRequest)
	if err != nil {
		logger.Log.Debug("error func CreateProfile, method Create by path internal/useCase/profile/profile.go", zap.Error(err))
		return nil, err
	}
	for _, i := range profileRequest.Images {
		image := &profileEntity.ImageProfile{
			ProfileID: profileRequest.ID,
			UUID:      i.UUID,
			Name:      i.Name,
			Url:       i.Url,
			Size:      i.Size,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
			IsDeleted: i.IsDeleted,
			IsEnabled: i.IsEnabled,
		}
		_, err := uc.db.AddImage(ctx, image)
		if err != nil {
			logger.Log.Debug(
				"error func CreateProfile, method AddImage by path internal/useCase/profile/profile.go",
				zap.Error(err))
			return nil, err
		}
	}
	foundedProfile, err := uc.db.FindByUUID(ctx, newProfile.UUID)
	if err != nil {
		logger.Log.Debug("error func CreateProfile, method FindByUUID by path internal/useCase/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	responseProfile := profileEntity.ResponseProfile{
		UUID:      foundedProfile.UUID,
		UserID:    foundedProfile.UserID,
		CreatedAt: foundedProfile.CreatedAt,
		UpdatedAt: foundedProfile.UpdatedAt,
		Images:    foundedProfile.Images,
	}
	return &responseProfile, nil
}

func (uc *UseCaseProfile) GetProfileByUUID(ctx *fiber.Ctx) (*profileEntity.ResponseProfile, error) {
	params := ctx.Params("uuid")
	paramsStr, err := uuid.Parse(params)
	if err != nil {
		logger.Log.Debug("error func GetProfileByUUID, method Parse by path internal/useCase/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	foundedProfile, err := uc.db.FindByUUID(ctx, paramsStr)
	if err != nil {
		logger.Log.Debug(
			"error func GetProfileByUUID, method FindByUUID by path internal/useCase/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	responseProfile := profileEntity.ResponseProfile{
		UUID:      foundedProfile.UUID,
		UserID:    foundedProfile.UserID,
		CreatedAt: foundedProfile.CreatedAt,
		UpdatedAt: foundedProfile.UpdatedAt,
		Images:    foundedProfile.Images,
	}
	return &responseProfile, nil
}
