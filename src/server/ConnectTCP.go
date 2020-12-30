package server

import(
	"lib"
	"net"
	"objects"
	"time"
	"fmt"
	"strconv"
)

var TCPCon net.Conn

var serverIP string

func ConnectTCP(){

	defer lib.Handlepanic()

	fmt.Println(string(objects.ColorGreen), "Connecting to tcp socket server: ", *objects.Conf.Server.Host+":"+strconv.Itoa(objects.Conf.Server.Port))

	var err error

	RECONNECT:

	TCPCon, err = net.Dial("tcp", *objects.Conf.Server.Host+":"+strconv.Itoa(objects.Conf.Server.Port))

    if err != nil {
        fmt.Println(string(objects.ColorRed), err)
        time.Sleep(5 * time.Second)
        goto RECONNECT
    }

    fmt.Println(string(objects.ColorGreen), "Connected to tcp socket server: ", *objects.Conf.Server.Host+":"+strconv.Itoa(objects.Conf.Server.Port))

    ifaces, err := net.Interfaces()

    if err != nil{
    	fmt.Println(string(objects.ColorRed), err)
    	return
    }
	
	for _, address := range ifaces {

        addrs, err := address.Addrs()

	    if err != nil {
	        fmt.Println(string(objects.ColorRed), err)
	        return
	    }
	    
	    for _, addr := range addrs {
	        serverIP += addr.String()+"|"
	    }
    }

	serverIP = serverIP[:len(serverIP) - 1]

	// receive messages from server

	go receive()
}	

// receiving messages from server
func receive(){

	defer lib.Handlepanic()

	if TCPCon == nil{
		ConnectTCP()
		return
	}

	// if socket disconnects then reconnect again
	reply := make([]byte, 4)

	for{

		_, err := TCPCon.Read(reply)

		if err != nil{
			fmt.Println(string(objects.ColorRed), err)
			break
		}

	}

	TCPCon = nil

	time.Sleep(5 * time.Second)

	ConnectTCP()
}