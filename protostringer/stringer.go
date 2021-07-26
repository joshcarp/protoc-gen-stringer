package protostringer

import (
	"fmt"
	"sync"

	"github.com/joshcarp/protoc-gen-stringer/stringer"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type EnumRegistry interface {
	String(protoreflect.Enum) string
	Enum(string, interface{})
	Register(descriptors ...protoreflect.Descriptor)
}

type enumRegistry struct {
	sync.RWMutex
	registry map[string]*enumMap
}

func (e *enumRegistry) String(enum protoreflect.Enum) string {
	desc := enum.Descriptor()
	descName := string(desc.FullName())

	// Find the relevant enum map
	// If it doesn't exist, register it
	e.RLock()
	enumMap, ok := e.registry[descName]
	e.RUnlock()
	if !ok {
		e.Register(desc)
		e.RLock()
		enumMap, ok = e.registry[descName]
		e.RUnlock()

		// If it still doesn't exist, this is bad
		if !ok {
			panic("bad things")
		}
	}

	if stringValue, ok := enumMap.enumToString[enum.Number()]; ok {
		return stringValue
	}
	return fmt.Sprint(enum)
}

func (e *enumRegistry) Enum(s string, enum interface{}) {
	panic("not implemented")
}

func (e *enumRegistry) Register(descriptors ...protoreflect.Descriptor) {
	// prevents writes while a new descriptor list is loaded
	e.Lock()
	defer e.Unlock()
	for _, desc := range descriptors {
		loadDescriptorToRegistry(desc, e)
	}
}

// enumMap stores two maps, one for mapping numbers to strings, and another
// for strings to enum values
type enumMap struct {
	stringToEnum map[string]protoreflect.EnumNumber
	enumToString map[protoreflect.EnumNumber]string
}

func (e *enumMap) String(value protoreflect.EnumNumber) string {
	return e.enumToString[value]
}

func (e *enumMap) Enum(s string) protoreflect.EnumNumber {
	return e.stringToEnum[s]
}

func loadDescriptorToRegistry(descriptor protoreflect.Descriptor, registry *enumRegistry) {
	switch desc := descriptor.(type) {
	case enumLister:
		loadEnumListToRegistry(desc, registry)
	case protoreflect.EnumDescriptor:
		loadEnumDescriptorToRegistry(desc, registry)
	}
}

type enumLister interface {
	Messages() protoreflect.MessageDescriptors
	Enums() protoreflect.EnumDescriptors
}

func loadEnumListToRegistry(descriptor enumLister, registry *enumRegistry) {
	messages := descriptor.Messages()
	for i := 0; i < messages.Len(); i++ {
		loadEnumListToRegistry(messages.Get(i), registry)
	}
	enums := descriptor.Enums()
	for i := 0; i < enums.Len(); i++ {
		loadEnumDescriptorToRegistry(enums.Get(i), registry)
	}
}

func loadEnumDescriptorToRegistry(descriptor protoreflect.EnumDescriptor, registry *enumRegistry) {
	enumMap := &enumMap{
		stringToEnum: map[string]protoreflect.EnumNumber{},
		enumToString: map[protoreflect.EnumNumber]string{},
	}
	values := descriptor.Values()
	for i := 0; i < values.Len(); i++ {
		valueDesc := values.Get(i)
		number := valueDesc.Number()
		options := valueDesc.Options().(*descriptorpb.EnumValueOptions)
		stringValue, ok := proto.GetExtension(options, stringer.E_Enum).(string)
		if !ok || stringValue == "" {
			stringValue = "TODO"
		}
		enumMap.enumToString[number] = stringValue
		enumMap.stringToEnum[stringValue] = number
	}
	registry.registry[string(descriptor.FullName())] = enumMap
}

var (
	gblRegistry = &enumRegistry{
		registry: make(map[string]*enumMap),
	}
)

func String(value protoreflect.Enum) string {
	return gblRegistry.String(value)
}
