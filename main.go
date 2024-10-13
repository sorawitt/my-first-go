package main

import (
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
	books = append(books, Book{ID: 1, Name: "kong", Author: "kong"})
	app.Listen(":8080")
}

func handleGetBooks(fiber *fiber.Ctx) error {
	return fiber.Status(200).JSON(books)
}
