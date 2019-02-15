----- install
1. go get -u github.com/gin-gonic/gin
2. go get -u github.com/golang/dep/cmd/dep
3. dep init
4. dep ensure -v


----- to test Assignment 1 on Postman
localhost:5000/api/comparecoin/:ticker_symbol_1/:ticker_symbol_2
For example: localhost:5000/api/comparecoin/BTC/ETH
//where BTC is symbol of bitcoin and ETH is symbol of ethereum


----- to test Assignment 2 on Postman
localhost:5000/api/transfersteller/:source_account/:destination_account/:amount

**** TO USE
////You can check balance of Source Account on Postman, before you'll test : https://horizon-testnet.stellar.org/accounts/GCSBBIE5OYXM4TDCLKTV3I4ISVY4MUX2CKJBFFWYP6OAUVPLREI3RR42
//where source_account = SDCLYLJXG6RGOVFWFB7AJQ5E7RNNNGEAEW7A4TSOWBADGHQZ7H4AUEPY
//where destination_account = GBGJRPOBWRIBIXWEWGMTD2PR2C7WBY2MCYI7GPUSD2T3PRCIZW65LGIT
//where amount = 100
//// For example: localhost:5000/api/SDCLYLJXG6RGOVFWFB7AJQ5E7RNNNGEAEW7A4TSOWBADGHQZ7H4AUEPY/GBGJRPOBWRIBIXWEWGMTD2PR2C7WBY2MCYI7GPUSD2T3PRCIZW65LGIT/100
//// You can check balance of Destination Account on Postman : https://horizon-testnet.stellar.org/accounts/GBGJRPOBWRIBIXWEWGMTD2PR2C7WBY2MCYI7GPUSD2T3PRCIZW65LGIT


Details
////keySource source_account Public Key = GCSBBIE5OYXM4TDCLKTV3I4ISVY4MUX2CKJBFFWYP6OAUVPLREI3RR42
////keySource source_account Private Key = SDCLYLJXG6RGOVFWFB7AJQ5E7RNNNGEAEW7A4TSOWBADGHQZ7H4AUEPY
////keySource destination_account Public Key = GBGJRPOBWRIBIXWEWGMTD2PR2C7WBY2MCYI7GPUSD2T3PRCIZW65LGIT
////keySource destination_account Private Key = SAXJZMCMZWBQIO2ZJUC47Z657AKMAPPJMFTHHMQ4BXVTZ6NUABXRIHNW
////Base Fee = 100 (0.00001 XLM/Lumen)

----- to test Assignment 2 on TestNet
//Source Account = GCSBBIE5OYXM4TDCLKTV3I4ISVY4MUX2CKJBFFWYP6OAUVPLREI3RR42
//Transaction Squenece Number = (Can check Squenece number by using)
                                // https://horizon-testnet.stellar.org/accounts/GCSBBIE5OYXM4TDCLKTV3I4ISVY4MUX2CKJBFFWYP6OAUVPLREI3RR42
                                // Example: 9587771458977798 + 1 
                                // Transaction Squenece Number = 9587771458977799
// Operation Type = Payment
// Destination = GBGJRPOBWRIBIXWEWGMTD2PR2C7WBY2MCYI7GPUSD2T3PRCIZW65LGIT
// Asset = native
// Amount = 100
//// This process you will get XDR, go to tab Transaction Signer
//// Put XDR Format
//// Sign with Private Key of Source Account = SDCLYLJXG6RGOVFWFB7AJQ5E7RNNNGEAEW7A4TSOWBADGHQZ7H4AUEPY
//// Submit Transaction


----- to test Bonus on Postman