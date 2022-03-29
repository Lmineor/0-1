# 虚拟货币钱包
钱包中是自己账户对应的key，而我的钱包里有多少钱是存储在区块链上的。
这样来说，虚拟货币钱包实际上是管理和存储key的工具。这把key就是我的私钥，而账户是从我的公钥衍生出来。

# BIP
> BIP全名是 **Bitcoin Improvement Proposals**，是提出Bitcoin的新功能或改进措施的文件。可以由任何人提出，经过审核后公布在[bitcoin/bips: Bitcoin Improvement Proposals (github.com)](https://github.com/bitcoin/bips)上。BIP 和 Bitcoin 的关系，就像是 RFC 之于 Internet。

而其中的 **BIP32, BIP39, BIP44 共同定义了目前被广泛使用的 HD Wallet**，包含其设计动机和理念、实作方式、实例等。
BIP32：定义**Hierarchical Deterministic wallet (简称 "HD Wallet")**，是一个系统可以从单一个seed产生一树状存储多组keypairs（私钥和公钥）。好处是可以方便的备份、转移到其他相容装置（因为都只需要seed），以及分层的权限控制等。