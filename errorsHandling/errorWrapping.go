package errorsHandling

import "errors"

func PublicError() {

}

func ErrorBase() error {
	return errors.New("error1")
}
