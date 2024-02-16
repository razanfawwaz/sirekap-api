package config

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Page struct {
	Size        int `json:"size"`
	TotalData   int `json:"totalData"`
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
}
