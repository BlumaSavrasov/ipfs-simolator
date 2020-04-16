import (
	"fmt"
	"github.com/reiver/go-telnet"
	"strings"
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
	starter=new(Starter)
	starter._nStarterId=g_nStarterGlobalId
	g_nStarterGlobalId++
	_mastersIP="127.0.0.1"
	_mastersPort="9000"
}
type Master struct{
	_nIP string
	_nPortNumber string
}

func newMaster(IP string, portNumber string){
	_nIP=IP
	_nPortNumber=portNumber
}

func startSimulation(obj var){
	switch  {
	case mode="starter":
        go starterMode()
	case mode="master":
		go masterMode()
	default:
	}
}
func masterMode(){
	handler:=createServer()
	for{
		
	}

}
func createServer(){
	handler := telnet.EchoHandler
		err := telnet.ListenAndServe(_nIP+":"+_nPortNumber, handler)
		if nil != err {
			panic(err)
		}
	return handler
}

func starterMode(){
	var starter telnet.Caller = telnet.StandardCaller
	telnet.DialToAndCall("127.0.0.1:9000", starter)
}


func main() {
	master = newMaster("127.0.0.1","9000")

	var starters[5]Starter
	for i := 0; i < 5; i++ {
		starters[i]=newStarter()
	}

}