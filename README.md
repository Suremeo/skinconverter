# Examples

### Skin Data to PNG
```go
	dat, err := ioutil.ReadFile("Skin.dat")
	if err != nil {
		panic(err)
	}

	img := SkinConverter.SkinDataToImage(dat)

	save := bytes.NewBuffer(nil)

	err = png.Encode(save, img)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("Skin.png", save.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
```

### PNG to Skin Data

```go
	data, err := ioutil.ReadFile("Skin.png")

	pngimage, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	skindata := SkinConverter.ImageToSkinData(pngimage)

	err = ioutil.WriteFile("Skin.dat", skindata, os.ModePerm)
	if err != nil {
		panic(err)
	}
```