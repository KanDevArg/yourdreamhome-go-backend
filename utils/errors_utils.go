package utils

type ApplicationError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e ApplicationError) Error() string {
	return e.Message
}
