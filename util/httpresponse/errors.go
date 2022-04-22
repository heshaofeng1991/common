package httpresponse

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
)

type CommonError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s CommonError) Error() string {
	return s.error
}

func (s CommonError) Slug() string {
	return s.slug
}

func (s CommonError) ErrorType() ErrorType {
	return s.errorType
}

func NewSlugError(err string, slug string) CommonError {
	return CommonError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeUnknown,
	}
}

func NewAuthorizationError(err string, slug string) CommonError {
	return CommonError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeAuthorization,
	}
}

func NewIncorrectInputError(err string, slug string) CommonError {
	return CommonError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeIncorrectInput,
	}
}
