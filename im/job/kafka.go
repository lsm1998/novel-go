package job

import (
	"github.com/Shopify/sarama"
	"im/config"
	"utils"
)

func Init() {
	imC := config.Config.Im
	utils.ConsumeGroup(imC.Adders, []string{imC.Topic}, imC.Group, func(msg *sarama.ConsumerMessage) {
	})
}
