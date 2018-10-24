# Blockchain
This is going to have different Blockchain concepts.
- PoW [X] Proof of Work


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

    ```

# License

MIT
