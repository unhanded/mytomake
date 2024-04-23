package util

import "strconv"

type RangeStringArgs struct {
	Min string
	Max string
}

func (rsargs RangeStringArgs) ToFloat() (float64, float64, error) {
	max, err := strconv.ParseFloat(rsargs.Max, 64)
	if err != nil {
		return 0, 0, err
	}
	min, err := strconv.ParseFloat(rsargs.Min, 64)
	if err != nil {
		return 0, 0, err
	}
	return min, max, nil
}

func (rsargs RangeStringArgs) ToInt() (int64, int64, error) {
	max, err := strconv.ParseInt(rsargs.Max, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	min, err := strconv.ParseInt(rsargs.Min, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return min, max, nil
}

type DictArgs struct {
	DictName string
}
