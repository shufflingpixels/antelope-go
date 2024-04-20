package chain

type BlockHeader struct {
	Timestamp        BlockTimestamp `json:"timestamp"`
	Producer         Name           `json:"producer"`
	Confirmed        uint16         `json:"confirmed"`
	Previous         Checksum256    `json:"previous"`
	TransactionMRoot Checksum256    `json:"transaction_mroot"`
	ActionMRoot      Checksum256    `json:"action_mroot"`
	ScheduleVersion  uint32         `json:"schedule_version"`

	// EOSIO 1.x
	NewProducersV1 *ProducerSchedule `json:"new_producers,omitempty" eosio:"optional"`

	HeaderExtensions []TransactionExtension `json:"header_extensions"`
}
