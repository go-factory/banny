# banny

![](https://github.com/go-factory/banny/blob/master/logo.png)

## What's banny?
banny is an identicon generator, here is an introduction about what is [identicon](https://en.wikipedia.org/wiki/Identicon).

## How to use?
Get the project:
```
go get github.com/go-factory/banny
```
Generate identicon, you can generate it in your own project directory just like this:
```golang
fi, err := os.Create("logo.png")
if err != nil {
  log.Fatal(err)
}
defer fi.Close()

png.Encode(fi, banny.Config.Generate(240))
```

## What else needs to be done?
- [ ] Support email convertion
- [ ] Add color scheme
- [ ] Customize identicon configuration(size, color etc.)
