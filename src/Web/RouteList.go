package Web

import(
	"cns"
	"net/http"
	"objects"
	"encoding/json"
)

func Routes(){

	httpApp := cns.Http{}

	httpApp.Get("/",func(req *http.Request,res http.ResponseWriter){
       	
       	msg, err := json.Marshal(objects.Adapters)

       	if err != nil{
       		return
       	}

       	res.Write(msg)
    })
}