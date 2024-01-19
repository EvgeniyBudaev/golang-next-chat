CREATE TABLE profiles (
                          id BIGSERIAL NOT NULL PRIMARY KEY,
                          user_id VARCHAR NOT NULL UNIQUE,
                          username VARCHAR NOT NULL UNIQUE,
                          first_name VARCHAR NOT NULL,
                          last_name VARCHAR,
                          email VARCHAR NOT NULL UNIQUE,
                          created_at TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP NOT NULL,
                          is_deleted bool NOT NULL,
                          is_enabled bool NOT NULL
);
