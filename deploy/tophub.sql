CREATE TABLE `data`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `host`        varchar(64) COLLATE utf8mb4_general_ci                         NOT NULL DEFAULT '' COMMENT '网站host',
    `rank`        tinyint unsigned NOT NULL COMMENT '资讯排名',
    `title`       varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '资讯标题',
    `description` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资讯描述',
    `image`       varchar(512) COLLATE utf8mb4_general_ci                        NOT NULL DEFAULT '' COMMENT '资讯配图',
    `url`         varchar(512) COLLATE utf8mb4_general_ci                        NOT NULL COMMENT '资讯地址',
    `extra`       json                                                           NOT NULL COMMENT '拓展字段',
    `spider_time` int unsigned NOT NULL COMMENT '抓取时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;