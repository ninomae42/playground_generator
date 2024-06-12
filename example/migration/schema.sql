CREATE TABLE `users` (
    id varchar(255) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `user_details` (
    user_id varchar(255),
    username varchar(255),
    crated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_details_users_id FOREIGN KEY (
        user_id
    ) REFERENCES users(id) ON DELETE CASCADE
);
