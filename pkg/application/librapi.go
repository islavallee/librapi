package application

type StorageRepository interface {
	Delete(key string) error
	Get(key string, value interface{}) error
	Put(key string, value interface{}) error
}

type Librapi struct {
	storage StorageRepository
}

func NewLibrapi(storage StorageRepository) *Librapi {
	return &Librapi{
		storage: storage,
	}
}
