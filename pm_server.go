package main

import(
	"lib"
	"flag"
	"fmt"
	"net"
	"cns"
	"Web"
	"net/http"
	"encoding/json"
	"encoding/binary"
	"objects"
)

func main(){

	defer lib.Handlepanic()

	hostName := flag.String("server", "0.0.0.0:8300", "a string")

	flag.Parse()

	httpApp := cns.Http{}

	Web.Routes()

	objects.Response = objects.Resp{}

	// create socket
	go createSocket(hostName)

    httpApp.DefaultMethod(func(req *http.Request,res http.ResponseWriter){
        res.Header().Set("Name", "Sudeep Dasgupta")
        res.Header().Set("Content-Type", "application/json")
        res.Header().Set("Access-Control-Allow-Origin", "*")
        res.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
    })

    cns.CreateHttpServer(":3600")
}

func createSocket(hostName *string){

	defer lib.Handlepanic()

	// host tcp
	Conn, err := net.Listen("tcp", *hostName)

    if err != nil {
        fmt.Println(err)
        return
    }
    // Close the listener when the application closes.
    defer Conn.Close()

    fmt.Println("Server hosted...")

    for {
        // Listen for an incoming connection.
        conn, err := Conn.Accept()

        if err != nil {
            fmt.Println(err)
            break
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn){

	defer lib.Handlepanic()

	fmt.Println("New client...")

	defer conn.Close()

	// Make a buffer to hold incoming data.
  	buf := make([]byte, 2)

  	var err error

  	for{

  		// Read the incoming connection into the buffer.
	  	_, err = conn.Read(buf)

	  	if err != nil {
		    fmt.Println(err)
		    break
	  	}

	  	msgLen := int16(binary.BigEndian.Uint16(buf))

	  	if msgLen < 0{
	  		continue
	  	}

	  	completePacket := make([]byte, msgLen)

	  	_, err = conn.Read(completePacket)

	  	if err != nil {
		    fmt.Println(err)
		    break
	  	}

	  	// Send a response back to person contacting us.
	  	conn.Write([]byte("PING"))

	  	err = json.Unmarshal(completePacket, &objects.Response)

	  	if err != nil{
	  		fmt.Println(err)
	  		continue
	  	}

	  	key := objects.Response.PsName+"---->"+objects.Response.Ip

	  	Web.Mtx.RLock()
	  	_, ok := objects.Adapters[key]
	  	Web.Mtx.RUnlock()

	  	if !ok{

	  		Web.Mtx.Lock()
	  		objects.Adapters[key] = make(map[string]interface{})
	  		Web.Mtx.Unlock()
	  	}

	  	Web.Mtx.Lock()
	  	objects.Adapters[key]["msg"] = objects.Response.Msg
	  	objects.Adapters[key]["psName"] = objects.Response.PsName
	  	objects.Adapters[key]["ip"] = objects.Response.Ip
	  	objects.Adapters[key]["lut"] = objects.Response.LUT
	  	objects.Adapters[key]["path"] = objects.Response.PsPath
	  	Web.Mtx.Unlock()
  	}  	
}