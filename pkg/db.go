package minify

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	ID          string
	Destination string
}

type Analytics struct {
	ID         string
	AccessedAt string
}

// type detailsClicks struct {
// 	day        string
// 	timestamps []string
// 	accesses   int
// }

type Details struct {
	ID          string
	Destination string
	// Details     []detailsClicks
	Analytics []Analytics
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/minify?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	r.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	ErrorCheck(err)
}

func (r *Repository) Create(destination string, id string) string {
	if id == "" {
		id = randSeq(10)
	}

	record := Record{
		ID:          id,
		Destination: destination,
	}

	result := r.db.Create(&record)
	ErrorCheck(result.Error)
	return record.ID
}

func (r *Repository) Read(id string, destination string) []Record {
	var records []Record

	if id != "" {
		r.db.First(&records, "id = ?", id)
		return records
	}

	if destination != "" {
		r.db.Find(&records, "destination = ?", destination)
		return records
	}

	r.db.Find(&records)
	return records
}

func (r *Repository) Update(rec Record) {
	result := r.db.Save(rec)
	ErrorCheck(result.Error)
}

func (r *Repository) Delete(id string) {
	result := r.db.Delete(Record{ID: id})
	ErrorCheck(result.Error)
}

func (r *Repository) addAnalytics(id string) {
	r.db.Create(&Analytics{ID: id})
}

func (r *Repository) GetDetails(id string) Details {

	var record Record
	var analytics []Analytics

	r.db.First(&record, "ID = ?", id)
	r.db.Find(&analytics, "ID = ?", id)

	return Details{
		ID:          record.ID,
		Destination: record.Destination,
		Analytics:   analytics,
	}
}
