package tests

import (
	"testing"

	"github.com/gsmerlin/minify/backend/internal/db"
)

var TestLink = db.Record{
	ID:          "test",
	Email:       "test@test.test",
	Destination: "http://test.test",
}

func TestNewLink(t *testing.T) {
	db.Start()
	_, err := db.NewLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	if err != nil {
		t.Error(err)
	}
	db.DeleteLink(TestLink.ID)
}
func TestGetLink(t *testing.T) {
	db.Start()
	db.NewLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	_, err := db.GetLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	if err != nil {
		t.Error(err)
	}
	db.DeleteLink(TestLink.ID)
}
func TestUpdateLink(t *testing.T) {
	db.Start()
	db.NewLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	_, err := db.UpdateLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	if err != nil {
		t.Error(err)
	}
	db.DeleteLink(TestLink.ID)

}
func TestAddAnalytics(t *testing.T) {
	db.Start()
	db.NewLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	err := db.AddAnalytics(TestLink.ID)
	if err != nil {
		t.Error(err)
	}
	db.DeleteLink(TestLink.ID)
}
func TestGetAnalytics(t *testing.T) {
	db.Start()
	db.NewLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	_, err := db.GetAnalytics(TestLink.ID)
	if err != nil {
		t.Error(err)
	}
	db.DeleteLink(TestLink.ID)
}

func TestDeleteLink(t *testing.T) {
	db.Start()
	db.NewLink(TestLink.ID, TestLink.Email, TestLink.Destination)
	_, err := db.DeleteLink(TestLink.ID)
	if err != nil {
		t.Error(err)
	}
}
