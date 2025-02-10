package mapper

import (
	"errors"
	"fetch-assessment/api"
	"fetch-assessment/rules"
	"reflect"
	"sync"
)

type MapperFactory struct {
	once           sync.Once            // assists in ensuring that registerMappers is idempotent
	ItemMappers    map[reflect.Type]any // covariance for generics not supported in Go so we use any instead of ItemMapper[any]
	ReceiptMappers map[reflect.Type]any
}

// creates and initialize a mapper factory
func NewMapperFactory() *MapperFactory {
	f := &MapperFactory{}
	f.registerMappers() // ensure initialization
	return f
}

var factory = NewMapperFactory() // global factor instance

// idempotent
func (f *MapperFactory) registerMappers() {
	f.once.Do(func() {
		apiItemMapper := &ApiItemMapper{}
		jsonItemMapper := &JsonItemMapper{}
		f.ItemMappers = make(map[reflect.Type]any)
		f.ItemMappers[reflect.TypeOf(api.Item{})] = apiItemMapper
		f.ItemMappers[reflect.TypeOf([]byte{})] = jsonItemMapper

		apiReceiptMapper := &ApiReceiptMapper{
			ItemMapper: apiItemMapper,
		}
		jsonReceiptMapper := &JsonReceiptMapper{
			ItemMapper: jsonItemMapper,
		}
		f.ReceiptMappers = make(map[reflect.Type]any)
		f.ReceiptMappers[reflect.TypeOf(api.Receipt{})] = apiReceiptMapper
		f.ReceiptMappers[reflect.TypeOf([]byte{})] = jsonReceiptMapper
	})
}

func MapToItem[T any](source T) (rules.Item, error) {
	sourceType := reflect.TypeOf(source)
	mapperObject, exists := factory.ItemMappers[sourceType]
	if !exists {
		// mapper doesn't exist, return empty obj and an error
		return rules.Item{}, errors.New("Mapper does not exist for type: " + sourceType.String())
	}
	// assert the expected mapper type
	mapper, ok := mapperObject.(ItemMapper[T])
	if !ok {
		return rules.Item{}, errors.New("Mapper type assertion failed for type: " + sourceType.String())
	}
	return mapper.ToDomain(source)
}

func MapToReceipt[T any](source T) (rules.Receipt, error) {
	sourceType := reflect.TypeOf(source)
	mapperObject, exists := factory.ReceiptMappers[sourceType]
	if !exists {
		// mapper doesn't exist, return empty obj and an error
		return rules.Receipt{}, errors.New("Mapper does not exist for type: " + sourceType.String())
	}
	// assert the expected mapper type
	mapper, ok := mapperObject.(ReceiptMapper[T])
	if !ok {
		return rules.Receipt{}, errors.New("Mapper type assertion failed for type: " + sourceType.String())
	}
	return mapper.ToDomain(source)
}
