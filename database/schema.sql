-- Drop existing tables to start fresh
DROP TABLE IF EXISTS transaction_details, transactions, products, categories CASCADE;

-- Create categories table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    total_amount INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create transaction_details table
CREATE TABLE IF NOT EXISTS transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id) ON DELETE SET NULL,
    quantity INTEGER NOT NULL,
    subtotal INTEGER NOT NULL
);

-- Seed initial data
INSERT INTO categories (id, name) VALUES (1, 'Minuman'), (2, 'Makanan') ON CONFLICT (id) DO NOTHING;
INSERT INTO products (name, price, stock, category_id) VALUES ('Kopi Tubruk', 15000, 100, 1), ('Kopi Susu', 20000, 50, 1), ('Nasi Goreng', 25000, 30, 2);
