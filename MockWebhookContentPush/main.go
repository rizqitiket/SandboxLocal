package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	agodaOTAHotelDescriptiveContentNotifRS struct {
		XMLName       xml.Name     `xml:"OTA_HotelDescriptiveContentNotifRS"`
		TimeStamp     time.Time    `xml:"TimeStamp,attr,omitempty"`
		Target        string       `xml:"Target,attr,omitempty"`
		Version       string       `xml:"Version,attr,omitempty"`
		PrimaryLangID string       `xml:"PrimaryLangID,attr,omitempty"`
		Error         *ErrorResult `xml:"Error"`
		CorrelationID string       `xml:"CorrelationID,attr"`
		Errors        *Errors      `xml:"Errors,omitempty"`
		Success       *string      `xml:"Success,omitempty"`
		UniqueID      *UniqueID    `xml:"UniqueID,omitempty"`
		Warnings      *Warnings    `xml:"Warnings,omitempty"`
	}

	agodaOTAHotelSummaryNotifRS struct {
		XMLName       xml.Name  `xml:"OTA_HotelSummaryNotifRS"`
		CorrelationID string    `xml:"CorrelationID,attr"`
		Errors        *Errors   `xml:"Errors,omitempty"`
		Success       *string   `xml:"Success,omitempty"`
		UniqueID      *UniqueID `xml:"UniqueID,omitempty"`
		Warnings      *Warnings `xml:"Warnings,omitempty"`
	}

	agodaOTAHotelInvNotifRS struct {
		XMLName            xml.Name                 `xml:"OTA_HotelInvNotifRS"`
		CorrelationID      string                   `xml:"CorrelationID,attr"`
		TimeStamp          time.Time                `xml:"TimeStamp,attr,omitempty"`
		Target             string                   `xml:"Target,attr,omitempty"`
		Version            string                   `xml:"Version,attr,omitempty"`
		PrimaryLangID      string                   `xml:"PrimaryLangID,attr,omitempty"`
		Error              *ErrorResult             `xml:"Error"`
		Errors             *Errors                  `xml:"Errors,omitempty"`
		Success            *string                  `xml:"Success,omitempty"`
		UniqueID           *UniqueID                `xml:"UniqueID,omitempty"`
		Warnings           *Warnings                `xml:"Warnings,omitempty"`
		InventoryCrossRefs *agodaInventoryCrossRefs `xml:"InventoryCrossRefs,omitempty"`
	}

	agodaInventoryCrossRefs struct {
		RequestInvCode  string `xml:"RequestInvCode,attr"`
		ResponseInvCode string `xml:"ResponseInvCode,attr"`
	}
)

type (
	Errors struct {
		Error []ErrorResult `xml:"Error"`
	}

	ErrorResult struct {
		Status    string `xml:"Status,attr"`
		ShortText string `xml:"ShortText,attr"`
		Code      string `xml:"Code,attr"`
	}

	Warnings struct {
		Warning []Warning `xml:"Warning"`
	}

	Warning struct {
		Status    string `xml:"Status,attr"`
		ShortText string `xml:"ShortText,attr"`
		Code      string `xml:"Code,attr"`
	}

	UniqueID struct {
		ID   string `xml:"ID,attr"`
		Type string `xml:"Type,attr"`
	}
)

const (
	error424  = `<OTA_HotelDescriptiveContentNotifRS CorrelationID="44173250-eedc-11ed-939b-6fbffb4c259c"><Errors><Error Status="Internal Server Error" Code="424" ShortText="something wrong"/></Errors></OTA_HotelDescriptiveContentNotifRS>`
	error9000 = `<OTA_HotelDescriptiveContentNotifRS CorrelationID="44173250-eedc-11ed-939b-6fbffb4c259c"><Errors><Error Status="Internal Server Error" Code="9000" ShortText="something wrong"/></Errors></OTA_HotelDescriptiveContentNotifRS>`
	succeed   = `<OTA_HotelDescriptiveContentNotifRS CorrelationID="44173250-eedc-11ed-939b-6fbffb4c259c"><Success/></OTA_HotelDescriptiveContentNotifRS>`
)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/push-content", handlerPushContent)
	e.Logger.Fatal(e.Start(":1323"))
}

func handlerPushContent(ec echo.Context) error {
	errorString := []string{
		error424,
		error9000,
		succeed,
	}
	idx := rand.Uint32() % 3

	var v agodaOTAHotelDescriptiveContentNotifRS
	err := xml.Unmarshal([]byte(errorString[idx]), &v)
	if err != nil {
		return ec.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(v)
	return ec.XML(http.StatusOK, v)
}
