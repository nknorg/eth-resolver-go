package ethresolver

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/imdario/mergo"
	"github.com/nknorg/eth-resolver-go/contracts"
	"github.com/patrickmn/go-cache"
	"github.com/wealdtech/go-ens/v3"
	"strings"
	"time"
)

const (
	// PREFIX Protocol prefix
	PREFIX = "ETH:"

	// RPC_SERVER RPC server url
	RPC_SERVER = ""

	// CONTRACT_ADDRESS Contract address
	CONTRACT_ADDRESS = ""
)

// Config is the Resolver configuration.
type Config struct {
	Prefix          string        // Protocol prefix
	RpcServer       string        // RPC server url
	ContractAddress string        // Contract address
	CacheTimeout    time.Duration // Seconds
	DialTimeout     int           // Milliseconds
}

// Resolver implement ETH resolver.
type Resolver struct {
	config *Config
	cache  *cache.Cache
}

// DefaultConfig is the default Resolver config.
var DefaultConfig = Config{
	Prefix:          PREFIX,
	RpcServer:       RPC_SERVER,
	ContractAddress: CONTRACT_ADDRESS,
	CacheTimeout:    cache.NoExpiration,
	DialTimeout:     5000,
}

// GetDefaultConfig returns the default Resolver config with nil pointer
// fields set to default.
func GetDefaultConfig() *Config {
	return &DefaultConfig
}

// MergeConfig merges a given Resolver config with the default Resolver config
// recursively. Any non zero value fields will override the default config.
func MergeConfig(config *Config) (*Config, error) {
	merged := GetDefaultConfig()
	if config != nil {
		err := mergo.Merge(merged, config, mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	}

	return merged, nil
}

// NewResolver creates a Resolver. If config is nil, the default Resolver config will be used.
func NewResolver(config *Config) (*Resolver, error) {
	config, err := MergeConfig(config)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	var cancel context.CancelFunc
	if config.DialTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(config.DialTimeout)*time.Millisecond)
		defer cancel()
	}
	conn, err := ethclient.DialContext(ctx, config.RpcServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = contracts.NewNKNAccount(common.HexToAddress(config.ContractAddress), conn)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		config: config,
		cache:  cache.New(config.CacheTimeout*time.Second, 60*time.Second),
	}, nil
}

// Resolve resolves the address and returns the mapping address.
func (r *Resolver) Resolve(address string) (string, error) {
	if !strings.HasPrefix(strings.ToUpper(address), r.config.Prefix) {
		return "", nil
	}
	address = address[len(r.config.Prefix):]
	addr := address
	addrCache, ok := r.cache.Get(address)
	if ok {
		addr = addrCache.(string)
		return addr, nil
	}

	ctx := context.Background()
	var cancel context.CancelFunc
	if r.config.DialTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(r.config.DialTimeout)*time.Millisecond)
		defer cancel()
	}
	conn, err := ethclient.DialContext(ctx, r.config.RpcServer)
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

	contract, err := contracts.NewNKNAccount(common.HexToAddress(r.config.ContractAddress), conn)
	if err != nil {
		return "", err
	}
	res, err := contract.QueryAddr(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}
	r.cache.Set(address, res, cache.DefaultExpiration)
	return res, nil
}
