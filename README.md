# Template Backend Golang

## setup framework chi
Chi sangat ringan dan cocok untuk membangun layanan mikro karena desainnya yang ramah terhadap pustaka standar net/http.
- setup router
```go
r := chi.NewRouter()
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome"))
})
```
- setup handler
Memisahkan fungsi handler ke file berbeda dengan signature:
```go 
func handler (w http.ResponseWriter, r *http.Request){}
```
- setup middleware
```go
r.Use(middleware.Logger)
r.Use(middleware.Recoverer)
```

## setup framework gin
Gin adalah framework paling populer di ekosistem Go karena performanya yang cepat dan fitur bawaan yang lengkap.
- setup router
```go
r := gin.Default()
r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
})
```
- setup handler
Setup Handler Menggunakan context Gin:
```go 
func handler (c *gin.Context){}
```
- setup middleware
```go
r.Use(gin.Recovery())
r.Use(CustomMiddleware())
```

## setup framework fiber
Fiber terinspirasi dari Express (Node.js) dan dibangun di atas Fasthttp, mesin HTTP tercepat untuk Go.
- setup router
```go
app := fiber.New()
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})
```
- setup handler
Setup Handler Menggunakan context Fiber:
```go 
func handler(c *fiber.Ctx) error{}
```
- setup middleware
```go
app.Use(logger.New())
```

## setup GRPC
GRPC adalah framework untuk membangun layanan RPC (Remote Procedure Call) yang cepat dan efisien.
- setup router
```go
grpcServer := grpc.NewServer()
pb.RegisterProductServiceServer(grpcServer, &productHandler{})
```
- setup handler
Mengimplementasikan interface yang dihasilkan oleh file proto.

- setup middleware
```go
grpc.UnaryInterceptor(myInterceptor)
```

- generate file proto
running in directory root (main)
```cmd
protoc --proto_path=proto --go_out=./proto --go-grpc_out=./proto proto/product.proto
```