package handlers

import "strconv"

func parseUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

func parseStr(i uint) string {
	return strconv.Itoa(int(i))
}
