package common

import (
	"encoding/json"
)

type Property struct {
	kv map[string]string
}

func NewProperty() *Property {
	return &Property{kv: make(map[string]string)}
}

func (this *Property) IsEmpty() bool {
	if len(this.kv) == 0 {
		return true
	}

	return false
}

func (this *Property) Set(k string, v string) {
	this.kv[k] = v
}

func (this *Property) Get(k string) (string, bool) {
	v, ok := this.kv[k]
	if ok {
		return string(v), true
	}
	return "", false
}

func (this *Property) Del(k string) {
	delete(this.kv, k)
}

func (this *Property) Contact(p *Property) {
	if p == nil {
		return
	}

	for k, v := range p.kv {
		this.Set(k, v)
	}
}

func (this *Property) Marshal() (string, error) {
	data, err := json.Marshal(this.kv)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (this *Property) Unmarshal(str string) error {
	err := json.Unmarshal([]byte(str), &this.kv)
	if err != nil {
		return err
	}

	return nil
}
