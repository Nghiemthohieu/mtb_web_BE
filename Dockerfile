# Sử dụng image Golang chính thức
FROM golang:1.24

# Set thư mục làm việc trong container
WORKDIR /app

# Sao chép go.mod và go.sum vào container
COPY go.mod go.sum ./

# Tải các dependency
RUN go mod tidy

# Sao chép toàn bộ mã nguồn vào container
COPY . .

# Build ứng dụng từ cmd/main.go
RUN go build -o main ./cmd/server

# Mở cổng (nếu cần, ví dụ 8080)
EXPOSE 8080

# Lệnh chạy khi container khởi động
CMD ["go", "run", "./cmd/server/main.go"]

