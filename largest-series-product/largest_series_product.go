package lsproduct

import (
	"errors"
	"slices"
	"unicode/utf8"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	len := utf8.RuneCountInString(digits)
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if len < span {
		return 0, errors.New("span must be smaller than string length")
	}

	products := make([]int64, len)
	for i := 0; i < len; i++ {
		digit := int64(digits[i] - '0')

		if digit < 0 || digit > 9 {
			return 0, errors.New("invalid digit")
		}
		if i < len-span+1 {
			products[i] = digit
		}

		for j := 0; j < span-1; j++ {
			if i > j {
				products[i-j-1] *= digit
			}
		}
	}
	return slices.Max(products), nil
}
