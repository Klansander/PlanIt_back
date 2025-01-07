package create_item

type createItem interface {
	CreateItem(item int) (int, error)
}
