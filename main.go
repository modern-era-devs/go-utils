package main

import (
	"fmt"
	"github.com/modern-era-devs/go-utils/authentication/jwt"
	"github.com/modern-era-devs/go-utils/config"
	"github.com/modern-era-devs/go-utils/database/postgres"
	"github.com/modern-era-devs/go-utils/queue/kafka"
	"github.com/spf13/viper"
	"runtime"
	"time"
)

type AppConfig struct {
	PostgresCfg   postgres.PostgresConfig `mapstructure:"POSTGRES"`
	KafkaProducer kafka.ProducerConfig    `mapstructure:"KAFKA_PRODUCER"`
	KafkaConsumer kafka.ConsumerConfig    `mapstructure:"KAFKA_CONSUMER"`
}

func (cfg AppConfig) Init(filename string) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (cfg AppConfig) Validate() error {
	return nil
}

func main() {
	//ctx := context.Background()
	fmt.Println("initiating the project")
	cfg := AppConfig{}
	err := cfg.Init("application")
	if err != nil {
		fmt.Println("viper init ", err.Error())
		return
	}
	err = config.Load(&cfg)
	if err != nil {
		fmt.Println("config load ", err.Error())
		return
	}

	fmt.Println(cfg)

	postgresConfig := postgres.GetPostgresConfig(cfg.PostgresCfg.Port, cfg.PostgresCfg.MaxPoolSize, cfg.PostgresCfg.MaxIdleConnections, cfg.PostgresCfg.Username, cfg.PostgresCfg.Password, cfg.PostgresCfg.Host, cfg.PostgresCfg.Name, cfg.PostgresCfg.SSLMode)

	conn, err := postgres.NewPostgres(postgresConfig)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err = conn.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	token, err := jwt.CreateHS256SignedJWT("something", map[string]interface{}{"exp": time.Now().Add(1 * time.Hour), "adm": false, "id": "LasdfsdafS123asdf2sfadsfdasfasdf34"})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(token)

	isValid, err := jwt.IsValidHS256JWTSignature("somethawsdfing", token)
	if err != nil {
		fmt.Println(err)
	}

	if isValid {
		fmt.Println("valid signature")
	} else {
		fmt.Println("invalid signature")
	}

	//producer, err := kafka.SetupProducerConnection(cfg.KafkaProducer)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//for i := 0; i < 100; i++ {`
	//	err = kafka.Produce(producer, cfg.KafkaProducer.Topic, []byte(fmt.Sprintf("go-utils testing %d", i)))
	//	if err != nil {
	//		fmt.Println("error while producing message to kafka: ", err.Error())
	//	}
	//}

	//consumer, err := kafka.SetupConsumerConnection(cfg.KafkaConsumer)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println("consumer: ", consumer)
	//
	//err = kafka.StartConsumption(ctx, consumer, cfg.KafkaConsumer.PollTimeout, executeMessage)
	//
	//if err != nil {
	//	fmt.Println("error while consuming from connection: ", err.Error())
	//}
	//err = kafka.Produce(producer, []byte("something"))
}

func executeMessage(value []byte) error {
	fmt.Println(string(value))
	fmt.Println("number of active goroutines: ", runtime.NumGoroutine())
	return nil
}
