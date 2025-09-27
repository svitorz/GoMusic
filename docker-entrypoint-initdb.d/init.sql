CREATE DATABASE IF NOT EXISTS gomusic;

CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  birthdate DATE,
);

CREATE TABLE IF NOT EXISTS albums(
  name VARCHAR(255),
  description text,
  user_id INT NOT NULL,
  CONSTRAIT FK_ALBUMS_USERS FOREIGN KEY(user_id) REFERENCES users(id)
);
