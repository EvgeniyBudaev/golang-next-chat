CREATE TABLE rooms (
                       id BIGSERIAL NOT NULL PRIMARY KEY,
                       uuid UUID NOT NULL UNIQUE,
                       room_name VARCHAR NOT NULL UNIQUE,
                       title VARCHAR NOT NULL
);