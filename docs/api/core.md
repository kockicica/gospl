---
title: "core"  
parent: API  
permalink: "/api/core"
---

## Core REST endpoints

### Get existing currencies

Request:

```http
  GET /core/currency
```

Request params:

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `currency-id` | `string` | currency id |
| `currency-code` | `int` | currency code |
| `currency-code-alpha` | `string` | currency code alpha char |

All parameters are optional

Response:

```json
[
  {
    "currencyID": "3940e111-1610-4155-b65f-22e2e0ffc14e",
    "currencyCode": 8,
    "currencyCodeNumChar": "008",
    "currencyCodeAlfaChar": "ALL",
    "unit": 1,
    "convertible": 0,
    "currencyNameSerCyrl": "Лек",
    "currencyNameSerLat": "Lek",
    "currencyNameEng": "Lek",
    "indicator": 1,
    "countryID": "6ec910eb-1999-4b2a-82b2-ade20276f913",
    "countryCode": 8,
    "countryCodeNumChar": "008",
    "countryCodeAlfaChar3": "ALB",
    "countryCodeAlfaChar2": "AL",
    "countryNameSerCyrl": "Албанија",
    "countryNameSerLat": "Albanija",
    "countryNameEng": "Albania"
  }
]
```

### Get bank

Request:

```http
  GET /core/bank
```

Request params:

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `bank-code` | `int` | bank code |

Response:

```json
[
  {
    "bankID": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "bankHistoryID": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "startDate": "2021-11-22T00:00:00+01:00",
    "endDate": "",
    "bankCode": 999,
    "nationalIdentificationNumber": 9999999,
    "taxIdentificationNumber": "",
    "name": "xxxx xxxxxxx xxxxxxxx",
    "logoSmall": "",
    "logoBig": "",
    "address": "xxxx xxxxxxxxxxxx xxxx",
    "city": "xxxxxxxx",
    "region": "",
    "state": "",
    "country": "",
    "postalCode": "xxxxx",
    "phone": "xxxxxxxxxxxx",
    "fax": "xxxxxxxxxxxx",
    "email": "xxxxxxxxxxxxxxxx",
    "webAddress": "xxxxxxxxxxxxx",
    "director": "xxxxxxxxxxxxxxxx",
    "description": "",
    "bankTypeID": 9,
    "bankStatusID": 9
  }
]
```

### Get bank status

Request:

```http
  GET /core/bank/status
```


Response:

```json
[
  {
    "bankStatusID": 1,
    "name": "Активан"
  },
  {
    "bankStatusID": 3,
    "name": "Ушла у састав"
  }
]
```


### Get bank type

Request:

```http
  GET /core/bank/type
```


Response:

```json
[
  {
    "bankTypeID": 1,
    "name": "Пословна банка"
  }
]
```
