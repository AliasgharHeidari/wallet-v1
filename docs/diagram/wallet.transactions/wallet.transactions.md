# wallet transactions

![wallet transactions](./wallet.transactions.png)

## Description

1. user put phone number in request url & calls `transactions` endpont from the wallet component.
2. wallet component checks if phone number exist and returns request's response.(wallet transactions including : value, date, ID for each transaction)

# Api contract

## wallet transactions

```
Name: transactions
Method: Get
Url: http://localhost:9898/wallet/transactions/:number
Headers:
Body:
Errors:
    - code: 400
      Name: bad request
      Body:
          {
            "error" : "invalid request url",
          }
    - code: 404
      Name: not found
      Body:
          {
            "error" : "wallet does not exist",
          }
    - code: 500
      Name: internal server error
      Body:
          {
            "error" : "internal error, please try again later",
          }
Responses:
    - code: 200
      Name: ok
      Body:
          {
            "transaction-ID" : (int),
            "value" :          (unit64),
            "date" :           (Time.time)
          }
```
