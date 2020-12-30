package Web

import(
	"cns"
	"net/http"
	"objects"
	"encoding/json"
  "sync"
)

var Mtx sync.RWMutex

func Routes(){

	httpApp := cns.Http{}

	httpApp.Get("/",func(req *http.Request,res http.ResponseWriter){
       	
    Mtx.RLock()
   	msg, err := json.Marshal(objects.Adapters)
    Mtx.RUnlock()

   	if err != nil{
   		return
   	}

   	res.Write(msg)
  })
}