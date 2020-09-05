package errors

// ServiceError should be used to return business error Message
type ServiceError struct {
	Message string `json:"message"`
}
