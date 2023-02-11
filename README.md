# configurator

Configurator helps to get configuration values from environment variables.

> This library is currently being worked on, and it shouldn't be considered stable. Only after version 1.0.0, you
> shouldn't expect breaking changes (based on semver).

## Usage

### Instantiation

#### Using an .env file

```go
config, err := configurator.NewConfiguratorFromFile(path)
```

#### Using OS environment variables

```go
config, err := configurator.NewConfiguratorFromFile(path)
```

### Getting values

#### String

```go
myString := config.GetString("FOO", "default string")
```

#### Boolean

```go
myBoolean := config.GetBoolean("FOO", false)
```

#### Required value

```go
myBoolean, err := config.MustGetString("FOO")
```
