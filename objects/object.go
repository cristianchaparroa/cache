package objects

import "time"

const (
	NeverExpires = 0
)

// Object is the domain object with will be stored in cache.
type Object struct {
	Data      string // This is the raw data
	TTL       int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewObject(data string) *Object {
	now := time.Now()
	return &Object{
		Data:      data,
		TTL:       NeverExpires,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewObjectWithTTL(data string, ttl int64) *Object {
	now := time.Now()
	return &Object{
		Data:      data,
		TTL:       ttl,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (o *Object) IsExpired() bool {

	if o.TTL == NeverExpires {
		return false
	}

	now := time.Now()
	expectTimeLived := o.CreatedAt.Add(time.Second * time.Duration(o.TTL))

	if now.After(expectTimeLived) {
		return true
	}

	return false
}
