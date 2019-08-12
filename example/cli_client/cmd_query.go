package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"github.com/urfave/cli"

	"github.com/the729/go-libra/client"
)

func cmdQueryLedgerInfo(ctx *cli.Context) error {
	c, err := client.New(ServerAddr, TrustedPeersFile)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	knownVersion, err := strconv.ParseUint(ctx.Args().Get(0), 10, 64)
	if err != nil {
		return err
	}
	ledgerInfo, err := c.QueryLedgerInfo(knownVersion)
	if err != nil {
		return err
	}

	log.Printf("Ledger info: version %d, time %d", ledgerInfo.GetVersion(), ledgerInfo.GetTimestampUsec())
	return nil
}

func cmdQueryAccountState(ctx *cli.Context) error {
	c, err := client.New(ServerAddr, TrustedPeersFile)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	wallet, err := LoadAccounts(WalletFile)
	if err != nil {
		log.Fatal(err)
	}

	addr := ctx.Args().Get(0)
	account, err := wallet.GetAccount(addr)
	if err != nil {
		return err
	}
	accountState, err := c.QueryAccountState(account.Address)
	if err != nil {
		return err
	}

	ledgerInfo := accountState.GetLedgerInfo()
	log.Printf("Ledger info: version %d, time %d", ledgerInfo.GetVersion(), ledgerInfo.GetTimestampUsec())

	if accountState.IsNil() {
		fmt.Println("Account is not present in the ledger.")
	} else {
		log.Printf("Account version: %d", accountState.GetVersion())
		resource, err := c.GetLibraCoinResourceFromAccountBlob(accountState.GetAccountBlob())
		if err != nil {
			fmt.Println("Account does not contain libra coin resource.")
			return nil
		}

		log.Printf("Balance (microLibra): %d", resource.GetBalance())
		log.Printf("Sequence Number: %d", resource.GetSequenceNumber())
		log.Printf("SentEventsCount: %d", resource.GetSentEventsCount())
		log.Printf("ReceivedEventsCount: %d", resource.GetReceivedEventsCount())
		log.Printf("DelegatedWithdrawalCapability: %v", resource.GetDelegatedWithdrawalCapability())
	}
	return nil
}

func cmdQueryTransactionRange(ctx *cli.Context) error {
	c, err := client.New(ServerAddr, TrustedPeersFile)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	start, err := strconv.Atoi(ctx.Args().Get(0))
	if err != nil {
		return err
	}
	limit, err := strconv.Atoi(ctx.Args().Get(1))
	if err != nil {
		return err
	}
	if limit > 100 {
		log.Printf("Limit>100, set to 100.")
		limit = 100
	}
	provenTxnList, err := c.QueryTransactionRange(uint64(start), uint64(limit), true)
	if err != nil {
		log.Fatal(err)
	}

	for _, provenTxn := range provenTxnList.GetTransactions() {
		log.Printf("Txn #%d:", provenTxn.GetVersion())
		rawTxn, _ := provenTxn.GetSignedTxn().UnmarshalRawTransaction()
		log.Printf("    Raw txn:")
		log.Printf("        Sender account: %v", hex.EncodeToString(rawTxn.SenderAccount))
		log.Printf("        Sender seq #%v", rawTxn.SequenceNumber)
		log.Printf("        Program: %v...", hex.EncodeToString(rawTxn.GetProgram().Code[:30]))
		log.Printf("        Arg 0: %v", hex.EncodeToString(rawTxn.GetProgram().Arguments[0].Data))
		log.Printf("        Arg 1: %v", hex.EncodeToString(rawTxn.GetProgram().Arguments[1].Data))
		log.Printf("    Gas used: %v", provenTxn.GetGasUsed())
		if provenTxn.GetWithEvents() {
			log.Printf("    Events: (%d total)", len(provenTxn.GetEvents()))
			for idx, ev := range provenTxn.GetEvents() {
				log.Printf("      #%d:", idx)
				log.Printf("        Seq #%d", ev.SequenceNumber)
				log.Printf("        Addr: %v", hex.EncodeToString(ev.AccessPath.Address))
				log.Printf("        Raw path: %v", hex.EncodeToString(ev.AccessPath.Path))
				log.Printf("        Raw data: %v", hex.EncodeToString(ev.Data))
			}
		} else {
			log.Printf("    Events not present")
		}
	}
	return nil
}

func cmdQueryTransactionByAccountSeq(ctx *cli.Context) error {
	c, err := client.New(ServerAddr, TrustedPeersFile)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	wallet, err := LoadAccounts(WalletFile)
	if err != nil {
		log.Fatal(err)
	}

	account, err := wallet.GetAccount(ctx.Args().Get(0))
	if err != nil {
		return err
	}

	sequence, err := strconv.Atoi(ctx.Args().Get(1))
	if err != nil {
		return err
	}

	provenTxn, err := c.QueryTransactionByAccountSeq(account.Address, uint64(sequence), true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Txn #%d:", provenTxn.GetVersion())
	rawTxn, _ := provenTxn.GetSignedTxn().UnmarshalRawTransaction()
	log.Printf("    Raw txn:")
	log.Printf("        Sender account: %v", hex.EncodeToString(rawTxn.SenderAccount))
	log.Printf("        Sender seq #%v", rawTxn.SequenceNumber)
	log.Printf("        Program: %v...", hex.EncodeToString(rawTxn.GetProgram().Code[:30]))
	log.Printf("        Arg 0: %v", hex.EncodeToString(rawTxn.GetProgram().Arguments[0].Data))
	log.Printf("        Arg 1: %v", hex.EncodeToString(rawTxn.GetProgram().Arguments[1].Data))
	log.Printf("    Gas used: %v", provenTxn.GetGasUsed())
	if provenTxn.GetWithEvents() {
		log.Printf("    Events: (%d total)", len(provenTxn.GetEvents()))
		for idx, ev := range provenTxn.GetEvents() {
			log.Printf("      #%d:", idx)
			log.Printf("        Seq #%d", ev.SequenceNumber)
			log.Printf("        Addr: %v", hex.EncodeToString(ev.AccessPath.Address))
			log.Printf("        Raw path: %v", hex.EncodeToString(ev.AccessPath.Path))
			log.Printf("        Raw data: %v", hex.EncodeToString(ev.Data))
		}
	} else {
		log.Printf("    Events not present")
	}
	return nil
}