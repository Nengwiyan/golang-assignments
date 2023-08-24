package request

type VariantRequest struct {
	VariantName string `form:"variant_name"`
	Quantity    int    `form:"quantity"`
	ProductID   string `form:"product_id"`
}
