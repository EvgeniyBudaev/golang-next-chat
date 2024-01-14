CREATE TABLE room_clients (
                                id BIGSERIAL NOT NULL PRIMARY KEY,
                                room_id BIGINT NOT NULL,
                                user_id VARCHAR NOT NULL UNIQUE,
                                username VARCHAR NOT NULL,
                                CONSTRAINT fk_room_id FOREIGN KEY (room_id) REFERENCES rooms (id)
);