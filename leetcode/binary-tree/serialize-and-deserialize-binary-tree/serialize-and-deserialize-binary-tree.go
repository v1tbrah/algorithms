package second_solution

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	nilNodeMark string

	openNodeMark  string
	lenNodeMacros string
	closeNodeMark string

	valueOpen  string
	valueClose string
}

func Constructor() Codec {
	return Codec{
		nilNodeMark: "\t",

		openNodeMark:  "q",
		lenNodeMacros: "{%d}",
		closeNodeMark: "p",

		valueOpen:  "VO",
		valueClose: "VC",
	}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var result string

	result += c.openNodeMark + c.lenNodeMacros
	if root == nil {
		result += c.nilNodeMark
		result += c.closeNodeMark
		result = fmt.Sprintf(result, 6)
		return result
	}

	result += c.valueOpen
	result += strconv.Itoa(root.Val)
	result += c.valueClose

	result += c.serialize(root.Left)
	result += c.serialize(root.Right)

	result += c.closeNodeMark

	strOfLenOfResult := strconv.Itoa(len(result))
	commLen := len(result) - 2 + len(strOfLenOfResult) // -2 for %d

	result = fmt.Sprintf(result, commLen)

	return result
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	lenNodeInStr := c.getLenNodeInStr(data)
	lenLenNodeStr := len(lenNodeInStr)
	lenNode, err := strconv.Atoi(lenNodeInStr)
	if err != nil {
		panic(fmt.Sprintf("parse value (%s) to int: %s", lenNodeInStr, err.Error()))
	}

	data = data[len(c.openNodeMark)+1+lenLenNodeStr+1 : lenNode] // 1 for "{ and 1 for "}"
	if strings.HasPrefix(data, c.nilNodeMark) {
		return nil
	}

	if len(data) == 0 {
		panic("unexpected behavior")
	}

	node := &TreeNode{}

	data = data[len(c.valueOpen):]
	var valStr string
	for i := 0; i < len(data) && data[i] != c.valueClose[0]; i++ {
		valStr += string(data[i])
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		panic(fmt.Sprintf("parse value (%s) to int: %s", valStr, err.Error()))
	}
	node.Val = val

	data = data[len(valStr)+len(c.valueClose):]

	lenLeftNodeInStr := c.getLenNodeInStr(data)
	lenLeftNode, err := strconv.Atoi(lenLeftNodeInStr)
	if err != nil {
		panic(fmt.Sprintf("parse value (%s) to int: %s", lenLeftNodeInStr, err.Error()))
	}
	node.Left = c.deserialize(data)

	data = data[lenLeftNode:]
	node.Right = c.deserialize(data)

	return node
}

func (c *Codec) getLenNodeInStr(data string) string {
	data = data[len(c.openNodeMark):]

	data = data[1:] // for '{'
	var lenNodeInStr string
	for i := 0; i < len(data) && string(data[i]) != "}"; i++ {
		lenNodeInStr += string(data[i])
	}

	return lenNodeInStr
}
