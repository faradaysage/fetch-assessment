package mapper

import (
	"fetch-assessment/api"
	"reflect"
	"testing"
)

func EnsureMappersAreRegistered(t *testing.T) {
	// ensure api.Item has an item mapper
	if _, exists := factory.ItemMappers[reflect.TypeOf(api.Item{})]; !exists {
		t.Errorf("Expected ItemMapper for api.Item, but it was not registered")
	}

	// ensure json has an item mapper
	if _, exists := factory.ItemMappers[reflect.TypeOf([]byte{})]; !exists {
		t.Errorf("Expected ItemMapper for []byte, but it was not registered")
	}

	// ensure api.Receipt has a receipt mapper
	if _, exists := factory.ReceiptMappers[reflect.TypeOf(api.Receipt{})]; !exists {
		t.Errorf("Expected ReceiptMapper for api.Receipt, but it was not registered")
	}

	// ensure json has a receipt mapper
	if _, exists := factory.ReceiptMappers[reflect.TypeOf([]byte{})]; !exists {
		t.Errorf("Expected ReceiptMapper for []byte, but it was not registered")
	}
}
