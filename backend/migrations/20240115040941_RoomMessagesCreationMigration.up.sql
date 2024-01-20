CREATE TABLE room_messages (
                                      id BIGSERIAL NOT NULL PRIMARY KEY,
                                      room_id BIGINT NOT NULL,
                                      user_id VARCHAR NOT NULL,
                                      type VARCHAR NOT NULL,
                                      created_at TIMESTAMP NOT NULL,
                                      updated_at TIMESTAMP NOT NULL,
                                      is_deleted bool NOT NULL,
                                      is_edited bool NOT NULL,
                                      is_joined bool NOT NULL,
                                      is_left bool NOT NULL,
                                      content VARCHAR NOT NULL,
                                      CONSTRAINT fk_room_id FOREIGN KEY (room_id) REFERENCES rooms (id)
);