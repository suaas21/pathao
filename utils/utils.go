package utils

import (
	"fmt"
	"time"
)

const (
	UserTagIds = "userTagIDs"
)

type ResolverKey struct {
	Key interface{}
}

func NewResolverKey(key interface{}) *ResolverKey {
	return &ResolverKey{
		Key: key,
	}
}

func (rk *ResolverKey) String() string {
	return fmt.Sprintf("%v", rk.Key)
}

func (rk *ResolverKey) Raw() interface{} {
	return rk.Key
}

func BoolP(boolValue bool) *bool {
	return &boolValue
}

func GetCurrentTime() string {
	return time.Now().UTC().Add(6 * time.Hour).Format(time.RFC3339)
}
