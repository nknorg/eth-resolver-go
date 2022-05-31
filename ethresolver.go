package ethresolver

import (
	"encoding/hex"
	"errors"
	"ethresolver/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/wealdtech/go-ens/v3"
	"log"
	"strings"
	"time"
)

const (
	PREFIX = "ETH:"
)

type EthResolverConfig struct {
	Prefix          string
	RpcServer       string
	ContractAddress string
	CacheTimeout    int // seconds
}

type EthResolver struct {
	config *EthResolverConfig
	cache  *cache.Cache
}

func NewEthSolver(config *EthResolverConfig) *EthResolver {
	if config.Prefix == "" {
		config.Prefix = PREFIX
	}
	return &EthResolver{
		config: config,
		cache:  cache.New(time.Duration(config.CacheTimeout)*time.Second, 60*time.Second),
	}
}

func (s *EthResolver) Resolve(address string) (string, error) {
	if !strings.HasPrefix(address, s.config.Prefix) {
		return "", errors.New("incorrect address format")
	}
	tail := address[len(s.config.Prefix):]

	conn, err := ethclient.Dial(s.config.RpcServer)
	if err != nil {
		log.Println("Dial err", err)
		return "", nil
	}
	defer conn.Close()
	contract, err := contracts.NewNKNAccount(common.HexToAddress(s.config.ContractAddress), conn)
	if err != nil {
		log.Println("NewNKNAccount err", err)
		return "", nil
	}
	addr := tail
	verify := common.IsHexAddress(tail)
	if !verify {
		ensAddr, ok := s.cache.Get(addr)
		if ok {
			addr = ensAddr.(string)
		} else {
			ensAddr, err := ens.Resolve(conn, addr)
			if err != nil {
				return "", err
			}
			addr = ensAddr.Hex()
			s.cache.Add(tail, addr, 0)
		}
	}

	nknAddrCache, ok := s.cache.Get(addr)
	var nknAddr string
	if ok {
		nknAddr = nknAddrCache.(string)
	} else {
		res, err := contract.GetNKNAddr(&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		}, common.HexToAddress(addr))
		if err != nil {
			log.Println("NewNKNAccount err", err)
			return "", err
		}
		nknAddr = hex.EncodeToString(res.PublicKey[:])
		if res.Identifier != "" {
			nknAddr = res.Identifier + "." + hex.EncodeToString(res.PublicKey[:])
		}
		s.cache.Add(addr, nknAddr, 0)
	}

	return nknAddr, nil
}
