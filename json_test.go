package kgoutil

import (
    "testing"
    // "fmt"
    // "reflect"
)

func TestJsonDump(t *testing.T) {
    m := map[string]interface{}{"haha":1 , "dada":"kaka"}
    JsonDump("/tmp/test.tmp" , m);
}
func TestJsonLoad(t *testing.T) {
    m := JsonLoad("/tmp/test.tmp");
    // fmt.Println(m)
    if m["dada"] != "kaka" {
        t.Fatal("error")
    }
    haha := m["haha"]
    // fmt.Println(reflect.TypeOf(m["haha"]))
    v , ok := haha.(float64)
    // fmt.Println(ok)
    // fmt.Println(v)
    if !ok {
        t.Fatal("error")
    }
    if v != 1 {
        t.Fatal("error")
    }
}

