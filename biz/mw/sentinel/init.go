package sentinel

import (
	"log"

	aliSentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func Load() {
	if err := aliSentinel.InitDefault(); err != nil {
		log.Println("sentinel init error: ", err)
		return
	}

	var newRule []*flow.Rule
	update(&newRule)

	if _, err := flow.LoadRules(newRule); err != nil {
		log.Println("load rules error: ", err)
		return
	}
}
