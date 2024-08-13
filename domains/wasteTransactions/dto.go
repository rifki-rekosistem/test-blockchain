package wasteTransactions

type WasteTransaction struct {
	Weight   float64 `json:"weight"`
	Category string  `json:"category"`
	WasteHub string  `json:"wastehub"`
}

type Block struct {
	Index        int                `json:"index"`
	Timestamp    int64              `json:"timestamp"`
	Transactions []WasteTransaction `json:"transactions"`
	PrevHash     string             `json:"prev_hash"`
	Hash         string             `json:"hash"`
	Validator    string             `json:"validator"`
}
