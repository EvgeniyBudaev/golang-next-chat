CREATE TABLE rooms (
                          id BIGSERIAL NOT NULL PRIMARY KEY,
                          name VARCHAR NOT NULL UNIQUE
);