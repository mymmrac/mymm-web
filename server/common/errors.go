package common

import "errors"

func FirstError(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}

	return nil
}

var ErrNotFound = errors.New("not found")
