package main

import (
	"fmt"
	"math/rand"
	"strings"
)
//*********************************CLASS CONNECTION****************************************
type Connection struct{
	_nFirstNodeId int
	_nSecondNodeId int
	_bUpSpeed int // speed of upload - in bytes/sec
	_bDownSpeed int // speed of download - in bytes/sec
	_dbLatency int // in microseconds
}
func NewConnection(nFirstNodeId int,nSecondNodeId int)*Connection{
	connection:=new(Connection)
	connection._nFirstNodeId=nFirstNodeId
	connection._nSecondNodeId=nSecondNodeId
	// SKA: The latency should be calculated based on ASes the First and Second nodes belong to
	// SKA: up/down speeds should be calculated based on initial nodes setup
	connection._dbLatency=rand.Intn(20)
	return connection
}
func (c Connection) String() string {
	return fmt.Sprintf("node: %d connected to node: %d with latency: %d",c._nFirstNodeId,c._nSecondNodeId,c._dbLatency)
}
//*******************************END OF CLASS CONNECTION**************************************

//******************************CLASS NODE****************************************************
var g_nGlobalId=0
type Node struct {
	_nId int
	_nMyASId int //++++++++++++++++++++++++++++++++++++++ // SKA: IMHO not needed
	_nMyCountryId int //+++++++++++++++++++++++++++++++++++++++need to had in constuctor // SKA: IMHO not needed
	_rgContent [] int
	_nNumOfBlocks int
	_rgConnections []Connection
	_cNumOfConnections int
}

func NewNode(numOfBlocks int) *Node {
	g_nGlobalId++
	fmt.Println(g_nGlobalId)
	node := new(Node)
	node._nId=g_nGlobalId
	node._rgContent=make([]int,numOfBlocks)
	for i:=0;i<numOfBlocks;i++{
		node._rgContent[i]=rand.Intn(1000)
	}
	node._cNumOfConnections=-1//rand.Intn(10-3)+3
	node._rgConnections=nil//make([]Connection,node._cNumOfConnections)
	//node.fnSetConnections()
	node._nNumOfBlocks=numOfBlocks
	return node
}

// SKA TDB: calculate node up/down speed and latency
// Do not create connections, they belong to runtime
// content is also something we should "generate" during runtime.
func (n Node) String() string {
	var node []string
	node=append(node,fmt.Sprintf("[node with id: %d,",n._nId))
	if n._rgConnections!= nil{node = append(node,fmt.Sprintf("connected to nodes:"))}
	for i:=0;i<n._cNumOfConnections;i++{
		node = append(node,n._rgConnections[i].String())
	}

	node=append(node,fmt.Sprintf("number of blocks: %d content: %d]", n._nNumOfBlocks,n._rgContent))
	return fmt.Sprintf(strings.Join(node,"\n"))
}
//*******************************************END OF CLASS NODE**********************************************

//*****************************************CLASS ASN********************************************************
var g_nGlobalASN=0
type AS struct{
	_rgSetOfNodes []Node
	_numOfNodes_n int
	_ASN int
}
// SKA TBD: add amount of nodes arg
func NewAS(numOfNodes_n int)*AS{
	fmt.Println("newAS")
	AS:=new(AS)
	g_nGlobalASN++
	// SKA TBD: create nodes as per args
	AS._rgSetOfNodes=make([]Node,numOfNodes_n)
	AS._ASN=g_nGlobalASN
	for  i:=0 ;i<numOfNodes_n ;i++  {
		AS._rgSetOfNodes[i]=*NewNode(rand.Intn(100))
	}
	AS._numOfNodes_n=numOfNodes_n
	return AS
}
func (a AS) String() string {
	var as []string
	as=append(as,fmt.Sprintf("AS with ASN:%d, Contains nodes:\n",a._ASN))
	for i:=0;i<a._numOfNodes_n ;i++  {
		as=append(as,a._rgSetOfNodes[i].String()+"\n")
	}
	return fmt.Sprintf(strings.Join(as,"\n"))
}
//*********************************END OF CLASS ASN***********************************************************

//*********************************CLASS COUNTRY**************************************************************
type Country struct{
	_Name string
	//_CentralAS AS //SKA: no need. Assume that all ASes in the same country have 10ms latency and unlimited bandwidth
	_rgAS []AS
	_numOfAS int
}
// SKA: TBD: add more variables
func newCountry(numOfAS int) *Country {
	fmt.Println("newCountry")
	country:=new(Country)
	country._Name=""
	country._numOfAS=numOfAS
//	country._CentralAS=*NewAS(rand.Intn(100))
	country._rgAS=make([]AS,numOfAS)
	for i:=0;i<country._numOfAS ;i++  {
		// SKA: TBD: create ASes as per args.
		country._rgAS[i]=*NewAS(rand.Intn(100))
	}
	return country
}
func (c Country) String() string {
	var country []string
	country = append(country,fmt.Sprintf("country's name: %d \n",c._Name))
	country=append(country,c._CentralAS.String()+"\n")
	for i:=0;i<c._numOfAS;i++{
		//world=append(world,w._rgSetOfNodes[i].String()+"\n")
		country=append(country,c._rgAS[i].String()+"\n")
	}

	return fmt.Sprintf(strings.Join(country,"\n"))
}
//*****************************END OF CLASS COUNTRY*********************************************************
var g_world []AS

func newWorld(){
	// Append to world country USA with 100 ASes with 1000 nodes in each, where 20% of nodes
	// are on low-latency colos and rest are ADSL users
	g_world=append(NewCountry("USA",100,1000,20))
	g_world=append(NewCountry("Israel",10, 100, 5))
}
func main() {
	// SKA: invoke create_world function
	// It then creates (in this order) Countries, ASes, Nodes with flexible distribution
	
}


