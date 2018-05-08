package blockchain

// Blockchain holds a list of blocks
type Blockchain struct {
	blocks []*Block
}

// GetBlocks returns a pinter to blocks slice
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}

// AddBlock adds specified Block to the chain
func (bc *Blockchain) AddBlock(data *BlockData) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a blockchain with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{newGenesisBlock()}}
}

func newGenesisBlock() *Block {
	return NewBlock(&BlockData{"Genesis Block"}, []byte{0})
}
