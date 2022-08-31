package geth10

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	ethTrie "github.com/ethereum/go-ethereum/trie"
	"github.com/stretchr/testify/require"
)

func TestStorageTrie(t *testing.T) {
	fmt.Println("RUNNING TEST: TestStorageTrie")
	fmt.Println("IN FILE: storage_proof_test.go")

	// slot indexes
	// slot0 := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000")
	// 0x4e46545475746f7269616c000000000000000000000000000000000000000016
	// slot1 := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001")
	// 0x4e46540000000000000000000000000000000000000000000000000000000006
	// slot6: common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000006")
	// 0x02
	// slot3: common.FromHex("0x679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c")
	// 0x2813736e6204ee248e79c26de69d49bddbe0f7d0
	// slot4: common.FromHex("0x211f3d93987b5218a32eac3af2b87ceafa2dad2bbbdbe3688f9c11e352c27cd8")
	// 0x02

	// contract creation transaction (goerli):
	// https://goerli.etherscan.io/tx/0x8de2944e0c7bd6e93753f6c055984cfab5f1e97fad36c580327bca3fe61457a1#statechange
	// mint NFT index 0:
	// https://goerli.etherscan.io/tx/0x27bfae83807487dbf4a24b944ec14cb5a430b5a4024e05c3e6ef8179af6d2299#statechange
	// mint NFT index 1:
	// https://goerli.etherscan.io/tx/0x628f3ae6afe1497c3c792c981adbfed4341911f0c99155af40bd4160a0c760f6#statechange

	nameSlot := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000")
	symbolSlot := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001")
	counterSlot := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000006")
	balancesSlot := common.FromHex("0x211f3d93987b5218a32eac3af2b87ceafa2dad2bbbdbe3688f9c11e352c27cd8")
	owner0Slot := common.FromHex("0xe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0")
	owner1Slot := common.FromHex("0x679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c")

	// name, err := rlp.EncodeToBytes([]byte("NFTTutorial")) // common.FromHex("0x4e46545475746f7269616c000000000000000000000000000000000000000016")
	namehex, err := hex.DecodeString("4e46545475746f7269616c000000000000000000000000000000000000000016") // need to use this! the appending of the length from RLP seems to be required. Cannot simply use the byte string array
	require.NoError(t, err)
	name, err := rlp.EncodeToBytes(namehex)
	require.NoError(t, err)
	// symbol, err := rlp.EncodeToBytes([]byte("NFT")) // common.FromHex("0x4e46540000000000000000000000000000000000000000000000000000000006")
	symbolhex, err := hex.DecodeString("4e46540000000000000000000000000000000000000000000000000000000006") // need to use this! the appending of the length from RLP seems to be required. Cannot simply use the byte string array
	require.NoError(t, err)
	symbol, err := rlp.EncodeToBytes(symbolhex)
	require.NoError(t, err)
	counter1, err := rlp.EncodeToBytes(common.FromHex("0x01"))
	require.NoError(t, err)
	counter2, err := rlp.EncodeToBytes(common.FromHex("0x02"))
	require.NoError(t, err)
	ownerAddress, err := rlp.EncodeToBytes(common.FromHex("0x2813736e6204ee248e79c26de69d49bddbe0f7d0"))
	require.NoError(t, err)
	balances1, err := rlp.EncodeToBytes(common.FromHex("0x01")) // need to use 0x02 instead of padded hex (identical to the one provided in the "Slot")
	require.NoError(t, err)
	balances2, err := rlp.EncodeToBytes(common.FromHex("0x02")) // need to use 0x02 instead of padded hex (identical to the one provided in the "Slot")
	require.NoError(t, err)

	fmt.Println("nameSlot: ", nameSlot)
	fmt.Println("symbolSlot: ", symbolSlot)
	fmt.Println("owner0Slot: ", owner0Slot)
	fmt.Println("owner1Slot: ", owner1Slot)
	fmt.Println("balanceSlot: ", balancesSlot)
	fmt.Println("counterSlot: ", counterSlot)

	fmt.Println("name: ", name)
	fmt.Println("symbol: ", symbol)
	fmt.Println("ownerAddress ", ownerAddress)
	fmt.Println("balances: ", balances1)
	fmt.Println("balances: ", balances2)
	fmt.Println("counter1: ", counter1)
	fmt.Println("counter2: ", counter2)

	fmt.Println("name to hex: ", common.Bytes2Hex(name))
	fmt.Println("symbol to hex: ", common.Bytes2Hex(symbol))
	fmt.Println("ownerAddress to hex: ", common.Bytes2Hex(ownerAddress))
	fmt.Println("balances to hex: ", common.Bytes2Hex(balances1))
	fmt.Println("balances to hex: ", common.Bytes2Hex(balances2))
	fmt.Println("counter1 to hex: ", common.Bytes2Hex(counter1))
	fmt.Println("counter2 to hex: ", common.Bytes2Hex(counter2))

	// create a trie and store the key-value pairs, the key needs to be hashed
	trie := ethTrie.NewEmpty(&ethTrie.Database{})
	trie.Update(crypto.Keccak256(nameSlot), name)
	trie.Update(crypto.Keccak256(symbolSlot), symbol)
	// trie.Put(crypto.Keccak256(counterSlot), counter1) // overwritten trie slots do not affect the hash
	// trie.Put(crypto.Keccak256(balancesSlot), balances1) // overwritten trie slots do not affect the hash
	trie.Update(crypto.Keccak256(owner0Slot), ownerAddress)
	trie.Update(crypto.Keccak256(counterSlot), counter2)
	trie.Update(crypto.Keccak256(balancesSlot), balances2)
	trie.Update(crypto.Keccak256(owner1Slot), ownerAddress)

	// compute the root hash and check if consistent with the storage hash of contract 0xcca577ee56d30a444c73f8fc8d5ce34ed1c7da8b
	rootHash := trie.Hash()
	storageHash := common.BytesToHash(common.FromHex("0xcf1e4b90f815964e5f79b713232d0cfb7bb54617e7775bededbc4bd9d96c0fad"))

	fmt.Println("storageHash: ", fmt.Sprintf("%+x", storageHash))
	fmt.Println("rootHash:", fmt.Sprintf("%+x", rootHash))

	require.Equal(t, storageHash, rootHash)
}
