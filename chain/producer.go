package chain

type ProducerKey struct {
	AccountName     Name      `json:"producer_name"`
	BlockSigningKey PublicKey `json:"block_signing_key"`
}

type ProducerSchedule struct {
	Version   uint32        `json:"version"`
	Producers []ProducerKey `json:"producers"`
}
