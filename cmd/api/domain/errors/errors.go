package errors

import "errors"

const (
	// Indica un registro no encontrado
	NotFound        = "NotFound"
	notFoundMessage = "registro no encontrado"
	// Indica un error desconocido
	UnknownError        = "UnknownError"
	unknownErrorMessage = "algo salió mal"
	// Indica un error en el repositorio (ej. base de datos)
	RepositoryError        = "RepositoryError"
	repositoryErrorMessage = "error en operación de repositorio"
	// Indica un error en el mapeo de datos
	MapError        = "MapError"
	mapErrorMessage = "error en mapeo de datos"
)

type AppError struct {
	Err  error
	Type string
}

func NewAppError(err error, errType string) *AppError {
	return &AppError{
		Err:  err,
		Type: errType,
	}
}

func NewAppErrorWithType(errType string) *AppError {
	var err error

	switch errType {
	case NotFound:
		err = errors.New(notFoundMessage)
	case UnknownError:
		err = errors.New(unknownErrorMessage)
	case RepositoryError:
		err = errors.New(repositoryErrorMessage)
	case MapError:
		err = errors.New(mapErrorMessage)
	}

	return &AppError{
		Err:  err,
		Type: errType,
	}
}

func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}
