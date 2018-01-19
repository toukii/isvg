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

<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="320px" height="321px" viewBox="0 0 320 321" enable-background="new 0 0 320 321" xml:space="preserve">  <image id="image0" width="320" height="321" x="0" y="0"
    xlink:href="data:image/png;base64,iV/>
</svg>
