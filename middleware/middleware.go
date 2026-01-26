package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

// Cara pakai saat inisialisasi server:
// s := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor))
func AuthUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// Ambil metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata not found")
		}

		// Ambil authorization
		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization required")
		}

		token := strings.TrimPrefix(authHeader[0], "Bearer ")
		if token == "" {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		// Contoh validasi token (dummy)
		if token != "my-secret-token" {
			return nil, status.Error(codes.Unauthenticated, "unauthorized")
		}

		// Simpan user info ke context
		ctx = context.WithValue(ctx, "userID", "123")

		// lanjut ke handler
		return handler(ctx, req)
	}
}
