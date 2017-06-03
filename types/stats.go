package types

import "time"

type StatSet struct {
	Ts time.Time `json:"time"`

	Load         float64          `json:"load"`
	IoRead       map[string]int64 `json:"io_read"`
	IoReadTotal  int64            `json:"io_read_total"`
	IoWrite      map[string]int64 `json:"io_write"`
	IoWriteTotal int64            `json:"io_write_total"`
	NetSent      map[string]int64 `json:"net_sent"`
	NetSentTotal int64            `json:"net_sent_total"`
	NetRecv      map[string]int64 `json:"net_recv"`
	NetRecvTotal int64            `json:"net_recv_total"`
}
