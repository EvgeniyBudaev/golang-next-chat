CREATE TABLE profiles (
                          id BIGSERIAL NOT NULL PRIMARY KEY,
                          uuid UUID NOT NULL UNIQUE,
                          user_id VARCHAR NOT NULL UNIQUE,
                          created_at TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP NOT NULL
);