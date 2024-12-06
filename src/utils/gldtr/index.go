package gldtr

import "github.com/golodash/galidator/v2"

var (
	G = galidator.New().
		CustomMessages(galidator.Messages{
			"required": "'$field' is required",
			"email":    "'$field' must be an email",
			// "min" : "'$field' $rule characters",
			// "time_format" : "'$field' must be a valid time",
		})
)
