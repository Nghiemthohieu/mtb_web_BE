package dto

import "mtb_web/internal/models"

type RedisCartItem struct {
	ProductID uint           `json:"product_id"`
	Product   models.Product `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Quantity  int            `json:"quantity"`
	Price     float64        `json:"price"`
}
