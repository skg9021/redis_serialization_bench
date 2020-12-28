package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/skg9021/proto_size_comparison/sample"
)

func main() {
	fmt.Printf("| %10s | %15s | %10s | %15s | %21s | %40s | \n", "json", " gzipped json", " proto", "gzipped proto", "proto size(%) of json", "gzipped proto size(%) of gzipped json")
	protoStruct := sample.Address{
		Address: "Chunchagatta Main Road, Jaraganahalli Ward, Bommanahalli Zone, Bengaluru, Bangalore South, Bangalore Urban, Karnataka, 560083, India",
		Uid:     "hZEMqQ3kBOgGMngqnl8Z4w",
	}
	jsonl, gzJsonlen, protol, gzProto := jsonProtoLengts(&protoStruct)

	fmt.Printf("| %10d | %15d | %10d | %15d | %21f | %40f | \n", jsonl, gzJsonlen, protol, gzProto, float32(gzProto)/float32(gzJsonlen), float32(protol)/float32(jsonl))
}


func jsonProtoLengts(protoSome *sample.Address) (jsonLen, gzipJSONLen, protoLen, gzpProto int) {
	data, _ := proto.Marshal(protoSome)
	protoLen = len(data)
	jsonified, _ := json.Marshal(protoSome)
	jsonLen = len(jsonified)
	gzipJSONLen = gzipLen(jsonified)
	gzpProto = gzipLen(data)
	return
}

func gzipLen(jsonData []byte) int {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(jsonData); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}

	return b.Len()
}