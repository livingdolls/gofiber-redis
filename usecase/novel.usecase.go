package usecase

import (
	"errors"
	"gofiber-redis/domain"
	"gofiber-redis/model"
)

type novelUseCase struct {
	novelRepo domain.NovelRepo
}

// DeleteNovel implements domain.NovelUseCase.
func (n *novelUseCase) DeleteNovel(id int) (model.Novel, error) {
	res, err := n.novelRepo.DeleteNovel(id)

	if err != nil {
		return model.Novel{}, errors.New("internal server error : " + err.Error())
	}

	return res, nil	
}

// GetAllNovel implements domain.NovelUseCase.
func (n *novelUseCase) GetAllNovel() ([]model.Novel, error) {
	res, err := n.novelRepo.GetAllNovel()

	if err != nil {
		return []model.Novel{}, errors.New("internal server error : " + err.Error())
	}

	return res, nil
}

// GetNovelById implements domain.NovelUseCase.
func (n *novelUseCase) GetNovelById(id int) (model.Novel, error) {
	res, err := n.novelRepo.GetNovelById(id)

	if err != nil {
		return model.Novel{}, errors.New("internal server error : " + err.Error())
	}

	return res, nil
}

func (n *novelUseCase) CreateNovel(createNovel model.Novel) error {
	err := n.novelRepo.CreateNovel(createNovel)
	return errors.New("Internal server error : " + err.Error())
}

func NewNovelUseCase(novelRepo domain.NovelRepo) domain.NovelUseCase {
	return &novelUseCase{
		novelRepo: novelRepo,
	}
}
