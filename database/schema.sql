-- Drop existing tables to start fresh
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

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
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);

-- Seed initial data
INSERT INTO categories (id, name) VALUES (1, 'Minuman'), (2, 'Makanan') ON CONFLICT (id) DO NOTHING;
INSERT INTO products (name, price, category_id) VALUES ('Kopi Tubruk', 15000, 1), ('Kopi Susu', 20000, 1), ('Nasi Goreng', 25000, 2);
