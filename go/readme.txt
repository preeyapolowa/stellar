----- install
1. go get -u github.com/gin-gonic/gin


----- run on port 5000 ------


----- to test Assignment 1 on Postman
Link API: localhost:5000/api/comparecoin/:ticker_symbol_1/:ticker_symbol_2
For example: localhost:5000/api/comparecoin/BTC/ETH
//where BTC is symbol of bitcoin and ETH is symbol of ethereum


----- to test Assignment 2 on Postman
Link API: localhost:5000/api/transfersteller/:source_account/:destination_account/:amount
For example: localhost:5000/api/transferstellar/SAKRBSS2RPQW7YZIJWAHC5IVAFWVOODUPTQBQEWI5ZZVAWW6JH63U5R7/GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW/100

**** TO USE
////You can check balance of Source Account on Postman, before you'll test : https://horizon-testnet.stellar.org/accounts/GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF
//where source_account = SAKRBSS2RPQW7YZIJWAHC5IVAFWVOODUPTQBQEWI5ZZVAWW6JH63U5R7
//where destination_account = GBGJRPOBWRIBIXWEWGMTD2PR2C7WBY2MCYI7GPUSD2T3PRCIZW65LGIT
//where amount = 100
//// You can check balance of Destination Account on Postman : https://horizon-testnet.stellar.org/accounts/GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW

Details
////keySource source_account Public Key = GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF
////keySource source_account Private Key = SAKRBSS2RPQW7YZIJWAHC5IVAFWVOODUPTQBQEWI5ZZVAWW6JH63U5R7
////keySource destination_account Public Key = GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW
////keySource destination_account Private Key = SDWAISJBW3X6T5CSJG5HPQGZ3PVMBZNGBSWWQOSWFXVGLQ3BIQ4BMINO


----- to test Assignment 2 on TestNet
//Source Account = GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF
//Transaction Squenece Number = (Can check Squenece number by using)
                                // https://horizon-testnet.stellar.org/accounts/GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF
                                // Example: 9648648325431302 + 1 
                                // Transaction Squenece Number = 9648648325431303
// Operation Type = Payment
// Destination = GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW
// Asset = native
// Amount = 100
// Click Sign in Transaction Signer Button on bottom or copy XDR, go to tab Transaction Signer and paste XDR
//// Sign with Private Key of Source Account = SAKRBSS2RPQW7YZIJWAHC5IVAFWVOODUPTQBQEWI5ZZVAWW6JH63U5R7
//// Click Submit to Post Transaction endpoint
//// Click Submit



----- to test Bonus on Postman
/// I already chanaged Trust limit to be 5000 to the Destination Account (GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW)
/// So I'm going to transfer my stellar token which is called "RYU" token to destination account

*** To USE
Link API: localhost:5000/api/bonus/:source_account/:destination_account/:amount
For example: localhost:5000/api/bonus/SAKRBSS2RPQW7YZIJWAHC5IVAFWVOODUPTQBQEWI5ZZVAWW6JH63U5R7/GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW/100
//// You can check balance of "RYU" token here : https://horizon-testnet.stellar.org/accounts/GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW
//// "RYU" token balance now: 1000.0000000

----- to test Bonus on TestNet
/// I used Public Key Source Account (GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF) and
    Destination Account (GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW)
/// Asset will be Alphanumeric 4 
/// Asset Code = RYU
/// Issuer Code = GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF
/// Amount = 100

Details
////keySource source_account Public Key = GAUYP7CGARUOL4S42OSWXLPBZHH36ZODE36CTD5DODYNFYD27LXWUJWF
////keySource source_account Private Key = SAKRBSS2RPQW7YZIJWAHC5IVAFWVOODUPTQBQEWI5ZZVAWW6JH63U5R7
////keySource destination_account Public Key = GDBGLWUFUH4IDFRNUQCLSB6WFC3PYPGHZSFDZ2RJ3SGHSXSIVSX5RALW
////keySource destination_account Private Key = SDWAISJBW3X6T5CSJG5HPQGZ3PVMBZNGBSWWQOSWFXVGLQ3BIQ4BMINO
