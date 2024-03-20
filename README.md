### Go CRUD API Application with Chi Router and SQLite3

This Go application offers a versatile CRUD API, equipped with the Chi router for streamlined routing and SQLite3 for efficient database management. The architecture is designed with an interface layer, facilitating seamless transitions between different database systems.

#### Features:
- **CRUD Operations:** Perform Create, Read, Update, and Delete operations on resources.
- **Chi Router:** Utilize the lightweight and flexible Chi router for efficient request handling and routing.
- **SQLite3 Integration:** Benefit from SQLite3 as the underlying database system, ensuring reliability and ease of use.
- **Interface Layer:** Implement an interface layer to support easy switching between various database implementations, enhancing scalability and flexibility.
- **Logging:** Enhance monitoring and troubleshooting capabilities with comprehensive logging functionality.

#### Usage:
1. Clone the repository.
2. Configure the database connection parameters.
3. Build and run the application.
4. Access the API endpoints to perform CRUD operations.

#### Dependencies:
- [Chi Router](https://github.com/go-chi/chi)
- [SQLite3 Driver](https://github.com/mattn/go-sqlite3)

#### Getting Started:
```bash
git clone https://github.com/schadokar/go-crud-app.git
cd go-crud-app
go build
./gocrudapp
```

#### API Endpoints:
- `GET /v1/user`: Retrieve all users.
- `GET /v1/user/{id}`: Retrieve a specific user by ID.
- `POST /v1/user`: Create a new user.
- `PUT /v1/user/{id}`: Update an existing user.
- `DELETE /v1/user/{id}`: Delete a user.
- `DELETE /v1/user`: Delete all users.

#### API Documentation:
Copy the `gocrud.postman_collection.json` content and import in postman for api collection.

#### Configuration:
Ensure to configure the database connection parameters in the application before running.

Create a new file `.env` for environment variable and add `DB_NAME` key.
```
DB_NAME="TestDB"
```

#### Contributing:
Contributions are welcome! Please open an issue or submit a pull request with any enhancements or fixes.

#### License:
This project is licensed under the [MIT License](https://github.com/schadokar/go-crud-app?tab=MIT-1-ov-file).

#### Acknowledgments:
Special thanks to the creators and maintainers of Chi Router and SQLite3 Driver for their invaluable contributions.

#### Contact:
For any inquiries or feedback, please reach out to [Shubham](https://schadokar.dev).

Connect here:
- [Youtube Practicego](https://youtube.com/@practicego)
- [X](https://twitter.com/schadokar1)
- [Linkedin](https://linkedin.com/in/schadokar)

---