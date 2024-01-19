CREATE TABLE rooms (
                       id BIGSERIAL NOT NULL PRIMARY KEY,
                       room_name VARCHAR NOT NULL UNIQUE,
                       title VARCHAR NOT NULL
);