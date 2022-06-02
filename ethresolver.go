package ethresolver

import (
	"encoding/hex"
	"ethresolver/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/wealdtech/go-ens/v3"
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

func NewEthResolver(config *EthResolverConfig) (*EthResolver, error) {
	if config.Prefix == "" {
		config.Prefix = PREFIX
	}
	conn, err := ethclient.Dial(config.RpcServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = contracts.NewNKNAccount(common.HexToAddress(config.ContractAddress), conn)
	if err != nil {
		return nil, err
	}

	return &EthResolver{
		config: config,
		cache:  cache.New(time.Duration(config.CacheTimeout)*time.Second, 60*time.Second),
	}, nil
}

func (s *EthResolver) Resolve(address string) (string, error) {
	if !strings.HasPrefix(address, s.config.Prefix) {
		return "", nil
	}
	address = address[len(s.config.Prefix):]
	addr := address
	addrCache, ok := s.cache.Get(address)
	if ok {
		addr = addrCache.(string)
		return addr, nil
	}

	conn, err := ethclient.Dial(s.config.RpcServer)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	verify := common.IsHexAddress(address)
	if !verify {
		ensAddr, err := ens.Resolve(conn, addr)
		if err != nil {
			return "", err
		}
		addr = ensAddr.Hex()
	}

	var nknAddr string
	contract, err := contracts.NewNKNAccount(common.HexToAddress(s.config.ContractAddress), conn)
	if err != nil {
		return "", err
	}
	res, err := contract.GetNKNAddr(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}
	nknAddr = hex.EncodeToString(res.PublicKey[:])
	if res.Identifier != "" {
		nknAddr = res.Identifier + "." + nknAddr
	}
	s.cache.Add(addr, nknAddr, cache.DefaultExpiration)

	return nknAddr, nil
}
