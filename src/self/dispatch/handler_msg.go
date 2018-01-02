/**
 * Created by shiyi on 2017/10/16.
 * Email: shiyi@fightcoder.com
 */

package dispatch

import (
	"encoding/json"

	"self/judger"

	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Topic string
}

func (this *Handler) HandleMessage(m *nsq.Message) error {
	judgerData := new(judger.Judger)
	if err := json.Unmarshal(m.Body, judgerData); err != nil {
		log.Errorf("unmarshal JudgerData from NsqMessage failed, err: %v, event:%s", err, m.Body)
		return nil
	}

	log.Infof("consume Message from dispatch: %#v", judgerData)

	handlerCount <- 1
	go this.doJudge(judgerData)

	return nil
}

func (this *Handler) doJudge(judgerData *judger.Judger) {
	defer func() {
		<-handlerCount
	}()

	judgerData.DoJudge()
}
