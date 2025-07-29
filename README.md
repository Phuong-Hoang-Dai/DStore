# 🛒 DStore API – Food Ordering System (Golang | Microservice | Clean Architecture)

## 🚀 Overview

**DStore** is a RESTful API service for managing food orders, built with **Go** using a **Microservice architecture** and inspired by the **Clean Architecture** principles. It is designed for scalability, maintainability, and clean separation of concerns.

This project demonstrates my skills in backend architecture, domain-driven design, and practical implementation using popular tools and patterns.

---

## 🧱 Project Structure

The project is organized into services:

- `product` – product catalog management  
- `order` – customer order processing  
- `user` – authentication(with JWT) & user management

Key design principles:
- SOLID
- Dependency Inversion via interfaces
- Domain-driven modularity
- Testability and maintainability

---

## 🔧 Tech Stack

- **Go**
- **Gin** – Web framework
- **GORM** – ORM for database interaction
- **MySQL** – Relational database
- **Docker** – Containerization
- **JWT** – Authentication

---

## ⚙️ Getting Started

### 🚢 Run with Docker

```bash
# 1. Clone the repository
git clone https://github.com/Phuong-Hoang-Dai/DStore.git

# 2. Start services
docker-compose up --build
```
📚 API Endpoints

- POST	 /api/v1/user	register new user
- PUT    /api/v1/user/:id update info user
- DELETE /api/v1/user/:id soft delete user
- GET    /api/v1/user/:id get user by id
- GET    /api/v1/user?limit=?&offset=? get list users
- POST	 /api/v1/user/login	user login to get token JWT

- GET	   /api/v1/product	get product list
- POST /api/v1/product/getstock decrease quantity for order
- POST /api/v1/product/restorestock increase quantity for canceled order
- GET /api/v1/product/:id get product info by id
- GET /api/v1/product?limit=?&offset=? get list products
- PUT /api/v1/product/:id update info product
- DELETE /api/v1/product/:id soft delete product

- POST	 /api/v1/order	Place a new order with userId in token JWT 
- PUT	 /api/v1/order/:id	update state of order
- GET /api/v1/order/history/:id?limit=?&offset=? get order history by id user
- GET /api/v1/order/:id get order by id order
- GET /api/v1/order?limit=?&offset=? get list orders for admin
- DELETE /api/v1/order/:id cancel order

🔐 Authentication
- Login returns a JWT token
- Protected routes require header:
  - Authorization: Bearer <your-token>
  - Middleware injects userID into request context
- Protected password with bcrypt

🧪 Testing
- Unit tests for UseCases 
- Mocked database interactions
- Api test with post man

📈 Future Improvements
 - Redis for caching
 - CI/CD with GitHub Actions
 - Message broker with rabbitmq
 - Microservice
