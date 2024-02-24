package datasets

type EventLogArg struct {
	Name  string
	Value any
}

type EventLog struct {
	LogIndex uint64
	Block    uint64
	Address  string
	Event    string
	Chain    string
	TxHash   [32]byte
	Args     []EventLogArg
}

type EventLogReport struct {
	EventLog
	Signature [48]byte
}

type BroadcastEventPacket struct {
	Info      EventLog
	Signature [48]byte
	Signers   [][]byte
}
