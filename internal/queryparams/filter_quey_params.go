package queryparams

type FilterQueryParams struct {
	Field    string
	Operator string
	Value    any
}

type Filters []FilterQueryParams
