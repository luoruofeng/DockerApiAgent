package consul

import (
	"encoding/json"

	"github.com/hashicorp/consul/api"
)

type KVManager interface {
	Get(key string) (string, interface{}, error)
	Put(key string, value interface{}) error
}

func (i *serviceInstance) Get(key string) (string, interface{}, error) {
	kvp, _, err := i.kv.Get(key, nil)
	if err != nil {
		return "", nil, err
	}
	return kvp.Key, kvp.Value, nil
}

func (i *serviceInstance) Put(key string, value interface{}) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	kvp := &api.KVPair{
		Key:   key,
		Value: v,
	}
	_, err = i.kv.Put(kvp, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}
