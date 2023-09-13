package main

import "gorm.io/gorm"

const tableSQL = `CREATE TABLE IF NOT EXISTS news
(
    id          bigint unsigned                                                NOT NULL AUTO_INCREMENT,
    spider_time datetime                                                       NOT NULL,
    site        varchar(32) COLLATE utf8mb4_general_ci                         NOT NULL DEFAULT '',
    ranking     tinyint unsigned                                               NOT NULL DEFAULT '0',
    title       varchar(128) COLLATE utf8mb4_general_ci                        NOT NULL DEFAULT '',
    url         varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (id),
    KEY news_spider_time_site_index (spider_time DESC, site)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='新闻资讯表';`

func prepare(db *gorm.DB) {
	db.Exec(tableSQL)
}
