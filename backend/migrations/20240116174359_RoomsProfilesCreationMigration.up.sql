CREATE TABLE rooms_profiles (
                                id BIGSERIAL NOT NULL PRIMARY KEY,
                                 room_id BIGINT NOT NULL,
                                 profile_id BIGINT NOT NULL,
                                 FOREIGN KEY (room_id) REFERENCES rooms(id),
                                 FOREIGN KEY (profile_id) REFERENCES profiles(id)
);