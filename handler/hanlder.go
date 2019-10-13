package handler

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/hiennguyen1610/api-yelp/service"
	"encoding/json"
)

func Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	srvc := service.NewAddressService()
	addresses := srvc.Find("", "");
	fmt.Printf("%v", addresses)
	
	output, err := json.MarshalIndent(&addresses, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	}
	w.Write(output)
}
