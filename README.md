# API-HUB
A simple Go-based microservice gateway that demonstrates authentication, external API integration, SFTP file handling, MongoDB connection, and Swagger documentation.

## ✨ Features

JWT-based authentication (/login)

Protected routes (/protected/*)

Fetching and saving LMS data from mock API to MongoDB

Reading file content from SFTP server

Rate limiting & logging middleware

Swagger API docs at /swagger/index.html

## ⚡ Tech Stack

Golang (Gin framework)

MongoDB Atlas (Cloud database)

JWT for secure authentication

SFTP for file operations

Swagger for API documentation

## 📂 Project Structure

api-hub/

├── cmd/

│   └── main.go 

├── internal/

│   ├── config/              

│   ├── handlers/            

│   ├── middleware/           

│   ├── models/             

│   └── services/             

├── docs/                    

└── go.mod

## ✅ How to Run

Clone the repo

Create `.env` file with:

`JWT_SECRET=""`


`MONGO_URI=mongodb+srv://...your_atlas_uri`

Run the app:

`go run cmd/main.go`

Open Swagger UI:

http://localhost:8080/swagger/index.html

##  🌐 API Endpoints

Public

    POST /login → Get JWT token

Protected (Require Bearer Token)

    GET /protected/lms → Fetch and store LMS data

    GET /protected/sftp → Download file from SFTP


📄 Authorize with JWT in Swagger

    Call /login and copy the token

    Click Authorize in Swagger

    Paste: Bearer <your_token>

    Now test protected endpoints

##  🏙️ Demo Credentials

username: admin

password: 1234

## 🚀 Notes

LMS API uses a mock server

MongoDB connection is cloud-based

SFTP uses Rebex test server

## 🎓 License

MIT

