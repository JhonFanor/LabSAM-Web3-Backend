package dto

type PaginationDTO struct {
	TotalRecords int `json:"totalRecords"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
	PageSize     int `json:"pageSize"`
}
