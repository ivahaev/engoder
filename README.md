# engoder
Tiny golang module for encrypting passwords
## Installation
--
    go get github.com/ivahaev/engoder

## Usage
--
    import "github.com/ivahaev/engoder"

#### func  Check

```go
func Check(input string, _hash, salt []byte) bool
```
Check will encrypt input string with salt provided and compare with _hash. If
equals, will return true.

#### func  Encrypt

```go
func Encrypt(input string, _salt ...[]byte) (hash []byte, salt []byte, err error)
```
Encrypt takes input string and options []byte salt and return encrypted hash,
salt and error.
