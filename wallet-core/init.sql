CREATE DATABASE IF NOT EXISTS wallet;

USE wallet;

CREATE TABLE
  IF NOT EXISTS clients (
    id VARCHAR(255),
    name VARCHAR(255),
    email VARCHAR(255),
    created_at DATE
  );

CREATE TABLE
  IF NOT EXISTS accounts (
    id VARCHAR(255),
    client_id VARCHAR(255),
    balance INT,
    created_at DATE
  );

CREATE TABLE
  IF NOT EXISTS transactions (
    id VARCHAR(255),
    account_id_from VARCHAR(255),
    account_id_to VARCHAR(255),
    amount INT,
    created_at DATE
  );

INSERT INTO
  clients (id, name, email, created_at)
VALUES
  (
    '810af69b-fafd-456e-9b31-efb851b82a18',
    'João',
    'jao@j.com',
    '2021-01-01'
  ),
  (
    '6f8d1914-5f85-42cf-b37e-6b1b3e150df8',
    'Maria',
    'maria@j.com',
    '2021-01-01'
  );

INSERT INTO
  accounts (id, client_id, balance, created_at)
VALUES
  (
    "f422b886-b5e8-4d70-b085-33f69611b0f3",
    "810af69b-fafd-456e-9b31-efb851b82a18",
    1000,
    "2021-01-01"
  ),
  (
    "158a877c-efda-4415-9107-102f03949a58",
    "6f8d1914-5f85-42cf-b37e-6b1b3e150df8",
    1000,
    "2021-01-01"
  );