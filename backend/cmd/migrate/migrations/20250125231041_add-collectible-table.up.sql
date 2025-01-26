CREATE TABLE IF NOT EXISTS collectible (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `is_legendary` BOOLEAN NOT NULL,
    `user_username` VARCHAR(255) NOT NULL,
    `image_url` VARCHAR(255) NOT NULL,
    `evolution` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_username`) REFERENCES users(`username`)
);