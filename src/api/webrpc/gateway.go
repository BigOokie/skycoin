package webrpc

import "github.com/skycoin/skycoin/src/visor"
import "github.com/skycoin/skycoin/src/cipher"

// Gatewayer provides interfaces for getting skycoin related info.
type Gatewayer interface {
	GetLastBlocks(num uint64) *visor.ReadableBlocks
	GetBlocks(start, end uint64) *visor.ReadableBlocks
	GetUnspentByAddrs(addrs []string) []visor.ReadableOutput
	GetUnspentByHashes(hashes []string) []visor.ReadableOutput
	GetTransaction(txid cipher.SHA256) (*visor.TransactionResult, error)
}
