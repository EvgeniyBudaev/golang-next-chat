CREATE TABLE profiles (
                          id BIGSERIAL NOT NULL PRIMARY KEY,
                          username VARCHAR NOT NULL UNIQUE
);