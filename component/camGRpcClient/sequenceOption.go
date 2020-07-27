package camGRpcClient

import "google.golang.org/grpc"

// sequence logic option
type sequenceOption struct {
	keys  []string
	len   int
	index int
}

func newSequenceOption(connDict *connDict) *sequenceOption {
	opt := new(sequenceOption)
	connDict.Range(func(key string, conn *grpc.ClientConn) bool {
		opt.keys = append(opt.keys, key)
		return true
	})
	opt.len = connDict.Len()
	opt.index = 0
	return opt
}

// get now key and move index to the next
func (opt *sequenceOption) nextKey() string {
	key := opt.keys[opt.index]
	if opt.len > 1 {
		opt.index++
		if opt.index == opt.len {
			opt.index = 0
		}
	}
	return key
}
