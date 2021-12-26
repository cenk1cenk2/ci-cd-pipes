package utils

import (
	"errors"
	"fmt"

	"github.com/creasty/defaults"
	validator "github.com/go-playground/validator/v10"
)

func ValidateAndSetDefaults(metadata TaskMetadata, s interface{}) error {
	log := Log.WithField("context", metadata.Context)

	if err := defaults.Set(s); err != nil {
		return errors.New(fmt.Sprintf("Can not set defaults: %s", err))
	}

	validate := validator.New()

	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error := fmt.Sprintf(
				"\"%s\" field failed validation: %s",
				err.Namespace(),
				err.Tag(),
			)

			log.Errorln(error)
		}

		return errors.New("Validation failed.")
	}

	return nil
}
