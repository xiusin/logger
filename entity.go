package logger

import "sync"

type LogEntity struct {
	*Logger

	sync.Mutex
}

func (entity *LogEntity) SetDataKV(key string, val interface{}) {
	entity.Lock()
	defer entity.Unlock()

	if entity.entityData == nil {
		entity.entityData = map[string]interface{}{}
	}
	entity.entityData[key] = val
}

func (entity *LogEntity) SetData(data map[string]interface{}) {
	entity.Lock()
	defer entity.Unlock()
	if entity.entityData == nil {
		entity.entityData = map[string]interface{}{}
	}

	for k, v := range data {
		entity.entityData[k] = v
	}
}

func (entity *LogEntity) getDatas() map[string]interface{} {
	return entity.entityData
}
