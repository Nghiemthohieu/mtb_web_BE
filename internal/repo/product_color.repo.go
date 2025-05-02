package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"
)

type ProductColorRepo struct {
}

func NewProductColorRepo() *ProductColorRepo {
	return &ProductColorRepo{}
}

func (r *ProductColorRepo) Create(productColor *models.ProductColor) (int, error) {
	if err := global.Mdb.Create(productColor).Error; err != nil {
		return 20032, err
	}
	return 20001, nil
}

func (r *ProductColorRepo) Update(productColor *models.ProductColor) (int, error) {
	var existingProductColor models.ProductColor
	if err := global.Mdb.Where("id = ?", productColor.ID).First(&existingProductColor).Error; err != nil {
		return 20033, err
	}
	productColor.CreatedAt = existingProductColor.CreatedAt
	if err := global.Mdb.Save(productColor).Error; err != nil {
		return 20033, err
	}
	return 20001, nil
}
func (r *ProductColorRepo) GetById(id int) (*models.ProductColor, int, error) {
	var productColor models.ProductColor
	if err := global.Mdb.Where("id = ?", id).First(&productColor).Error; err != nil {
		return nil, 20037, err
	}
	return &productColor, 20001, nil
}
func (r *ProductColorRepo) GetAll() ([]*models.ProductColor, int, error) {
	var productColors []*models.ProductColor
	if err := global.Mdb.Find(&productColors).Error; err != nil {
		return nil, 20038, err
	}
	return productColors, 20001, nil
}

func (r *ProductColorRepo) DeleteById(id int) (int, error) {
	var productColor models.ProductColor
	if err := global.Mdb.Where("id = ?", id).Delete(&productColor).Error; err != nil {
		return 20034, err
	}
	return 20001, nil
}
func (r *ProductColorRepo) DeleteByIDs(ids []int) (int, error) {
	var productColor models.ProductColor
	if err := global.Mdb.Where("id IN (?)", ids).Delete(&productColor).Error; err != nil {
		return 20035, err
	}
	return 20001, nil
}
func (r *ProductColorRepo) DeleteAll() (int, error) {

	if err := global.Mdb.Delete(&models.ProductColor{}).Error; err != nil {
		return 20036, err
	}
	return 20001, nil
}
