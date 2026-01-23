# Cashier System REST API

A simple REST API for a cashier system built with Go.

**Production URL:** [https://learn-api-golang-production.up.railway.app/](https://learn-api-golang-production.up.railway.app/)

## Features

- **Product Management**: CRUD (Create, Read, Update, Delete) for products.
- **Category Management**: CRUD (Create, Read, Update, Delete) for categories.
- **In-Memory Storage**: Data is stored in memory for demonstration purposes.

## API Endpoints

### Categories

- `GET /api/categories`: Get all categories.
- `POST /api/categories`: Create a new category.
- `GET /api/categories/{id}`: Get category by ID.
- `PUT /api/categories/{id}`: Update category by ID.
- `DELETE /api/categories/{id}`: Delete category by ID.

### Products

- `GET /api/produk`: Get all products.
- `POST /api/produk`: Create a new product.
- `GET /api/produk/{id}`: Get product by ID.
- `PUT /api/produk/{id}`: Update product by ID.
- `DELETE /api/produk/{id}`: Delete product by ID.

## How to Run Locally

1. Make sure you have Go installed.
2. Clone the repository.
3. Run the following command:
   ```bash
   go run main.go
   ```
4. The server will start at `http://localhost:8080`.

## Deployment

This project is configured for deployment on **Railway**.
