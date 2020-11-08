package model

import "errors"

var ErrNotFound = errors.New("element not found in repo")
var ErrDuplicatedValue = errors.New("duplicated value in repo")
var ErrDuplicatedIndexValue = errors.New("duplicated index value in repo")
var ErrIndexNotFound = errors.New("index value not found while updating")
