package dto

import "errors"

var ConflictError = errors.New("resource version not match")
var UnavailableVersionError = errors.New("current version unavailable: current_version > version")

var NotFoundError = errors.New("resource not found")
var AlreadyExistsError = errors.New("resource already exists")
