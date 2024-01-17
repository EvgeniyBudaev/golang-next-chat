CREATE TABLE rooms (
                       id BIGSERIAL NOT NULL PRIMARY KEY,
                       uuid UUID NOT NULL UNIQUE
);