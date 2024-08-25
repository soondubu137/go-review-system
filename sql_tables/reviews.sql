CREATE TABLE `reviews` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_by` BIGINT NOT NULL COMMENT 'User ID',
    `updated_by` BIGINT NOT NULL COMMENT 'User ID',
    `version` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Optimistic locking',
    `deleted_at` DATETIME DEFAULT NULL COMMENT 'Soft delete',
    `review_id` BIGINT NOT NULL,
    `content` TEXT NOT NULL,
    `rating` TINYINT UNSIGNED NOT NULL COMMENT 'Overall rating',
    `service_rating` TINYINT UNSIGNED NOT NULL COMMENT 'Seller service rating',
    `delivery_rating` TINYINT UNSIGNED NOT NULL COMMENT 'Delivery rating',
    `has_media` BOOLEAN NOT NULL DEFAULT FALSE,
    `pictures` JSON NOT NULL COMMENT 'Review pictures',
    `videos` JSON NOT NULL COMMENT 'Review videos',
    `order_id` BIGINT NOT NULL,
    `sku_id` BIGINT NOT NULL,
    `spu_id` BIGINT NOT NULL,
    `seller_id` BIGINT NOT NULL,
    `buyer_id` BIGINT NOT NULL,
    `is_anonymous` BOOLEAN NOT NULL DEFAULT FALSE,
    `status` ENUM('PENDING', 'APPROVED', 'REJECTED', 'HIDDEN') NOT NULL DEFAULT 'PENDING',
    `tags` JSON NOT NULL COMMENT 'Related tags',
    `is_default` BOOLEAN NOT NULL DEFAULT FALSE,
    `has_reply` BOOLEAN NOT NULL DEFAULT FALSE,
    `op_reject_reason` TEXT DEFAULT NULL COMMENT 'Operator reject reason',
    `op_reject_at` DATETIME DEFAULT NULL COMMENT 'Operator reject time',
    `op_reject_by` BIGINT DEFAULT NULL COMMENT 'Operator reject user ID',
    `op_note` TEXT DEFAULT NULL COMMENT 'Operator note',
    `snapshot` JSON NOT NULL COMMENT 'Snapshot of the review',
    `ext_json` JSON NOT NULL COMMENT 'Extended information',
    `ctrl_json` JSON NOT NULL COMMENT 'Control information',
    PRIMARY KEY (`id`),
    KEY `idx_review_id` (`review_id`),
    KEY `idx_order_id` (`order_id`),
    KEY `idx_user_id` (`buyer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;