CREATE TABLE room_client_messages (
                                      id BIGSERIAL NOT NULL PRIMARY KEY,
                                      room_id BIGINT NOT NULL,
                                      client_id BIGINT NOT NULL,
                                      content VARCHAR NOT NULL,
                                      CONSTRAINT fk_room_id FOREIGN KEY (room_id) REFERENCES rooms (id),
                                      CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES room_clients (id)
);