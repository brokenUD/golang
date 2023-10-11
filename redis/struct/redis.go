package redis

// import "go/types"

type redisObj struct {
	rType uint64
	encoding uint64
	lru uint64
	rfcount int64
	void *
}
