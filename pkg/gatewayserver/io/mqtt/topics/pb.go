package topics

import (
	"context"
)

const topicV3 = "v3"

type v3 struct{}

func (v3 *v3) BirthTopic(uid string) []string {
	return nil
}

func (v3 *v3) IsBirthTopic(path []string) bool {
	return false
}

func (v3 *v3) LastWillTopic(uid string) []string {
	return nil
}

func (v3 *v3) IsLastWillTopic(path []string) bool {
	return false
}

func (v3 *v3) UplinkTopic(uid string) []string {
	return v3.createTopic(uid, []string{"up"})
}

func (v3 *v3) IsUplinkTopic(path []string) bool {
	return len(path) == 3 && path[0] == topicV3 && path[2] == "up"
}

func (v3 *v3) StatusTopic(uid string) []string {
	return v3.createTopic(uid, []string{"status"})
}

func (v3 *v3) IsStatusTopic(path []string) bool {
	return len(path) == 3 && path[0] == topicV3 && path[2] == "status"
}

func (v3 *v3) TxAckTopic(uid string) []string {
	return v3.createTopic(uid, []string{"down", "ack"})
}

func (v3 *v3) IsTxAckTopic(path []string) bool {
	return len(path) == 4 && path[0] == topicV3 && path[2] == "down" && path[3] == "ack"
}

func (v3 *v3) DownlinkTopic(uid string) []string {
	return v3.createTopic(uid, []string{"down"})
}

func (v3 *v3) createTopic(uid string, path []string) []string {
	inTopicIdentifier := uid
	return append([]string{topicV3, inTopicIdentifier}, path...)
}

// New returns the default layout.
func New(ctx context.Context) Layout {
	return &v3{}
}
