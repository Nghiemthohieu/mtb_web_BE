package dto

type CategorySizeRequest struct {
	CategoryId int   `json:"category_id"`
	SizeIds    []int `json:"size_ids"`
}
