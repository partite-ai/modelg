package models

type OrderBy int

const (
	OrderByAsc OrderBy = iota + 1
	OrderByDesc
)

func (o OrderBy) SQLText() string {
	switch o {
	case OrderByAsc:
		return "ASC"
	case OrderByDesc:
		return "DESC"
	}
	return ""
}
