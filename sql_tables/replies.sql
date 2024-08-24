CREATE TABLE replies (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_by` BIGINT NOT NULL COMMENT 'User ID',
    `updated_by` BIGINT NOT NULL COMMENT 'User ID',
    `version` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Optimistic locking',
    `deleted_at` DATETIME DEFAULT NULL COMMENT 'Soft delete',

    `reply_id` BIGINT NOT NULL,
    `review_id` BIGINT NOT NULL,
    `content` TEXT NOT NULL,
    `seller_id` BIGINT NOT NULL,
    `pictures` JSON NOT NULL DEFAULT '[]' COMMENT 'Reply pictures',
    `videos` JSON NOT NULL DEFAULT '[]' COMMENT 'Reply videos',

    `ext_json` JSON NOT NULL DEFAULT '{}' COMMENT 'Extended information',
    `ctrl_json` JSON NOT NULL DEFAULT '{}' COMMENT 'Control information',

    PRIMARY KEY (`id`),
    KEY `idx_reply_id` (`reply_id`),
    KEY `idx_review_id` (`review_id`),
    KEY `idx_seller_id` (`seller_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;