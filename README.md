# OnceCell

A thread-safe value that can only be set once

---

Have you ever wanted to use a constant in your program, but did not have access to its value until runtime?

Have you ever wanted to use a global variable to store command line arguments, but wanted to ensure the variable is never changed?

Have you every wanted to enforce immutability with data shared by multiple threads?

Well, this is the library for you!


## Installation

```bash
go get github.com/Xavientois/once-cell
```

## Usage

```go
import oncecell "github.com/Xavientois/once-cell"

var myString oncecell.OnceString

// Setters
myString.SetValue("hello") // Set the value once
err := myString.SetValue("world") // Set the value again and an error will be returned
myString.MustSetValue("world") // Will panic if the value is already set

// Getters
s, err := myString.Value() // Returns an error if the value has not been set
s := myString.MustValue() // Panics if the value has not been set

```

## Custom Types

The library is currently only implemented for primitive types, at least until generics are introduced into the language.
If you wish to use custom types in the meantime, you can use the `once-cell-gen.sh` script to generate an implementation for your custom type.

```bash
./once-cell-gen.sh YourTypeName
```

The second command-line argument is the value to be used in the generated tests

## License

Apache 2.0


*Inspired by [once_cell](https://doc.rust-lang.org/std/lazy/struct.OnceCell.html) from Rust*
