package topics

// Layout represents an MQTT topic layout.
type Layout interface {
	BirthTopic(uid string) []string
	IsBirthTopic(path []string) bool
	LastWillTopic(uid string) []string
	IsLastWillTopic(path []string) bool

	UplinkTopic(uid string) []string
	IsUplinkTopic(path []string) bool
	StatusTopic(uid string) []string
	IsStatusTopic(path []string) bool
	TxAckTopic(uid string) []string
	IsTxAckTopic(path []string) bool

	DownlinkTopic(uid string) []string
}
