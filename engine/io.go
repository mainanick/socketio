package engine

import "strings"

// SEPARATOR delimiter
const SEPARATOR string = "\x1e" // see https://en.wikipedia.org/wiki/Delimiter#ASCII_delimited_text

// PROTOCOL Engine IO protocal version
const PROTOCOL int = 4

type PacketType int

const (
	PacketOpen PacketType = iota
	PacketClose
	PacketPing
	PacketPong
	PacketMessage
)

type DataType int

const (
	BinaryData DataType = iota
	StringData
)

type Packet struct {
	Type     PacketType
	DataType DataType
	Data     []interface{}
}

func PacketSplit(s string) []string {
	return strings.Split(s, SEPARATOR)
}

func IsBinaryData(p string) bool {
	return string(p[0]) == "b"
}
