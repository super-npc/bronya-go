package commons

import "strconv"

func StringsToUints(ss []string) ([]uint, error) {
	res := make([]uint, 0, len(ss))
	for _, s := range ss {
		v, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			return nil, err
		}
		res = append(res, uint(v))
	}
	return res, nil
}
