package fetcher

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Fetcher struct {
	httpClient           http.Client
	scheme               string
	host                 string
	resourcesUrlTemplate url.URL
}

func NewFetcher(scheme, host string) *Fetcher {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 30 * time.Second,
	}
	httpClient := http.Client{Transport: transport}
	resourcesUrlTemplate := url.URL{Scheme: scheme, Host: host}
	return &Fetcher{
		httpClient:           httpClient,
		scheme:               scheme,
		host:                 host,
		resourcesUrlTemplate: resourcesUrlTemplate,
	}
}

func (f *Fetcher) SetHeader(req *http.Request) {
	req.Header.Set("authority", "www.locrating.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("subscriptiondomain", "https://www.locrating.com/school_catchment_areas.aspx")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sticket", "FgKJt5HmLwDL1B4XHrGNRc8rGvxMtpdQbY+3Jp+tB/79CgH8h6ivcmEHRpMT/wsfDu2cPfoxLtlgPFYPykPl0jr3PRdZ8EMW+lreOwOn5ohEk3qn5OnTSJXASHSiS0ogYkQcUBZCuIpNyjPSwskuQemFDOg/qfiV5OIJfABfrgLC6/rXfYU65TtSqVej4FuC2i70R1tBZ/ayIyj53aIjNPx1zE7Eip2J2AtpiVVl/qB6/Kwfs05FV0ZttiHqwbb/Wg0mQrK7+zDARvtB+la1fekdSa0hUrMrvOXN9fkF4+zX4SIyNlnjSvQGOKD6iz//ZnBvfkJgMiPFmO4Z8bIoQH9HSopFWaTBtkJ7c5sW72XA1jNpZqnSlV0QR6weX5UsE+Y9aTNF1+J1qco8UNnSfLEc1HEUX3Re+jzKS30PcbWFpxMeUZ4zeeGLCwZuEOu1zlLj5XqV0edlnacj+/rvwz4+llEfw0kEpODpAGroWIomAohTOXa3piU67ZlOGfnuGBzfM3hAZDU0A5VUdwwIB1Orxh7Of5AnVJLyvprQ656z7VpEAJdCL+VL3xDow2i+oth2DAEmVloIciBBFZxrtAeTFCVSoh2Hn/6xYARkMNc=")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36")
	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("origin", "https://www.locrating.com")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://www.locrating.com/html5/plugin/schoolsmap_osm.aspx?&lat=53.4794892&lng=-2.2451148&zoom=15&showfullscreenbutton=false&showschoolslist=false&id=schoolsmap_frame&pluginmode=website")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8,zh-TW;q=0.7,zh;q=0.6")
	req.Header.Set("cookie", "__cfduid=d0f01e16fcd5ad5ecd5688bf96eb9a60a1598842259; __utmz=261041168.1598842261.5.4.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); amember_ru=anguslou; amember_rp=c3dc12f3c4c964ba7b9ee606ae236547d54e3375; __utma=261041168.658841990.1595914015.1599037305.1599193972.7; __utmc=261041168; __utmt=1; ASP.NET_SessionId=obukspcp30k4c1egpxni4gu1; PHPSESSID=6dilgn7t9tgei56ckkum2a3dj5; wwwlocratingcom-schoolplugin-settings=www.locrating.com%2Fschoolsmap.aspx%3Fschooltype%3D2%26religion%3D0%26rating%3D5%26showindis%3D0%26showspecial%3D0%26schoolgender%3D0%26grammar%3D0%26performancepct%3D100%26ks2progress%3D0%26ks4progress%3D0%26ks5progress%3D0%26oversubscribedpct%3D-1%26shortlist%3D%26maxschools%3D0%26mylocation%3D%26catchmentyear%3D0%26user%3D%26usertype%3D-1%26showonlyincatchment%3D0%26forsale%3Dtrue%26sale_price_min%3D%26sale_price_max%3D350000%26rent_price_min%3D600%26rent_price_max%3D1750%26minbeds%3D1%26distancefromschool%3D0%26distancefromstation%3D0%26propertytype%3Da%26furnished%3D%26isretirementhome%3D%26issharedownership%3D%26petsallowed%3D%26includesold%3D%26isnew%3D%26ischainfree%3D%26propertyfilter%3D%26shortlist%3D%26fetchurl%3Dfalse%26lastupdated%3D0%26zoom%3D16%26lat%3D53.480180266581655%26lng%3D-2.241243124008179%26maptype%3Dcolour%26traffic%3Dfalse%26showproperties%3Dfalse%26schoolsshortlist%3Durn136273%252curn135122%252curn136378%252curn135071%252curn110069%252curn110071%252curn110068%252curn116433%252curn132268%252curn133580%252curn141844%252curn101928%252curn135242%252curn138667%252curn110060%252curn136756%252curn140703%252curn105581%252curn136944%252curn104018%252curn144509%252curn141042%252curn114607%252curn136164%252curn142166%252curn145126%252curn100049%252curn140884%252curn138134%252curn138614%252curn106376%252curn143104%252curn136377%252curn138123%252curn136498%252curn106375%252curn106368%252curn106370%252curn105574%252curn139148%252curn135296%252curn143260%252curn142509%252curn138682%252curn138681%252curn137940%252curn139276%26propertiesshortlist%3D54831833%26showcatchments%3Dtrue%26showcatchmenttrends%3Dtrue%26showonlyincatchments%3Dfalse%26showstations%3Dtrue%26showschoolslist%3Dfalse%26showschoolsshortlist%3Dfalse%26showpropertiesshortlist%3Dfalse%26schoolsinviewsort%3D1%26catchmentyear%3D0%26schoolyear%3D0%26showadvancedcatchments%3Dtrue%26showsocioeconomic%3Dfalse%26showpropertybar%3Dtrue%26showsocioeconomicprops%3Dfalse%26showvideohelp%3Dtrue; __utmb=261041168.54.7.1599194042515")
}

func (f *Fetcher) Fetch(path string, body interface{}) ([]byte, error) {
	f.resourcesUrlTemplate.Path = path
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, f.resourcesUrlTemplate.String(), buffer)
	if err != nil {
		return nil, err
	}
	f.SetHeader(req)
	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK:
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return bytes, err
	case http.StatusNotModified:
		return nil, errors.New(string(res.StatusCode))
	case http.StatusTeapot:
		return nil, errors.New(string(res.StatusCode))
	default:
		return nil, errors.New(string(res.StatusCode))
	}
}
