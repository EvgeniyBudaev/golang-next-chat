package profile

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/profile"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type CreateProfileRequest struct {
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	IsEnabled string `json:"isEnabled"`
	Image     []byte `json:"image"`
}

type GetProfileRequest struct {
	UUID uuid.UUID `json:"uuid"`
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
	isEnabled, err := strconv.ParseBool(req.IsEnabled)
	if err != nil {
		logger.Log.Debug(
			"error func CreateProfile, method ParseBool by path internal/useCase/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	profileRequest := &profileEntity.Profile{
		UUID:      uuid.New(),
		UserID:    req.UserID,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
		IsEnabled: isEnabled,
		Images:    imagesCatalog,
	}
	newProfile, err := uc.db.Create(ctx, profileRequest)
	if err != nil {
		logger.Log.Debug("error func CreateProfile, method Create by path internal/useCase/profile/profile.go",
			zap.Error(err))
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
		Username:  foundedProfile.Username,
		Firstname: foundedProfile.Firstname,
		Lastname:  foundedProfile.Lastname,
		Email:     foundedProfile.Email,
		CreatedAt: foundedProfile.CreatedAt,
		UpdatedAt: foundedProfile.UpdatedAt,
		IsDeleted: foundedProfile.IsDeleted,
		IsEnabled: foundedProfile.IsEnabled,
		Images:    foundedProfile.Images,
	}
	return &responseProfile, nil
}

func (uc *UseCaseProfile) GetProfileByUUID(ctx *fiber.Ctx, req GetProfileRequest) (*profileEntity.ResponseProfile, error) {
	foundedProfile, err := uc.db.FindByUUID(ctx, req.UUID)
	if err != nil {
		logger.Log.Debug(
			"error func GetProfileByUUID, method FindByUUID by path internal/useCase/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	responseProfile := profileEntity.ResponseProfile{
		UUID:      foundedProfile.UUID,
		UserID:    foundedProfile.UserID,
		Username:  foundedProfile.Username,
		Firstname: foundedProfile.Firstname,
		Lastname:  foundedProfile.Lastname,
		Email:     foundedProfile.Email,
		CreatedAt: foundedProfile.CreatedAt,
		UpdatedAt: foundedProfile.UpdatedAt,
		IsDeleted: foundedProfile.IsDeleted,
		IsEnabled: foundedProfile.IsEnabled,
		Images:    foundedProfile.Images,
	}
	return &responseProfile, nil
}
