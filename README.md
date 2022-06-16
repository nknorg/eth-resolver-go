# eth-resolver-go

## Usage

* Add Resolver, ETH address to NKN address
```
account, err := NewAccount(nil)
ethResolver, err := ethresolver.NewResolver(&ethresolver.Config{
    RpcServer:       "API Server",
    ContractAddress: "Contract Address",
})
if err != nil {
    return err
}

conf := &nkn.ClientConfig{
    Resolvers: nkngomobile.NewResolverArray(ethResolver),
}
client, err := NewMultiClient(account, "identifier", 3, true, conf)
client.Send(nkn.NewStringArray("ETH:0x123..."), "Hello world.", nil)
```
