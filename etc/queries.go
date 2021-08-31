package etc

import (
	"fmt"
	"github.com/bektosh03/pkg/errs"
	"strconv"
)

func ParsePage(pageValue string) (uint64, error) {
	page, err := strconv.ParseUint(pageValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%v: %w", err, errs.ErrBadPageValue)
	}
	return page, nil
}

func ParseLimit(limitValue string) (uint64, error) {
	limit, err := strconv.ParseUint(limitValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%v: %w", err, errs.ErrBadLimitValue)
	}
	return limit, nil
}
