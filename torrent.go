package main

import (
	"io"
	"log"
	"strconv"
)

type Info struct {
	Pieces      string
	PieceLength int
	Length      int
	Name        string
}

type Torrent struct {
	Announce string
	Info     Info
}

// TODO: Maybe we can use seek to avoid reading the whole file into memory?
func decode(r io.Reader) (*Torrent, error) {
	t := &Torrent{}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	d, _ := readDict(0, b)
	log.Printf("%+v", d)
	announce := d["announce"].([2]int)
	t.Announce = string(b[announce[0]:announce[1]])
	// info
	info := d["info"].(map[string]Token)
	pieces := info["pieces"].([2]int)
	t.Info.Pieces = string(b[pieces[0]:pieces[1]])
	pieceLength := info["piece length"].([2]int)
	pl, err := strconv.Atoi(string(b[pieceLength[0]:pieceLength[1]]))
	if err != nil {
		return nil, err
	}
	t.Info.PieceLength = pl
	length := info["length"].([2]int)
	l, err := strconv.Atoi(string(b[length[0]:length[1]]))
	if err != nil {
		return nil, err
	}
	t.Info.Length = l
	name := info["name"].([2]int)
	t.Info.Name = string(b[name[0]:name[1]])
	return t, nil
}
