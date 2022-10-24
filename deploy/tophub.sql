CREATE TABLE `data`
(
    `id`          bigint        NOT NULL AUTO_INCREMENT,
    `tab`         varchar(32)   NOT NULL COMMENT '标识',
    `position`    tinyint unsigned NOT NULL COMMENT '资讯位置',
    `title`       varchar(512)  NOT NULL DEFAULT '' COMMENT '资讯标题',
    `url`         varchar(512)  NOT NULL COMMENT '资讯地址',
    `extra`       json          NOT NULL COMMENT '拓展字段',
    `spider_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '抓取时间',
    PRIMARY KEY (`id`),
    KEY           `idx_spider_time` (`spider_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;