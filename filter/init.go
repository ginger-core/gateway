package filter

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type ctxKey string

const (
	ProcessedFieldsMapKey ctxKey = "fields"
)
