package models

import (
	"strings"

	"recipes.krogowski.dev/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func (f Filters) SortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort param:" + f.Sort)
}

func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}

	return "ASC"
}

func (f Filters) Limit() int {
	return f.PageSize
}

func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000, "page", "must be a maximum of 10 000")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
