
package main

import (
	"net";
	"fmt";
	"bufio"
	"os"
)

var g_nStarterGlobalId=0
type Starter struct{
   _nStarterId int
   _listOfOtherStarters []int
   _listOfExecuters []int
   _mastersIP string
   _mastersPort string

}

func newStarter()*Starter{
	starter:=new(Starter)
	starter._nStarterId=g_nStarterGlobalId
	g_nStarterGlobalId++
	starter._mastersIP="127.0.0.1"
	starter._mastersPort="9000"
	return starter
}
type Master struct{
	_nIP string
	_nPortNumber string
}

func newMaster(IP string, portNumber string)*Master{
	master:=new(Master)
	master._nIP=IP
	master._nPortNumber=portNumber
	return master
}

/*func startSimulation(obj var){
	switch  {
	case mode="starter":
        go starterMode()
	case mode="master":
		go masterMode()
	default:
	}
}*/
func masterMode(){
	for{
     //master awaits connections from starters
     //prints the details
     fmt.Println("waiting for connections:)")
     // listen on all interfaces
     ln, _ := net.Listen("tcp", ":9000")
     // accept connection on port
     conn, _ := ln.Accept()
     // run loop forever (or until ctrl-c)
     for{
     	// will listen for message to process ending in newline (\n)
     	message, _ := bufio.NewReader(conn).ReadString('\n')
     	// output message received
     	fmt.Print("got connections from: ", string(message))
     	// sample process for string received
     	conn.Write([]byte("waiting for more!\n"))
     }
	}

}


func starterMode(){
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:9000")
	for { 
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("connected to server")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
  }
}


func main() {
	/*master = newMaster("127.0.0.1","9000")
	var starters[5]Starter
	for i := 0; i < 5; i++ {
		starters[i]=newStarter()
	}*/
	masterMode();
	starterMode();
}
