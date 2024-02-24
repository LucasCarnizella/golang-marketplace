CREATE ROLE market_user WITH
  LOGIN
  NOSUPERUSER
  NOCREATEROLE
  NOREPLICATION;

ALTER USER market_user WITH 
    ENCRYPTED PASSWORD 'dev123';

CREATE DATABASE marketplace WITH
    OWNER = master
    ENCODING = 'UTF8'
    LOCALE_PROVIDER = 'libc'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

GRANT ALL PRIVILEGES ON DATABASE marketplace TO market_user;

\connect marketplace;

CREATE TABLE products (
	id SERIAL PRIMARY KEY,
	name VARCHAR,
	description VARCHAR,
	price DECIMAL,
	quantity integer
)