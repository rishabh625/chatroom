CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  admin BOOLEAN NOT NULL,
  enabled BOOLEAN NOT NULL,
  username VARCHAR(255) UNIQUE,
  password VARCHAR(255),
  email VARCHAR (50) UNIQUE
);

CREATE TABLE chathistory (
	id serial PRIMARY KEY,
	data text,
    room VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);