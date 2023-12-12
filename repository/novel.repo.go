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

// Get All novel

func (n *novelRepo) GetAllNovel() ([]model.Novel, error) {
	var novels []model.Novel
	var c = context.Background()

	// Check available Redis

	result,err := n.rdb.Get(c, "getAllNovel").Result()

	if err != nil && err != redis.Nil {
		return novels, err
	}

	// If data available in redis, decode it from json and return it

	if len(result) > 0 {
		fmt.Printf("data get from redis")

		err := json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	// if data not available in redis, get it from database
	err = n.db.Model(&novels).Select("id", "name", "description", "author").Find(&novels).Error
	if err != nil {
		return novels, err
	}

	// Encode that slice into json before saving json
	jsonBytes, err := json.Marshal(novels)

	if err != nil {
		return novels, err
	}

	jsonString := string(jsonBytes)

	err = n.rdb.Set(c, "getAllNovel", jsonString, 24*time.Hour).Err()
	if err != nil {
		return novels, err
	}

	fmt.Println("save redis")

	return novels, nil
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

// Delete Novel Implements domain.novelRepo
func (n *novelRepo) DeleteNovel(id int) (model.Novel, error) {
	var novels model.Novel
	var c = context.Background();


	// Check available redis;

	result, err := n.rdb.Get(c, "novel"+strconv.Itoa(id)).Result();

	fmt.Println(result)

	if err != nil && err != redis.Nil {
		return novels, err
	}

	if len(result) > 0 {
		// Delete from redis
		_, err := n.rdb.Del(c, "novel"+strconv.Itoa(id)).Result()
		if err != nil {
			
		}

		fmt.Println("Delete Ok")
	}

	err = n.db.Model(model.Novel{}).Where("id = ?", id).Find(&novels).Error;

	if err != nil {
		return novels, err
	}

	// Delete

	errDelete := n.db.Delete(&novels).Error

	if errDelete != nil {
		return novels, err
	}

	return novels, err
}

func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db: db,
		rdb: rdb,
	}
}