package create_item

type createItem interface {
	SaveItem(int) (int, error)
}
