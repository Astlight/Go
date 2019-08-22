package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

/*一个简单的区块链创建*/
type Block struct {
	//创建一个区块的结构体
	Time         int64  //时间戳
	Data         []byte //数据信息
	PreviousHash []byte //前一个哈希值
	Hash         []byte //当前的哈希
}
type Blockchain struct { //创建一个区块链类型
	blocks []*Block //一系列区块储存，这里用切片来保存
}

func (blockchain *Blockchain) Addblock(data string) { //添加区块的方法
	newblock := Block{}                                                      //一个新区块
	newblock.Data = []byte(data)                                             //初始化数据
	newblock.PreviousHash = blockchain.blocks[len(blockchain.blocks)-1].Hash //得到前一个区块的哈希
	newblock.Sethash()                                                       //得到哈希
	blockchain.blocks = append(blockchain.blocks, &newblock)                 //新区块添加到这条链当中

}
func Newblockchain() *Blockchain { // 创建一条新的区块链
	blos := Blockchain{[]*Block{Firstblock()}} //初始化
	return &blos
}
func Firstblock() *Block {
	firstblock := Newblock("firstblock", []byte{})
	return firstblock
}
func (block *Block) Sethash() {
	timer := []byte(strconv.FormatInt(block.Time, 10))
	herds := bytes.Join([][]byte{timer, []byte(block.Data), block.PreviousHash}, []byte{})
	hash := sha256.Sum256(herds)
	block.Hash = hash[:]
}
func Newblock(data string, prevhash []byte) *Block {
	block := Block{}
	block.Time = time.Now().Unix()
	block.Data = []byte(data)
	block.PreviousHash = prevhash
	block.Sethash()

	return &block
}
func main() {
	//创建一个区块链
	blocks := Newblockchain()
	blocks.Addblock("seng one BTC to sary")   //信息（数据）为seng one BTC to sary 添加到区块链中
	blocks.Addblock("send two ETH to wuman")  //信息（数据）为send two ETH to wuman 添加到区块链中
	blocks.Addblock("send one ADA to zijian") //信息（数据）为send one ADA to zijian 添加到区块链中
	for _, v := range blocks.blocks {         //循环遍历打印一下看结果
		fmt.Println("=======================================")
		fmt.Printf("data=:%s\n", v.Data)
		fmt.Printf("prevhash:=%x\n", v.PreviousHash)
		fmt.Printf("hash:=%x\n", v.Hash)
	}
}
