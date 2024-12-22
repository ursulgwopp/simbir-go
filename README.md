# SimbirGO

SimbirGO is a vehicle rental service application that allows users to rent cars, bikes, and scooters. Users can choose from various rental types, making it easy to find the right option for their needs. The application is built as a monolithic architecture, simplifying deployment and management.

## Features

- **User  Registration and Authentication**: Securely register and log in to your account.
- **Profile Management**: Easily manage your user profile and account settings.
- **Vehicle Management**: Browse, add, update, and delete vehicles available for rent.
- **Rental Operations**: Start and stop rentals with a simple interface.
- **RESTful API Design**: Interact with the application through a well-defined API.
- **Database Migrations**: Manage database schema changes effectively.
- **JSON Responses**: Receive data in a structured JSON format.

## Technologies Used

- **Go**: The programming language used for the backend.
- **Gin**: A web framework for building the API.
- **PostgreSQL**: A relational database for data storage.
- **Swagger**: API documentation for easy reference.
- **JWT**: For authorization and authentication.

## Getting Started

Follow these steps to run the application locally:

1. **Clone the Repository**:
```bash
git clone https://github.com/ursulgwopp/simbir-go
```
2. **Install Required Go Packages:**
```bash
go mod tidy
```
3. **Run the Application:**
```bash
go run cmd/main/main.go
```
4. **Access the API:**
The API will be available at http://localhost:2024/swagger/index.html.

## Contributing

Contributions are welcome! If you have suggestions for improvements or want to report a bug, please open an issue or submit a pull request.