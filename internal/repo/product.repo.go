package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type ProductRepo struct{}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (r *ProductRepo) GetAll() ([]models.Product, int, error) {
	var products []models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		Find(&products).Error; err != nil {
		return nil, 20058, err
	}
	return products, 20001, nil
}

func (r *ProductRepo) GetById(id int) (models.Product, int, error) {
	var product models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		First(&product, id).Error; err != nil {
		return models.Product{}, 20057, err
	}
	return product, 20001, nil
}

func (r *ProductRepo) Create(product *models.Product) (int, error) {
	if err := global.Mdb.Create(product).Error; err != nil {
		return 20052, err
	}
	return 20001, nil
}

func (r *ProductRepo) Update(product *models.Product) (int, error) {
	var existing models.Product
	if err := global.Mdb.
		Preload("ProductImages").
		Preload("Size").
		Preload("Color").
		Preload("Categories").
		Where("id = ?", product.ID).
		First(&existing).Error; err != nil {
		return 20053, err
	}

	product.CreatedAt = existing.CreatedAt

	// ====================
	// ✅ Lọc dữ liệu rỗng
	// ====================
	validImages := []models.ProductImage{}
	for _, img := range product.ProductImages {
		if img.ImageURL != "" {
			// Đảm bảo gán đúng ProductID
			img.ProductID = &product.ID
			validImages = append(validImages, img)
		}
	}
	product.ProductImages = validImages

	validSizes := []models.ProductSize{}
	for _, s := range product.Size {
		if s.Size != "" {
			validSizes = append(validSizes, s)
		}
	}
	product.Size = validSizes

	validColors := []models.ProductColor{}
	for _, c := range product.Color {
		if c.Color != "" {
			validColors = append(validColors, c)
		}
	}
	product.Color = validColors

	// ===============================
	// ✅ XÓA TOÀN BỘ ẢNH CŨ trước khi cập nhật
	// ===============================
	if err := global.Mdb.Where("product_id = ?", product.ID).
		Delete(&models.ProductImage{}).Error; err != nil {
		return 20054, err
	}

	// ====================================
	// ✅ Save toàn bộ sản phẩm và liên kết
	// ====================================
	if err := global.Mdb.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(product).Error; err != nil {
		return 20053, err
	}

	return 20001, nil
}

func (r *ProductRepo) DeleteById(id int) (int, error) {
	var product models.Product
	if err := global.Mdb.First(&product, id).Error; err != nil {
		return 20054, err
	}
	if err := global.Mdb.Delete(&product).Error; err != nil {
		return 20054, err
	}
	return 20001, nil
}

func (r *ProductRepo) DeleteByIDs(ids []int) (int, error) {
	var products []models.Product
	if err := global.Mdb.Where("id IN ?", ids).Find(&products).Error; err != nil {
		return 2019, err
	}
	if err := global.Mdb.Where("id IN ?", ids).Delete(&products).Error; err != nil {
		return 20055, err
	}
	return 20001, nil
}

func (r *ProductRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Product{}).Error; err != nil {
		return 20056, err
	}
	return 20001, nil
}

func (r *ProductRepo) GetByCategoryID(id int) ([]models.Product, int, error) {
	var products []models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		Where("category_id = ?", id).
		Find(&products).Error; err != nil {
		return nil, 20057, err
	}
	return products, 20001, nil
}

func (r *ProductRepo) GetByCaregoryIDAndColorID(categoryID int, colorID int) ([]models.Product, int, error) {
	var products []models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		Where("category_id = ?", categoryID).
		Where("color_id = ?", colorID).
		Find(&products).Error; err != nil {
		return nil, 20057, err
	}
	return products, 20001, nil
}

func (r *ProductRepo) GetByCaregoryIDAndStyleID(categoryID int, styleID int) ([]models.Product, int, error) {
	var products []models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		Where("category_id = ?", categoryID).
		Where("product_style_id = ?", styleID).
		Find(&products).Error; err != nil {
		return nil, 20057, err
	}
	return products, 20001, nil
}

func (r *ProductRepo) GetByCaregoryIDAndMaterialID(categoryID int, materialID int) ([]models.Product, int, error) {
	var products []models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		Where("category_id = ?", categoryID).
		Where("product_material_id = ?", materialID).
		Find(&products).Error; err != nil {
		return nil, 20057, err
	}
	return products, 20001, nil
}

func (r *ProductRepo) GetByCaregoryIDAndSizeIDs(categoryID int, sizeIDs []int) ([]models.Product, int, error) {
	var products []models.Product
	if err := global.Mdb.Preload("Size").
		Preload("Color").
		Preload("Categories").
		Preload("ProductStyle").
		Preload("ProductMaterial").
		Preload("ProductImages").
		Where("category_id = ?", categoryID).
		Where("size_id IN ?", sizeIDs).
		Find(&products).Error; err != nil {
		return nil, 20057, err
	}
	return products, 20001, nil
}
