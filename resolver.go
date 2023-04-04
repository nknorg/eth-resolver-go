package ethresolver

import (
	"context"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/imdario/mergo"
	"github.com/nknorg/eth-resolver-go/contracts"
	"github.com/nknorg/nkngomobile"
	"github.com/patrickmn/go-cache"
	"github.com/wealdtech/go-ens/v3"
)

const (
	// DefaultDialTimeout Dial timeout
	DefaultDialTimeout = 5000
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

// DefaultEthereumConfig is the default Ethereum Resolver config.
var DefaultEthereumConfig = Config{
	Prefix:          "ETH:",
	RpcServer:       "https://rpc.ankr.com/eth",
	ContractAddress: "0x7BfFaF65698ecA3187CEE7651d0678127Bd7e1e2",
	CacheTimeout:    cache.NoExpiration,
	DialTimeout:     DefaultDialTimeout,
}

// DefaultHarmonyConfig is the default Harmony Resolver config.
var DefaultHarmonyConfig = Config{
	Prefix:          "ONE:",
	RpcServer:       "https://api.harmony.one",
	ContractAddress: "0x5969aC08B88819201A30CdBaA9D1c5a04Dc0C52d",
	CacheTimeout:    cache.NoExpiration,
	DialTimeout:     DefaultDialTimeout,
}

// DefaultIotexConfig is the default IoTeX Resolver config.
var DefaultIotexConfig = Config{
	Prefix:          "IOTX:",
	RpcServer:       "https://babel-api.mainnet.iotex.io",
	ContractAddress: "0xFE9Ca78B57D72226266113660e92B111a5D2E316",
	CacheTimeout:    cache.NoExpiration,
	DialTimeout:     DefaultDialTimeout,
}

// DefaultThetaConfig is the default Theta Resolver config.
var DefaultThetaConfig = Config{
	Prefix:          "TFUEL:",
	RpcServer:       "https://eth-rpc-api.thetatoken.org/rpc",
	ContractAddress: "0x748f7CeF212ce30e6Ce8c176D2b581a3E4EbD729",
	CacheTimeout:    cache.NoExpiration,
	DialTimeout:     DefaultDialTimeout,
}

func GetDefaultEthereumConfig() *Config {
	cfg := DefaultEthereumConfig
	return &cfg
}

func GetDefaultHarmonyConfig() *Config {
	cfg := DefaultHarmonyConfig
	return &cfg
}

func GetDefaultIotexConfig() *Config {
	cfg := DefaultIotexConfig
	return &cfg
}

func GetDefaultThetaConfig() *Config {
	cfg := DefaultThetaConfig
	return &cfg
}

// MergeConfig merges a given Resolver config with the default Resolver config
// recursively. Any non zero value fields will override the default config.
func MergeConfig(config *Config) (*Config, error) {
	merged := GetDefaultEthereumConfig()
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

// NewDefaultResolvers creates default Resolvers.
func NewDefaultResolvers() (*nkngomobile.ResolverArray, error) {
	ethResolver, err := NewResolver(GetDefaultEthereumConfig())
	if err != nil {
		return nil, err
	}
	oneResolver, err := NewResolver(GetDefaultHarmonyConfig())
	if err != nil {
		return nil, err
	}
	iotxResolver, err := NewResolver(GetDefaultIotexConfig())
	if err != nil {
		return nil, err
	}
	thetaResolver, err := NewResolver(GetDefaultThetaConfig())
	if err != nil {
		return nil, err
	}
	resolvers := nkngomobile.NewResolverArray(ethResolver, oneResolver, iotxResolver, thetaResolver)
	return resolvers, nil
}

// Resolve wraps ResolveContext with background context.
func (r *Resolver) Resolve(address string) (string, error) {
	return r.ResolveContext(context.Background(), address)
}

// ResolveContext resolves the address and returns the mapping address.
func (r *Resolver) ResolveContext(ctx context.Context, address string) (string, error) {
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
		Context:     ctx,
	}, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}
	r.cache.Set(address, res, cache.DefaultExpiration)
	return res, nil
}
