package entity

// Body ...
type Body struct {
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	CurrentPage  interface{} `json:"current_page,omitempty"`
	Limit        interface{} `json:"limit,omitempty"`
	TotalEntries interface{} `json:"total_entries,omitempty"`
}
