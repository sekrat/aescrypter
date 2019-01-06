aescrypter is a `Crypter` implementation for [sekrat](https://github.com/sekrat/sekrat) that uses the AES-256-GCM algorithm to encrypt/decrypt the data handed to it.

[![GoDoc](https://godoc.org/github.com/sekrat/aescrypter?status.svg)](https://godoc.org/github.com/sekrat/aescrypter)

## Installation ##

```
go get github.com/sekrat/aescrypter
```

## Usage ##

To use this `Crypter` with a `sekrat.Manager`, you do something like this:

```go
package main

import (
  "github.com/sekrat/sekrat"
  "github.com/sekrat/aescryptor"
)

func main() {
  warehouse := sekrat.NewMemoryWarehouse()
  crypter := aescrypter.New()

  confidant := sekrat.New(warehouse, crypter)

  // Do more stuff
}
```

Now, when you `confidant.Put(...)` secrets, they will be encrypted pretty strongly.

***NOTE: The encrypted payloads that result from this algorithm are basically made of binary data, so if you're using a storage mechanism that doesn't do well with that sort of thing, it might not be a bad idea to make sure that said storage mechanism either is capable of dealing with binary data or transforms the data to a non-binary format for storage.***

## Contributing ##

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request against the "develop" branch

## History ##

* v1.0.0 - Initial release

## License ##

Sekrat is released under the Apache 2.0 license. See [LICENSE](https://github.com/sekrat/sekrat/blob/master/LICSENSE)
