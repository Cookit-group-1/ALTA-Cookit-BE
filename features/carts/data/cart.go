package data

import (
	"alta-cookit-be/features/carts"
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	_cartModel "alta-cookit-be/features/carts/models"

	"gorm.io/gorm"
)

type CartData struct {
	db             *gorm.DB
	recipeData     recipes.RecipeData_
	ingredientData ingredients.IngredientData_
}

func New(db *gorm.DB, recipeData recipes.RecipeData_, ingredientData ingredients.IngredientData_) carts.CartData_ {
	return &CartData{
		db:             db,
		recipeData:     recipeData,
		ingredientData: ingredientData,
	}
}

func (d *CartData) ActionValidator(id, userId uint) bool {
	tempGorm := _cartModel.Cart{}
	d.db.Model(&tempGorm).Where("id = ? AND user_id = ?", id, userId).Find(&tempGorm)

	return tempGorm.ID != 0
}

func (d *CartData) SelectCartsByUserId(entity *carts.CartEntity) (*[]carts.CartEntity, error) {
	gorms := []_cartModel.Cart{}

	tx := d.db.Where("user_id = ?", entity.UserID).Limit(entity.DataLimit).Offset(entity.DataOffset).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []carts.CartEntity{}
	for _, gorm := range gorms {
		recipeGorm := d.recipeData.SelectRecipeByIngredientId(gorm.IngredientID)
		ingredientGorm := d.ingredientData.SelectIngredientById(gorm.IngredientID)
		entities = append(entities, *ConvertToEntity(&gorm, recipeGorm, ingredientGorm))
	}

	return &entities, nil
}

func (d *CartData) InsertCart(entity *carts.CartEntity) (*carts.CartEntity, error) {
	gorm := ConvertToGorm(entity)
	tx := d.db.Create(gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "ingredient_id") {
			return nil, errors.New(consts.INGREDIENT_InvalidIngredient)
		}
		return nil, tx.Error
	}

	recipeGorm := d.recipeData.SelectRecipeByIngredientId(gorm.IngredientID)
	ingredientGorm := d.ingredientData.SelectIngredientById(gorm.IngredientID)
	return ConvertToEntity(gorm, recipeGorm, ingredientGorm), nil
}

func (d *CartData) UpdateCartById(entity *carts.CartEntity) error {
	tx := d.db.Where("id = ?", entity.ID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *CartData) DeleteCartById(entity *carts.CartEntity) error {
	tx := d.db.Unscoped().Where("id = ?", entity.ID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
