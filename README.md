# Blockchain

This is going to have different Blockchain concepts.

- PoW [X] Proof of Work
- PoW [X] Proof of Stake

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
            Id:  1  Address: AB1BCCBD99BC6E8F  Stake: 20
            Id:  2  Address: ED4212D87FB2CE5A  Stake: 30
            Id:  3  Address: A15E974200E5F4B8  Stake: 50


    Lets Carry out few Transactions ...
    Adding New Block=> Data: Ravin -- 5 --> Binod
            Winner => Id:3 Stake:50
            TS: 1553408322151083  Hash: dd5d304cba74cf57810939c7f6144f97  vId: 3  vAddr: A15E974200E5F4B8 Data: Ravin -- 5 --> Binod

    Adding New Block=> Data: Binod -- 3 --> Suresh
            Winner => Id:2 Stake:30
            TS: 1553408322152196  Hash: 5dc299c1235bfbc8e76f76db36a5112d  vId: 2  vAddr: ED4212D87FB2CE5A Data: Binod -- 3 --> Suresh

    Adding New Block=> Data: Suresh -- 1 --> Ravin
            Winner => Id:3 Stake:60
            TS: 1553408322152769  Hash: b6f4d56346dcb552658be4c7ff50526a  vId: 3  vAddr: A15E974200E5F4B8 Data: Suresh -- 1 --> Ravin


    Lets Show/Print all Transactions in the Blockchain ...
    Block GEN>      TS: 1553408322149448  Hash: 406b2673c2f13abace944e04a386752d  vId: 0  vAddr:  Data:
    Block 1 >       TS: 1553408322151083  Hash: dd5d304cba74cf57810939c7f6144f97  vId: 3  vAddr: A15E974200E5F4B8 Data: Ravin -- 5 --> Binod
    Block 2 >       TS: 1553408322152196  Hash: 5dc299c1235bfbc8e76f76db36a5112d  vId: 2  vAddr: ED4212D87FB2CE5A Data: Binod -- 3 --> Suresh
    Block 3 >       TS: 1553408322152769  Hash: b6f4d56346dcb552658be4c7ff50526a  vId: 3  vAddr: A15E974200E5F4B8 Data: Suresh -- 1 --> Ravin
    ```

# License

MIT
