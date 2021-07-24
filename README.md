# protoc-gen-stringer

## About <a name = "about"></a>
A proto plugin to generate string values as defined by proto options.

This proto plugin was created because protobufs use c++ style naming scopes: [github.com/protocolbuffers/protobuf/issues/5425](https://github.com/protocolbuffers/protobuf/issues/5425).
This means that it's forbidden to have the following in protoc:
```proto
enum Foo {
    Unknown = 1;
}

enum Bar {
    Unknown = 1; // 'Unknown' is already defined in scope 'example'
}

```
The solution to this is to prefix the enum with a stringval, but this means that a custom stringer logic needs to be implemented, which is annoying and unmaintainable.

```proto
enum Foo {
    Foo_Unknown = 1;
}

enum Bar {
    Bar_Unknown = 1; // No error but now custom stringer logic needs to be implemented
}

```

## Getting Started <a name = "getting_started"></a>

### Installing

```
go install github.com/joshcarp/protoc-gen-stringer
```

## Usage <a name = "usage"></a>

```
protoc -I example/ example/example.proto --go_out=paths=source_relative:example --stringer_out=source_relative:example
```