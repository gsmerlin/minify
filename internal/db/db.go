package db

import (
	"log"

	"github.com/gsmerlin/minify/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	ID          string
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

func InitDB() {
	log.Println("Initializing database connection...")
	dsn := "root:root@tcp(127.0.0.1:3306)/minify?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	r, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.ErrorCheck(err)
	log.Println("Database connection successfully initialized")
}

func NewLink(id string, destination string) string {
	if id == "" {
		id = utils.RandSeq(10)
	}

	record := Record{
		ID:          id,
		Destination: destination,
	}

	result := r.Create(&record)
	utils.ErrorCheck(result.Error)
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

	r.Find(&records)
	return records
}

func Update(rec Record) {
	result := r.Save(rec)
	utils.ErrorCheck(result.Error)
}

func DeleteLink(id string) {
	result := r.Delete(Record{ID: id})
	utils.ErrorCheck(result.Error)
}

func AddAnalytics(id string) error {
	result := r.Save(&Analytics{ID: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
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
