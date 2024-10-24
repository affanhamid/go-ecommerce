CREATE TABLE IF NOT EXISTS Users
(
  email VARCHAR(255) NOT NULL,
  country_code VARCHAR(4) NOT NULL,
  phone_number VARCHAR(15) NOT NULL,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  PRIMARY KEY (email),
  UNIQUE (phone_number)
);


CREATE TABLE IF NOT EXISTS Addresses
(
  country VARCHAR(100) NOT NULL,
  state VARCHAR(100) NOT NULL,
  city VARCHAR(100) NOT NULL,
  postcode VARCHAR(20) NOT NULL,
  address_id SERIAL PRIMARY KEY,
  user_email VARCHAR(255) NOT NULL,
  FOREIGN KEY (user_email) REFERENCES Users(email)
);