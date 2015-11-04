# String reader

Simple StringReader which implements Golang io.Reader interface. 

# Usage

``` GO
    str := "Some nice string example!"
    r := &StringReader{Val: str}

    bytes, err := ioutil.ReadAll(r)
    if str != string(bytes) {
        fmt.Printf("Read wrong bytes! Got: %s, Expected: %s \n", string(bytes), str)
    }
```

