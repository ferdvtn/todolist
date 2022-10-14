CREATE TABLE `todolist`.`todos` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL,
	`description` TEXT,
	`priority` INT NOT NULL DEFAULT 1,
	PRIMARY KEY (`id`)
);