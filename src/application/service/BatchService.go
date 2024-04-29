package service

type BatchService struct {
}

func NewBatchService(service *BatchService) *BatchService {
	return &BatchService{}
}

func (b *BatchService) Migrate(path string) {

}
