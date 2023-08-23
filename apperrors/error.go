package apperrors

type AppError struct {
	ErrCode `json:"errCode"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (err AppError) Error() string {
	return err.Err.Error()
}

func (err AppError) Unwrap() error {
	return err.Err
}
