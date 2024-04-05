package rabbitmq

type RabbitMQConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	ExchangeName string
	Kind         string
}

//func NewRabbitMQConn(cfg *RabbitMQConfig, ctx context.Context) (*amqp091.Connection, error) {
//	connAddr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
//
//	var conn *amqp091.Connection
//
//}
