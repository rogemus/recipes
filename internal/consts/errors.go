package consts

import "errors"

var ErrorNoEntry = errors.New("model: no entry")

var ErrInvalidCredentials = errors.New("auth: invalid credentials")

var ErrorEmailInUse = errors.New("model: email already in use")

var ErrorNoTemplate = errors.New("template: template does not exist")
