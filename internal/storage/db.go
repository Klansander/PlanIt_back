package storage

type Repository interface {
	FetchItems() ([]int, error)
	SaveItem(item int) (int, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FetchItems() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (r *repository) SaveItem(item int) (int, error) {
	return 1, nil
}
