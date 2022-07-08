package db

import (
	"github.com/gsmerlin/minify/internal/logger"
	"github.com/gsmerlin/minify/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	ID          string
	Email       string
	Destination string
}

type Analytics struct {
	ID string
}
type Details struct {
	ID          string
	Destination string
	Analytics   []Analytics
}

var r *gorm.DB

func Start() {
	logger.Info("Initializing database connection...")
	dsn := "root:root@tcp(127.0.0.1:3306)/minify?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	r, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Database connection successfully initialized")
}

func NewLink(id string, email string, destination string) string {
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
		logger.Error(result.Error.Error())
	}
	return record.ID
}

func GetLink(id string, destination string) []Record {
	var records []Record

	if id != "" {
		r.First(&records, "id = ?", id)
		return records
	}

	if destination != "" {
		r.Find(&records, "destination = ?", destination)
		return records
	}

	result := r.Find(&records)
	if result.Error != nil {
		logger.Error(result.Error.Error())
	}
	return records
}

func UpdateLink(rec Record) {
	result := r.Save(rec)
	if result.Error != nil {
		logger.Error(result.Error.Error())
	}
}

func DeleteLink(id string) {
	result := r.Delete(Record{ID: id})
	if result.Error != nil {
		logger.Error(result.Error.Error())
	}
}

func AddAnalytics(id string) error {
	result := r.Save(&Analytics{ID: id})
	if result.Error != nil {
		logger.Error(result.Error.Error())
	}
	return result.Error
}

func GetDetails(id string) Details {

	var record Record
	var analytics []Analytics

	r.First(&record, "ID = ?", id)
	r.Find(&analytics, "ID = ?", id)

	return Details{
		ID:          record.ID,
		Destination: record.Destination,
		Analytics:   analytics,
	}
}
