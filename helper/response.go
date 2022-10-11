package helper

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  Errors `json:"errors"`
}

type Errors interface{}
