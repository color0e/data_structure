package main

import "fmt"

type Blockchain struct {
	b       []*Block
	lastnum int
}

type Block struct {
	data     []string
	prevhash int
	hash     int
	num      int
}

func (bc *Blockchain) AddBlock(data string) {
	bc.b = append(bc.b, &Block{[]string{data}, 0, 0, bc.b[bc.lastnum].num + 1})
	bc.lastnum = bc.lastnum + 1
}

func NewGenesisblock() *Block {
	return &Block{[]string{"genesisblock"}, 0, 0, 0}
}

func Newblockchain() *Blockchain {
	bc := &Blockchain{[]*Block{NewGenesisblock()}, 0}
	return bc
}

func main() {
	bc := Newblockchain()
	bc.AddBlock("data")
	for _, v := range bc.b {
		fmt.Println("data:", v.data)
		fmt.Println("prevhash:", v.prevhash)
		fmt.Println("hash:", v.hash)
		fmt.Println("num:", v.num)
	}
	fmt.Println("\nlastnum:", bc.lastnum)
}
