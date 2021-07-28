package controller

import (
	"api/test/config"
	"api/test/database"
	"api/test/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func testDeleteBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	reqBody, err := json.Marshal(map[string]string{
		"title": "Abc",
	})
	if err != nil {
		t.Error(err)
		return
	}
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var book model.Book
	if err := json.Unmarshal(body, &book); err != nil {
		t.Error(err)
	}
	if book.Title != "Abc" {
		t.Errorf("expected Abc, got: %#v", book.Title)
	}
}

func TestDBDeleteBookController(t *testing.T) {
	// bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Book{})
	db.Delete(&model.Book{}, "1=1")
	m := model.NewGormBookModel(db)
	// bikin controller
	bookController := CreatePostBookController(m)
	if err != nil {
		t.Error(err)
	}
	// test controller
	testPostBookController(t, bookController)
	db.Delete(&model.Book{}, "1=1")
}

func TestMockDeleteBookController(t *testing.T) {
	m := model.NewMockBookModel()
	bookController := CreatePostBookController(m)
	// test controller
	testPostBookController(t, bookController)
}
