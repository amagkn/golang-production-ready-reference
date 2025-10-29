package main

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err)
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             "my-topic", // имя топика
			NumPartitions:     4,          // количество партиций в топике
			ReplicationFactor: 3,          // количество реплик у каждой партиции
			ConfigEntries: []kafka.ConfigEntry{
				// Минимальное количество реплик для синхронизации
				{"min.insync.replicas", "2"},

				// Сегменты удаляются по времени или размеру
				{"cleanup.policy", "delete"},

				// Максимальный размер сообщения в байтах
				{"max.message.bytes", "1048588"},

				// Минимальное время хранения сообщений. Сегменты удаляются, если все их сообщения старше этого времени
				{"retention.ms", "604800000"},

				// Максимальный общий размер топика. Удаляются самые старые сегменты при превышении этого размера
				{"retention.bytes", "1073741824"},

				// Время, после которого сегмент закрывается. Даже если сегмент не заполнен, он всё равно будет закрыт
				{"segment.ms", "604800000"},

				// Максимальный размер файла сегмента. При достижении этого размера Kafka создаёт новый сегмент
				{"segment.bytes", "1073741824"},

				// CreateTime - время создания сообщения продюсером, LogAppendTime - время добавления в лог брокером
				{"message.timestamp.type", "CreateTime"},
			},
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err)
	}
}
