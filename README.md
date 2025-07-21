# 📲 Send SMS Service

This microservice is responsible for sending SMS notifications, such as appointment confirmations, cancellations, and alerts. Built with **Go**, it follows a layered and modular architecture. The service is fully containerized with Docker and supports automated deployment via GitHub Actions to an EC2 instance.

---

## 🧩 Features

- Send SMS messages via integrated providers  
- Modular Go structure (config, middleware, services, etc.)  
- JWT-protected endpoints  
- Docker-ready  
- GitHub Actions for CI/CD  
- AWS EC2 deployable

---

## ⚙️ Tech Stack

- Go 1.20+  
- REST API  
- Twilio or custom SMS providers (extensible via `service/`)  
- Docker  
- GitHub Actions  

---

## 📁 Main Components

- `cmd/main.go`: Application entry point  
- `internal/config`: Loads environment variables  
- `internal/controller`: Handles API requests  
- `internal/middleware`: Middleware for JWT auth  
- `internal/service`: SMS business logic  
- `internal/repository`: Abstraction for provider/storage  
- `.env`: Environment variable config  
- `Dockerfile`: Container build setup  

---

## 🔐 Environment Variables (`.env`)

```env
PORT=4008
SMS_API_KEY=your_twilio_api_key
SMS_API_SECRET=your_twilio_secret
SMS_SENDER=+1234567890
JWT_SECRET=supersecretkey

🐳 Run with Docker
1. Build the Docker image

docker build -t send-sms .

2. Run the container

docker run -d -p 4008:4008 --env-file .env send-sms

🔌 API Endpoints
Method	Route	Description
POST	/sms/send	Send an SMS message

    Example JSON payload:

{
  "to": "+593987654321",
  "message": "Your appointment has been confirmed."
}

All endpoints require:

Authorization: Bearer <jwt_token>

☁️ EC2 Deployment
1. SSH into EC2

ssh -i key.pem ec2-user@<YOUR_EC2_PUBLIC_IP>

2. Pull Docker image

docker pull jeffri1997/send-sms:qa

3. Run the service

docker run -d -p 4008:4008 --env-file .env jeffri1997/send-sms:qa

✅ Make sure port 4008 is open in the EC2 security group.
🤖 GitHub Actions CI/CD

Automatically builds and pushes to Docker Hub on changes to qa branch:

name: Build and Push SMS Service

on:
  push:
    branches: [qa]

jobs:
  build-and-push:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Docker Login
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build and Push Docker Image
        run: |
          docker build -t jeffri1997/send-sms:qa .
          docker push jeffri1997/send-sms:qa

🧪 Testing

If test files are added, run:

go test ./...

🧰 Use Cases

    Appointment confirmation via SMS

    Cancellation alerts

    System-wide broadcast notifications via phone

👤 Author

Jefferson Marcalla
GitHub: @Jeff97ares