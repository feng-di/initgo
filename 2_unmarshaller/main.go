package main

import "encoding/xml"
import "fmt"

const sampleXml = `
<mydoc>
  <soap:test>
	 <testchild></testchild>
  </soap:test>
  <foo>Foo</foo>
  <bar>Bar</bar>
  <foo>Another Foo</foo>
  <foo>Foo #3</foo>
  <bar>Bar 2</bar>
</mydoc>
`

type MyDoc struct {
	XMLName xml.Name `xml:"mydoc"`
	Items   []Item
}

func (md *MyDoc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	md.XMLName = start.Name
	// grab any other attrs

	// decode inner elements
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}

		var i Item
		switch tt := t.(type) {

		case xml.StartElement:

			fmt.Printf("---------------%s\n", tt.Name)

			switch tt.Name.Local {
			case "foo":
				i = new(Foo)
			case "bar":
				i = new(Bar)
				// default: ignored for brevety
			}
			if i != nil {
				err = d.DecodeElement(i, &tt)
				fmt.Println(i)
				if err != nil {
					return err
				}
				md.Items = append(md.Items, i)
				i = nil
			}
		case xml.EndElement:
			if tt == start.End() {
				return nil
			}
		}

	}
	return nil
}

type Item interface {
	IsItem()
	String() string
}

type Foo struct {
	XMLName xml.Name `xml:"foo"`
	Name    string   `xml:",chardata"`
}

func (f Foo) IsItem()        {}
func (f Foo) String() string { return "Foo:<" + f.Name + ">" }

type Bar struct {
	XMLName xml.Name `xml:"bar"`
	Nombre  string   `xml:",chardata"`
}

func (b Bar) IsItem()        {}
func (b Bar) String() string { return "Bar:<" + b.Nombre + ">" }

func main() {
	//doMarshal()
	doUnmarshal()
}

func doMarshal() {
	myDoc := MyDoc{
		Items: []Item{
			Foo{Name: "Foo"},
			Bar{Nombre: "Bar"},
			Foo{Name: "Another Foo"},
			Foo{Name: "Foo #3"},
			Bar{Nombre: "Bar 2"},
		},
	}
	bytes, err := xml.MarshalIndent(myDoc, "", "  ")
	if err != nil {
		panic(err)
	}
	// Prints an XML document just like "sampleXml" above.
	fmt.Println(string(bytes))
}

func doUnmarshal() {
	myDoc := MyDoc{}
	err := xml.Unmarshal([]byte(sampleXml), &myDoc)
	if err != nil {
		panic(err)
	}
	// Fails to unmarshal the "Item" elements into their respective structs.
	fmt.Printf("OK: %v", myDoc)
}
