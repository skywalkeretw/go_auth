CREATE TABLE users (
                       id serial PRIMARY KEY ,
                       password VARCHAR ( 128 ) NOT NULL,
                       email VARCHAR ( 255 ) UNIQUE NOT NULL,
                       first_name varchar (255) NOT NUll,
                       last_name varchar (255) NOT NUll,
                       type VARCHAR ( 255 ) NOT NULL DEFAULT 'user',
                       confirmed BOOLEAN NOT NULL DEFAULT false ,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL
);