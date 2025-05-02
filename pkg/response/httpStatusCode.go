package response

const (
	ErrCodeSuccess                            = 20001
	ErrCodeProductSizeNotFound                = 20011
	ErrCodeProductSizeCreateFailed            = 20012
	ErrCodeProductSizeUpdateFailed            = 20013
	ErrCodeProductSizeDeleteFailed            = 20014
	ErrCodeProductSizeDeleteByIDsFailed       = 20015
	ErrCodeProductSizeDeleteAllFailed         = 20016
	ErrCodeProductSizeGetByIDFailed           = 20017
	ErrCodeProductSizeGetAllFailed            = 20018
	ErrCodeProductSizeGetByIDsFailed          = 20019
	ErrCoderProductStyleNotFound              = 20021
	ErrCoderProductStyleCreateFailed          = 20022
	ErrCoderProductStyleUpdateFailed          = 20023
	ErrCoderProductStyleDeleteFailed          = 20024
	ErrCoderProductStyleDeleteByIDsFailed     = 20025
	ErrCoderProductStyleDeleteAllFailed       = 20026
	ErrCoderProductStyleGetByIDFailed         = 20027
	ErrCoderProductStyleGetAllFailed          = 20028
	ErrCoderProductStyleGetByIDsFailed        = 20029
	ErrCoderProductColorNotFound              = 20031
	ErrCoderProductColorCreateFailed          = 20032
	ErrCoderProductColorUpdateFailed          = 20033
	ErrCoderProductColorDeleteFailed          = 20034
	ErrCoderProductColorDeleteByIDsFailed     = 20035
	ErrCoderProductColorDeleteAllFailed       = 20036
	ErrCoderProductColorGetByIDFailed         = 20037
	ErrCoderProductColorGetAllFailed          = 20038
	ErrCoderProductMaterialsNotFount          = 20041
	ErrCoderProductMaterialsCreateFailed      = 20042
	ErrCoderProductMaterialsUpdateFailed      = 20043
	ErrCoderProductMaterialsDeleteFailed      = 20044
	ErrCoderProductMaterialsDeleteByIDsFailed = 20045
	ErrCoderProductMaterialsDeleteAllFailed   = 20046
	ErrCoderProductMaterialsGetByIDFailed     = 20047
	ErrCoderProductMaterialsGetAllFailed      = 20048
	ErrCoderProductNotFount                   = 20051
	ErrCoderProductCreateFailed               = 20052
	ErrCoderProductUpdateFailed               = 20053
	ErrCoderProductDeleteFailed               = 20054
	ErrCoderProductDeleteByIDsFailed          = 20055
	ErrCoderProductDeleteAllFailed            = 20056
	ErrCoderProductGetByIDFailed              = 20057
	ErrCoderProductGetAllFailed               = 20058
	ErrCoderProductStockNotFount              = 20059
	ErrCoderCategoryNotFount                  = 20061
	ErrCoderCategoryCreateFailed              = 20062
	ErrCoderCategoryUpdateFailed              = 20063
	ErrCoderCategoryDeleteFailed              = 20064
	ErrCoderCategoryDeleteByIDsFailed         = 20065
	ErrCoderCategoryDeleteAllFailed           = 20066
	ErrCoderCategoryGetByIDFailed             = 20067
	ErrCoderCategoryGetAllFailed              = 20068
	ErrCoderProductVariantNotFount            = 20071
	ErrCoderProductVariantCreateFailed        = 20072
	ErrCoderProductVariantUpdateFailed        = 20073
	ErrCoderProductVariantDeleteFailed        = 20074
	ErrCoderProductVariantDeleteByIDsFailed   = 20075
	ErrCoderProductVariantDeleteAllFailed     = 20076
	ErrCoderProductVariantGetByIDFailed       = 20077
	ErrCoderProductVariantGetAllFailed        = 20078
	ErrCoderOrderStatusNotFount               = 20081
	ErrCoderOrderStatusCreateFailed           = 20082
	ErrCoderOrderStatusUpdateFailed           = 20083
	ErrCoderOrderStatusDeleteFailed           = 20084
	ErrCoderOrderStatusDeleteByIDsFailed      = 20085
	ErrCoderOrderStatusDeleteAllFailed        = 20086
	ErrCoderOrderStatusGetByIDFailed          = 20087
	ErrCoderOrderStatusGetAllFailed           = 20088
	ErrCoderOrderNotFount                     = 20091
	ErrCoderOrderCreateFailed                 = 20092
	ErrCoderOrderUpdateFailed                 = 20093
	ErrCoderOrderDeleteFailed                 = 20094
	ErrCoderOrderDeleteByIDsFailed            = 20095
	ErrCoderOrderDeleteAllFailed              = 20096
	ErrCoderOrderGetByIDFailed                = 20097
	ErrCoderOrderGetAllFailed                 = 20098
	ErrCoderUserNotFount                      = 20101
	ErrCoderUserCreateFailed                  = 20102
	ErrCoderUserUpdateFailed                  = 20103
	ErrCoderUserDeleteFailed                  = 20104
	ErrCoderUserDeleteByIDsFailed             = 20105
	ErrCoderUserDeleteAllFailed               = 20106
	ErrCoderUserGetByIDFailed                 = 20107
	ErrCoderUserGetAllFailed                  = 20108
	ErrCoderRoleNotFount                      = 20111
	ErrCoderRoleCreateFailed                  = 20112
	ErrCoderRoleUpdateFailed                  = 20113
	ErrCoderRoleDeleteFailed                  = 20114
	ErrCoderRoleDeleteByIDsFailed             = 20115
	ErrCoderRoleDeleteAllFailed               = 20116
	ErrCoderRoleGetByIDFailed                 = 20117
	ErrCoderRoleGetAllFailed                  = 20118
	ErrCoderPermissionNotFount                = 20121
	ErrCoderPermissionCreateFailed            = 20122
	ErrCoderPermissionUpdateFailed            = 20123
	ErrCoderPermissionDeleteFailed            = 20124
	ErrCoderPermissionDeleteByIDsFailed       = 20125
	ErrCoderPermissionDeleteAllFailed         = 20126
	ErrCoderPermissionGetByIDFailed           = 20127
	ErrCoderPermissionGetAllFailed            = 20128
	ErrCoderPaymentMethodNotFount             = 20131
	ErrCoderPaymentMethodCreateFailed         = 20132
	ErrCoderPaymentMethodUpdateFailed         = 20133
	ErrCoderPaymentMethodDeleteFailed         = 20134
	ErrCoderPaymentMethodDeleteByIDsFailed    = 20135
	ErrCoderPaymentMethodDeleteAllFailed      = 20136
	ErrCoderPaymentMethodGetByIDFailed        = 20137
	ErrCoderPaymentMethodGetAllFailed         = 20138
	ErrCoderCustomerNotFount                  = 20141
	ErrCoderCustomerCreateFailed              = 20142
	ErrCoderCustomerUpdateFailed              = 20143
	ErrCoderCustomerDeleteFailed              = 20144
	ErrCoderCustomerDeleteByIDsFailed         = 20145
	ErrCoderCustomerDeleteAllFailed           = 20146
	ErrCoderCustomerGetByIDFailed             = 20147
	ErrCoderCustomerGetAllFailed              = 20148
	ErrCoderStaffNotFount                     = 20151
	ErrCoderStaffCreateFailed                 = 20152
	ErrCoderStaffUpdateFailed                 = 20153
	ErrCoderStaffDeleteFailed                 = 20154
	ErrCoderStaffDeleteByIDsFailed            = 20155
	ErrCoderStaffDeleteAllFailed              = 20156
	ErrCoderStaffGetByIDFailed                = 20157
	ErrCoderStaffGetAllFailed                 = 20158
	ErrCoderNotificationNotFount              = 20161
	ErrCoderNotificationGetByUserIDFailed     = 20162
	ErrCoderNotificationAdminFailed           = 20163
	ErrCoderSlideShowNotFount                 = 20171
	ErrCoderSlideShowCreateFailed             = 20172
	ErrCoderSlideShowUpdateFailed             = 20173
	ErrCoderSlideShowDeleteFailed             = 20174
	ErrCoderSlideShowDeleteByIDsFailed        = 20175
	ErrCoderSlideShowDeleteAllFailed          = 20176
	ErrCoderSlideShowGetByIDFailed            = 20177
	ErrCoderSlideShowGetAllFailed             = 20178
)

var msg = map[int]string{
	ErrCodeSuccess:                            "Success", // ✅ Sửa lỗi chính tả
	ErrCodeProductSizeNotFound:                "Product size not found",
	ErrCodeProductSizeCreateFailed:            "Failed to create product size",
	ErrCodeProductSizeUpdateFailed:            "Failed to update product size",
	ErrCodeProductSizeDeleteFailed:            "Failed to delete product size",
	ErrCodeProductSizeDeleteByIDsFailed:       "Failed to delete product sizes by IDs",
	ErrCodeProductSizeDeleteAllFailed:         "Failed to delete all product sizes",
	ErrCodeProductSizeGetByIDFailed:           "Failed to get product size by ID",
	ErrCodeProductSizeGetAllFailed:            "Failed to get all product sizes",
	ErrCodeProductSizeGetByIDsFailed:          "Failed to get product sizes by IDs",
	ErrCoderProductStyleNotFound:              "Product style not found",
	ErrCoderProductStyleCreateFailed:          "Failed to create product style",
	ErrCoderProductStyleUpdateFailed:          "Failed to update product style",
	ErrCoderProductStyleDeleteFailed:          "Failed to delete product style",
	ErrCoderProductStyleDeleteByIDsFailed:     "Failed to delete product styles by IDs",
	ErrCoderProductStyleDeleteAllFailed:       "Failed to delete all product styles",
	ErrCoderProductStyleGetByIDFailed:         "Failed to get product style by ID",
	ErrCoderProductStyleGetAllFailed:          "Failed to get all product styles",
	ErrCoderProductStyleGetByIDsFailed:        "Failed to get product styles by IDs",
	ErrCoderProductColorNotFound:              "Product color not found",
	ErrCoderProductColorCreateFailed:          "Failed to create product color",
	ErrCoderProductColorUpdateFailed:          "Failed to update product color",
	ErrCoderProductColorDeleteFailed:          "Failed to delete product color",
	ErrCoderProductColorDeleteByIDsFailed:     "Failed to delete product colors by IDs",
	ErrCoderProductColorDeleteAllFailed:       "Failed to delete all product colors",
	ErrCoderProductColorGetByIDFailed:         "Failed to get product color by ID",
	ErrCoderProductColorGetAllFailed:          "Failed to get all product colors",
	ErrCoderProductMaterialsNotFount:          "Product Materials not found",
	ErrCoderProductMaterialsCreateFailed:      "Failed to create product Materials",
	ErrCoderProductMaterialsUpdateFailed:      "Failed to update product Materials",
	ErrCoderProductMaterialsDeleteFailed:      "Failed to delete product Materials",
	ErrCoderProductMaterialsDeleteByIDsFailed: "Failed to delete product Materials by IDs",
	ErrCoderProductMaterialsDeleteAllFailed:   "Failed to delete all product Materials",
	ErrCoderProductMaterialsGetByIDFailed:     "Failed to get product Materials by ID",
	ErrCoderProductMaterialsGetAllFailed:      "Failed to get all product Materials",
	ErrCoderProductNotFount:                   "Product not found",
	ErrCoderProductCreateFailed:               "Failed to create product ",
	ErrCoderProductUpdateFailed:               "Failed to update product ",
	ErrCoderProductDeleteFailed:               "Failed to delete product ",
	ErrCoderProductDeleteByIDsFailed:          "Failed to delete product by IDs",
	ErrCoderProductDeleteAllFailed:            "Failed to delete all product ",
	ErrCoderProductGetByIDFailed:              "Failed to get product by ID",
	ErrCoderProductGetAllFailed:               "Failed to get all product ",
	ErrCoderProductStockNotFount:              "Product Stock not found",
	ErrCoderCategoryNotFount:                  "Category not found",
	ErrCoderCategoryCreateFailed:              "Failed to create Category ",
	ErrCoderCategoryUpdateFailed:              "Failed to update Category ",
	ErrCoderCategoryDeleteFailed:              "Failed to delete Category ",
	ErrCoderCategoryDeleteByIDsFailed:         "Failed to delete Category by IDs",
	ErrCoderCategoryDeleteAllFailed:           "Failed to delete all Category ",
	ErrCoderCategoryGetByIDFailed:             "Failed to get Category by ID",
	ErrCoderCategoryGetAllFailed:              "Failed to get all Category ",
	ErrCoderProductVariantNotFount:            "Product Variant not found",
	ErrCoderProductVariantCreateFailed:        "Failed to create product Variant",
	ErrCoderProductVariantUpdateFailed:        "Failed to update product Variant",
	ErrCoderProductVariantDeleteFailed:        "Failed to delete product Variant",
	ErrCoderProductVariantDeleteByIDsFailed:   "Failed to delete product Variant by IDs",
	ErrCoderProductVariantDeleteAllFailed:     "Failed to delete all product Variant",
	ErrCoderProductVariantGetByIDFailed:       "Failed to get product Variant by ID",
	ErrCoderProductVariantGetAllFailed:        "Failed to get all product Variant",
	ErrCoderOrderStatusNotFount:               "Order Status not found",
	ErrCoderOrderStatusCreateFailed:           "Failed to create Order Status",
	ErrCoderOrderStatusUpdateFailed:           "Failed to update Order Status",
	ErrCoderOrderStatusDeleteFailed:           "Failed to delete Order Status",
	ErrCoderOrderStatusDeleteByIDsFailed:      "Failed to delete Order Status by IDs",
	ErrCoderOrderStatusDeleteAllFailed:        "Failed to delete all Order Status",
	ErrCoderOrderStatusGetByIDFailed:          "Failed to get Order Status by ID",
	ErrCoderOrderStatusGetAllFailed:           "Failed to get all Order Status",
	ErrCoderOrderNotFount:                     "Order not found",
	ErrCoderOrderCreateFailed:                 "Failed to create Order ",
	ErrCoderOrderUpdateFailed:                 "Failed to update Order ",
	ErrCoderOrderDeleteFailed:                 "Failed to delete Order ",
	ErrCoderOrderDeleteByIDsFailed:            "Failed to delete Order by IDs",
	ErrCoderOrderDeleteAllFailed:              "Failed to delete all Order ",
	ErrCoderOrderGetByIDFailed:                "Failed to get Order by ID",
	ErrCoderOrderGetAllFailed:                 "Failed to get all Order ",
	ErrCoderUserNotFount:                      "User not found",
	ErrCoderUserCreateFailed:                  "Failed to create User ",
	ErrCoderUserUpdateFailed:                  "Failed to update User ",
	ErrCoderUserDeleteFailed:                  "Failed to delete User ",
	ErrCoderUserDeleteByIDsFailed:             "Failed to delete User by IDs",
	ErrCoderUserDeleteAllFailed:               "Failed to delete all User ",
	ErrCoderUserGetByIDFailed:                 "Failed to get User by ID",
	ErrCoderUserGetAllFailed:                  "Failed to get all User ",
	ErrCoderRoleNotFount:                      "Role not found",
	ErrCoderRoleCreateFailed:                  "Failed to create Role ",
	ErrCoderRoleUpdateFailed:                  "Failed to update Role ",
	ErrCoderRoleDeleteFailed:                  "Failed to delete Role ",
	ErrCoderRoleDeleteByIDsFailed:             "Failed to delete Role by IDs",
	ErrCoderRoleDeleteAllFailed:               "Failed to delete all Role ",
	ErrCoderRoleGetByIDFailed:                 "Failed to get Role by ID",
	ErrCoderRoleGetAllFailed:                  "Failed to get all Role ",
	ErrCoderPermissionNotFount:                "Permission not found",
	ErrCoderPermissionCreateFailed:            "Failed to create Permission ",
	ErrCoderPermissionUpdateFailed:            "Failed to update Permission ",
	ErrCoderPermissionDeleteFailed:            "Failed to delete Permission ",
	ErrCoderPermissionDeleteByIDsFailed:       "Failed to delete Permission by IDs",
	ErrCoderPermissionDeleteAllFailed:         "Failed to delete all Permission ",
	ErrCoderPermissionGetByIDFailed:           "Failed to get Permission by ID",
	ErrCoderPermissionGetAllFailed:            "Failed to get all Permission ",
	ErrCoderPaymentMethodNotFount:             "PaymentMethod not found",
	ErrCoderPaymentMethodCreateFailed:         "Failed to create PaymentMethod ",
	ErrCoderPaymentMethodUpdateFailed:         "Failed to update PaymentMethod ",
	ErrCoderPaymentMethodDeleteFailed:         "Failed to delete PaymentMethod ",
	ErrCoderPaymentMethodDeleteByIDsFailed:    "Failed to delete PaymentMethod by IDs",
	ErrCoderPaymentMethodDeleteAllFailed:      "Failed to delete all PaymentMethod ",
	ErrCoderPaymentMethodGetByIDFailed:        "Failed to get PaymentMethod by ID",
	ErrCoderPaymentMethodGetAllFailed:         "Failed to get all PaymentMethod ",
	ErrCoderCustomerNotFount:                  "Customer not found",
	ErrCoderCustomerCreateFailed:              "Failed to create Customer ",
	ErrCoderCustomerUpdateFailed:              "Failed to update Customer ",
	ErrCoderCustomerDeleteFailed:              "Failed to delete Customer ",
	ErrCoderCustomerDeleteByIDsFailed:         "Failed to delete Customer by IDs",
	ErrCoderCustomerDeleteAllFailed:           "Failed to delete all Customer ",
	ErrCoderCustomerGetByIDFailed:             "Failed to get Customer by ID",
	ErrCoderCustomerGetAllFailed:              "Failed to get all Customer ",
	ErrCoderStaffNotFount:                     "Staff not found",
	ErrCoderStaffCreateFailed:                 "Failed to create Staff ",
	ErrCoderStaffUpdateFailed:                 "Failed to update Staff ",
	ErrCoderStaffDeleteFailed:                 "Failed to delete Staff ",
	ErrCoderStaffDeleteByIDsFailed:            "Failed to delete Staff by IDs",
	ErrCoderStaffDeleteAllFailed:              "Failed to delete all Staff ",
	ErrCoderStaffGetByIDFailed:                "Failed to get Staff by ID",
	ErrCoderStaffGetAllFailed:                 "Failed to get all Staff ",
	ErrCoderNotificationNotFount:              "Notification not found",
	ErrCoderNotificationGetByUserIDFailed:     "Failed to get Notification by User ID",
	ErrCoderNotificationAdminFailed:           "Failed to get Notification admin",
	ErrCoderSlideShowNotFount:                 "SlideShow not found",
	ErrCoderSlideShowCreateFailed:             "Failed to create SlideShow ",
	ErrCoderSlideShowUpdateFailed:             "Failed to update SlideShow ",
	ErrCoderSlideShowDeleteFailed:             "Failed to delete SlideShow ",
	ErrCoderSlideShowDeleteByIDsFailed:        "Failed to delete SlideShow by IDs",
	ErrCoderSlideShowDeleteAllFailed:          "Failed to delete all SlideShow ",
	ErrCoderSlideShowGetByIDFailed:            "Failed to get SlideShow by ID",
	ErrCoderSlideShowGetAllFailed:             "Failed to get all SlideShow ",
}
