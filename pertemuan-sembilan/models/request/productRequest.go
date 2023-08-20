package request

import "mime/multipart"

type ProductRequest struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       int64                 `form:"price" binding:"required"`
	Stock       int                   `form:"stock" binding:"required"`
	Image       *multipart.FileHeader `form:"file"`
}
