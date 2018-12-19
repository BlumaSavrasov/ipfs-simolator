package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var GlobalId=0
type Node struct {
	Id int
	Content [] int
	NumOfBlocks int
}

func NewNode(numOfBlocks int) *Node {
	GlobalId++
	node := new(Node)
	node.Id=GlobalId
	for i:=0;i<numOfBlocks;i++{
		node.Content=append(node.Content,rand.Intn(1000))
	}
	node.NumOfBlocks=numOfBlocks
	return node
}

func (n Node) String() string {
	return fmt.Sprintf("[node id: %d,number of blocks: %d content: %d]", n.Id,n.NumOfBlocks,n.Content)
}
type World struct {
	NumOfNodes int
	SetOfNodes []Node
}

func NewWorld(numOfNodes int) *World {
	world := new(World)
	world.NumOfNodes=numOfNodes
	for i:=0;i<numOfNodes;i++{
		numOfBlocks:= rand.Intn(100)
		world.SetOfNodes=append(world.SetOfNodes, *NewNode(numOfBlocks))
	}
	return world
}

func (w World) String() string {
	 var setOfNodes []string
	for i:=0;i<w.NumOfNodes;i++{
		setOfNodes=append(setOfNodes,w.SetOfNodes[i].String()+"\n")
	}
	return fmt.Sprintf(strings.Join(setOfNodes,"\n"))
}
func main() {
	w:=NewWorld(4)
	fmt.Println(w)
}


