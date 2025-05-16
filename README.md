
## Installation

### Build from source (Linux)
Go version 1.24.2
mr mint
```bash
git clone https://github.com/Arunchauhan2000/mrmint.git
cd mrmint
go build -o mrmint
sudo cp ./mrmint /usr/local/bin/

Initialize chain

mrmintchain init mynode --chain-id os_9000-1 --home ~/.mrmint

Open genesis - nano ~/.mrmint/config/genesis.json

Edit below parameters:
Staking.params.bond_denom  →  replace “stake” with “aphoton”
Crisis.constant_fee.denom  →  replace “stake” with “aphoton”
Gov.deposit_params.denom  →  replace “stake” with “aphoton”
mint.params.mint_denom →  replace “stake” with “aphoton”
block_gas  → replace “block_gas” with “20000000”
max_gas  → replace “max_gas” with “100000000”

 Open config.toml - nano 	

Edit below parameters:
allow_duplicate_ip = false →true 		
addr_book_strict = true → false
Add in [consensus] section  block_gas_limit = 20000000
1.Maximum number of transactions in the mempool
size = 5000 → 1000000


2.  Create validator key 
mrmint keys add validator1--algo eth_secp256k1 --keyring-backend test

3.  Add genesis account 
mrmint add-genesis-account validator1 100000000000000000000000000aphoton --keyring-backend test

4. Create genesis trx
mrmint gentx validator1 1000000aphoton --chain-id os_9000-1 --keyring-backend test

Chain id should match


5. Add above validator to genesis
mrmint collect-gentxs

6. Start chain
mrmint start --home ~/.mrmint
