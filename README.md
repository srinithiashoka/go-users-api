# Go Users API 🚀

A RESTful backend service built with **Go (Golang)** and **Fiber** to manage users with their Date of Birth (DOB).  
The API dynamically calculates and returns the user’s age when fetching data.

---

## 📌 Features

- Create, Read, Update, Delete (CRUD) users
- Store Date of Birth in database
- Dynamically calculate age using Go `time` package
- Clean layered architecture (Handler, Service, Repository)
- MySQL database integration
- Logging using Uber Zap
- RESTful API design

---

## 🗂️ Project Structure

go-users-api/
├── cmd/
│ └── server/
│ └── main.go
├── config/
│ └── db.go
├── db/
│ └── migrations/
│ └── users.sql
├── internal/
│ ├── handler/
│ │ └── user_handler.go
│ ├── repository/
│ │ └── user_repository.go
│ ├── service/
│ │ └── age.go
│ ├── routes/
│ │ └── routes.go
│ ├── models/
│ │ └── user.go
│ └── logger/
│ └── logger.go
├── .gitignore
├── go.mod
├── go.sum
└── README.md

yaml
Copy code

---

## 🔧 Tech Stack

- **Go (Golang)**
- **Fiber** – Web framework
- **MySQL** – Database
- **Uber Zap** – Logging
- **Go standard library** (`time`, `database/sql`)

---

## 🗄️ Database Schema

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);
⚙️ Environment Variables (Optional)
Create a .env file:

env
Copy code
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=usersdb
▶️ Run the Application
1️⃣ Install dependencies
bash
Copy code
go mod tidy
2️⃣ Start the server
bash
Copy code
go run cmd/server/main.go
Server will run at:

arduino
Copy code
http://localhost:8080
🔄 API Endpoints
➕ Create User
POST /users

json
Copy code
{
  "name": "Alice",
  "dob": "1990-05-10"
}
Response

json
Copy code
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 34
}
📄 Get User by ID
GET /users/:id

json
Copy code
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 34
}
📋 List All Users
GET /users

json
Copy code
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  }
]
✏️ Update User
PUT /users/:id

json
Copy code
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
❌ Delete User
DELETE /users/:id

Response

css
Copy code
204 No Content
🧠 Age Calculation Logic
go
Copy code
func CalculateAge(dob time.Time) int {
    now := time.Now()
    age := now.Year() - dob.Year()

    if now.Month() < dob.Month() ||
       (now.Month() == dob.Month() && now.Day() < dob.Day()) {
        age--
    }
    return age
}
👩‍💻 Author
Srinithi A
GitHub: https://github.com/srinithiashoka

yaml
Copy code

---

## ✅ FINAL COMMANDS (RUN NOW)

```bash
git add README.md
git commit -m "Fix README formatting"
git push
