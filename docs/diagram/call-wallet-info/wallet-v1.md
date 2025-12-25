# get wallet info

![wallet](./wallet-v1-doc-image.drawio.png)

## Description

1. user call `WalletInfo` endpoint from the wallet component
2. wallet component return response

## wallet

```
Name:   getWalletInfo
method: Get
Url:    http://localhost:9898/wallet/:number
Header: no content
Body:   no content
Errors:
    - code: 400
      Name: bad request
      Body:
         {
            "error" : "invalid number",
         }
    - code: 404
      Name: not found
      Body:
         {
            "error" : "number not found",
         }
    - code: 500
      Name: internal server error
      Body:
         {
            "error" : "internal server error",
         }
Responses:
      - code: 200
        Name: accepted
        Body:
            {
                "Balance" :          (int),
                "lase transaction" : (time.Time),
            }

```
