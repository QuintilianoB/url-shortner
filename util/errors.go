package util

// ErrorMessage and its method allow us to return json messages for the http requests.
type ErrorMessage struct {
	Message string `json:"message"`
}

func (e ErrorMessage) Error() string {
	return e.Message
}
