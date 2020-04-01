package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"

	"github.com/the729/go-libra/crypto"
	"github.com/the729/go-libra/types"
)

type Account struct {
	PrivateKey     ed25519.PrivateKey
	Address        types.AccountAddress
	AuthKey        ed25519.PublicKey
	SequenceNumber uint64
}

type AccountConfig struct {
	PrivateKey crypto.PrivateKey `toml:"private_key"`
	AuthKey    crypto.PublicKey  `toml:"auth_key"`
	Address    []byte            `toml:"address"`
}

type WalletConfig struct {
	Accounts []*AccountConfig `toml:"accounts"`
}

type SimpleWallet struct {
	Accounts map[string]*Account
}

func LoadAccounts(file string) (*SimpleWallet, error) {
	walletConf := &WalletConfig{}
	_, err := toml.DecodeFile(file, walletConf)
	if err != nil {
		return nil, fmt.Errorf("toml decode file error: %v", err)
	}

	wallet := &SimpleWallet{
		Accounts: make(map[string]*Account),
	}
	for _, accountConf := range walletConf.Accounts {
		account := &Account{
			PrivateKey: ed25519.PrivateKey(accountConf.PrivateKey),
			AuthKey:    ed25519.PublicKey(accountConf.AuthKey),
		}
		if len(accountConf.Address) == types.AccountAddressLength {
			copy(account.Address[:], accountConf.Address)
		}
		if accountConf.PrivateKey != nil {
			pubkey := account.PrivateKey.Public().(ed25519.PublicKey)
			hasher := sha3.New256()
			hasher.Write(pubkey)
			hasher.Write([]byte{0})
			account.AuthKey = hasher.Sum([]byte{})
			if len(accountConf.Address) != types.AccountAddressLength {
				copy(account.Address[:], account.AuthKey[16:])
			}
		}
		wallet.Accounts[hex.EncodeToString(account.Address[:])] = account
	}
	return wallet, nil
}

func (w *SimpleWallet) GetAccount(prefix string) (*Account, error) {
	var seen *Account
	for addr, account := range w.Accounts {
		if strings.HasPrefix(string(addr[:]), prefix) {
			if seen != nil {
				return nil, fmt.Errorf("more than 1 accounts have prefix %s", prefix)
			}
			seen = account
		}
	}
	if seen != nil {
		return seen, nil
	}

	newAddr, err := hex.DecodeString(prefix)
	if err == nil && len(newAddr) == types.AccountAddressLength {
		a := &Account{}
		copy(a.Address[:], newAddr)
		return a, nil
	}

	return nil, fmt.Errorf("account not present in local config file")
}
