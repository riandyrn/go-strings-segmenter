# Go Strings Segmenter

This is library to break slice of strings into segments iteratively.

For example we have following slice of strings:

```golang
strs := []string{"0", "1", "2", "3", "4"}
```

We want to break this slice of strings into segment that has maximum length of `2`. So we would have following segments iteratively:

```golang
[]string{"0", "1"} // first iteration
[]string{"2", "3"} // second iteration
[]string{"4"} // third iteration
```

We could use this library to do just that.

## Sample Usage

```golang
func main() {
    sgmntr := segmenter.NewSegmenter(segmenter.Configs{
        Strings:       []string{"0", "1", "2", "3", "4"},
        SegmentLength: 2
    })
    for sgmntr.HasNext() {
        log.Println(sgmntr.Next())
    }
    // The result would be:
    //
    // []string{"0", "1"}
    // []string{"2", "3"}
    // []string{"4"}
}
```

For more examples on usage please check `segmenter_test.go`.

[Back to Top](#go-strings-segmenter)

---
