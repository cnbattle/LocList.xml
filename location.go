package main

import "encoding/xml"

type Location struct {
	XMLName       xml.Name `xml:"Location"`
	Text          string   `xml:",chardata"`
	CountryRegion []struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"Name,attr"`
		Code  string `xml:"Code,attr"`
		State []struct {
			Text string `xml:",chardata"`
			Name string `xml:"Name,attr"`
			Code string `xml:"Code,attr"`
			City []struct {
				Text   string `xml:",chardata"`
				Name   string `xml:"Name,attr"`
				Code   string `xml:"Code,attr"`
				Region []struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name,attr"`
					Code string `xml:"Code,attr"`
				} `xml:"Region"`
			} `xml:"City"`
		} `xml:"State"`
	} `xml:"CountryRegion"`
}
