package routes

import (
	"api-backend-go/handlers"
	"api-backend-go/handlers/grpchandler"
	"api-backend-go/middleware"
	"api-backend-go/proto/pb"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// server api Native golang
func ServerApiNative() {
	mux := http.NewServeMux()

	mux.Handle("GET /user", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserGet)))
	mux.Handle("POST /user", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserPost)))
	mux.Handle("PUT /user/{id}", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserPut)))
	mux.Handle("DELETE /user/{id}", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserDelete)))

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("error server")
	}
}

// server api chi golang
func ServerApiChi() {
	r := chi.NewRouter()
	// middleware implem
	r.With(middleware.ChiMiddleware).Get("/", handlers.HandlerExampleChi)
	r.Post("/", handlers.HandlerExampleChi)
	r.Put("/", handlers.HandlerExampleChi)
	r.Delete("/", handlers.HandlerExampleChi)

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}
}

// server api GIN Golang
func ServerApiGin() {
	r := gin.Default()
	// middleware implem
	r.GET("/ping", middleware.GinMiddleware(), handlers.HandlerExampleGin)
	r.POST("/ping", handlers.HandlerExampleGin)
	r.PUT("/ping", handlers.HandlerExampleGin)
	r.DELETE("/ping", handlers.HandlerExampleGin)

	fmt.Println("server running on port 8080")
	r.Run(":8080")
}

// server api framework fiber golang
func ServerApiFiber() {
	app := fiber.New()
	// middleware implem
	app.Get("/ping", middleware.FiberMiddleware, handlers.HandlerExampleFiber)
	app.Post("/ping", handlers.HandlerExampleFiber)
	app.Put("/ping", handlers.HandlerExampleFiber)
	app.Delete("/ping", handlers.HandlerExampleFiber)

	fmt.Println("server running on port 8080")
	log.Fatal(app.Listen(":3000"))
}

// server api GRPC
func ServerApiGRPCProduct(db *gorm.DB) {
	// gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ProductService := grpchandler.NewProductServiceServer(db)
	grpcServer := grpc.NewServer()
	// implem middleware
	// grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthUnaryInterceptor()))
	pb.RegisterProductServiceServer(grpcServer, &ProductService)

	log.Printf("Product Service is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func ServerApiGRPCUser(db *gorm.DB) {
	// gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// implem middleware
	// grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthUnaryInterceptor()))
	pb.RegisterUserServiceServer(grpcServer, &grpchandler.UserServiceServer{DB: db})

	log.Printf("User Service is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
