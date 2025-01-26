CREATE TABLE IF NOT EXISTS collectible (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `rarity` VARCHAR(100) NOT NULL,
    `type` VARCHAR(100) NOT NULL,
    `image_url` VARCHAR(255) NOT NULL,
    `user_username` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_username`) REFERENCES users(`username`)
);
