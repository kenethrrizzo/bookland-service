package books

import (
	"errors"

	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

var allowedGenres = [...]string{"TER", "COM", "MIS", "POL", "DRA"}

func validateGenre(genres []string) error {
	for _, genre := range genres {
		if genre == "" {
			return nil
		}
	
		for _, g := range allowedGenres {
			if g != genre {
				return domainErrors.NewAppError(errors.New("genre not allowed"), domainErrors.MapError)
			}
		}
	}
	return nil
}
