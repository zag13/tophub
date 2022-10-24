package service

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	v1 "tophub/api/interface/v1"
	taskv1 "tophub/api/task/v1"
	"tophub/app/interface/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInterfaceService, NewDiscovery, NewRegistrar, NewDB, NewRedisDB, NewTaskClient)

// InterfaceService is an Interface service.
type InterfaceService struct {
	v1.UnimplementedInterfaceServer

	tc  taskv1.TaskClient
	db  *gorm.DB
	rdb *redis.Client
	log *klog.Helper
}

// NewInterfaceService new an interface service.
func NewInterfaceService(logger klog.Logger, db *gorm.DB, rdb *redis.Client, tc taskv1.TaskClient) (*InterfaceService, func(), error) {
	cleanup := func() {
		klog.NewHelper(logger).Info("closing the data resources")
	}
	return &InterfaceService{
		tc:  tc,
		db:  db,
		rdb: rdb,
		log: klog.NewHelper(klog.With(logger, "module", "interface/service")),
	}, cleanup, nil
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

func NewTaskClient(r registry.Discovery, tp *tracesdk.TracerProvider) taskv1.TaskClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tophub.task"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return taskv1.NewTaskClient(conn)
}
