CREATE TABLE IF NOT EXISTS chore_instance (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `chore_id` INT UNSIGNED NOT NULL,
    `user_username` VARCHAR(255) NOT NULL,
    `due_date` DATE NOT NULL,
    `completed` BOOLEAN NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`chore_id`) REFERENCES chore(`id`),
    FOREIGN KEY (`user_username`) REFERENCES users(`username`)
);