package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gofiber-redis/domain"
	"gofiber-redis/model"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type novelRepo struct {
	db *gorm.DB
	rdb *redis.Client
}

// GetNovelById Implements domain.novelRepo
func (n *novelRepo) GetNovelById(id int) (model.Novel, error) {
	var novels model.Novel
	var c = context.Background()

	// check available redis

	result, err := n.rdb.Get(c, "novel"+strconv.Itoa(id)).Result()

	if err != nil && err != redis.Nil {
		return novels, err
	}

	// if data available in redis, decode it from json and return it

	if len(result) > 0 {
		fmt.Printf("get data from redis")
		err := json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	// if data not available in redis, get it from database
	err = n.db.Model(model.Novel{}).Select("id", "name", "description", "author").Where("id = ?", id).Find(&novels).Error
	if err != nil {
		return novels, err
	}

	// Encode taht slice into json before saving redis
	jsonBytes, err := json.Marshal(novels)

	if err != nil {
		return novels, err
	}

	jsonString := string(jsonBytes)

	// set the json-encoded value in redis

	err = n.rdb.Set(c, "novel"+strconv.Itoa(id), jsonString, 24*time.Hour).Err();
	if err != nil {
		return novels, err
	}

	fmt.Printf("save redis")

	return novels, nil
}

// CreateNovel Implements domain.novelRepo
func (n *novelRepo) CreateNovel(createNovel model.Novel) error {
	if err := n.db.Create(&createNovel).Error; err != nil {
		return errors.New("Internal server error: cannot create novel")
	}

	return nil;
}

func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db: db,
		rdb: rdb,
	}
}