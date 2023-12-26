package mqttsnclient

import (
	"fmt"
	"time"

	"github.com/energomonitor/bisquitt/client"
	"github.com/energomonitor/bisquitt/topics"
	"github.com/energomonitor/bisquitt/util"
	"github.com/ryodeushii/mqttsn-bombardier/utils"
)

func Connect(log utils.ILogger, username string, password string, ip string, port int, keepAlive *int) error {
	brokerAddress := fmt.Sprintf("%s:%d", ip, port)

	if keepAlive == nil {
		ka := 60
		keepAlive = &ka
	}

	logger := util.NewProductionLogger("mqttsn-client")
	activityTopic := fmt.Sprintf("messages/%s/activity", username)
	clientID := fmt.Sprintf("%s%s", username, password)

	predefinedTopics := topics.PredefinedTopics{}

	my_client := client.NewClient(logger, &client.ClientConfig{
		// User:           username,
		// Password:       []byte(password),
		RetryDelay:     1 * time.Second,
		RetryCount:     1,
		CleanSession:   true,
		Insecure:       true,
		ConnectTimeout: 10 * time.Second,
		// PredefinedTopics: predefinedTopics,
		ClientID:  clientID,
		KeepAlive: time.Duration(*keepAlive) * time.Second,
	})

	err := my_client.Dial(brokerAddress)
	defer my_client.Close()
	if err != nil {
		return err
	}

	_ = my_client.Connect()

	for i := 0; i < 15; i++ {
		topicID, isPredefinedTopic := predefinedTopics.GetTopicID(clientID, activityTopic)

		payload := []byte(fmt.Sprintf("some activity with id %v", i))
		if isPredefinedTopic {
			_ = my_client.PublishPredefined(topicID, payload, 0, true)

		} else {
			_ = my_client.Register(activityTopic)
			_ = my_client.Publish(activityTopic, payload, 0, true)
		}
	}

	return nil
}
