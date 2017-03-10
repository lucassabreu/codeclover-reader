package parser

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

type Coverage struct {
    Generated string `xml:"generated,attr"`
    Projects []Project `xml:"project>"`
}

type Project struct {
    Timestamp int `xml:"timestamp,attr"`
}

func GetObjectFromXML(fileName *string) (cover Coverage, err error) {
    f, err := os.Open(*fileName)
    if err != nil {
        return cover, err
    }
    defer f.Close()

    b, _ := ioutil.ReadAll(f)

    xml.Unmarshal(b, &cover)

    fmt.Printf("Generated: %s\n", cover.Generated)
    for _, p := range cover.Projects {
        fmt.Printf("\t%d\n", p.Timestamp)
    }

    return cover, err
}

