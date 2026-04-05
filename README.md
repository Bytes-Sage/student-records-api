# Project structure

student-api/
├── main.go
├── db/
│   └── db.go           # SQLite connection + setup
├── models/
│   └── student.go      # Student struct
├── repository/
│   └── student.go      # All database queries
├── handlers/
│   └── student.go      # HTTP request handlers
├── middleware/
│   └── logger.go       # Request logging
├── go.mod
└── students.db         # Auto-generated SQLite file