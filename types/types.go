package types

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

type APIResponse struct {
	Payload string `json:"d"`
}

type POI struct {
	POIs [][]string
}

func NewPOI() *POI {
	return &POI{}
}

func (p *POI) DecodeFromString(payload string) {
	p1 := strings.Split(payload, "addPOIMarker(")
	for _, v := range p1 {
		p2 := strings.Split(v, ");")
		if len(p2) > 0 {
			p3 := strings.Split(p2[0], ",")
			if len(p3) == 7 {
				p.POIs = append(p.POIs, p3)
			}
		}
	}
}

type Javascript struct {
	Javascripts []string
}

func NewJavascript() *Javascript {
	return &Javascript{}
}

func (j *Javascript) DecodeFromString(payload string) {
	p1 := strings.Split(payload, "addSchoolMarker(")
	for _, v := range p1 {
		p2 := strings.Split(v, ");")
		if len(p2) > 0 {
			j.Javascripts = append(j.Javascripts, p2[0])
		}
	}
}

type InfoWindowDetails struct {
	Details []string
}

func NewInfoWindowDetails() *InfoWindowDetails {
	return &InfoWindowDetails{}
}

func (i *InfoWindowDetails) DecodeFromString(payload string) {
	p1 := strings.Split(payload, "popUpInfoWindow(")
	if len(p1) > 1 {
		r := strings.NewReader(p1[1])
		n, err := html.Parse(r)
		if err != nil {
			log.Print(err)
		}
		i.DecodeFromHTML(n)
	}
}

func (i *InfoWindowDetails) DecodeFromHTML(n *html.Node) (element *html.Node) {
	if n.Type == html.NodeType(1) && len(strings.Replace(n.Data, " ", "", -1)) > 0 {
		i.Details = append(i.Details, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		element = i.DecodeFromHTML(c)
	}
	return
}
