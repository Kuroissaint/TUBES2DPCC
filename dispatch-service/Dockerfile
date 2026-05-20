# Stage 1: Builder
FROM golang:1.25-alpine AS builder

# Install git jika ada dependency yang memerlukannya
RUN apk add --no-cache git

WORKDIR /app

# Copy file dependency agar layer cache efisien
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code ke dalam container
COPY . .

# Build binary aplikasi
# CGO_ENABLED=0 agar binary bersifat statis dan bisa jalan di alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Stage 2: Final Image
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy hasil build dari stage builder
COPY --from=builder /main .

# Port yang dibuka sesuai jatah service Location
EXPOSE 8002

# Jalankan aplikasi saat container dimulai
CMD ["./main"]