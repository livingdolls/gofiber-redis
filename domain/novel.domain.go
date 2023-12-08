package domain

import "gofiber-redis/model"

type NovelRepo interface {
	CreateNovel(createNovel model.Novel) error
	GetNovelById(id int) (model.Novel, error)
}

type NovelUseCase interface {
	CreateNovel(createNovel model.Novel) error
	GetNovelById(id int) (model.Novel, error)
}