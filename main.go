package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int
	Name   string
	Author string
}

var books []Book

func main() {
	app := fiber.New()
	app.Get("books", handleGetBooks)
	app.Get("book/:ID", handleGetBook)
	books = append(books, Book{ID: 1, Name: "kong", Author: "kong"})
	app.Listen(":8080")
}

func handleGetBooks(c *fiber.Ctx) error {
	return c.Status(200).JSON(books)
}

func handleGetBook(c *fiber.Ctx) error {
	bookID, error := strconv.Atoi(c.Params("ID"))
	if error != nil {
		return c.Status(400).SendString(fiber.ErrBadRequest.Error())
	}
	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusBadRequest).SendString("Book not found")
}
