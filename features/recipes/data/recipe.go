package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/images"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/features/users"
	"alta-cookit-be/utils/consts"
	"errors"
	"fmt"
	"strings"

	_imageModel "alta-cookit-be/features/images/models"
	_recipeModel "alta-cookit-be/features/recipes/models"

	"gorm.io/gorm"
)

type RecipeData struct {
	db        *gorm.DB
	userData  users.UserData_
	imageData images.ImageData_
}

func New(db *gorm.DB, userData users.UserData_, imageData images.ImageData_) recipes.RecipeData_ {
	return &RecipeData{
		db:        db,
		userData:  userData,
		imageData: imageData,
	}
}

func (d *RecipeData) SelectRecipesByUserId(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	gorms := []_recipeModel.Recipe{}

	tx := d.db.Preload("Recipe").Where("user_id = ?", entity.UserID).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []recipes.RecipeEntity{}
	userGorm := d.userData.SelectUserById(entity.UserID)
	for index, gorm := range gorms {
		entities = append(entities, *ConvertToEntity(&gorm, userGorm))
		if entities[index].Recipe != nil {
			subUserGorm := d.userData.SelectUserById(entities[index].Recipe.UserID)
			entities[index].Recipe.UserName = subUserGorm.Username
			entities[index].Recipe.UserRole = subUserGorm.Role
			entities[index].Recipe.ProfilePicture = subUserGorm.ProfilePicture
		}
	}

	for index, entity := range entities {
		d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(lk.recipe_id) as total_like").Joins("left join likes lk on lk.recipe_id = recipes.id").Where("lk.recipe_id = ?", entity.ID).Find(&entities[index].TotalLike)
		if entities[index].Recipe != nil {
			d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(lk.recipe_id) as total_like").Joins("left join likes lk on lk.recipe_id = recipes.id").Where("lk.recipe_id = ?", entities[index].Recipe.ID).Find(&entities[index].Recipe.TotalLike)
		}
	}

	for index, entity := range entities {
		d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(cs.recipe_id) as total_comment").Joins("left join comments cs on cs.recipe_id = recipes.id").Where("cs.recipe_id = ?", entity.ID).Find(&entities[index].TotalComment)
		if entities[index].Recipe != nil {
			d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(cs.recipe_id) as total_comment").Joins("left join comments cs on cs.recipe_id = recipes.id").Where("cs.recipe_id = ?", entities[index].Recipe.ID).Find(&entities[index].Recipe.TotalComment)
		}
		fmt.Println(entity.TotalComment)
	}
	return &entities, nil
}

func (d *RecipeData) InsertRecipe(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	gorm := *ConvertToGorm(entity)

	txTransaction := d.db.Begin()
	if txTransaction.Error != nil {
		txTransaction.Rollback()
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	tx := txTransaction.Omit("recipe_id").Create(&gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		return nil, tx.Error
	}

	for index, file := range entity.Image {
		urlImage, err := storage.GetStorageClient().UploadFile(file, entity.ImageName[index])
		if err != nil {
			return nil, err
		}
		gorm.Images = append(gorm.Images, _imageModel.Image{
			UrlImage: urlImage,
		})
	}

	tx = txTransaction.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	userGorm := d.userData.SelectUserById(entity.UserID)
	return ConvertToEntity(&gorm, userGorm), nil
}

func (d *RecipeData) UpdateRecipeById(entity *recipes.RecipeEntity) error {
	gorm := ConvertToGorm(entity)

	tx := d.db.Omit("recipe_id").Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Updates(&gorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *RecipeData) DeleteRecipeById(entity *recipes.RecipeEntity) error {
	gorm, imageGorms := ConvertToGorm(entity), d.imageData.SelectImagesByRecipeId(entity.ID)

	for _, imageGorm := range *imageGorms {
		err := storage.GetStorageClient().DeleteFile(imageGorm.UrlImage)
		if err != nil {
			return err
		}

		tx := d.db.Delete(&imageGorm)
		if tx.Error != nil {
			return tx.Error
		}
	}

	tx := d.db.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Delete(&gorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
