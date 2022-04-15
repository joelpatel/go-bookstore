package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joelpatel/go-bookstore/pkg/models"
	"github.com/joelpatel/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetAllBooks(res_w http.ResponseWriter, req *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books) // convert into json
	res_w.Header().Set("Content-Type", "application/json")
	res_w.WriteHeader(http.StatusOK)
	res_w.Write(res)
}

func GetBookById(res_w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	res_w.Header().Set("Content-Type", "application/json")
	res_w.WriteHeader(http.StatusOK)
	res_w.Write(res)
}

func CreateBook(res_w http.ResponseWriter, req *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(req, newBook)
	b := newBook.CreateBook() // this create book function is of models package
	res, _ := json.Marshal(b)
	res_w.Header().Set("Content-Type", "application/json")
	res_w.WriteHeader(http.StatusOK)
	res_w.Write(res)
}

func DeleteBook(res_w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	res_w.Header().Set("Content-Type", "application/json")
	res_w.WriteHeader(http.StatusOK)
	res_w.Write(res)
}

func UpdateBook(res_w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	updateValues := &models.Book{}
	utils.ParseBody(req, updateValues)
	updatedBook := models.UpdateBook(ID, updateValues)
	res, _ := json.Marshal(updatedBook)
	res_w.Header().Set("Content-Type", "application/json")
	res_w.WriteHeader(http.StatusOK)
	res_w.Write(res)
}
