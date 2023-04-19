package handler

import (
	"challenge_10/entity"
	"challenge_10/helper"
	"challenge_10/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}

func (h *handler) CreateBook(c *gin.Context) {
	var payloadBook *entity.PayloadBook

	userData := c.MustGet("userData").(jwt.MapClaims)
	user_id := userData["user_id"].(string)

	err := c.ShouldBindJSON(&payloadBook)
	if err != nil {
		error := helper.FormatErrors(err)
		errorMessage := gin.H{"errors": error}

		response := helper.APIResponse("Unprocessable Entity!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	errorMessage := []string{}
	if payloadBook.Book_name == "" {
		errorMessage = append(errorMessage, "Book Name must be Fill!")
	}

	if payloadBook.Author == "" {
		errorMessage = append(errorMessage, "Author Name must be Fill!")
	}

	if payloadBook.Book_name == "" || payloadBook.Author == "" {
		response := helper.APIResponse("Failed to Create a Book", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	payloadBook.User_id = uuid.Must(uuid.FromString(user_id))

	newBook, err := h.service.CreateBook(payloadBook)
	if err != nil {
		response := helper.APIResponse("Failed to Create a Book", http.StatusBadRequest, "Error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Succes Create a Book", http.StatusCreated, "Success", newBook)
	c.JSON(http.StatusCreated, response)
}

func (h *handler) GetBooks(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	user_id := userData["user_id"].(string)
	access_level := userData["access_level"].(string)

	userUUID := uuid.Must(uuid.FromString(user_id))

	switch access_level {
	case "client":
		books, err := h.service.GetBooksByClientID(userUUID)
		if err != nil {
			response := helper.APIResponse("Failed to Get All Books", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("Success Get All Books", http.StatusOK, "Success", books)
		c.JSON(http.StatusOK, response)

	case "admin":
		books, err := h.service.GetBooks()
		if err != nil {
			response := helper.APIResponse("Failed to Get All Books", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("Success Get All Books", http.StatusOK, "Success", books)
		c.JSON(http.StatusOK, response)
	default:
		response := helper.APIResponse("Access Level not Found!", http.StatusNotFound, "Error", nil)
		c.JSON(http.StatusNotFound, response)
	}

}

func (h *handler) GetBookByID(c *gin.Context) {
	param := c.Param("bookID")

	paramUUID, _ := uuid.FromString(param)

	userData := c.MustGet("userData").(jwt.MapClaims)

	user_id := userData["user_id"].(string)
	access_level := userData["access_level"].(string)

	userUUID := uuid.Must(uuid.FromString(user_id))

	if param == ":bookID" {
		errorMessage := "UUID is empty!"
		response := helper.APIResponse("Failed to Get All Books", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !helper.IsValidUUID(param) {
		errorMessage := "UUID is not Valid!"
		response := helper.APIResponse("Failed to Get All Books", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	switch access_level {
	case "client":
		book, err := h.service.GetBookClientIDByID(paramUUID, userUUID)

		if book == nil {
			errorMessage := "Book not Found!"
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusNotFound, "Error", errorMessage)
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err != nil {
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Success to Get Book by ID", http.StatusOK, "Success", book)
		c.JSON(http.StatusOK, response)

	case "admin":
		book, err := h.service.GetBookByID(paramUUID)

		if book == nil {
			errorMessage := "Book not Found!"
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusNotFound, "Error", errorMessage)
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err != nil {
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Success to Get Book by ID", http.StatusOK, "Success", book)
		c.JSON(http.StatusOK, response)
	default:
		response := helper.APIResponse("Access Level not Found!", http.StatusNotFound, "Error", nil)
		c.JSON(http.StatusNotFound, response)
	}

}

func (h *handler) UpdateBookByID(c *gin.Context) {
	param := c.Param("bookID")
	var payloadBook entity.PayloadBook

	paramUUID, _ := uuid.FromString(param)

	userData := c.MustGet("userData").(jwt.MapClaims)
	access_level := userData["access_level"].(string)

	if param == ":bookID" {
		errorMessage := "UUID is empty!"
		response := helper.APIResponse("Failed Update a Books", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !helper.IsValidUUID(param) {
		errorMessage := "UUID is not Valid!"
		response := helper.APIResponse("Failed Update a Books", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&payloadBook)
	if err != nil {
		errors := helper.FormatErrors(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Format Form Book not Appripriate!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	switch access_level {
	case "client":
		response := helper.APIResponse("Access Denied!", http.StatusForbidden, "Error", "Access Denied!")
		c.JSON(http.StatusOK, response)

	case "admin":
		book, err := h.service.GetBookByID(paramUUID)
		if book == nil {
			errorMessage := "Book not Found!"
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusNotFound, "Error", errorMessage)
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err != nil {
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		updateBook, err := h.service.UpdateBookByID(paramUUID, &payloadBook)
		if err != nil {
			response := helper.APIResponse("Failed to Update Book by ID", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Succes Update Book by ID", http.StatusOK, "Success", updateBook)
		c.JSON(http.StatusOK, response)

	default:
		response := helper.APIResponse("Access Level not Found!", http.StatusNotFound, "Error", nil)
		c.JSON(http.StatusNotFound, response)
	}

}

func (h *handler) DeleteBookByID(c *gin.Context) {
	param := c.Param("bookID")

	paramUUID, _ := uuid.FromString(param)

	userData := c.MustGet("userData").(jwt.MapClaims)
	access_level := userData["access_level"].(string)

	if param == ":bookID" {
		errorMessage := "UUID is empty!"
		response := helper.APIResponse("Failed Update a Books", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !helper.IsValidUUID(param) {
		errorMessage := "UUID is not Valid!"
		response := helper.APIResponse("Failed Update a Books", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	switch access_level {
	case "client":
		response := helper.APIResponse("Access Denied!", http.StatusForbidden, "Error", "Access Denied!")
		c.JSON(http.StatusOK, response)

	case "admin":
		book, err := h.service.GetBookByID(paramUUID)
		if book == nil {
			errorMessage := "Book not Found!"
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusNotFound, "Error", errorMessage)
			c.JSON(http.StatusNotFound, response)
			return
		}

		if err != nil {
			response := helper.APIResponse("Failed to Get Book by ID", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		err = h.service.DeleteBookByID(paramUUID)
		if err != nil {
			response := helper.APIResponse("Failed to Delete Book by ID!", http.StatusBadRequest, "Error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Success Deleted a Book", http.StatusOK, "Success", nil)
		c.JSON(http.StatusOK, response)

	default:
		response := helper.APIResponse("Access Level not Found!", http.StatusNotFound, "Error", nil)
		c.JSON(http.StatusNotFound, response)
	}

}
