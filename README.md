# randstr - a random string generator

A super simple random string generator suitable for passwords **(cryptographic strength entropy)** and the creation of
all sorts of random jibberish.

A quick look on [Kasperski Password Check](https://password.kaspersky.com/) _**using a default generated string**_ gives
the following

![password check result!](etc/password_check.jpeg "wow")

## Source (this stuff)

This is a very basic Go module with all logic comtained within `main.go`. _Comments omitted since it should be pretty
much self explanatory.

## Installation

**Go**

```shell 
go install github.com/stuartinnes/randstr
```

**Homebrew**

```shell 
brew install stuartinnes/tap/randstr
```

## Arguments

```shell 
--Upper value, -u value    Upper case characters (default: 4)
--Lower value, -l value    Lower case characters (default: 4)
--Digits value, -d value   Digits (default: 4)
--Symbols value, -s value  symbols (default: 4)
--help, -h                 show help (default: false)
```

## Example

```shell 
#generate a string with 8 upper and lower case characters and no digits or symbols
randstr -u 8 -l 8 -d 0 -s 0
```
