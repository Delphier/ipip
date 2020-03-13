# ipip
Golang wrapper for ipip.net ipdb file and free API

## Install
```
go get github.com/delphier/ipip
```

## Example: ipdb file
```golang
if err := ipip.Open("./ipipfree.ipdb"); err != nil {
    log.Fatal(err)
}
// or
// ipip.FatalOpen("./ipipfree.ipdb")

fmt.Println(ipip.Get("118.28.8.8")) // Output: 天津
```

## Example: free API
```golang
fmt.Println(ipip.APIGet("211.143.12.218")) // Output: 湖南 邵阳
```
