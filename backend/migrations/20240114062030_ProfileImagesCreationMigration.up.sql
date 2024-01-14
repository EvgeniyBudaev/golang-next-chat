CREATE TABLE profile_images (
                                id BIGSERIAL NOT NULL PRIMARY KEY,
                                profile_id BIGINT NOT NULL,
                                uuid UUID NOT NULL UNIQUE,
                                name VARCHAR NOT NULL,
                                url VARCHAR NOT NULL,
                                size INTEGER NOT NULL,
                                created_at TIMESTAMP NOT NULL,
                                updated_at TIMESTAMP NOT NULL,
                                is_deleted bool NOT NULL,
                                is_enabled bool NOT NULL,
                                CONSTRAINT fk_profile_id FOREIGN KEY (profile_id) REFERENCES profiles (id)
);
