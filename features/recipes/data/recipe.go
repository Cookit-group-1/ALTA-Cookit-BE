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
	_userModel "alta-cookit-be/features/users/data"

	"gorm.io/gorm"
)

type RecipeData struct {
	db        *gorm.DB
	userData  users.UserData
	imageData images.ImageData_
}

func New(db *gorm.DB, userData users.UserData, imageData images.ImageData_) recipes.RecipeData_ {
	return &RecipeData{
		db:        db,
		userData:  userData,
		imageData: imageData,
	}
}

func (d *RecipeData) ActionValidator(recipeId, userId uint) bool {
	tempGorm := _recipeModel.Recipe{}

	d.db.Where("id = ? AND user_id = ?", recipeId, userId).Find(&tempGorm)

	return tempGorm.ID != 0
}

func (d *RecipeData) SelectRecipeByIngredientId(ingredientId uint) *_recipeModel.Recipe {
	tempGorm := _recipeModel.Recipe{}

	subQuery := d.db.Table("ingredients").Where("id = ?", ingredientId).Select("recipe_id")
	d.db.Where("id IN (?)", subQuery).Find(&tempGorm)

	return &tempGorm
}

func (d *RecipeData) SelectRecipes(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	gorms := []_recipeModel.Recipe{}

	qString := ""
	for key, val := range entity.ExtractedQueryParams {
		if qString != "" {
			qString += " AND "
		}
		if key == "name" {
			qString += fmt.Sprintf("%s LIKE %s%s%s ", key, "'%", val, "%'")
		} else {
			qString += fmt.Sprintf("%s = '%s'", key, val)
		}
	}

	tx := d.db.Preload("Recipe").Preload("Recipe.Images").Preload("Images").Preload("Ingredients").Where(qString).Limit(entity.DataLimit).Offset(entity.DataOffset).Order("created_at desc").Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []recipes.RecipeEntity{}
	for index, gorm := range gorms {
		userEntity := d.userData.SelectUserById(users.Core{ID: gorm.UserID})
		userModel := _userModel.CoreToModel(*userEntity)
		entities = append(entities, *ConvertToEntity(&gorm, &userModel))
		if entities[index].Recipe != nil {
			userEntity := d.userData.SelectUserById(users.Core{ID: entities[index].Recipe.UserID})
			entities[index].Recipe.UserName = userEntity.Username
			entities[index].Recipe.UserRole = userEntity.Role
			entities[index].Recipe.ProfilePicture = userEntity.ProfilePicture
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
	}
	return &entities, nil
}

func (d *RecipeData) InsertRecipe(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	gorm := *ConvertToGorm(entity)

	for index, file := range entity.Image {
		urlImage, err := storage.GetStorageClient().UploadFile(file, entity.ImageName[index])
		if err != nil {
			return nil, err
		}
		gorm.Images = append(gorm.Images, _imageModel.Image{
			UrlImage: urlImage,
		})
	}

	tx := d.db
	if *gorm.RecipeID != 0 {
		tx = d.db.Create(&gorm)
	} else {
		tx = d.db.Omit("recipe_id").Create(&gorm)
	}
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return nil, errors.New(consts.RECIPE_InvalidRecipe)
		}
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		return nil, tx.Error
	}

	tx = d.db.Preload("Recipe").Preload("Recipe.Steps").Preload("Recipe.Ingredients").Preload("Recipe.Images").Preload("Recipe.Ingredients.IngredientDetails").Preload("Steps").Preload("Ingredients").Preload("Images").Preload("Ingredients.IngredientDetails").Find(&gorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	userEntity := d.userData.SelectUserById(users.Core{ID: entity.UserID})
	userModel := _userModel.CoreToModel(*userEntity)
	entity = ConvertToEntity(&gorm, &userModel)
	if entity.Recipe != nil {
		subUserEntity := d.userData.SelectUserById(users.Core{ID: entity.Recipe.UserID})
		entity.Recipe.UserName = subUserEntity.Username
		entity.Recipe.UserRole = subUserEntity.Role
		entity.Recipe.ProfilePicture = subUserEntity.ProfilePicture
	}
	return entity, nil
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

	tx := d.db.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Unscoped().Delete(&gorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *RecipeData) SelectRecipesTimeline(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	gorms := []_recipeModel.Recipe{}

	subQuery := d.db.Table("followers").Distinct("to_user_id").Where("from_user_id = ?", entity.UserID).Select("to_user_id")
	tx := d.db.Preload("Recipe").Where("user_id IN (?)", subQuery).Order("created_at desc").Limit(entity.DataLimit).Offset(entity.DataOffset).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []recipes.RecipeEntity{}
	for index, gorm := range gorms {
		userEntity := d.userData.SelectUserById(users.Core{ID: gorm.UserID})
		userModel := _userModel.CoreToModel(*userEntity)
		entities = append(entities, *ConvertToEntity(&gorm, &userModel))
		if entities[index].Recipe != nil {
			userEntity := d.userData.SelectUserById(users.Core{ID: entities[index].Recipe.UserID})
			entities[index].Recipe.UserName = userEntity.Username
			entities[index].Recipe.UserRole = userEntity.Role
			entities[index].Recipe.ProfilePicture = userEntity.ProfilePicture
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
	}
	return &entities, nil
}

func (d *RecipeData) SelectRecipesTrending(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	gorms := []_recipeModel.Recipe{}

	tx := d.db.Preload("Recipe").Select("*, recipes.id as id, recipes.user_id as user_id, COUNT(DISTINCT lk.id) as total_like, COUNT(DISTINCT cs.id) as total_comment").Joins("left join likes lk on lk.recipe_id = recipes.id").Joins("left join comments cs on cs.recipe_id = recipes.id").Order("total_like desc, total_comment desc").Group("recipes.id").Where("recipes.type = 'Original' AND recipes.created_at >= DATE(NOW() - INTERVAL 7 DAY)").Order("recipes.created_at desc").Limit(entity.DataLimit).Offset(entity.DataOffset).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []recipes.RecipeEntity{}
	for index, gorm := range gorms {
		userEntity := d.userData.SelectUserById(users.Core{ID: gorm.UserID})
		userModel := _userModel.CoreToModel(*userEntity)
		entities = append(entities, *ConvertToEntity(&gorm, &userModel))
		if entities[index].Recipe != nil {
			userEntity := d.userData.SelectUserById(users.Core{ID: entities[index].Recipe.UserID})
			entities[index].Recipe.UserName = userEntity.Username
			entities[index].Recipe.UserRole = userEntity.Role
			entities[index].Recipe.ProfilePicture = userEntity.ProfilePicture
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
	}
	return &entities, nil
}

func (d *RecipeData) SelectRecipeDetailById(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	gorm := _recipeModel.Recipe{}

	tx := d.db.Preload("Recipe").Preload("Recipe.Steps").Preload("Recipe.Ingredients").Preload("Recipe.Images").Preload("Recipe.Ingredients.IngredientDetails").Preload("Steps").Preload("Ingredients").Preload("Images").Preload("Ingredients.IngredientDetails").Where("id = ?", entity.ID).First(&gorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	userEntity := d.userData.SelectUserById(users.Core{ID: gorm.UserID})
	userModel := _userModel.CoreToModel(*userEntity)
	entity = ConvertToEntity(&gorm, &userModel)
	if entity.Recipe != nil {
		subUserEntity := d.userData.SelectUserById(users.Core{ID: entity.Recipe.UserID})
		entity.Recipe.UserName = subUserEntity.Username
		entity.Recipe.UserRole = subUserEntity.Role
		entity.Recipe.ProfilePicture = subUserEntity.ProfilePicture
	}

	d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(lk.recipe_id) as total_like").Joins("left join likes lk on lk.recipe_id = recipes.id").Where("lk.recipe_id = ?", entity.ID).Find(&entity.TotalLike)
	if entity.Recipe != nil {
		d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(lk.recipe_id) as total_like").Joins("left join likes lk on lk.recipe_id = recipes.id").Where("lk.recipe_id = ?", entity.Recipe.ID).Find(&entity.Recipe.TotalLike)
	}

	d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(cs.recipe_id) as total_comment").Joins("left join comments cs on cs.recipe_id = recipes.id").Where("cs.recipe_id = ?", entity.ID).Find(&entity.TotalComment)
	if entity.Recipe != nil {
		d.db.Model(&_recipeModel.Recipe{}).Select("COUNT(cs.recipe_id) as total_comment").Joins("left join comments cs on cs.recipe_id = recipes.id").Where("cs.recipe_id = ?", entity.Recipe.ID).Find(&entity.Recipe.TotalComment)
	}

	return entity, nil
}
