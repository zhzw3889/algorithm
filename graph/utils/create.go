package utils

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// adjacency list
// const MAXVEX = 20

type ArcNode struct {
	adjvex int
	weight int
	next   *ArcNode
}
type VertexNode struct {
	vexdata int
	head    *ArcNode
}
type AdjList struct {
	vertex []VertexNode
	vexnum int
	arcnum int
}

// adjacency matrix
const INFINITY = 32767

type AdjMatrix struct {
	ares   [][]int //边的信息
	vex    []int   //顶点信息
	vexnum int     //顶点数目
	arcnum int     //边的数目
}
