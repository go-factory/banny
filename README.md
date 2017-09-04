# banny ![](https://travis-ci.org/go-factory/banny.svg?branch=master)

![](https://github.com/go-factory/banny/blob/master/logo.png)

## What's banny?
banny is an identicon generator, here is an introduction about what is [identicon](https://en.wikipedia.org/wiki/Identicon).

## How to use?
Get this library:
```
go get github.com/go-factory/banny
```
Generate identicon, `support ip analysis and email analysis`, you can generate it in your own project directory just like this:
```golang
fi, err := os.Create("logo.png")
if err != nil {
  log.Fatal(err)
}
defer fi.Close()

png.Encode(fi, banny.Config.Generate(240, ""))
png.Encode(fi, banny.Config.Generate(240, "124.2.45.230"))
png.Encode(fi, banny.Config.Generate(240, "zztcc@foxmail.com"))

Output:
Input is neither an ip address nor email address, use local ip address instead.
This is an ip address 124.2.45.230
This is an email zztcc@foxmail.com
```
