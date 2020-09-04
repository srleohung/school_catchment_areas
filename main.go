package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	. "school_catchment_areas/fetcher"
	"school_catchment_areas/types"
)

func GetPOI(f *Fetcher) *types.POI {
	b, err := f.Fetch("/LocratingWebService.asmx/GetPOI_plugin", map[string]string{
		"center":          "(51.506232, -0.134497)",
		"bounds":          "((51.479298, -0.216894),(51.533149, -0.052099))",
		"latestRequestId": "20",
		"poiTypes":        "Station",
	})
	if err != nil {
		log.Print(err)
		return nil
	}
	res := &types.APIResponse{}
	json.Unmarshal(b, &res)
	GetPOI := types.NewPOI()
	GetPOI.DecodeFromString(res.Payload)
	return GetPOI
}

func GetJavascript(f *Fetcher) *types.Javascript {
	b, err := f.Fetch("/LocratingWebService.asmx/GetJavascript_plugin", map[string]string{
		"center":              "(51.507621, -0.131063)",
		"bounds":              "((51.480688, -0.213461),(51.534537, -0.048666))",
		"schoolType":          "0",
		"religion":            "0",
		"rating":              "5",
		"showIndis":           "0",
		"showSpecial":         "0",
		"schoolGender":        "0",
		"grammar":             "0",
		"performancePct":      "100",
		"ks2Progress":         "0",
		"ks4Progress":         "0",
		"ks5Progress":         "0",
		"oversubscribedPct":   "-1",
		"shortList":           "",
		"latestRequestId":     "4",
		"maxSchools":          "0",
		"myLocation":          "",
		"catchmentYear":       "0",
		"user":                "",
		"userType":            "-1",
		"showOnlyInCatchment": "0",
		"GUID":                "BCB692BD-E5DD-4971-8B1A-E406E494C23F",
	})
	if err != nil {
		log.Print(err)
		return nil
	}
	res := &types.APIResponse{}
	json.Unmarshal(b, &res)
	GetJavascript := types.NewJavascript()
	GetJavascript.DecodeFromString(res.Payload)
	return GetJavascript
}

func GetInfoWindowDetails(f *Fetcher) *types.InfoWindowDetails {
	b, err := f.Fetch("/LocratingWebService.asmx/GetInfoWindowDetails_plugin", map[string]string{
		"id":       "urn101131",
		"GUID":     "BCB692BD-E5DD-4171-8B1A-E406E492C33F",
		"user":     "",
		"userType": "-1",
	})
	if err != nil {
		log.Print(err)
		return nil
	}
	res := &types.APIResponse{}
	json.Unmarshal(b, &res)
	GetInfoWindowDetails := types.NewInfoWindowDetails()
	GetInfoWindowDetails.DecodeFromString(res.Payload)
	return GetInfoWindowDetails
}

func getPOI(w http.ResponseWriter, req *http.Request) {
	fetcher := NewFetcher("https", "www.locrating.com")
	pOI := GetPOI(fetcher)
	b, err := json.Marshal(pOI)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(b))
}

func getJavascript(w http.ResponseWriter, req *http.Request) {
	fetcher := NewFetcher("https", "www.locrating.com")
	pOI := GetJavascript(fetcher)
	b, err := json.Marshal(pOI)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(b))
}

func getInfoWindowDetails(w http.ResponseWriter, req *http.Request) {
	fetcher := NewFetcher("https", "www.locrating.com")
	pOI := GetInfoWindowDetails(fetcher)
	b, err := json.Marshal(pOI)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(b))
}

func main() {
	/*
		fetcher := NewFetcher("https", "www.locrating.com")
		pOI := GetPOI(fetcher)
		log.Printf("pOI.POIs %v", len(pOI.POIs))
		for _, v := range pOI.POIs {
			log.Print(v)
		}
		infoWindowDetails := GetInfoWindowDetails(fetcher)
		log.Printf("infoWindowDetails.Details %v", len(infoWindowDetails.Details))
		for _, v := range infoWindowDetails.Details {
			log.Print(v)
		}
		javascript := GetJavascript(fetcher)
		log.Printf("javascript.Javascripts %v", len(javascript.Javascripts))
		for _, v := range javascript.Javascripts {
			log.Print(v)
		}
	*/
	http.HandleFunc("/getPOI", getPOI)
	http.HandleFunc("/getJavascript", getJavascript)
	http.HandleFunc("/getInfoWindowDetails", getInfoWindowDetails)
	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal(err)
	}
}
