CREATE TABLE `pythons_stats` (
  `python_id` VARCHAR(50) NOT NULL,
  `hp` INT UNSIGNED NOT NULL DEFAULT 0,
  `attack` INT UNSIGNED NOT NULL DEFAULT 0,
  `defense` INT UNSIGNED NOT NULL DEFAULT 0,
  `speed` INT UNSIGNED NOT NULL DEFAULT 0,
  `skill1` BIGINT UNSIGNED DEFAULT NULL,
  `skill2` BIGINT UNSIGNED DEFAULT NULL,
  `skill3` BIGINT UNSIGNED DEFAULT NULL,
  `skill4` BIGINT UNSIGNED DEFAULT NULL,
  PRIMARY KEY (`python_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;