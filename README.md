# Cashier System REST API

A simple REST API for a cashier system built with Go.

**Production URL:** [https://learn-api-golang-production.up.railway.app/](https://learn-api-golang-production.up.railway.app/)

## Features

- **Product Management**: CRUD with **Search By Name** (Optional).
- **Category Management**: CRUD for product categorization.
- **Checkout System**: Multi-item checkout with automatic stock deduction and transaction recording.
- **Sales Reports**: Summary of today's revenue, transactions, and best-selling products.
- **Database Storage**: Integrated with PostgreSQL (Supabase).

## API Endpoints

### Categories

- `GET /api/categories`: Get all categories.
- `POST /api/categories`: Create a new category.
- `GET /api/categories/{id}`: Get category by ID.

### Products

- `GET /api/produk`: Get all products.
- `GET /api/produk?name=indom`: Search products by name (case-insensitive).
- `POST /api/produk`: Create a new product (include `stock` field).
- `GET /api/produk/{id}`: Get product by ID.

### Checkout

- `POST /api/checkout`: Perform a checkout.
  **Payload:**
  ```json
  {
    "items": [
      { "product_id": 1, "quantity": 2 },
      { "product_id": 2, "quantity": 1 }
    ]
  }
  ```

### Reports

- `GET /api/report/hari-ini`: Get today's sales summary.
- `GET /api/report?start_date=2026-01-01&end_date=2026-02-01`: Get sales summary for a specific date range.

## How to Run Locally

1. Make sure you have Go installed.
2. Clone the repository.
3. Configure `.env` file with your `PORT` and `DB_CONN`.
4. Run the following command:
   ```bash
   go run main.go
   ```
5. The server will start at `http://localhost:8080`.

## Deployment

This project is configured for deployment on **Railway**.
