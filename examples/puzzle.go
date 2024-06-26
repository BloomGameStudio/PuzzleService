package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/**
 * See "MitochondriaSmartContracts" README.md for the private key and corresponding address data used here.
 * https://github.com/BloomGameStudio/MitochondriaSmartContracts
 */
func main() {
	// Create a signer account that matches the account set on the contract
	signerAddress := common.HexToAddress("0x0376AAc07Ad725E01357B1725B5ceC61aE10473c")
	signerPrivateKey, _ := crypto.HexToECDSA(
		strings.TrimPrefix("0x0000000000000000000000000000000000000000000000000000000000000b0b", "0x"),
	)

	// Encode raw solution data to get puzzle identifier
	puzzleType := crypto.Keccak256Hash([]byte("LOCATION")).Bytes()
	x, _ := hex.DecodeString(fmt.Sprintf("%02x", 100))
	y, _ := hex.DecodeString(fmt.Sprintf("%02x", 100))
	z, _ := hex.DecodeString(fmt.Sprintf("%02x", 50))

	locationPuzzleId := crypto.Keccak256Hash(
		append(
			append(
				append(
					puzzleType,
					common.LeftPadBytes(x, 32)...,
				),
				common.LeftPadBytes(y, 32)...,
			),
			common.LeftPadBytes(z, 32)...,
		),
	)

	// Encode player address and puzzle identifier and hash to the create message digest
	playerAddress, _ := hex.DecodeString(
		strings.TrimPrefix("0xe05fcC23807536bEe418f142D19fa0d21BB0cfF7", "0x"),
	)

	// The message digest the Solidity equivalent of `abi.encode(address player_address, bytes32 puzzle_identifier)`
	data := crypto.Keccak256Hash(
		append(common.LeftPadBytes(playerAddress, 32), locationPuzzleId.Bytes()...),
	)

	// Ethereum signed messages must be prefixed with "\x19Ethereum Signed Message:\n" and the length of the message
	message := fmt.Sprintf(
		"\x19Ethereum Signed Message:\n%d%s",
		len(data),
		string(data[:]),
	)

	// Hash the full message message and sign the hash with the private key
	hash := crypto.Keccak256([]byte(message))
	signature, _ := crypto.Sign(hash, signerPrivateKey)

	fmt.Printf("Signature: %x\n", signature)

	// Verify the signature
	recoveredPublicKey, _ := crypto.Ecrecover(hash, signature)
	recoveredPublicKeyTyped, _ := crypto.UnmarshalPubkey(recoveredPublicKey)
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPublicKeyTyped)

	if recoveredAddress != signerAddress {
		log.Fatalf("Signature verification failed. Recovered address: %s", recoveredAddress.Hex())
	}

	fmt.Println("Signature verified successfully!")
}
