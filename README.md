# Restaurant Management System

A comprehensive backend system for restaurant management built with Go, Gin framework, MongoDB, and Docker.

## 🚀 Features

- **Menu Management**: Create, update, delete, and view menu items
- **Order Processing**: Handle customer orders and order status tracking
- **Table Management**: Manage restaurant tables and reservations
- **Staff Management**: Handle staff information and roles
- **Inventory Tracking**: Monitor ingredient stock levels
- **Customer Management**: Store customer information and preferences
- **Payment Processing**: Handle various payment methods
- **Analytics & Reporting**: Generate sales and performance reports

## 🛠 Tech Stack

- **Backend Framework**: Go with Gin Gonic
- **Database**: MongoDB
- **Containerization**: Docker & Docker Compose
- **Authentication**: JWT (JSON Web Tokens)
- **API Documentation**: Swagger/OpenAPI
- **Environment Management**: Viper
- **Logging**: Logrus
- **Validation**: Go Playground Validator

## 📋 Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- MongoDB (if running locally without Docker)
- Git

## 🔧 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/restaurant-management-system.git
cd restaurant-management-system
```

### 2. Environment Configuration

Create a `.env` file in the root directory:

```env
# Server Configuration
PORT=8080
GIN_MODE=release

# Database Configuration
MONGO_URI=mongodb://mongodb:27017
MONGO_DATABASE=restaurant_db

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRE_HOURS=24

# Other Configuration
LOG_LEVEL=info
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
```

### 3. Docker Setup (Recommended)

#### Using Docker Compose

```bash
# Build and start all services
docker-compose up --build

# Run in detached mode
docker-compose up -d --build

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

#### Manual Docker Setup

```bash
# Build the Go application
docker build -t restaurant-api .

# Run MongoDB
docker run -d --name mongodb -p 27017:27017 mongo:latest

# Run the application
docker run -d --name restaurant-api -p 8080:8080 --link mongodb:mongodb restaurant-api
```

### 4. Local Development Setup

```bash
# Install dependencies
go mod download

# Install Air for hot reloading (optional)
go install github.com/cosmtrek/air@latest

# Run with hot reloading
air

# Or run normally
go run main.go
```

## 📁 Project Structure

```
restaurant-management-system/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   ├── auth.go
│   │   ├── menu.go
│   │   ├── orders.go
│   │   ├── tables.go
│   │   └── staff.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   ├── models/
│   │   ├── user.go
│   │   ├── menu.go
│   │   ├── order.go
│   │   └── table.go
│   ├── repositories/
│   │   ├── user_repo.go
│   │   ├── menu_repo.go
│   │   └── order_repo.go
│   ├── routes/
│   │   └── routes.go
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── menu_service.go
│   │   └── order_service.go
│   └── utils/
│       ├── database.go
│       ├── jwt.go
│       └── validator.go
├── pkg/
│   └── logger/
│       └── logger.go
├── docs/
│   ├── swagger.json
│   └── swagger.yaml
├── docker-compose.yml
├── Dockerfile
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## 🔌 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout
- `GET /api/v1/auth/profile` - Get user profile

### Menu Management
- `GET /api/v1/menu` - Get all menu items
- `GET /api/v1/menu/:id` - Get specific menu item
- `POST /api/v1/menu` - Create new menu item
- `PUT /api/v1/menu/:id` - Update menu item
- `DELETE /api/v1/menu/:id` - Delete menu item

### Order Management
- `GET /api/v1/orders` - Get all orders
- `GET /api/v1/orders/:id` - Get specific order
- `POST /api/v1/orders` - Create new order
- `PUT /api/v1/orders/:id` - Update order
- `DELETE /api/v1/orders/:id` - Cancel order

### Table Management
- `GET /api/v1/tables` - Get all tables
- `GET /api/v1/tables/:id` - Get specific table
- `POST /api/v1/tables` - Add new table
- `PUT /api/v1/tables/:id` - Update table
- `DELETE /api/v1/tables/:id` - Remove table

## 📖 API Documentation

Access the Swagger documentation at: `http://localhost:8080/swagger/index.html`

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./internal/controllers
```

## 🚀 Deployment

### Docker Production Deployment

```bash
# Build production image
docker build -t restaurant-api:prod -f Dockerfile.prod .

# Run with docker-compose for production
docker-compose -f docker-compose.prod.yml up -d
```

### Environment Variables for Production

```env
GIN_MODE=release
PORT=8080
MONGO_URI=mongodb://your-mongo-host:27017
JWT_SECRET=your-production-jwt-secret
LOG_LEVEL=warn
```

## 📊 Database Schema

### Collections

1. **users** - User authentication and profiles
2. **menu_items** - Restaurant menu items
3. **orders** - Customer orders
4. **tables** - Restaurant table management
5. **staff** - Staff information
6. **inventory** - Ingredient inventory
7. **payments** - Payment transactions

## 🔐 Security Features

- JWT-based authentication
- Password hashing with bcrypt
- Input validation and sanitization
- CORS protection
- Rate limiting
- SQL injection prevention
- XSS protection

## 📈 Monitoring & Logging

- Structured logging with Logrus
- Request/response logging middleware
- Error tracking and reporting
- Performance monitoring
- Health check endpoints

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Your Name** - *Initial work* - [YourGitHub](https://github.com/yourusername)

## 🙏 Acknowledgments

- Gin Gonic framework community
- MongoDB documentation
- Docker community
- Go community

## 📞 Support

For support, email support@restaurant-system.com or create an issue on GitHub.

## 🔄 Changelog

### v1.0.0 (Latest)
- Initial release
- Basic CRUD operations
- Authentication system
- Docker support
- API documentation

---

**Happy Coding! 🍽
