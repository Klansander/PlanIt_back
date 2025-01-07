package get_item

type GetItemsService struct {
	repo getItem
}

func NewGetItemsService(repo getItem) *GetItemsService {
	return &GetItemsService{repo: repo}
}

func (s *GetItemsService) GetItems() ([]int, error) {
	return s.repo.FetchItems()
}
