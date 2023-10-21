package bootstrap

import (
	"github.com/zag13/tophub/server/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDatabase(config *Config) *dal.Query {
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	prepare(db)

	return dal.Use(db)
}

func prepare(db *gorm.DB) {
	db.Exec(TableSQL)
}

const TableSQL = "CREATE TABLE IF NOT EXISTS tops " +
	"(" +
	"    id          bigint unsigned                                                NOT NULL AUTO_INCREMENT," +
	"    spider_time datetime                                                       NOT NULL," +
	"    site        varchar(32) COLLATE utf8mb4_general_ci                         NOT NULL DEFAULT ''," +
	"    `rank` 	 tinyint unsigned                                               NOT NULL DEFAULT '0'," +
	"    title       varchar(128) COLLATE utf8mb4_general_ci                        NOT NULL DEFAULT ''," +
	"    url         varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''," +
	"    extra		json                                                            NOT NULL," +
	"    PRIMARY KEY (id)," +
	"    KEY idx_spider_time_site (spider_time DESC, site)" +
	") ENGINE = InnoDB" +
	"  DEFAULT CHARSET = utf8mb4" +
	" COLLATE = utf8mb4_general_ci COMMENT ='热点资讯表';"
