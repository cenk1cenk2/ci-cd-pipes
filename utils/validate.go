package utils

import (
	"fmt"
	"sync"

	"github.com/creasty/defaults"
	validator "github.com/go-playground/validator/v10"
)

func ValidateAndSetDefaults(metadata TaskMetadata, s ...interface{}) {
	log := Log.WithField("context", metadata.Context)

	var wg sync.WaitGroup
	wg.Add(len(s))

	var validationError bool

	for i, v := range s {
		go func(i int, v interface{}) {
			defer wg.Done()

			pointer := &s[i]

			if err := defaults.Set(pointer); err != nil {
				log.Fatalln(fmt.Sprintf("Can not set defaults: %s", err))
			}

			validate := validator.New()

			err := validate.Struct(v)

			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					error := fmt.Sprintf(
						"\"%s\" field failed validation: %s",
						err.Namespace(),
						err.Tag(),
					)

					log.Errorln(error)

				}

				validationError = true
			}
		}(i, v)
	}

	wg.Wait()

	if validationError {
		log.Fatalln("Validation failed.")
	}
}
