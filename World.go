package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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
	Connections [][]int
	Seed string
}

func NewWorld(numOfNodes int) *World {
	world := new(World)
	world.SeedGanerator()
	world.NumOfNodes=numOfNodes
	world.SetOfNodesGenerator()
	world.SetConnections()
	return world
}
func (w *World)SeedGanerator(){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<50;i++{
		w.Seed+=string(rand.Intn(127-33)+33)
	}
}
func (w *World) SetConnections(){
	w.Connections=make([][]int,w.NumOfNodes,w.NumOfNodes)
	for i:=0;i<w.NumOfNodes;i++{
		w.Connections[i]=make([]int,w.NumOfNodes)
	}
	for i:=0;i<w.NumOfNodes;i++{
		numOfConnections:=rand.Intn(10-3)+3
		for j:=i;j<numOfConnections;j++ {
			connect:=rand.Intn(w.NumOfNodes-1)+1
			for w.Connections[i][connect]==1 || connect==i{
				connect=rand.Intn(w.NumOfNodes-1)+1
			}
			w.Connections[i][connect]=1
		}
	}


}
//Set
func (w *World) SetOfNodesGenerator(){
	for i:=0;i<w.NumOfNodes;i++{
		w.SetOfNodes=append(w.SetOfNodes, *NewNode(rand.Intn(100)))
	}
}
func (w World) String() string {
	var world []string
	world = append(world,w.Seed+"\n")
	for i:=0;i<w.NumOfNodes;i++{
		world=append(world,w.SetOfNodes[i].String()+"\n")
	}
	for i:=0;i<w.NumOfNodes;i++{
		for j:=i;j<w.NumOfNodes;j++ {
			if w.Connections[i][j]==1 {
				world = append(world,fmt.Sprintf("[%d,%d]",i,j))
			}
		}
	}
	return fmt.Sprintf(strings.Join(world,"\n"))
}
func main() {
	w:=NewWorld(10)
	fmt.Println(w)
}


