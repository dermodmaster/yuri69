package database

import (
	"encoding/json"
	"strings"

	"github.com/xujiajun/nutsdb"
	. "github.com/zekrotja/yuri69/pkg/models"
)

const (
	bucketSounds = "sounds"
	bucketGuilds = "guilds"

	keySeparator = ":"
)

type NutsConfig struct {
	Location string
}

type Nuts struct {
	db *nutsdb.DB
}

var _ (IDatabase) = (*Nuts)(nil)

func NewNuts(c NutsConfig) (*Nuts, error) {
	var (
		t   Nuts
		err error
	)

	opts := nutsdb.DefaultOptions
	opts.Dir = c.Location
	t.db, err = nutsdb.Open(opts)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Nuts) Close() error {
	return t.db.Close()
}

func (t *Nuts) PutSound(sound Sound) error {
	data, err := marshal(sound)
	if err != nil {
		return err
	}

	return t.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(bucketSounds, []byte(sound.Uid), data, 0)
	})
}

func (t *Nuts) RemoveSound(uid string) error {
	return t.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Delete(bucketSounds, []byte(uid))
		return t.wrapErr(err)
	})
}

func (t *Nuts) GetSounds() ([]Sound, error) {
	var entries nutsdb.Entries
	err := t.db.View(func(tx *nutsdb.Tx) error {
		var err error
		entries, err = tx.GetAll(bucketSounds)
		return t.wrapErr(err)
	})
	if err != nil {
		return nil, err
	}

	sounds := make([]Sound, 0, len(entries))
	for _, e := range entries {
		sound, err := unmarshal[Sound](e.Value)
		if err != nil {
			return nil, err
		}
		sounds = append(sounds, sound)
	}

	return sounds, nil
}

func (t *Nuts) GetSound(uid string) (Sound, error) {
	var e *nutsdb.Entry
	err := t.db.View(func(tx *nutsdb.Tx) error {
		var err error
		e, err = tx.Get(bucketSounds, []byte(uid))
		return t.wrapErr(err)

	})
	if err != nil {
		return Sound{}, err
	}

	sound, err := unmarshal[Sound](e.Value)
	return sound, err
}

func (t *Nuts) SetGuildVolume(guildID string, volume int) error {
	data, err := marshal(volume)
	if err != nil {
		return err
	}
	return t.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(bucketGuilds, key(guildID, "volume"), data, 0)
	})
}

func (t *Nuts) GetGuildVolume(guildID string) (int, error) {
	var e *nutsdb.Entry
	err := t.db.View(func(tx *nutsdb.Tx) error {
		var err error
		e, err = tx.Get(bucketGuilds, key(guildID, "volume"))
		return t.wrapErr(err)
	})
	if err != nil {
		return 0, err
	}
	v, err := unmarshal[int](e.Value)
	return v, err
}

// --- Internal ---

func (t *Nuts) wrapErr(err error) error {
	if err == nil {
		return nil
	}
	if err == nutsdb.ErrKeyNotFound ||
		err == nutsdb.ErrNotFoundKey ||
		err == nutsdb.ErrBucketNotFound ||
		strings.HasPrefix(err.Error(), "bucket not found:") ||
		err == nutsdb.ErrBucketEmpty {
		return ErrNotFound
	}
	return err
}

func marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func unmarshal[T any](data []byte) (v T, err error) {
	err = json.Unmarshal(data, &v)
	return v, err
}

func key(elements ...string) []byte {
	return []byte(strings.Join(elements, keySeparator))
}
