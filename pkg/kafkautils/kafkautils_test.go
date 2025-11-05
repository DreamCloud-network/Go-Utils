package kafkautils

import (
	"fmt"
	"log"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestProducer(t *testing.T) {
	// Create new kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19095", // Endereço do broker Kafka
	})

	if err != nil {
		t.Errorf("Erro ao criar o producer: %v", err)
		return
	}
	defer p.Close()

	// Canal para capturar eventos de entrega (delivery reports)
	deliveryChan := make(chan kafka.Event)

	// Enviar mensagens
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Mensagem %d", i)

		// Publicar a mensagem no tópico 'test-topic'
		topic := "test-topic"
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, deliveryChan)

		if err != nil {
			fmt.Printf("Erro ao produzir mensagem: %v\n", err)
		}
	}

	// Capturar eventos de entrega
	go func() {
		for e := range deliveryChan {
			m := e.(*kafka.Message)
			if m.TopicPartition.Error != nil {
				fmt.Printf("Erro na entrega da mensagem: %v\n", m.TopicPartition)
			} else {
				fmt.Printf("Mensagem entregue em %v\n", m.TopicPartition)
			}
		}
	}()

	// Aguardar a entrega das mensagens
	p.Flush(15 * 1000)
	close(deliveryChan)
}

func TestConsumer(t *testing.T) {
	// Configurar o Consumer Kafka
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19095",      // Endereço do broker Kafka
		"group.id":          "meu-grupo-de-consumo", // ID do grupo de consumidores
		"auto.offset.reset": "earliest",             // Definir para consumir desde o início se não houver offsets
	})

	if err != nil {
		log.Fatalf("Erro ao criar o consumer: %v", err)
	}
	defer c.Close()

	// Subscribir-se a um tópico
	topic := "test-topic"
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Erro ao subscrever ao tópico: %v", err)
	}

	fmt.Println("Consumer esperando mensagens...")

	// Loop para consumir mensagens
	for {
		msg, err := c.ReadMessage(-1) // Bloqueia até uma mensagem ser recebida ou ocorrer erro
		if err != nil {
			fmt.Printf("Erro ao consumir mensagem: %v (%v)\n", err, msg)
			continue
		}
		fmt.Printf("Mensagem recebida no tópico %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}
