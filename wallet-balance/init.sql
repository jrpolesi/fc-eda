CREATE DATABASE IF NOT EXISTS balance;

USE balance;

CREATE TABLE
  IF NOT EXISTS accounts (
    id VARCHAR(255),
    balance INT,
    created_at DATE,
    updated_at DATE
  );

INSERT INTO
  accounts (id, balance, created_at, updated_at)
VALUES
  (
    "f422b886-b5e8-4d70-b085-33f69611b0f3",
    1000,
    "2021-01-01",
    "2021-01-01"
  ),
  (
    "158a877c-efda-4415-9107-102f03949a58",
    1000,
    "2021-01-01",
    "2021-01-01"
  );