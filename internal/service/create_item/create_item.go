package create_item

type CreateItemService struct {
	repo createItem
}

func NewCreateItemService(repo createItem) *CreateItemService {
	return &CreateItemService{repo: repo}
}

func (s *CreateItemService) CreateItem(item int) (int, error) {
	return s.repo.SaveItem(item)
}
