package cluster

import (
	"hash/fnv"

	"github.com/Saad7890-web/neurocache/internal/kv"
)


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

func (m *ShardManager) getShard(key string) *kv.Store {

	hash := fnv.New32a()

	hash.Write([]byte(key))

	index := hash.Sum32() % m.count

	return m.shards[index]
}


func (m *ShardManager) Set(key, value string, ttl int) {

	shard := m.getShard(key)

	shard.Set(key, value, ttl)
}

func (m *ShardManager) Get(key string) (string, bool) {

	shard := m.getShard(key)

	return shard.Get(key)
}

func (m *ShardManager) Del(key string) bool {

	shard := m.getShard(key)

	return shard.Del(key)
}