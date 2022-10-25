package service

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	pb "tophub/api/task/v1"
	"tophub/app/task/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTaskService, NewDiscovery, NewRegistrar, NewDB, NewRedisDB, NewCron)

type TaskService struct {
	pb.UnimplementedTaskServer

	db   *gorm.DB
	rdb  *redis.Client
	cron *cron.Cron
	log  *klog.Helper
}

func NewTaskService(db *gorm.DB, rdb *redis.Client, cron *cron.Cron, logger klog.Logger) (*TaskService, func(), error) {
	cleanup := func() {
		klog.NewHelper(logger).Info("closing the data resources")
	}
	return &TaskService{
		db:   db,
		rdb:  rdb,
		cron: cron,
		log:  klog.NewHelper(klog.With(logger, "module", "task/service")),
	}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDB(c *conf.Data) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			LogLevel:      logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Source,
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		klog.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	return db
}

func NewRedisDB(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		klog.Errorf("ping redis database error: %v", err)
	}

	return rdb
}

func NewCron(db *gorm.DB) *cron.Cron {
	c := cron.New(cron.WithSeconds())
	c.Start()

	if _, err := c.AddJob("* */10 * * * *", &Spider{DB: db}); err != nil {
		klog.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to add Spider Job")
	}

	return c
}
