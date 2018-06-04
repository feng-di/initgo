package main

import (
	"encoding/xml"
	"fmt"
)

type soapRequestEnvelope struct {
	XMLName xml.Name `xml:"http://www.globalprocessing.ae/HyperionWeb soap:Envelope"`
	AttrEnv string   `xml:"xmlns:soap,attr,omitempty"`

	UserName string `xml:"soap:Header>AuthSoapHeader>strUserName,omitempty"`
	Password string `xml:"soap:Header>AuthSoapHeader>strPassword,omitempty"`

	Body interface{}
}

type CreateCardQueryData struct {
	XMLName       xml.Name `xml:"soap:Body><Ws_CreateCard"`
	WSID          string   `xml:"WSID,omitempty"`
	IssCode       string   `xml:"IssCode,omitempty"`
	TxnCode       string   `xml:"TxnCode,omitempty"`
	Title         string   `xml:"Title,omitempty"`
	LastName      string   `xml:"LastName,omitempty"`
	FirstName     string   `xml:"FirstName,omitempty"`
	Addrl1        string   `xml:"Addrl1,omitempty"`
	City          string   `xml:"City,omitempty"`
	PostCode      string   `xml:"PostCode,omitempty"`
	Country       string   `xml:"Country,omitempty"`
	Mobile        string   `xml:"Mobile,omitempty"`
	CardDesign    string   `xml:"CardDesign,omitempty"`
	ExternalRef   string   `xml:"ExternalRef,omitempty"`
	LocDate       string   `xml:"LocDate,omitempty"`
	LocTime       string   `xml:"LocTime,omitempty"`
	LoadValue     string   `xml:"LoadValue,omitempty"`
	CurCode       string   `xml:"CurCode,omitempty"`
	ItemSrc       string   `xml:"ItemSrc,omitempty"`
	LoadFee       string   `xml:"LoadFee,omitempty"`
	CreateImage   string   `xml:"CreateImage,omitempty"`
	CreateType    string   `xml:"CreateType,omitempty"`
	ActivateNow   string   `xml:"ActivateNow,omitempty"`
	ExpDate       string   `xml:"ExpDate,omitempty"`
	CardName      string   `xml:"CardName,omitempty"`
	PERMSGroup    string   `xml:"PERMSGroup,omitempty"`
	ProductRef    string   `xml:"ProductRef,omitempty"`
	Replacement   bool     `xml:"Replacement,omitempty"`
	Delv_AddrL1   string   `xml:"Delv_AddrL1,omitempty"`
	Delv_City     string   `xml:"Delv_City,omitempty"`
	Delv_PostCode string   `xml:"Delv_PostCode,omitempty"`
	Delv_Country  string   `xml:"Delv_Country,omitempty"`
	Lang          string   `xml:"Lang,omitempty"`
	Sms_Required  string   `xml:"Sms_Required,omitempty"`
	PBlock        string   `xml:"PBlock,omitempty"`
	Email         string   `xml:"Email,omitempty"`
	ThermalLine1  string   `xml:"ThermalLine1,omitempty"`
	ThermalLine2  string   `xml:"ThermalLine2,omitempty"`
}

func main() {

	reqEnvelop := soapRequestEnvelope{
		AttrEnv:  "http://schemas.xmlsoap.org/soap/envelope/",
		UserName: "username",
		Password: "password",

		Body: &CreateCardQueryData{
			WSID:      "111111111111111",
			IssCode:   "TZR",
			TxnCode:   "10",
			Title:     "Test-Title",
			LastName:  "",
			FirstName: "",
			Addrl1:    "",
		},
	}

	xmlString, err := xml.Marshal(reqEnvelop)

	fmt.Println(string(xmlString))
	fmt.Printf("---------- : `%+#v`, \n\n\n", string(xmlString))
	fmt.Printf("---------- : `%+#v`, \n\n\n", err)
}