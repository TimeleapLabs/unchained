package evmlog

type EventKey struct {
	Chain    string
	LogIndex uint64
	TxHash   [32]byte
}

type SupportKey struct {
	Chain   string
	Address string
	Event   string
}

type Event struct {
	Key   [32]byte
	Value [32]byte
}
