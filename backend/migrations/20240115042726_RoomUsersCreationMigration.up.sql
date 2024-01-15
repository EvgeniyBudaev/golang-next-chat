CREATE TABLE room_users (
                              id BIGSERIAL NOT NULL PRIMARY KEY,
                              room_id BIGINT NOT NULL,
                              user_id VARCHAR NOT NULL UNIQUE,
                              CONSTRAINT fk_room_id FOREIGN KEY (room_id) REFERENCES rooms (id)
);