// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.pantheon.tech/stonework/proto/mockcnf"
)

////////// type-safe key-value pair with metadata //////////

type MockCnf1KVWithMetadata struct {
	Key      string
	Value    *mockcnf.MockCnf1
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type MockCnf1Descriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *mockcnf.MockCnf1) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *mockcnf.MockCnf1) error
	Create               func(key string, value *mockcnf.MockCnf1) (metadata interface{}, err error)
	Delete               func(key string, value *mockcnf.MockCnf1, metadata interface{}) error
	Update               func(key string, oldValue, newValue *mockcnf.MockCnf1, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *mockcnf.MockCnf1, metadata interface{}) bool
	Retrieve             func(correlate []MockCnf1KVWithMetadata) ([]MockCnf1KVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *mockcnf.MockCnf1) []KeyValuePair
	Dependencies         func(key string, value *mockcnf.MockCnf1) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type MockCnf1DescriptorAdapter struct {
	descriptor *MockCnf1Descriptor
}

func NewMockCnf1Descriptor(typedDescriptor *MockCnf1Descriptor) *KVDescriptor {
	adapter := &MockCnf1DescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *MockCnf1DescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castMockCnf1Value(key, oldValue)
	typedNewValue, err2 := castMockCnf1Value(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *MockCnf1DescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castMockCnf1Value(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *MockCnf1DescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castMockCnf1Value(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *MockCnf1DescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castMockCnf1Value(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castMockCnf1Value(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castMockCnf1Metadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *MockCnf1DescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castMockCnf1Value(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castMockCnf1Metadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *MockCnf1DescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castMockCnf1Value(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castMockCnf1Value(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castMockCnf1Metadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *MockCnf1DescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []MockCnf1KVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castMockCnf1Value(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castMockCnf1Metadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			MockCnf1KVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *MockCnf1DescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castMockCnf1Value(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *MockCnf1DescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castMockCnf1Value(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castMockCnf1Value(key string, value proto.Message) (*mockcnf.MockCnf1, error) {
	typedValue, ok := value.(*mockcnf.MockCnf1)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castMockCnf1Metadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
