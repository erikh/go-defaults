## Use struct tags to determine struct defaults

Similar to rust's `Default` trait; just make sure you call the `defaults.Default()` function against pointer versions of your structs. It will only fill values that are not already at their golang-defined defaults, meaning you can use this after setting struct elements and it will do the right thing.

Note that you can overwrite the `CONVERSIONS` map in the package after initialization to adjust how things are parsed. Reasonable defaults have been supplied, but custom types are not supported. Only those that support `reflect.Type` as a constant.

You can also define custom types by writing a `func (t *Type) Default() error` function which modifies your struct directly. It will be called if it exists for all structs in the tree including the top level. This is very similar to the Rust `Default` trait.

Example:

```go
import (
    "github.com/erikh/go-defaults"
)

var DefaultDBConfig = DatabaseConfig{
    Username: "scott",
    Password: "tiger",
    Host: "localhost:1234",
}

type Config struct {
    Listen string `default:"localhost:3000"`
    Listeners uint `default:"5"`
    DB *DatabaseConfig
}

type DatabaseConfig {
    Username string
    Password string
    Host string
}

func (d *DatabaseConfig) Default() error {
    *d = DefaultDBConfig
    return nil
}

func main() {
    config := &Config{}
    if err := defaults.Default(config); err != nil {
        panic(err)
    }

    if config.DB.Username == "scott" && config.DB.Password == "tiger" {
        fmt.Println("success")
    }
}
```

For hopefully obvious reasons, **defaults only works on public struct members**.

## Usage

```
$ go get github.com/erikh/go-defaults
<sprinkle in code>
```

```go
import (
    "github.com/erikh/go-defaults"
)

type Config struct {
    Listen string `default:"localhost:3000"`
    Listeners uint `default:"5"`
}

func main() {
    config := &Config{}
    // for config files, pepper your defaults in, then parse your document
    if err := defaults.Default(config); err != nil {
        panic(err)
    }

    if config.Listeners == 5 {
        fmt.Println("it works!")
    }
}
```

## Author

Erik Hollensbe <erik@hollensbe.org>

## Stability

This library is considered stable and complete, only situations which panic the library or result in obviously broken behavior will be fixed.

## License

BSD 3 Clause
