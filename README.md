# simbir-go
SimbirGO is a rental service that allows users to rent vehicles such as cars, bikes or scooters and choose between different rental types. App is built as monolithic application.

## Features
- User registration and authentication

- Profile management

- Vehicles management

- Starting and stopping rents

- RESTful API design

- Database migrations

- JSON responses

## Technologies Used
- **Go**: Programming language used for the backend.

- **Gin**: Web framework for building the API.

- **PostgreSQL**: Relational database for storing data.

- **Swagger**: Documention of API.

- **JWT**: Authorization and authentication.

<!-- - **Docker**: For containerization. -->

## Running the Application
1. Clone the repository
```bash
git clone https://github.com/ursulgwopp/go-market-app
```
2. Install the required Go packages:
```bash
go mod tidy
```

3. Run the application
```bash
go run cmd/main/main.go
```
<!-- 3. Run `docker-compose.yml`
```bash
docker compose up
``` -->

4. The API will be available at `localhost:2024/swagger/index.html`

## Contributing
Contributions are welcome! If you have suggestions for improvements or want to report a bug, please open an issue or submit a pull request.