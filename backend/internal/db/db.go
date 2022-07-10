package db

import (
	"github.com/gsmerlin/minify/backend/internal/logger"
	"github.com/gsmerlin/minify/backend/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	ID          string
	Email       string
	Destination string
}

type Analytics struct {
	ID        string
	Timestamp string `sql:"-"`
}

type Details struct {
	ID        string
	Analytics []Analytics
}

var r *gorm.DB

func Start() {
	if r != nil {
		return
	}
	logger.Info("Initializing database connection...")
	dsn := "root:root@tcp(127.0.0.1:3306)/minify?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	r, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Database connection successfully initialized")
}

func NewLink(id string, email string, destination string) (Record, error) {
	if id == "" {
		id = utils.RandSeq(3)
	}

	record := Record{
		ID:          id,
		Email:       email,
		Destination: destination,
	}

	result := r.Create(&record)
	if result.Error != nil {
		return Record{}, result.Error
	}
	return record, nil
}

func GetLink(id string, email string, destination string) ([]Record, error) {
	var records []Record

	if id != "" && destination == "" && email == "" {
		result := r.First(&records, "id = ?", id)
		if result.Error != nil {
			return nil, result.Error
		}
		return records, nil
	}

	if destination != "" && id == "" && email == "" {
		result := r.Find(&records, "destination = ?", destination)
		if result.Error != nil {
			return nil, result.Error
		}

		return records, nil
	}

	if email != "" && id == "" && destination == "" {
		result := r.Find(&records, "email = ?", email)
		if result.Error != nil {
			return nil, result.Error
		}
		return records, nil
	}

	result := r.Find(&records)

	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func UpdateLink(id string, email string, destination string) (string, error) {
	record := Record{
		ID:          id,
		Email:       email,
		Destination: destination,
	}
	result := r.Save(&record)
	if result.Error != nil {
		return "", result.Error
	}
	return record.ID, nil
}

func DeleteLink(id string) (string, error) {
	result := r.Delete(Record{ID: id})
	if result.Error != nil {
		return "", result.Error
	}

	return id, nil
}

func AddAnalytics(id string) error {
	type Analytics struct {
		ID string
	}
	result := r.Create(&Analytics{ID: id})
	if result.Error != nil {
		logger.Error(result.Error.Error())
	}
	return result.Error
}

func GetAnalytics(id string) (Details, error) {

	var analytics []Analytics

	if res := r.Find(&analytics, "ID = ?", id); res.Error != nil {
		return Details{}, res.Error
	}

	return Details{
		ID:        id,
		Analytics: analytics,
	}, nil
}
