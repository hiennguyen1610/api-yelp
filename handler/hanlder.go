package handler

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/hiennguyen1610/api-yelp/service"
	"encoding/json"
)

func Log(handle httprouter.Handle) httprouter.Handle{
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Printf("%v", r)
		handle(w, r, p)
	}
}

// Request: /search?name=abc&location=xyz
func Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	srvc := service.NewAddressService()
	addresses := srvc.Find(r.URL.Query().Get("name"), r.URL.Query().Get("location"));
	fmt.Printf("%v", addresses)
	
	output, err := json.MarshalIndent(&addresses, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	}
	w.Write(output)
}
