CREATE TABLE IF NOT EXISTS users (
    `username` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `household_id` INT UNSIGNED,
    `title` VARCHAR(255) NOT NULL,
    `level` INT UNSIGNED NOT NULL,
    `xp` INT UNSIGNED NOT NULL,


    PRIMARY KEY (`username`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `email` (`email`)
);