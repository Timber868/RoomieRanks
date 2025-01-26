CREATE TABLE IF NOT EXISTS chore_instance (
    `chore_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `due_date` DATE NOT NULL,
    `completed` BOOLEAN NOT NULL,
    PRIMARY KEY (`chore_id`, `user_id`),
    FOREIGN KEY (`chore_id`) REFERENCES chore(`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);