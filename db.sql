CREATE TABLE users (
                       id serial PRIMARY KEY ,
                       password VARCHAR ( 128 ) NOT NULL,
                       email VARCHAR ( 255 ) UNIQUE NOT NULL,
                       type VARCHAR ( 255 ) NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL
);

CREATE TABLE accounts (
                          user_id serial PRIMARY KEY,
                          password VARCHAR ( 128 ) NOT NULL,
                          email VARCHAR ( 255 ) UNIQUE NOT NULL,
                          type VARCHAR ( 255 ) NOT NULL DEFAULT ,
                          created_on TIMESTAMP NOT NULL,
                          last_login TIMESTAMP
);