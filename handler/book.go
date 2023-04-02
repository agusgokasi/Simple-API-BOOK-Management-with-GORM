package handler

import (
	"ninth-learn/helper"
	"ninth-learn/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary		 Creates a new Book
// @Description  Create a new book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 		 book_request body model.BookRequest true "Book request object"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /books [post]
func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetBookById godoc
// @Summary      Show a book
// @Description  get detail book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /books/{id} [get]
func (h HttpServer) GetBookById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helper.BadRequest(c, "Invalid book ID")
		return
	}

	// call service
	res, err := h.app.GetBookById(id)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetBooks godoc
// @Summary      Show all book
// @Description  get all book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /books [get]
func (h HttpServer) GetBooks(c *gin.Context) {
	// call service
	res, err := h.app.GetBooks()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// UpdateBook godoc
// @Summary      Update book
// @Description  Update book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Param 		 book_request body model.BookRequest true "Book request object"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /books/{id} [put]
func (h HttpServer) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helper.BadRequest(c, "Invalid book ID")
		return
	}

	in := model.Book{}
	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	in.ID = id
	// call service
	res, err := h.app.UpdateBook(in)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// DeleteBook godoc
// @Summary      Delete book
// @Description  Delete book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /books/{id} [delete]
func (h HttpServer) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helper.BadRequest(c, "Invalid book ID")
		return
	}

	// call service
	err = h.app.DeleteBook(id)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Book deleted successfully")
}
