package app

import (
	"github.com/deliangyang/test-wire/internal/app/app"
	"github.com/pkg/errors"
)

var application *Application

type Application app.Application

func InitApplication() error {
	v, err := app.NewApplication()
	if err != nil {
		return errors.Wrap(err, "init application")
	}

	application = (*Application)(v)
	return nil
}

func GetApplication() *Application {
	return application
}
