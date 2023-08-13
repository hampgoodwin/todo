package pagination

// HasNextPage calculate if there are more pages.
func HasNextPage(limit, totalItems int64) bool {
	// if limit is 0, it means that there is no limit :)
	if limit == 0 {
		return false
	}
	// if totalItems from the last page is 0, we probably don't have another page
	if totalItems == 0 {
		return false
	}
	return limit <= totalItems
}
