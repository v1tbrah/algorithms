package second_solution

import (
	"bytes"
	"encoding/binary"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	nilNodeMark byte

	openNodeMark     byte
	lenNodeMacrosVal byte
	closeNodeMark    byte

	valueOpen  byte
	valueClose byte
}

func Constructor() Codec {
	return Codec{
		nilNodeMark: 128,

		openNodeMark:     112,
		lenNodeMacrosVal: 124,
		closeNodeMark:    113,

		valueOpen:  123,
		valueClose: 125,
	}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) []byte {
	var result []byte

	result = append(result, c.openNodeMark, c.lenNodeMacrosVal)
	if root == nil {
		result = append(result, c.valueOpen, c.nilNodeMark, c.valueClose, c.closeNodeMark)
		return c.expandLenNodeMacros(result)
	}

	result = append(result, c.valueOpen)

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, int64(root.Val)); err != nil {
		panic(err)
	}
	result = append(result, buf.Next(buf.Len())...)

	result = append(result, c.valueClose)

	result = append(result, c.serialize(root.Left)...)
	result = append(result, c.serialize(root.Right)...)

	result = append(result, c.closeNodeMark)

	return c.expandLenNodeMacros(result)
}

func (c *Codec) expandLenNodeMacros(data []byte) []byte {
	lenNode := int64(len(data))
	lenNodeInBytes := binary.AppendVarint(nil, lenNode)

	okCountCellsForLenNodeInBytes := 1
	for okCountCellsForLenNodeInBytes != len(lenNodeInBytes) {
		okCountCellsForLenNodeInBytes += len(lenNodeInBytes) - 1
		lenNode += int64(okCountCellsForLenNodeInBytes) - 1
		lenNodeInBytes = binary.AppendVarint(nil, lenNode)
	}

	var leftOfLenNode, rightOfLenNode []byte

	for i := 0; i < len(data); i++ {
		if data[i] == c.lenNodeMacrosVal {
			leftOfLenNode = data[:i]
			rightOfLenNode = data[i+1:]
			break
		}
	}

	result := make([]byte, 0, len(leftOfLenNode)+len(lenNodeInBytes)+len(rightOfLenNode))
	result = append(result, leftOfLenNode...)
	result = append(result, lenNodeInBytes...)
	result = append(result, rightOfLenNode...)

	return result
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data []byte) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	lenNodeInBytes := c.getLenNodeInBytes(data)

	data = data[1:]                   // c.openNodeMark
	data = data[len(lenNodeInBytes):] //
	data = data[1:]                   // c.valueOpen

	if data[0] == c.nilNodeMark {
		return nil
	}

	node := &TreeNode{}

	var valueBuf bytes.Buffer
	for i := 0; i < len(data) && data[i] != c.valueClose; i++ {
		valueBuf.WriteByte(data[i])
	}

	lenValueBuf := valueBuf.Len()

	var val int64
	if err := binary.Read(&valueBuf, binary.BigEndian, &val); err != nil {
		panic(err)
	}
	node.Val = int(val)

	data = data[lenValueBuf:] //
	data = data[1:]           // c.valueClose

	lenLeftNodeInBytes := c.getLenNodeInBytes(data)
	node.Left = c.deserialize(data)

	data = data[bytesToInt(lenLeftNodeInBytes):]
	node.Right = c.deserialize(data)

	return node
}

func (c *Codec) getLenNodeInBytes(data []byte) []byte {
	var buf bytes.Buffer

	for i, iAfterOpenNodeMark := 0, false; i < len(data) && data[i] != c.valueOpen; i++ {
		if iAfterOpenNodeMark {
			buf.WriteByte(data[i])
			continue
		}
		if data[i] == c.openNodeMark {
			iAfterOpenNodeMark = true
			continue
		}
	}

	return buf.Bytes()
}

func bytesToInt(srcBytes []byte) (result int64) {
	result, _ = binary.Varint(srcBytes)
	return result
}
