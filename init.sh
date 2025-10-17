#!/bin/bash

# Create project root
mkdir -p sp500-shariah && cd sp500-shariah

# -----------------------------
# 1️⃣ job-service
# -----------------------------
mkdir -p job-service/{cmd,internal/{domain/{stock,earnings,notification},application,infrastructure/{yahoo,persistence,mail,scheduler,logger},interfaces/job},configs}

# Main entry
cat <<'EOF' > job-service/cmd/main.go
package main

import "fmt"

func main() {
    fmt.Println("Job Service started...")
}
EOF

# Go module
cd job-service
go mod init github.com/rama378/playground-go/sp500-shariah/job-service
cd ..

# -----------------------------
# 2️⃣ api-service
# -----------------------------
mkdir -p api-service/{cmd,internal/{domain/{stock},application,infrastructure/{persistence,auth,logger},interfaces/{api,middleware}},configs}

# Main entry
cat <<'EOF' > api-service/cmd/main.go
package main

import "fmt"

func main() {
    fmt.Println("API Service started...")
}
EOF

# Go module
cd api-service
go mod init github.com/rama378/playground-go/sp500-shariah/api-service
cd ..

# -----------------------------
# 3️⃣ shared library
# -----------------------------
mkdir -p shared/{config,logger,utils}

# Example shared logger
cat <<'EOF' > shared/logger/logger.go
package logger

import "log"

func Info(msg string) {
    log.Println("[INFO]", msg)
}
EOF

cd shared
go mod init github.com/rama378/playground-go/sp500-shariah/shared
cd ..

# -----------------------------
# 4️⃣ Root Docker Compose (for later)
# -----------------------------
cat <<'EOF' > docker-compose.yml
version: "3.9"

services:
  job-service:
    build: ./job-service
    command: go run ./cmd
    volumes:
      - ./job-service:/app
    working_dir: /app

  api-service:
    build: ./api-service
    command: go run ./cmd
    volumes:
      - ./api-service:/app
    working_dir: /app
    ports:
      - "8080:8080"
EOF
