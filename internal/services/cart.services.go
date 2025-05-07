package services

import (
	"mtb_web/global"
	"mtb_web/internal/dto"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type CartService struct {
	CartRepo repo.CartRepo
}

func NewCartService() *CartService {
	return &CartService{
		CartRepo: *repo.NewCartRepo(),
	}
}

func (s *CartService) CreateRedis(sessionId string, cart *dto.RedisCartItem) error {
	return s.CartRepo.CreateRedis(sessionId, cart)
}

func (s *CartService) SyncCartFromRedisToDb(sessionId string, userId uint) error {
	items, err := s.CartRepo.GetRedis(sessionId)
	if err != nil {
		return err
	}

	// 1. Chuẩn bị danh sách CartItem
	var cartItems []models.CartItem
	var totalPrice float64

	for _, item := range *items {
		totalPrice += float64(item.Product.DiscountPrice) * float64(item.Quantity)

		cartItems = append(cartItems, models.CartItem{
			ProductID: &item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	// 2. Tạo Cart và gắn CartItem
	cart := &models.Cart{
		UserID:   int(userId),
		Price:    totalPrice,
		CartItem: cartItems,
	}

	if err := global.Mdb.Create(cart).Error; err != nil {
		return err
	}

	// 3. Redis cleanup
	return s.CartRepo.DeleteRedis(sessionId)
}

func (s *CartService) GetRedis(sessionId string) (*[]dto.RedisCartItem, error) {
	return s.CartRepo.GetRedis(sessionId)
}

func (s *CartService) SyncCartFromDbToRedis(sessionId string, userId uint) error {
	cart, err := s.CartRepo.GetCartByUserId(int(userId))
	if err != nil {
		return err
	}

	var redisItems []dto.RedisCartItem
	for _, item := range cart.CartItem {
		redisItems = append(redisItems, dto.RedisCartItem{
			ProductID: *item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Product:   *item.Product, // Nếu bạn có gán sẵn
		})
	}

	for _, item := range redisItems {
		if err := s.CartRepo.CreateRedis(sessionId, &item); err != nil {
			return err
		}
	}

	return nil
}
