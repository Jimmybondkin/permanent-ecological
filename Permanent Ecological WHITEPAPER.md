# Permanent Ecological WHITEPAPER

**PEE, A Highly Customizable Blockchain Infrastructure**

Version: 1.3 (091319)

_NOTE: If you can read this on GitHub, then we're still actively developing this
document.  Please check regularly for updates!_

## Directory ###########################################################

  * 1.0 Preface Abstract
  * 2.0 Permanent Ecological Overview
  * 3.0 Overview Of The Core Technology Of Permanent Ecological
    * 3.1 Tendermint Consensus Algorithm
      * 3.1.1 Tendermintbft
      * 3.1.2 The Verifier
      * 3.1.3 Consensus Agreement
      * 3.1.4 Lightweight
      * 3.1.5 ABCI
    * 3.2 Constitutional System
    * 3.3 Regions And Dispersers
    * 3.4 Block Chain Communication -Ibc
    * 3.5 Distributed Exchange
    * 3.6 Ethereum Zoom
    * 3.7 Multi-Application Integration
    * 3.8 Network Partition Mitigation
    * 3.9 Bridge To Other Cryptocurrencies
      * 3.9.1 The Token Will Be Sent To The Permanent Ecological Distributor
      * 3.9.2 Extract The Token From The Permanent Ecological Distributor
      * 3.9.3 Overall Responsibility System In Bridge Area
  * 4.0 Introduction To The Permanent Ecological Application
    * 4.1 Cloud Computing Power Light Client
    * 4.2 Distributed Digital Currency Exchange
    * 4.3 Multi-Chain Blockchain Wallet
    * 4.4 Autonomous Business Alliance
    * 4.5 Decentralized Financial Banking
  * 5.0 Permanent Ecological Release Overview
    * 5.1 Basic Parameters Of Pee Token
    * 5.2 Basic Parameters Of Pet Token
    * 5.3 Initial Allocation Information Of Pee Tokens
  * 6.0 Permanent Ecological Disclaimer
  * 7.0 Permanent Ecological Relevant Appendices And Literature
 

## 1.0 Introduction In This Paper ###########################################################

> Permanent Ecological Changes Will Shape The Future Of Business-Pee Global Financial Autonomy

A thought called **"Decentralization"** is innate from the beginning of human society. Human desire to remove the center, to achieve direct communication between people, direct transactions, direct transmission. Humans believe that one day we may no longer need centralised institutions. In the long course of human development, this kind of thinking has repeatedly collided with different organizations, environments and carriers, thus resulting in different social changes. In today's space-time dimension, we can still see decentralized religious churches, as well as decentralized political systems and social organizations.

The emergence and large-scale use of information technology has brought technology carriers to the idea of decentralization. Tracing the thread of this combination of ideas and technologies has produced a number of popular, even disruptive, innovations. P2P download, CND (p-cnd), distributed computing (cloud computing), social media (we media), P2P lending, crowdfunding, sharing, block chain, self-organizing DAO, etc. Along this line of innovation, we will see the footprints of many great innovative Internet. Throughout the development history of decentralized thinking, whether the blockchain technology will eventually win, or whether it will compromise with other technologies and integrate with each other. However, it is precisely because each technology route has felt great pressure from the market that it has also generated great development momentum. We believe that whatever the outcome, the world will be transformed.

Permanent Ecological has been looking for a new design of the future business form of blockchain, which is not only limited to the one-way business expansion, but also lies in the innovation of decentralized finance, decentralized expansion of business centers and the blockchain network architecture. The Permanent Ecological connection connects many networks of independent blockchains, and the interoperability between blockchains is realized with the help of the Permanent Ecological BFT consensus agreement and the Tendermint consensus algorithm, which fully realizes the value potential of the blockchain network.

In this paper, the Permanent Ecological public chain will be constantly updated, which includes but is not limited to key governance and key technologies,_**The mission of the Permanent Ecological public chain is to break the traditional business form and reshape a new business form for the future from the four dimensions of block chain technology, business model, economic model and governance structure, as well as create a new economic form, new organizational form and new social form.**_

## 2.0 Permanent Ecological Concise Overview ###########################################################

> The Term Permanent Ecological Refers To Creating a Sustainable Commercial Ecological Field Through Consensus Building And Win-Win Cooperation

Contractual relationships are an essential part of human society, and the importance of blockchain technology lies in providing a very efficient and low-cost way to achieve reliable contractual relationships. The contractual relationship lies in the relative consensus and trust reached by multiple parties when they participate in complex business interactions, and the relationship has always been maintained in business interactions, while the contractual relationship of block chain technology is very low cost. That is to say, blockchain technology provides an important element for distributed commerce (improving network benefits with very low transaction costs). More and more people realize that blockchain is a new value Internet influence, and will gradually transform the current business model into a more efficient distributed network. In particular, the general authentication mechanism built into most modern blockchains emphasizes the rights of each network participant and will revolutionize existing business models.

However, blockchain technology is still in its early stages. As with other new technologies, there are also shortcomings, including limited performance and governance mechanisms that have not yet been developed, and the lack of a sound business model and economic model for the majority of public chains, which make it difficult for the majority of public chains to support real distributed business collaboration. Organizations such as Hyperledger Fabric and R3 Corda, as well as the Ethereum Enterprise Alliance, are trying to address these performance and governance issues through consortium chains to make blockchain technology more enterprise-ready. However, today's chain of alliances is dominated by large corporations, whose closed up-and-down governance model is very inefficient. The alliance chain may lack vitality because of the lack of the general economic model of public chain and its openness and incentive. **We hope that through Permanent Ecological link technology, thousands of Small Medium Businesses (SMBs), even individual freelancers, can provide their services and enjoy the rewards in an open network.**

Ethereum offers Turing's complete virtual machine running intelligence contract, bringing many hopes for developing distributed applications. However, smart contracts can only deal with deterministic logic (so each node can reach the same state after processing the same transaction and block), whereas a large amount of existing business logic is uncertain and may change at different times and under different environmental parameters. On the other hand, some real-world business logic should run down the chain and should not be executed as a type of smart contract such as a repeatable operation. Utilizing services and resources under distributed ledger integration and collaborative chain is the key to further promote the application of block chain technology in more real scenarios.

It is not feasible to use a single public chain to handle all use cases. Each day a different blockchain comes online, each focusing on one aspect of the problem, such as distributed storage, asset ownership or market forecasting. According to coinmarketcap.com, there are currently more than 1,000 cryptocurrencies active on different trading platforms. Building business applications involves dealing with storage and sources of different data sources, and another motivation for our work is how to reuse some of the existing work, such as storage (IPFS, SIA, storj.io, etc.), data sending (Augur, Gnosis, Oraclize, etc.), and these dedicated blockchains provided by the Internet of things (IOTA, etc.).

In addition, there are many real-time business transactions that do require closer alliance/license/private chains to handle performance issues, security issues, and business governance requirements. Therefore, our vision for a distributed business infrastructure is the ability to interoperate across multiple heterogeneous chains, including public/alliance/license/private chains.

Cross-chain technology is a natural solution to this need. However, so far, existing cross-chain technologies are mainly aimed at providing interoperability within existing blockchains and focusing on value transfer through authentication. Permanent Ecological
It was designed based on the Tendermint consensus engine, achieving interoperability and scalability. We also built a multi-hub and multi-partition model, each partition being an independent block chain. In addition, we have separate governance structures that provide a very appropriate architecture for real businesses to model real-world complexity in the form of a Seperation of Concern (SOC).

Below, we focus on how the Permanent Ecological network can build the technical infrastructure for distributed business applications. If you can read this on GitHub, we're still working on this document. Please check for updates regularly! This Chinese translation was translated by Permanent Ecological community volunteers according to the English version of the white paper.

## 3.0 Permanent Ecological Overview Of Core Technologies ###########################################################

Permanent Ecological is an independent parallel blockchain network, each of which is supported by the classical BFT consensus algorithm. The Permanent Ecological connection to many other blockchains (or regions) is made possible by novel inter-blockchain communication protocols. It also keeps track of various token types, keeping track of the total number of tokens in each connection area. Tokens can be safely and quickly transferred from one area to another without the need for liquid exchange between areas because all inter-area coin transfers are made by the Permanent Ecological.

This architecture addresses many of the issues facing the blockchain space today, such as application interoperability, scalability, and seamless upgradability. For example, an area from Bitcoind, go-ethereum, CryptoNote, ZCash, or any blockchain system can be inserted with the Permanent Ecological. These areas allow the Permanent Ecological expansion to meet the needs of global trade. Regions are also well suited for distributed exchange and will be supported.

The Permanent Ecological is more than just a single distributed ledger. We are designing protocols for an open distributed ledger network based on cryptographic principles, rational economics, consensus theory, transparency, and accountability that can serve as a new foundation for future business systems.

### 3.1 Tendermint Consensus Algorithm

Tendermint full block chain is an open source implementation, and can be used for chain or chain alliance, its official position is the engine block chain consensus for developers, with other block chain platform such as Ethernet lane or EOS than Tendermint the biggest characteristic is its differentiated positioning, although contains a complete implementation of block chain, but it is in the form of the SDK will provide, the core functionality for developers to easily customize their proprietary block chain.

The Tendermint is roughly similar to two types of software. The first category includes the use of BFT consensus distributed key-value stores, such as the Zookeeper, etcd and consul. The second category, known as "blockchain technology", consists of cryptocurrencies such as bitcoin and ethereum and alternative distributed ledger designs such as Hyperledger's Burrow.

  * **3.1.1 TendermintBFT**

The Tendermint open source project was launched in 2014, aiming to solve the problems of bitcoin workload proving the speed, scalability and environment of consensus algorithms. By using and improving the proven BFT algorithm developed by MIT in 1988 [20], the Tendermint team was the first to conceptively demonstrate the equity proof cryptocurrency, which addresses the equity abuse suffered by the first generation of certificates, using cryptocurrencies such as NXT and BitShares1.0.

Today, nearly all bitcoin mobile wallets use trusted servers to authenticate their transactions. This is because workload proof needs to wait for many confirmations before the deal is seen as an irreversible commitment. Double-flower attacks have been demonstrated on services like CoinBase.

Unlike other blockchain consensus systems, the Tendermint provides instant and provably secure mobile client payment verification. Since the Tendermint design never bifurcates, mobile wallets can receive instant confirmation of transactions, making it possible to make distrustful and practical payments on smartphones. This has also had a major impact on iot adoption.

The validators in the Tendermint play a similar role to bitcoin miners, but use encrypted signatures to conduct the voting. The validator is the security-specific machine responsible for submitting the block. Non-verifiers can delegate their mortgage tokens to any verifier to earn a portion of the block fee and quantum reward, but they risk being penalized if the delegate verifier is hacked or breaches the agreement.

The Tendermint guarantees that it will never violate its security, that is, the verifier will never submit a conflict block at the same level. To achieve this, it introduces Locking rules that modularize the path in the flowchart. Once a verifier pre-commits a block, it is "locked" on that block.And then,
1. It must prevote for the locked block
2. Only in a later round, with a boka of that block, can it be unlocked and precommitted for a new block.

The effective security guarantees of the Tendermint BFT consensus, as well as the collateral deposits of the stakeholders (verifiers and principals), provide provable and quantifiable security for the nodes and light customers.

  * **3.1.2 The Verifier**
 
In the Tendermint, each node has the same weight, and the nodes have non-negative voting rights, while the nodes with positive voting rights are called validators. The verifier signs or votes on the consensus agreement by broadcasting the password to reach agreement on the next block.

The participants in the agreement are called "validators." They take turns proposing blocks and voting on them. Blocks are committed to the chain, with each block occupying a "height." The commit block may fail, and if it fails, the protocol starts the next commit, and a new verifier continues to commit the block at that height. To successfully commit a block, there are two phases of voting: "pre-vote" and "pre-commit".In the same round of submission, the same block can only be submitted to the chain if more than 2/3 validators precommit the same block.

For some reason, the verifier may fail to commit a block: the current proposer may be offline, or the network may be very slow. The Tendermint allowed them to verify that a verifier should be skipped. The verifier waits a short time to receive a complete proposal block from the proposer before proceeding to the next round of voting. This reliance on timeouts makes the Tendermint a weak synchronization protocol rather than an asynchronous one. However, the rest of the protocol is asynchronous, and the verifier will only take the next step if more than 2/3 verifier sets are received. One reason the Tendermint is able to simplify is that it USES the same mechanism to submit a block and skip straight to the next round.

Note: scores such as ⅔ and like represent the total voting power score, not the total number of verifiers, unless all verifiers have the same weight. >⅔ means "greater than ⅔", and ≥⅓ means "at least ⅓".

  * **3.1.3 Consensus Agreement**
 
In terms of technology, the Tendermint prides itself on its consensus algorithm -- the world's first Byzantine fault-tolerant algorithm that can be applied to a public chain. Tendermint is famous for its simplicity, high performance and bifurcated sense of responsibility. The protocol requires a fixed set of known validators, each identified by its public key. The validator tries to agree on one block at a time, one of which is a list of transactions. A full vote on whether to reach a consensus. Each turn has a round leader or proposer proposing a block. The verifier will then vote in stages on whether to accept the proposed block or proceed to the next round. The proposers of the round are determined from an ordered list of verifiers in proportion to their voting power.

The security of the Tendermint stems from its use of optimal Byzantine fault tolerance through super majority voting (> a) and locking mechanisms. Together they ensure that:

Voting power ≥⅓ must be Byzantine in order to violate security by submitting more than two values.

If any group of validators successfully violate security, or even attempt to do so, they can be identified by protocol. This includes voting on disputed parts and broadcasting unreasonable votes.

In spite of the strong guarantee, the Tendermint still has excellent performance. In a benchmark of 64 nodes spread across seven data centers on five continents, the Tendermint consensus was able to handle thousands of transactions per second in the case of commercial transactions, with a submission delay of one to two seconds. It's worth noting that, even in harsh confrontational conditions, the verifier can crash or spread maliciously crafted ballots, maintaining the performance of more than a thousand transactions per second.

![Figure of Tendermint throughput performance](https://raw.githubusercontent.com/gnuclear/atom-whitepaper/master/images/tendermint_throughput_blocksize.png)

  * **3.1.4 Lightweight**
 
The main advantage of the Tendermint consensus algorithm is that it simplifies lightweight client security, making it an ideal choice for mobile and Internet of things applications. Although the bitcoin lightweight client must synchronize the blockhead chain and find the chain with the most work, the Tendermint lightweight client only needs to keep up with the changes to the verifier set and then verify the >⅔PreCommits in the latest block to determine the latest status. The compact lightweight client proof also enables communication between blockchains.

  * **3.1.5 ABCI**
 
The Tendermint consensus algorithm is implemented in a program called Tendermint Core. Tendermint Core is an app-independent "consensus engine" that turns any deterministic black box application into a distributed, replicated blockchain. The Tendermint Core connects to the blockchain application via the application block link port (ABCI). Thus, ABCI allows block chain applications to be programmed in any language, not just in a programming language written by the consensus engine. In addition, ABCI makes it easy to swap out the consensus layer of any existing blockchain stack.

ABCI allows Byzantine fault-tolerant replication of applications that can be written in any programming language. Tendermint Core (" consensus engine ") communicates with applications through an abci-compliant socket protocol. Take a familiar example, bitcoin. Bitcoin is a cryptocurrency blockchain in which each node maintains a fully audited UTXO database. If someone wants to create a bitcoin-like system on top of ABCI,

**Tendermint Core will be responsible for:**

    Sharing blocks and transactions among nodes;
    
    Establish a standard/immutable order for transactions (blockchain);

**The application will be responsible for:**

    Maintain the UTXO database;
    
    Verify the encrypted signature of the transaction;
    
    Block transactions where costs do not yet exist;
    
    Allows clients to query the UTXO database;
    
ABCI contains three main message types, which are sent by core to the application, and the application will respond to the message accordingly:

The DeliverTx message is the main part of the application. Every transaction in the chain is transmitted through this message. The application needs to validate each transaction that receives the DeliverTx message based on the current state, application protocol, and encrypted certificate of the transaction. A validated transaction then needs to update the application status-either by binding a value to a key-value store, or by updating the UTXO database.

The CheckTx message is similar to DeliverTx, but it is only used to validate the transaction. The Tendermint Core's memory pool first checks the validity of a transaction by CheckTx and relays only valid transactions to other nodes. For example, an application might check for serial Numbers that are growing in the transaction, and if the serial Numbers are out of date, CheckTx returns an error. Or, they might use a volume-based system that requires recapacity for each transaction.

The Commit message is used to calculate a cryptographic commitment for the current application status, which is placed in the next block header. Here are some handy properties. Inconsistencies in status updates are now considered branches of the blockchain, which catch all programming errors. This also simplifies the development of a secure light-node client, since the merk-hash proof can be verified by checking on the block hash, which is signed by a quorum.

An application may have multiple ABCI socket connections. The Tendermint Core created three ABCI connections for the application: one for transaction validation when the memory pool is broadcasting, one for running the consensus engine when submitting blocks, and one for checking the status of the application.

### 3.2 Constitutional System

The distributed public ledger shall have constitutional and governance systems. Bitcoin relies on the bitcoin foundation and mining to coordinate the upgrade, but it is a slow process. Ethereum split into ETH and ETC after bifurcating to resolve TheDAO's hacking, largely because there was no prior social contract or decision-making mechanism.

Permanent Ecological verifiers and representatives can vote on proposals that automatically change the system's preset parameters (for example, limit gas restrictions), coordinate upgrades, and vote on human-readable constitutional amendments governing the policy. The center of the universe. The constitution allows stakeholders to work together on issues such as theft and errors (such as the TheDAO incident) so that problems can be solved more quickly and cleanly.

Each region could also have its own constitution and governance mechanisms. For example, the Permanent Ecological can have a construct that forces invariance (no rollback, except for the error of the Permanent Ecological node implementation), and each region can set its own rollback policy.

By achieving interoperability between different policy areas, the Permanent Ecological network provides users with maximum freedom and the potential for unlicensed experimentation.

### 3.3 Zones And Dispersers

The Permanent Ecological is a network of many blockchains supported by the Tendermint. Existing proposals aim to create a "single blockchain" with an overall global order of transactions, while the Permanent Ecological allows many blockchains to run concurrently while maintaining interoperability. Here, we describe a new model of decentralization and scalability.

Permanent Ecological Disperser management by many independent block chain, known as the "Zones" (sometimes called "fragments", reference database extensions technology called "debris"). The continuous flow of the latest block lift from the regions published on the disperser enables the disperser to keep up with the state of each region. Again, each region is synchronized with the state of the disperser (but not indirectly through the disperser). The packets are then passed from one region to another by publishing Merkle proofs as evidence of the sending and receiving of information. This mechanism is called block chain to block communication, or IBC.

### 3.4 Block Chain Communication -IBC

Now let's look at how the disperser communicates with the region. Chain, for example, if there are three blocks "Zone1", "Zone2" and "Disperser", we hope the Zone1 packets sent to "Zone2", and through the "Disperser". To move packets from one block chain to another, a certificate is published on the receiving chain. Evidence suggests that the sending chain publishes packets of so-called destinations. In order for the receiving chain to be able to check this proof, it must be able to keep up with the sender's bulk. This mechanism is similar to that used by side chains, which require two interacting chains to know each other through two-way presence to prove the datagram (transaction) flow.

Naturally, there are two types of transactions that define the IBC protocol: an IBCBlockCommitTx transaction that allows the block chain to prove its most recent hash value to any observer; Another IBCPacketTx transaction allows the blockchain to prove to any observer that a given packet is indeed published by the sender's application by Merkle's proof of the nearest block hash.

By splitting the IBC mechanism into two separate transactions, we allow the capital market mechanism of the receiving chain to determine the submitted (or acknowledged) packets, while allowing the sending chain complete freedom to allow as many outbound packets as possible.

### 3.5 Distributed Switching

Just as bitcoin improves security by becoming a distributed, massively copied ledger, we can reduce the vulnerability of exchanges to external and internal hackers by running them on the blockchain. We call this a distributed exchange.

oday, the cryptocurrency community's so-called decentralized exchanges are based on something called "atomic cross-chain" (AXC) trading, defined as quantum cross-chain [QXC]. With QXC transactions, two users on two different chains can make two transfer transactions committed together in two ledgers, or not committed at all (that is, quantum). For example, even if bitcoin and ethereum are not connected to each other, two users can use QXC transactions to exchange bitcoins for ethereum (or any two tokens in two different ledgers). The advantage of running exchanges on QXC transactions is that users do not need to trust each other or trade match services. The downside is that both parties must be online to trade.

Another type of decentralized exchange is a large-scale replicated distributed exchange that runs on its own blockchain. Users who make such trades can submit limit orders and shut down their computers, and the trades can be executed without the user being online. The blockchain represents the trader matching and completing the transaction.

The Tendermint consensus offers other benefits of faster transaction submission. By giving priority to fast certainty without sacrificing consistency, the region in the Permanent Ecological can quickly complete the transaction - the transaction of the exchange order and the transfer of the IBC token to and from other regions.

Given the current status of cryptocurrency exchange, an important application of Permanent Ecological is distributed exchange (aka the Permanent Ecological DEX). Transaction throughput and commit latency are comparable to centralized exchange. Traders can submit limit orders that can be executed without both parties being online. With the help of the Tendermint, the Permanent Ecological hub and IBC, traders can quickly move money in and out of the exchange to other areas.

### 3.6 Ethereum Zoom

Resolving the expansion issue is an open question for ethereum. Currently, ethereum nodes process each single transaction and store all states. Since the Tendermint can demonstrate faster block submission than ethereum's workload, the EVM area supported by the Tendermint consensus and running on the bridge ether can provide higher performance for ethereum's blockchain. In addition, although the scatterer and IBC packet mechanisms themselves do not allow arbitrary execution of contract logic, they can be used to coordinate token movement between ethereum contracts running on different regions, providing the basis for toker-centric ethereum scaling through fragmentation.

### 3.7 Multi-Application Integration

The Permanent Ecological region runs any application logic that is defined at the beginning of the region's life cycle and may be updated over time through governance. This flexibility allows the Permanent Ecological region to act as a bridge to other cryptocurrencies, such as ethereum or bitcoin, and also allows those blockchain derivatives that use the same code base but have a different set of validators and initial distribution. This allows a number of existing cryptocurrency frameworks (e.g., etim, Zerocash, bitcoin, CryptoNote, etc.) to be used in conjunction with Tendermint Core, which is a high performance consensus engine that can be used on a generic network, providing a huge opportunity for cross-platform interoperability. In addition, as a multi-asset block chain, a single transaction may contain multiple inputs and outputs, each of which can be any token type, so that the Permanent Ecological can be directly used as a platform for decentralized transactions, although it is assumed that the order is matched through other platforms. Alternatively, the area can be used as a distributed fault-tolerant exchange (with an order book), which can be a serious improvement on the existing centralized cryptocurrency exchange, which is often hacked over time.

Area can also be used as a support block chain version of enterprise and government system, which is traditionally organized by an organization or group of running some of the specific services as ABCI applications that run on an area, which makes it can inherit Permanent Ecological public network security and interoperability, without sacrificing control of the basic services. Thus, Permanent Ecological may provide the best of both worlds for organizations that want to take advantage of blockchain technology but give up control entirely to distributed third parties.

### 3.8 Network Partition Mitigation

It has been claimed that a major problem with consensus algorithms that favor consistency, such as the Tendermint, is that any network partition that results in no voting rights greater than (for example, ≥offline) will completely terminate the consensus. The Permanent Ecological system can alleviate this problem by using global hubs with regional autonomous regions, where voting rights are allocated for each region based on a common geographical area. For example, a common example might be for cities or regions to operate their own areas while sharing a common distributor (such as a Permanent Ecological distributor), thus sustaining municipal activities when the distributor is stopped due to temporary network partitioning. Note that this allows real geological, political, and network topological characteristics to be considered when designing robust joint fault-tolerant systems.

### 3.9 Bridge To Other Cryptocurrencies

A privileged area can act as a source for a bridge token for another cryptocurrency. Bridges are similar to the relationship between the Permanent Ecological hub and the region. Both must follow the latest block of the other to verify that the token has been moved from one to the other. The bridge area on the Permanent Ecological network keeps pace with hubs and other cryptocurrencies. The logic of indirectly allowing the disperser through the bridging area is kept simple and independent of other blockchain consensus strategies (such as bitcoin workload proof mining).

  * **3.9.1 The Token Is Sent To The Permanent Ecological Distributor**

Each bridge validator will run a blockchain with the Tendermint, a special ABCI bridge application, and a full node of the "origin" blockchain.

When a new block is mined at the origin, the bridge area validator will agree on the committed block by signing and sharing its local view of the origin's blockchain tricks. When the bridge area receives the payment at the origin (for a PoW chain like ethereum or bitcoin, it is confirmed that sufficient confirmation has been made), a corresponding account with that balance is created on the bridge area.

In the case of ethereum, the bridge area can share the same set of validators as the distributor. On the ethereum side (origin), the bridge contract will allow ethereum holders to transfer the ether to the bridge area by sending the ethereum currency to the bridge contract on ethereum. Once the Ethernet is received by the bridge contract, the Ethernet cannot be extracted unless the bridge contract receives the appropriate IBC packets from the bridge region. The bridge contract tracks the set of validators for the bridge area, which may be the same as the set of validators for the distributor.

In the case of bitcoin, the concept is similar, except that each UTXO will be controlled by threshold multi-signature P2SH releases, rather than a single bridge contract. Due to the limitations of the P2SH system, the signer cannot be the same as the disperser validator set.

  * **3.9.2 Extract The Token From The Permanent Ecological Distributor**

Ethereum on the bridge area (" bridging ethereum ") can be transmitted back and forth between the distributors and then destroyed by transactions that send it to specific withdrawal addresses on ethereum. The IBC packet can be used to prove that the transaction occurred in the bridge area and then published into the ethereum bridge contract to allow the ethereum currency to be extracted.

In the case of bitcoin, the restricted scripting system makes it difficult to mirror the IBC coin transfer mechanism. Each UTXO has its own independent distribution, so each UTXO must be migrated to the new UTXO when the set of bitcoin managed signers changes. One solution is to compress and decompress the UTXO set as needed to reduce the total number of UTXO.

  * **3.9.3 Overall Accountability In The Bridge Area**
 
The risk of such a shrunken contract lies with rogue verifiers. Byzantine voting power ≥⅓ May cause the fork to withdraw the etheric COINS from the bridge contract over ethereum while keeping the etheric over the bridge area. Worse, the >⅔ Byzantine voting rights could deviate from the original bridge logic of the bridge belt, thereby stealing the ether outright from those who sent it to the bridge contract.

These problems can be solved by designing fully responsible Bridges. For example, all IBC packets from the distributor and the originating place may require the bridge region validation so that the distributor or originator's bridge can effectively challenge and validate all state transition contracts in the bridge region. The distributors and sources should allow the bridge area verifier to post collateral and should delay the transfer of tokens from the bridge contract (and the collateral is unsecured for a sufficiently long period) to allow any challenge from the independent auditor. We will open the normative design and the realization of this system as the Suggestions for Permanent Ecological improvement in the future, so as to pass the governance system of the Permanent Ecological distributor.

## 4.0 Permanent Ecological Overview Of Applied Ecology ###########################################################

> We Have Been Trying To Build The Permanent Ecological Chain From The Four Dimensions Of Block Chain Technology, Business Model, Economic Model And Governance Structure. Business Model And Economic Model Are Also The Core And Key Of The Permanent Ecological Chain. This Paper Is a White Paper On Technology, Which Mainly Elaborates On The Technical Infrastructure Of The Permanent Ecological Public Link.

### 4.1 CloudHash&Light Client 

We provide users with the cloud to calculate force light Client [CloudHash&Light Client], does not need to be synchronized Permanent Ecological data blocks all node, just change of synchronous validator set and verify the latest blocks of > two-thirds PreCommits can determine the latest status. This simplification makes it an ideal choice for mobile and Internet of things use cases. Users can directly use the light client of the Permanent Ecological cloud to set up nodes, and the light client supports starting the mining mode at present (there is a limit on the number and opening time of the service of the light client of mining mode at present, and the end is dominated by the latest Permanent Ecological dynamic). Each light client is only allowed to set up one mining node for mining, and must log on the Permanent Ecological cloud every day to ensure that the node network data is synchronized to the latest state and the node access is normal. The cloud power light client also supports DAPP, an intelligent contract distributed by the Permanent Ecological network, to ensure the comprehensive nature of the Permanent Ecological application.

### 4.2 Distributed Digital Currency Exchanges

The Permanent Ecological public chain is building the emerging open financial ecosystem to build the liquidity infrastructure. The network is a binding block chain of equity proof based on the Tendermint consensus.

The base state machine is responsible for executing a simple set of access control rules and a basic order booking process. The network utilizes the Shared security model, in which all the mortgage and token based logic is implemented on the Permanent Ecological basis. This model provides certainty for specific changes in the Permanent Ecological system by utilizing the Tendermint, thus providing one-way communication and inspection between the two networks. A dynamic set of independent validators and voting token holders support core protocol and relay network. Through local token systems and simple economic mechanisms, network validators, users, and voters are coordinated and motivated to achieve common goals of network security and value.

The Permanent Ecological chain has also developed an unmanaged request quotation system, whose unique interaction model makes it one of the most competitive DEX models. Because the current DEX architecture makes access priority too high due to high latency and interaction scenarios, this is a price disadvantage. However, the PE quotation system introduced the inquiry model, which forced the signatory of the market maker agreement to provide the quotation in the form of order and required the maker to implement the commitment. The above protocols can provide users with a safe, fast and convenient distributed financial system.

### 4.3 Multi-Chain BlockChain Wallet

The Permanent Ecological wallet is HD wallet managed with a mnemonic. HD wallet is now the "premium" deterministic wallet. Based on BIP44 protocol implementation (currently commonly used BIP39, BIP44 protocol). BIP44 protocol defined, HD purse with tree structure derived keys (private and public keys), which means that the seed to generate the private key (or master key), then through the private key, you can derive a series of private key, and each child private key can be derived from a series of key sun again, and so on, can be derived. The process of deriving the child private key from the primary private key cannot be reversed. The child private key cannot derive the parent private key upward, nor can the private key of the same level be derived horizontally. Therefore, it can manage the private key in a very secure layer and effectively protect the financial privacy. Users can easily manage multiple cryptographic assets with a single wallet. Only need to back up the seed, do not need to back up the private key, to provide users with great convenience, reduce the burden of managing multi-chain assets. Light wallet currently only supports access and transaction verification.

### 4.4 Autonomous Business Alliance

Block chain autonomous business alliance is the new era of Internet commodity trade in the future, it will be a better choice. In today's business form, the global small and medium-sized enterprises and the real economy has been facing a severe test, how to block chain technology is helping them get through this cold winter became the topic under discussion, extremely capital chain chain scission, unsalable inventory products, product research and development problems such as low energy efficiency, product quality traceability are still with their development. With the increasing maturity of block chain technology, the features of distributed storage, point-to-point transmission and consensus mechanism provide a feasible way to solve these pain points. The Permanent Ecological pattern provides a new model of autonomous business alliance for the future business pattern. The supply chain management and information asymmetry brought by the centralization of e-commerce have been restricting the further development of commodity trade. The Permanent Ecological pattern gradually changes the existing business pattern by breaking this bottleneck.

### 4.5 Decentralized Financial Banking

Banks aim to solve two major problems that cannot be solved by centralized services: unequal financial services; Financial review. Unequal financial services. This refers to individuals' access to financial services such as loans, mortgages and insurance. People who have little or no access to financial services are often called "unbanked." Decentralised financial applications aim to improve the problem and ensure that they are accessible to people; All it takes is a smartphone and an Internet connection. Financial censorship, where governments, financial institutions, or third parties close the accounts of individuals or companies and restrict their trading for specific interests. For example, if a company dares to disagree publicly about government policies, the government can silence the company by restricting its access to basic financial services. Take the bank account service, the company needs to pay employees and other expenses through the bank account, without which the company will go bankrupt. However, the Permanent Ecological decentralized financial Banks break these two problems through effective blockchain application. For example, small and medium-sized enterprises and the real economy chain transformation and chain transformation Token bond transformation, or C.D.P.S mortgage debt position smart contract to achieve the value of the long-term Ecological decentralization of financial Banks.

## 5.0 Permanent Ecological Summary Of Issue ###########################################################

The Permanent Ecological public chain issues two tokens in total, which is a dual token system. The tokens issued by the main network are PEE; The side chain issue token is PET (stable token).

### 5.1 Basic Parameters Of PEE Token

| ***Total Circulation***        | ***3,33,0000,000 PEE***                     |
| :----------------------------- | :------------------------------------------ |
| Method Of Proof Pf Interest    | ***POS is based on the Tendermint engine*** |
| Block Out Interval             | ***6 Seconds***                             |
| Out Of The Block Reward Amount | ***40 PEE***                                |
| Production Cycle               | ***2635200 Blocks***                        |
| Cuts                           | ***3%***                                    |

### 5.2 Basic Parameters Of PET Token

| ***Total Circulation***        | ***56,0000,000 PET*** |
| :----------------------------- | :-------------------- |
| Block Out Interval             | ***21 Seconds***      |
| Out Of The Block Reward Amount | ***2 PET***           |

### 5.3 Initial Allocation Information For PEE Tokens

PEE total circulation for a total of 3,33,0000,000 PEE, triggered by quantum content marketing, light cloud to calculate force clients to dig version 1.0, the cloud to calculate force light client version 2.0, all nodes mining and mining, technical maintenance, Ecological contribution, community maintenance, all nodes of reward and incentive mechanism, because the Permanent Ecological USES is POS mechanism, in the creation stage triggered PEE will be allocated in the center of the several sections to node, triggered by intelligent contract, and all data will be made public audit in the blocks in the browser. And it is regulated and supervised by Permanent Ecological AI.

| ***De-Centralizing Node Properties***   | ***Allocation Proportion*** | ***Distribution Number***  |
| -------------------------- | ------------------ | ------------------- |
| Quantum Trigger Content Promotion  | ***10%***      | ***333000000*** |
| CloudHash&Light Client version 1.0 | ***10%***      | ***333000000*** |
| CloudHash&Light Client version 2.0 | ***15%***      | ***499500000*** |
| All The Nodes Mining               | ***25%***      | ***832500000*** |
| Maintenance                        | ***10%***      | ***333000000*** |
| Ecological Contribution            | ***5%***       | ***166500000*** |
| Community Maintenance              | ***10%***      | ***333000000*** |
| All The Nodes Reward               | ***10%***      | ***333000000*** |
| Incentive Mechanism                | ***5%***       | ***166500000*** |

## 6.0 Permanent Ecological Disclaimer ###########################################################

### 6.1 This Article Describes a Project Under Development

This white paper and related documents are used for the development and application of the Permanent Ecological network. For information dissemination purposes only and subject to change.

The Permanent Ecological network envisaged in this paper is under development and will be constantly updated, including but not limited to key governance and key technologies. The development of a test platform (software) and technology using the Permanent Ecological network or associated test platform (software) may not achieve or fully achieve the goals described in this white paper.

If the Permanent Ecological is accomplished, it may not be the same as described in this article. Nothing in this article represents or warrants the success or reasonableness of any future plans, projections or prospects, and nothing in this article shall be construed as a promise or representation of the future.

### 6.2 Not An Offer For Regulatory Products

Permanent Ecological does not represent any product under judicial supervision. This article does not constitute an offer or invitation to make an inquiry for any regulated product, nor does it constitute any promotion, invitation or inquiry for investment purposes. The terms of purchase are not documents providing financial services or prospectuses of any kind.

### 6.3 Is Not Recommended

This white paper does not constitute any purchase advice for PEE. Please do not rely on this white paper to make any contract or purchase decisions.

### 6.4 Risk Warning

Buying PEE and participating in the Permanent Ecological network comes with great risks. Before buying a PEE, you should carefully evaluate and consider the risks.

### 6.5 You Must Obtain All Necessary Professional Advice

It is important that you consult with an attorney, accountant and/or tax professional, as well as other professional advisers, before deciding whether to purchase a PEE or participate in a Permanent Ecological network program.

## 7.0 Permanent Ecological Related Literature ###########################################################

* BTC：[bitcoin.org/bitcoin.pdf][1]
* ETH：[github.com/ethereum/wiki/wiki/White-Paper][2]
* TheDAO：[download.slock.it/public/DAO/WhitePaper.pdf][3]
* BitcoinNG：[arxiv.org/pdf/1510.02037v2.pdf][4]
* Lightning Network：[lightning.network/lightning-network-paper-DRAFT-0.5.pdf][5]
* Tendermint：[github.com/tendermint/tendermint/wiki][6]
* PBFT：[pmg.csail.mit.edu/papers/osdi99.pdf][7]
* BitShares：[bitshares.org/technology/delegated-proof-of-stake-consensus][8]
* Interledger：[interledger.org/rfcs/0001-interledger-architecture][9]
* Side Chain：[blockstream.com/sidechains.pdf][10]
* ABCI：[github.com/tendermint/abci][11]
* DLS：[groups.csail.mit.edu/tds/papers/Lynch/jacm88.pdf][12]
* Thin Client Security：[en.bitcoin.it/wiki/Thin_Client_Security][13]
* Ethereum 2.0 Lavenderpaper：[vitalik.ca/files/mauve_paper.html][14]


[1]: https://bitcoin.org/bitcoin.pdf
[2]: https://github.com/ethereum/wiki/wiki/White-Paper
[3]: https://download.slock.it/public/DAO/WhitePaper.pdf
[4]: https://arxiv.org/pdf/1510.02037v2.pdf
[5]: https://lightning.network/lightning-network-paper-DRAFT-0.5.pdf
[6]: https://github.com/tendermint/tendermint/wiki
[7]: https://pmg.csail.mit.edu/papers/osdi99.pdf
[8]: https://bitshares.org/technology/delegated-proof-of-stake-consensus
[9]: https://interledger.org/rfcs/0001-interledger-architecture
[10]: https://blockstream.com/sidechains.pdf
[11]: https://github.com/tendermint/abci
[12]: https://groups.csail.mit.edu/tds/papers/Lynch/jacm88.pdf
[13]: https://en.bitcoin.it/wiki/Thin_Client_Security
[14]: https://vitalik.ca/files/mauve_paper.html
