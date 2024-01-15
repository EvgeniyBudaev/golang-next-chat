CREATE TABLE room_messages (
                                      id BIGSERIAL NOT NULL PRIMARY KEY,
                                      room_id BIGINT NOT NULL,
                                      user_id VARCHAR NOT NULL,
                                      content VARCHAR NOT NULL,
                                      CONSTRAINT fk_room_id FOREIGN KEY (room_id) REFERENCES rooms (id)
);