-- accounts domain

CREATE TABLE users(
  id BIGSERIAL PRIMARY KEY,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  username TEXT UNIQUE DEFAULT NULL,
  password TEXT,
  is_superuser BOOLEAN NOT NULL DEFAULT FALSE,
  is_staff BOOLEAN NOT NULL DEFAULT FALSE,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ,
  last_login_at TIMESTAMPTZ,
  group_id BIGINT REFERENCES groups(id) DEFAULT NULL
);

CREATE TYPE permission_role AS ENUM ('VIEWER', 'EDITOR', 'ADMIN');

CREATE TABLE permissions(
  id BIGSERIAL PRIMARY KEY,
  role permission_role NOT NULL,
  topic TEXT NOT NULL,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ,
  group_id BIGINT REFERENCES groups(id) DEFAULT NULL
);

CREATE TABLE groups(
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   created_at TIMESTAMPTZ,
   updated_at TIMESTAMPTZ
);

-- appointments domain TBC

-- ecommerce domain

CREATE TABLE customers(
    id BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    user_id BIGINT REFERENCES users(id) DEFAULT NULL
);

CREATE TABLE shipping_addresses(
    id BIGSERIAL PRIMARY KEY,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    zipcode TEXT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    customer_id BIGINT REFERENCES customers(id) NOT NULL
);

CREATE TYPE order_status AS ENUM ('PAID', 'PROCESSING', 'COMPLETED');

CREATE TABLE orders(
    id BIGSERIAL PRIMARY KEY,
    transaction_id TEXT NOT NULL,
    status order_status NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    customer_id BIGINT REFERENCES customers(id) NOT NULL
);