package services

import (
	"errors"
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	util "mtb_web/pkg/utils"
)

type ProductService struct {
	ProductRepo *repo.ProductRepo
}

func NewProductService() *ProductService {
	return &ProductService{
		ProductRepo: repo.NewProductRepo(),
	}
}

func (s *ProductService) GetAll() ([]models.Product, int, error) {
	return s.ProductRepo.GetAll()
}

func (s *ProductService) GetByID(id int) (models.Product, int, error) {
	return s.ProductRepo.GetById(id)
}

func (s *ProductService) Create(product *models.Product) (int, error) {
	var product_images []models.ProductImage

	// Duyệt qua từng ảnh, decode base64 và upload lên S3
	for _, img := range product.ProductImages {
		// Giải mã Base64 thành file ảnh (byte array)
		imgData, err := util.DecodeBase64Image(img.ImageBase64)
		if err != nil {
			return 20052, fmt.Errorf("failed to decode base64 image: %v", err)
		}
		imgHover, err := util.DecodeBase64Image(img.ImageHorver)
		if err != nil {
			return 20052, fmt.Errorf("failed to decode base64 image: %v", err)
		}

		// Upload ảnh lên AWS S3 và nhận URL
		imgURL, err := util.UpLoadFile(imgData)
		if err != nil {
			return 20052, fmt.Errorf("failed to upload image to S3: %v", err)
		}
		imgHoverURL, err := util.UpLoadFile(imgHover)
		if err != nil {
			return 20052, fmt.Errorf("failed to upload image to S3: %v", err)
		}

		// Cập nhật ImageURL trong ProductImage
		img.ImageURL = imgURL
		img.ImageHorver = imgHoverURL

		// Thêm ProductImage đã cập nhật vào danh sách
		product_images = append(product_images, img)
	}

	// Gán lại danh sách ProductImages vào product
	product.ProductImages = product_images

	// Tạo sản phẩm mới trong database
	return s.ProductRepo.Create(product)
}

func (s *ProductService) Update(product *models.Product) (int, error) {
	if product.ID == 0 {
		return 20052, errors.New("missing product ID")
	}

	var product_images []models.ProductImage

	// Duyệt qua từng ảnh, decode base64 và upload lên S3
	for _, img := range product.ProductImages {
		// Giải mã Base64 thành file ảnh (byte array)
		imgData, err := util.DecodeBase64Image(img.ImageBase64)
		if err != nil {
			return 20052, fmt.Errorf("failed to decode base64 image: %v", err)
		}
		imgHover, err := util.DecodeBase64Image(img.ImageHorver)
		if err != nil {
			return 20052, fmt.Errorf("failed to decode base64 image: %v", err)
		}

		// Upload ảnh lên AWS S3 và nhận URL
		imgURL, err := util.UpLoadFile(imgData)
		if err != nil {
			return 20052, fmt.Errorf("failed to upload image to S3: %v", err)
		}
		imgHoverURL, err := util.UpLoadFile(imgHover)
		if err != nil {
			return 20052, fmt.Errorf("failed to upload image to S3: %v", err)
		}

		// Cập nhật ImageURL trong ProductImage
		img.ImageURL = imgURL
		img.ImageHorver = imgHoverURL

		// Thêm ProductImage đã cập nhật vào danh sách
		product_images = append(product_images, img)
	}

	// Gán lại danh sách ProductImages vào product
	product.ProductImages = product_images

	return s.ProductRepo.Update(product)
}

func (s *ProductService) DeleteById(id int) (int, error) {
	return s.ProductRepo.DeleteById(id)
}

func (s *ProductService) DeleteByIDs(ids []int) (int, error) {
	return s.ProductRepo.DeleteByIDs(ids)
}

func (s *ProductService) DeleteAll() (int, error) {
	return s.ProductRepo.DeleteAll()
}

func (s *ProductService) GetByCategoryID(id int) ([]models.Product, int, error) {
	return s.ProductRepo.GetByCategoryID(id)
}

func (s *ProductService) GetByCaregoryIDAndColorID(categoryID int, colorID int) ([]models.Product, int, error) {
	return s.ProductRepo.GetByCaregoryIDAndColorID(categoryID, colorID)
}

func (s *ProductService) GetByCaregoryIDAndStyleID(categoryID int, styleID int) ([]models.Product, int, error) {
	return s.ProductRepo.GetByCaregoryIDAndStyleID(categoryID, styleID)
}

func (s *ProductService) GetByStyleID(styleID int) ([]models.Product, int, error) {
	return s.ProductRepo.GetByStyleID(styleID)
}

func (s *ProductService) GetByCaregoryIDAndMaterialID(categoryID int, materialID int) ([]models.Product, int, error) {
	return s.ProductRepo.GetByCaregoryIDAndMaterialID(categoryID, materialID)
}

func (s *ProductService) GetByCaregoryIDAndSizeIDs(categoryID int, sizeIDs []int) ([]models.Product, int, error) {
	return s.ProductRepo.GetByCaregoryIDAndSizeIDs(categoryID, sizeIDs)
}
