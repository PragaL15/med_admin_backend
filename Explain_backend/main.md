## main.go

 `import (
	"log"
	"net/http"
	"github.com/PragaL15/med_admin_backend/database"
	"github.com/PragaL15/med_admin_backend/src/routers/user"
	"github.com/gorilla/handlers"
)`

**import** --> keyword used to import the external packages and libraries.

**log** --> To print the logging messages.

**net/http** --> To handle the HTTP request and response.

**github.com/...** --> Custom packages which is our project components which we've to use.

---
##### Main Function --> Entry point of the program
---
##### Database Initilization 

`	db, err := database.InitializeDB()`
 
  This is to initilize the database connection.

---
 
 ##### Closing the database

 `	defer func() {
		sqlDB, err := db.DB()
    		if err != nil {
			log.Fatalf("Failed to get raw database connection: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()
`
1. **defer** --> Will excecute once the surrounding function is completed. It's commonly used for clean up task and closing database connection.
2. `sqlDB, err := db.DB()` 
 - `db` is instance of ***gorm.DB** object.
 - `DB()` is a method of `gorm.DB`, provides access to the underlying `*sql.DB` instance.
 - `sqlDB` is a raw database connection pool that GORM uses internally for exceuting SQL queries.
- `sqlDB.Close()` is must because closing connection is important to release resources like open connections and leads to connection leaks.

---
##### Router setup

`	router := routers.SetupRoutes(db)`

- **routers** is the package which has all the routes along with their handlers.

- **route** is variable that store the router instance.

---

##### CORS Middleware

`	corsOrigin := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"})
`
1. `corsOrigin`, `corsMethods`, `corsHeaders`: Variables to configure CORS policies.

2. `handlers.AllowedOrigins`, `handlers.AllowedMethods`, `handlers.AllowedHeaders`: Functions from the gorilla/handlers package that specify allowed origins, HTTP methods, and headers.
- In this , Handlers are package which have the logics and behavior of each function.
- `AllowedOrigins`, `AllowedMethods`, `AllowedHeaders` are the inbuilt functions of `cors.go`. 

`corsMiddleware := handlers.CORS(corsOrigin, corsMethods, corsHeaders)
` 
This will combine the CORS policies into middleware.

`corsMiddleware(router)` Applies the CORS middleware to the router for handling requests.
