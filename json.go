package kgoutil
import (
    "encoding/json"
    "io/ioutil"
)

func JsonLoad(file string) map[string]interface{} {
    b, err := ioutil.ReadFile(file)
    if err != nil { panic(err) }
    result := map[string]interface{}{}
    json.Unmarshal(b , &result)
    return result
}

func JsonDump(file string , m map[string]interface{}) {
    data , _ := json.Marshal(m)
    err := ioutil.WriteFile(file , data , 0644)
    if err != nil { panic(err) }
}

