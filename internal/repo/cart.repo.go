package repo

import (
	"context"
	"mtb_web/global"
	"mtb_web/internal/dto"
	"mtb_web/internal/models"

	"encoding/json"
)

type CartRepo struct{}

func NewCartRepo() *CartRepo {
	return &CartRepo{}
}

func (r *CartRepo) Create(cart *models.Cart) (int, error) {
	if err := global.Mdb.Create(cart).Error; err != nil {
		return 20042, err
	}
	return 20001, nil
}

func (r *CartRepo) CreateRedis(sessionId string, cart *dto.RedisCartItem) error {
	key := "Cart:session:" + sessionId
	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return global.Rdb.Set(context.Background(), key, data, 7*24*3600).Err()
}

func (r *CartRepo) GetRedis(sessionId string) (*[]dto.RedisCartItem, error) {
	key := "Cart:session:" + sessionId
	data, err := global.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	var cart []dto.RedisCartItem
	err = json.Unmarshal([]byte(data), &cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepo) DeleteRedis(sessionId string) error {
	key := "Cart:session:" + sessionId
	return global.Rdb.Del(context.Background(), key).Err()
}

func (r *CartRepo) UpdateRedis(sessionId string, cart *dto.RedisCartItem) error {
	key := "Cart:session:" + sessionId
	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return global.Rdb.Set(context.Background(), key, data, 7*24*3600).Err()
}

func (r *CartRepo) GetCartByUserId(userId int) (*models.Cart, error) {
	var cart models.Cart
	if err := global.Mdb.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepo) GetAllRedis() (*[]dto.RedisCartItem, error) {
	keys, err := global.Rdb.Keys(context.Background(), "Cart:session:*").Result()
	if err != nil {
		return nil, err
	}
	var carts []dto.RedisCartItem
	for _, key := range keys {
		data, err := global.Rdb.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		var cart dto.RedisCartItem
		err = json.Unmarshal([]byte(data), &cart)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return &carts, nil
}

func (r *CartRepo) DeleteAllRedis() error {
	keys, err := global.Rdb.Keys(context.Background(), "Cart:session:*").Result()
	if err != nil {
		return err
	}
	return global.Rdb.Del(context.Background(), keys...).Err()
}

func (r *CartRepo) DeleteCartByUserId(userId int) (int, error) {
	if err := global.Mdb.Where("user_id = ?", userId).Delete(&models.Cart{}).Error; err != nil {
		return 20043, err
	}
	return 20001, nil
}

func (r *CartRepo) UpdateCartByUserId(userId int, cart *models.Cart) (int, error) {
	if err := global.Mdb.Where("user_id = ?", userId).Updates(cart).Error; err != nil {
		return 20044, err
	}
	return 20001, nil
}
