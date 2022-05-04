package database

import (
	"fmt"
	"strings"
	"sync"

	. "github.com/zekrotja/yuri69/pkg/models"
)

const (
	cacheKeySeparator = ":"
)

type DatabaseCache struct {
	IDatabase

	cache sync.Map
}

var _ IDatabase = (*DatabaseCache)(nil)

func WrapCache(db IDatabase, err error) (IDatabase, error) {
	if err != nil {
		return nil, err
	}

	var t DatabaseCache
	t.IDatabase = db

	return &t, nil
}

func (t *DatabaseCache) GetSounds() ([]Sound, error) {
	var err error
	key := ckey("sounds")

	vi, _ := t.cache.Load(key)
	v, ok := vi.([]Sound)
	fmt.Println(v, ok)
	if !ok {
		v, err = t.IDatabase.GetSounds()
		if err != nil {
			return nil, err
		}
		t.cache.Store(key, v)
	}

	r := make([]Sound, len(v))
	copy(r, v)

	return r, nil
}

func (t *DatabaseCache) SetGuildVolume(guildID string, volume int) error {
	t.cache.Store(ckey("guilds", guildID, "volume"), volume)
	return t.IDatabase.SetGuildVolume(guildID, volume)
}

func (t *DatabaseCache) GetGuildVolume(guildID string) (int, error) {
	var err error
	key := ckey("guilds", guildID, "volume")

	vi, _ := t.cache.Load(key)
	v, ok := vi.(int)
	if !ok {
		v, err = t.IDatabase.GetGuildVolume(guildID)
		if err != nil {
			return 0, err
		}
		t.cache.Store(key, v)
	}

	return v, nil
}

// --- Felpers ---

func ckey(elements ...string) string {
	return strings.Join(elements, cacheKeySeparator)
}
