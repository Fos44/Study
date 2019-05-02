package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"
)

const targetBits  = 24
const maxNonce  = math.MaxInt64


func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork  {   //Создаем ProofOfWork
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int)[]byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestump),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
		)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte)  {   //Реализация ядра алгоритма Proof-Of-Work
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1{
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate()bool  {		//Проверка  доказательства работы
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

type Block struct {
	Timestump int64			//Дата создания блока
	Data []byte				//Полезная информация в блоке
	PrevBlockHash []byte	//Хэш предыдущего блока
	Hash []byte 			//Хэш блока
	Nonce int				//Счетчик
}

func (b *Block) SetHash()  {		//Соединяем все поля блока и хэшируем их
	timestamp := []byte(strconv.FormatInt(b.Timestump, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string,prevBlockHash []byte) *Block  {		// Создаем новый блок
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

type Blockchain struct {   //Блокчейн
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string)  {			//Добавление блоков в блокчейн
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block{			//Первый блок
	return NewBlock("Genesis block", []byte{})
}

func NewBlockChain() *Blockchain  {			//Создание нового блокчейна
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Serj")
	bc.AddBlock("Send 2 more BTC to Serj")

	for _, block := range bc.blocks{
		fmt.Printf("Perv. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		pow := NewProofOfWork(block)
		fmt.Printf("Pow %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
	
}
