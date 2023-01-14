package enum

type SortType string

const (
	CreatedAtAsc  SortType = "created"
	CreatedAtDesc SortType = "-created"
)

// this will return an empty string of the provided sort-type isn't one we support
func StringToSortType(s string) SortType {
	switch SortType(s) {
	case CreatedAtAsc, CreatedAtDesc:
		return SortType(s)
	}
	return ""
}
