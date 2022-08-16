package data

import (
	"github.com/Shopify/sarama"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"tophub/app/job/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKafkaConsumer, NewDiscovery, NewJobRepo)

// Data .
type Data struct {
	kc  sarama.Consumer
	log *log.Helper
}

// NewData .
func NewData(consumer sarama.Consumer, logger log.Logger) (*Data, func(), error) {
	zlog := log.NewHelper(log.With(logger, "module", "data/job"))
	d := &Data{
		kc:  consumer,
		log: zlog,
	}
	return d, func() {
		d.kc.Close()
	}, nil
}

func NewKafkaConsumer(conf *conf.Data) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
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
