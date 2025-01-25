CREATE TABLE IF NOT EXISTS chore (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `xp` INT UNSIGNED NOT NULL,
    'difficulty' INT UNSIGNED NOT NULL,
    'time_estimate' INT UNSIGNED NOT NULL,
    'completion_time' INT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
);