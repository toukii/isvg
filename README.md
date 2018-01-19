# ISVG

svg to image(png)

## usage


 - __Decode__
 
```
bs := goutils.ReadFile("github.svg")

dbs, err := Decode(bs)
if err != nil {
	t.Errorf("%s", err)
}
if err := icat.ICatRead(bytes.NewReader(dbs), os.Stdout); err != nil {
	t.Errorf("%s", err)
}
```

 - __Display__

```
if err := Display(bs); err != nil {
	t.Errorf("%s", err)
}
```


![](https://raw.githubusercontent.com/toukii/isvg/master/github.svg)

![](https://raw.githubusercontent.com/toukii/isvg/master/github.png)