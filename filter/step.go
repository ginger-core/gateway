package filter

type step int

const (
	stepType step = iota + 1
	stepWhere
	stepWhereField
	stepWhereOperator
	stepWhereValue
	stepWhereMatchDone
	stepOperation
	stepGroupStart
	stepGroupEnd
)
