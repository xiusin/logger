package logger

import "sync"

type LogEntity struct {
	*Logger
	sync.Mutex

	LoggerIdField string
	id            string
}

func (entity *LogEntity) Id(id string) {
	if len(id) > 0 {
		entity.Lock()
		if entity.entityData == nil {
			entity.entityData = map[string]interface{}{}
		}
		entity.Unlock()
		entity.id = id
		entity.SetDataKV(entity.LoggerIdField, id)
	}
}

func (entity *LogEntity) SetDataKV(key string, val interface{}) *LogEntity {
	entity.Lock()
	defer entity.Unlock()

	entity.entityData[key] = val
	return entity
}

func (entity *LogEntity) SetData(data map[string]interface{}) *LogEntity {
	entity.Lock()
	defer entity.Unlock()

	for k, v := range data {
		entity.entityData[k] = v
	}
	return entity
}

func (entity *LogEntity) ClearData() {
	entity.Lock()
	defer entity.Unlock()
	entity.entityData = map[string]interface{}{}
	if len(entity.id) > 0 {
		entity.SetDataKV(entity.LoggerIdField, entity.id)
	}
}

func (entity *LogEntity) getDatas() map[string]interface{} {
	return entity.entityData
}
