package handler

import "strconv"

func CheckID(id string) (uint, error) {
	valid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(valid), nil
}