package cluster

import "github.com/Saad7890-web/neurocache/internal/kv"


type ShardManager struct {
	shards []*kv.Store
	count uint32
}

func NewShardManager(numShards int) *ShardManager {

	shards := make([]*kv.Store, numShards)

	for i := 0; i < numShards; i++ {

		store := kv.NewStore()
		store.StartExpirationWorker()

		shards[i] = store
	}

	return &ShardManager{
		shards: shards,
		count:  uint32(numShards),
	}
}