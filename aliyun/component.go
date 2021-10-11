package aliyun

import "errors"

var(
	errEmptyResult = errors.New("Empty result")
)

type Component struct {
	config *config

}

