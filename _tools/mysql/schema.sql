CREATE TABLE `words`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '専門用語の識別子',
    `title` VARCHAR(128) NOT NULL COMMENT '専門用語',
    `description` TEXT COMMENT '専門用語の説明',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='専門用語';