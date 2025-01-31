CREATE TABLE `words`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '専門用語の識別子',
    `title` VARCHAR(128) NOT NULL COMMENT '専門用語',
    `description` TEXT COMMENT '専門用語の説明',
    `lab` VARCHAR(128) COMMENT '研究室分野',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='専門用語';

CREATE TABLE `articles`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '論文の識別子',
    `title` VARCHAR(256) NOT NULL COMMENT '論文名',
    `author` VARCHAR(256) NOT NULL COMMENT '著者情報',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='論文';

CREATE TABLE `article_words`
(
    `article_id` BIGINT UNSIGNED NOT NULL COMMENT '論文の識別子',
    `word_id` BIGINT UNSIGNED NOT NULL COMMENT '専門用語の識別子',
    `created` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`article_id`, `word_id`),
    CONSTRAINT `fk_article` FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_word` FOREIGN KEY (`word_id`) REFERENCES `words`(`id`) ON DELETE CASCADE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='論文と専門用語の関連';
