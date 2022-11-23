## compress

## fjson

github.com/json-iterator/go是比标准json包更快的json序列化/反序列化包。

```go
//高效的json序列化函数
JSONFastMarshal

//高效的json反序列化函数
JSONFastUnmarshal

//高效的json序列化函数，并将json数据进行snappy压缩
JSONFastMarshalSnappy

//高效的json反序列化函数，可以反序列snappy压缩的json数据，也可以反序列化普通的json数据
JSONFastUnmarshalSnappy

//读取json文件并反序列化成对象或map
LoadJSONFile2Obj 
```

压缩模块，现在主流的压缩有snappy和zstd，通过对这两者的测试对比，snappy的综合表现不错，
在CPU时间和压缩比上有不错的平衡。

使用方法:

```go
//压缩
var cd = compress.Snappy.Compress(data)

//解压
var dd, err = compress.Snappy.DeCompress(data)

//带前缀压缩，有时我们希望压缩数据内容，但保留数据头，即数据头不用被压缩。
//这样的好处是在某些情况下，可以根据数据头做出一些逻辑，以免所有情况都需要把数据解压。
var hcd = compress.Snappy.CompressWithPrefix(data, yourHead)
```
