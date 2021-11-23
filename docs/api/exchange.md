---
title: "exchange"  
parent: API  
permalink: "/api/exchange"
---

## Exchange REST endpoints

### Get exchange list types

Request:

```html
    GET /exchange/list-type
```

Response:

```json
[
  {
    "name": "xxxxxxxxx",
    "description": "",
    "exchangeRateListTypeID": 1
  }
]
```

### Get current exchange rate

Request:

```html
    GET /exchange/current
```

Request params:

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `list-type` | `int` | exchange list type |

Exchange list types may be queried using ```exchange/list-type``` request

Response:

```json
[
  {
    "exchangeRateListNumber": 225,
    "date": "23.11.2021",
    "exchangeRateListTypeID": 1,
    "currencyGroupID": 2,
    "currencyCode": 978,
    "currencyCodeNumChar": "978",
    "currencyCodeAlfaChar": "EUR",
    "currencyNameSerCyrl": "Евро",
    "currencyNameSerLat": "Evro",
    "currencyNameEng": "Euro",
    "countryNameSerCyrl": "ЕМУ",
    "countryNameSerLat": "EMU",
    "countryNameEng": "EMU",
    "unit": 1,
    "buyingRate": 117.2266,
    "middleRate": 0,
    "sellingRate": 117.932,
    "fixingRate": 0
  }
]
```

### Get current RSD - EUR exchange rate

Request:

```html
    GET /exchange/rsd-eur/current
```

Response:

```json
{
  "date": "2021-11-23T00:00:00+01:00",
  "typeID": 2,
  "amount": 117.584,
  "validityDate": "2021-11-24T00:00:00+01:00"
}
```

### Get exchange rate by currency

Request:

```html
    GET /exchange/currency
```

Request params:

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `list-type` | `int` | exchange list type |
| `currency-code` | `int` | currency code |
| `date-from` | `string` | start date interval, format: 'YYYYMMDD' **required** |
| `date-to` | `string` | start date interval, format: 'YYYYMMDD' **required** |


Response:

```json
{
        "exchangeRateListNumber": 210,
        "date": "01.11.2011",
        "createDate": "01.11.2011",
        "dateTo": "02.11.2011",
        "exchangeRateListTypeID": 1,
        "currencyGroupID": 2,
        "currencyCode": 978,
        "currencyCodeNumChar": "978",
        "currencyCodeAlfaChar": "EUR",
        "currencyNameSerCyrl": "Евро",
        "currencyNameSerLat": "Evro",
        "currencyNameEng": "Euro",
        "countryNameSerCyrl": "ЕМУ",
        "countryNameSerLat": "EMU",
        "countryNameEng": "EMU",
        "unit": 1,
        "BuyingRate": 100.6092,
        "MiddleRate": 0,
        "sellingRate": 101.2146,
        "fixingRate": 0
    }
```