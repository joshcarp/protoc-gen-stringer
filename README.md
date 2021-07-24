# protoc-gen-stringer

## About <a name = "about"></a>

A proto plugin to generate string values for go as defined by proto options.

This proto plugin was created because protobufs use c++ style naming scopes: [github.com/protocolbuffers/protobuf/issues/5425](https://github.com/protocolbuffers/protobuf/issues/5425). 
This means that it's forbidden to have the following in ptotobuf:

```proto

enum Foo {
    Foo1 = 1;
    Foo2 = 2;
    Unknown = 3;
}

enum Bar {
    Bar1 = 1;
    Bar2 = 2;
    Unknown = 3; // 'Unknown' is already defined in scope 'example'
}

```

The solution to this is to prefix the enum with a stringval, but this means that a custom stringer logic needs to be
implemented, which is annoying and unmaintainable.

```proto
enum Foo {
    Foo1 = 1;
    Foo2 = 2;
    Foo_Unknown = 3;
}

enum Bar {
    Bar1 = 1;
    Bar2 = 2;
    Bar_Unknown = 3; // No error but now custom stringer logic needs to be implemented
}
```

## Getting Started <a name = "getting_started"></a>

### Installing <a name = "installing"></a>

```
go install github.com/joshcarp/protoc-gen-stringer
```

## Usage <a name = "usage"></a>

```proto
enum Foo {
    Foo1 = 0 [(string_val) = "Foo1"];
    Foo2 = 1 [(string_val) = "Foo2"];
    Foo_Unknown = 2 [(string_val) = "Unknown"];
}

enum Bar {
    Bar1 = 0 [(string_val) = "Bar1"];
    Bar2 = 1 [(string_val) = "Bar2"];
    Bar_Unknown = 2 [(string_val) = "Unknown"]; // No error but now custom stringer logic needs to be implemented
}
```

```
protoc -I example/ example/example.proto --go_out=paths=source_relative:example --stringer_out=source_relative:example
```

This will generate a `StringVal()` method for each enum, and a `func StringToEnum(s string) Enum` function:
```go
fmt.Println(example.Foo_Unknown.StringVal())
// Unknown
fmt.Println(example.Bar_Unknown.StringVal())
// Unknown
a := example.StringToFoo("Unknown") // gives example.Foo_Foo_Unknown
```
