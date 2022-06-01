# eth-resolver-go

## Usage

* Add Resolver, ETH address to NKN address
```
account, err := NewAccount(nil)
conf := &ClientConfig {
    Resolver: []nkn.ResolverInterface {
		ethresolver.NewEthSolver(&ethresolver.EthResolverConfig {
			RpcServer:       "API Server",
			ContractAddress: "Contract Address",
		}),
	},
}
client, err := NewMultiClient(account, "identifier", 3, true, conf)
client.Send(nkn.NewStringArray("ETH:0x123..."), "Hello world.", nil)
```
