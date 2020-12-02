package request

// Paging common input parameter structure
type PageInfo struct {
	Page int `json:"page" validate:"required,gte=1"`
	PageSize int `json:"page_size" validate:"required,gte=1,lte=50"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" validate:"required,gte=1"`
}

type IdsReq struct {
	Ids []int `json:"ids" validate:"required"`
}
