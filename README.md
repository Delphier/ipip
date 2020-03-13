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

fmt.Println(ipip.Get("118.28.8.8"))     // Output: 天津
fmt.Println(ipip.Get("211.143.12.218")) // Output: 湖南 邵阳
fmt.Println(ipip.Get("67.230.161.196")) // Output: 美国
fmt.Println(ipip.Get("::1"))            // Output: 
```

## Example: free API
```golang
fmt.Println(ipip.APIGet("118.28.8.8"))     // Output: 天津
fmt.Println(ipip.APIGet("211.143.12.218")) // Output: 湖南 邵阳
fmt.Println(ipip.APIGet("67.230.161.196")) // Output: 美国 加利福尼亚州 洛杉矶
fmt.Println(ipip.APIGet("::1"))            // Output: 保留地址
```
