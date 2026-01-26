package handlers

import (
	"api-backend-go/dto"
	"api-backend-go/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
)

// handler Native GO
func HandlerExample(w http.ResponseWriter, r *http.Request) {
	// 1. Path Parameter
	id := r.PathValue("id")

	// 2. Query Parameter
	role := r.URL.Query().Get("role")

	// 3. Body (JSON)
	var req dto.RequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "ID: %s, Role: %s, Name: %s", id, role, req.Name)

	// response success
	utils.ResponseSuccess(w, http.StatusOK, "success", req)
}

// Handler With Chi
func HandlerExampleChi(w http.ResponseWriter, r *http.Request) {
	// 1. Path Parameter
	id := chi.URLParam(r, "id")

	// 2. Query Parameter
	role := r.URL.Query().Get("role")

	// 3. Body (JSON)
	var req dto.RequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "ID: %s, Role: %s, Name: %s", id, role, req.Name)

	// response success
	utils.ResponseSuccess(w, http.StatusOK, "success", req)
}

// Handler With Gin
func HandlerExampleGin(c *gin.Context) {

	// 1. Path Parameter
	id := c.Param("id")

	// 2. Query Parameter
	role := c.DefaultQuery("role", "guest")

	// 3. Body (JSON Binding)
	var req dto.RequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id, "role": role, "user": req})
}

// Handler With Fiber
func HandlerExampleFiber(c *fiber.Ctx) error {
	// 1. Path Parameter
	id := c.Params("id")

	// 2. Query Parameter
	role := c.Query("role")

	// 3. Body (Parsing)
	user := new(dto.RequestDTO)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"role": role,
		"user": user,
	})
}
