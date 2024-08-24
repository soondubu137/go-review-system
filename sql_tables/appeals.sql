CREATE TABLE appeals (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_by` BIGINT NOT NULL COMMENT 'User ID',
    `updated_by` BIGINT NOT NULL COMMENT 'User ID',
    `version` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Optimistic locking',
    `deleted_at` DATETIME DEFAULT NULL COMMENT 'Soft delete',

    `appeal_id` BIGINT NOT NULL,
    `review_id` BIGINT NOT NULL,
    `reason` VARCHAR(255) NOT NULL COMMENT 'Appeal reason',
    `content` TEXT NOT NULL COMMENT 'Appeal content',
    `seller_id` BIGINT NOT NULL,
    `pictures` JSON NOT NULL DEFAULT '[]' COMMENT 'Appeal pictures',
    `videos` JSON NOT NULL DEFAULT '[]' COMMENT 'Appeal videos',
    `status` ENUM('PENDING', 'APPROVED', 'REJECTED') NOT NULL DEFAULT 'PENDING',

    `op_note` TEXT DEFAULT NULL COMMENT 'Operator note',
    `op_process_at` DATETIME DEFAULT NULL COMMENT 'Operator process time',
    `op_process_by` BIGINT DEFAULT NULL COMMENT 'Operator process user ID',

    `ext_json` JSON NOT NULL DEFAULT '{}' COMMENT 'Extended information',
    `ctrl_json` JSON NOT NULL DEFAULT '{}' COMMENT 'Control information',

    PRIMARY KEY (`id`),
    KEY `idx_appeal_id` (`appeal_id`),
    KEY `idx_review_id` (`review_id`),
    KEY `idx_seller_id` (`seller_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;