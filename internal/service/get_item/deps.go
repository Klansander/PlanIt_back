package get_item

type getItem interface {
	FetchItems() ([]int, error)
}
