package consts

const (
	RelatedFieldFilterSeparator = "."
	OperatorFilterSeparator     = ":"
	InOptionsSeparator          = ","
)

// Filter operators

const (
	FilterOperatorEq                       = "eq"
	FilterOperatorNeq                      = "neq"
	FilterOperatorLt                       = "lt"
	FilterOperatorGt                       = "gt"
	FilterOperatorLeq                      = "leq"
	FilterOperatorGeq                      = "geq"
	FilterOperatorIn                       = "in"
	FilterOperatorNin                      = "nin"
	FilterOperatorContains                 = "contains"
	FilterOperatorNContains                = "ncontains"
	FilterOperatorContainsCaseSensitive    = "containss"
	FilterOperatorNContainsCaseSensitive   = "ncontainss"
	FilterOperatorStartsWith               = "startswith"
	FilterOperatorEndsWith                 = "endswith"
	FilterOperatorStartsWithCaseSensitive  = "startswiths"
	FilterOperatorEndsWithCaseSensitive    = "endswiths"
	FilterOperatorNStartsWith              = "nstartswith"
	FilterOperatorNEndsWith                = "nendswith"
	FilterOperatorNEndsWithCaseSensitive   = "nendswiths"
	FilterOperatorNStartsWithCaseSensitive = "nstartswiths"
)

// Query params

const (
	PageQueryParam   = "page"
	LimitQueryParam  = "limit"
	SortByQueryParam = "sort"
)

// Sort operators
const (
	SortOperatorAsc  = "asc"
	SortOperatorDesc = "desc"
)
