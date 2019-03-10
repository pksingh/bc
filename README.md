# Blockchain

This is going to have different Blockchain concepts.

- PoW [X] Proof of Work
- PoW [ ] Proof of Stake

## PoW /*Proof of Work*/

- **BUILD**
    Lets fire `go build -o bc_pow.exe .\bc_pow.go` on the console  
    This will generate binary/executable --> bc_pow or bc_pow.exe (in my case on Windows)

- **RUN**
    Lets RUN this `bc_pow.exe` or `bc_pow` based on the OS  

    *Output*:

    ```
    C:\Users\Home\Desktop\bc>bc_pow.exe
    Init Blockchain with Difficulty  4

    Adding Block:  Ravin -- 5 --> Binod
    PoW: 10170      Hash: 000011b524fb87c259f629ab59eabf00ac43f1cb, Data:map[amount:5 from:Ravin to:Binod]

    Adding Block:  Binod -- 3 --> Suresh
    PoW: 10404      Hash: 00006d21e563c8f58137401c0491dc7010591807, Data:map[amount:3 from:Binod to:Suresh]

    Adding Block:  Suresh -- 1 --> Ravin
    PoW: 22836      Hash: 0000247637cfaf9eda97949153ef43fbd6f3a2ff, Data:map[amount:1 from:Suresh to:Ravin]

    Validate Blockchain:  true

    Show Blockchain:
    0       TS:1546505999923189700  Hash:0  pHash:  PoW:0   Data:map[]
    1       TS:1546505999923784300  Hash:000011b524fb87c259f629ab59eabf00ac43f1cb   pHash:0         PoW:10170       Data:map[amount:5 from:Ravin to:Binod]
    2       TS:1546506000937132900  Hash:00006d21e563c8f58137401c0491dc7010591807   pHash:000011b524fb87c259f629ab59eabf00ac43f1cb  PoW:10404       Data:map[amount:3 from:Binod to:Suresh]
    3       TS:1546506002015701300  Hash:0000247637cfaf9eda97949153ef43fbd6f3a2ff   pHash:00006d21e563c8f58137401c0491dc7010591807  PoW:22836       Data:map[amount:1 from:Suresh to:Ravin]

    C:\Users\Home\Desktop\bc>
    ```

    *Animating Mining Output*:

    ```
    C:\Users\Home\Desktop\bc>bc_pow.exe
    Init Blockchain with Difficulty  4

    Adding Block:  Ravin -- 5 --> Binod
    Mining ... 000011b524fb87c259f629ab59eabf00ac43f1cb     PoW: 10170      Hash: 000011b524fb87c259f629ab59eabf00ac43f1cb, Data:map[amount:5 from:Ravin to:Binod]

    Adding Block:  Binod -- 3 --> Suresh
    Mining ... 00006d21e563c8f58137401c0491dc7010591807     PoW: 10404      Hash: 00006d21e563c8f58137401c0491dc7010591807, Data:map[amount:3 from:Binod to:Suresh]

    Adding Block:  Suresh -- 1 --> Ravin
    Mining ... 0000247637cfaf9eda97949153ef43fbd6f3a2ff     PoW: 22836      Hash: 0000247637cfaf9eda97949153ef43fbd6f3a2ff, Data:map[amount:1 from:Suresh to:Ravin]

    Validate Blockchain:  true

    Show Blockchain:
    0       TS:1546505999923189700  Hash:0  pHash:  PoW:0   Data:map[]
    1       TS:1546505999923784300  Hash:000011b524fb87c259f629ab59eabf00ac43f1cb   pHash:0         PoW:10170       Data:map[amount:5 from:Ravin to:Binod]
    2       TS:1546506000937132900  Hash:00006d21e563c8f58137401c0491dc7010591807   pHash:000011b524fb87c259f629ab59eabf00ac43f1cb  PoW:10404       Data:map[amount:3 from:Binod to:Suresh]
    3       TS:1546506002015701300  Hash:0000247637cfaf9eda97949153ef43fbd6f3a2ff   pHash:00006d21e563c8f58137401c0491dc7010591807  PoW:22836       Data:map[amount:1 from:Suresh to:Ravin]

    C:\Users\Home\Desktop\bc>
    ```

    ---

## PoS /*Proof of Stake*/

- **BUILD**
    Lets fire `go build -o bc_pos.exe ./pos/bc_pos.go` on the console  
    This will generate binary/executable --> bc_pos or bc_pos.exe (in my case on Windows)

- **RUN**
    Lets RUN this `bc_pos.exe` or `bc_pos` based on the OS  

    *Output*:
    ```
    Init PoSNetwork DONE
    InitValidators =>
            Id:  1  Address: DC961BADD322AD87  Stake: 20
            Id:  2  Address: 984419B79083B705  Stake: 30
            Id:  3  Address: DF29C24AC533C357  Stake: 50


    Lets Carry out few Transactions ...
    Adding New Block=> Data: Ravin -- 5 --> Binod
            Winner => Id:1 Stake:20
            TS: 1552232595107374  Hash: d5012f7f30d58e4b8f5612856338c649  vId: 1  vAddr: DC961BADD322AD87 Data: Ravin -- 5 --> Binod


    Lets Show/Print all Transactions in the Blockchain ...
    Block GEN>      TS: 1552232595103746  Hash: 1097b936847c3c997486c7b8bf426bf0  vId: 0  vAddr:  Data:
    Block 1 >       TS: 1552232595107374  Hash: d5012f7f30d58e4b8f5612856338c649  vId: 1  vAddr: DC961BADD322AD87 Data: Ravin -- 5 --> Binod
    ```

# License

MIT
