package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// Cara pakai:
// mux.Handle("/", LoggerMiddleware(myHandler))
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Native: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Lanjut ke handler berikutnya
	})
}

// Cara pakai:
// r := chi.NewRouter()
// r.Use(ChiMiddleware)
func ChiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Chi: Request started")
		next.ServeHTTP(w, r)
		fmt.Println("Chi: Request finished")
	})
}

// Cara pakai:
// r := gin.Default()
// r.Use(GinMiddleware())
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Gin: Before Request")

		c.Next() // Eksekusi handler utama

		fmt.Println("Gin: After Request")
	}
}

// Cara pakai:
// app := fiber.New()
// app.Use(FiberMiddleware)
func FiberMiddleware(c *fiber.Ctx) error {
	fmt.Println("Fiber: Logic before handler")

	err := c.Next() // Lanjut ke handler berikutnya

	fmt.Println("Fiber: Logic after handler")
	return err
}
