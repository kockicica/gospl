---
title: API  
has_children: true  
permalink: "/api"
---

## gospl proxy

gospl can be used as NBS web services proxy, exposing CLI commands as REST endpoints. To start server,
use [serve command](serve/index.md)

If using default configuration file for authorizations params, server can be started with

```shell
gospl serve
```

Expected console output should be something like:

```text
Using config file: /home/gospler/.gospl.yaml
I: 15:27:08 Running server on port: 31100
I: 15:27:08 Starting service:NBS proxy
I: 15:27:08 Started service:NBS proxy
```

Use curl to test is everything ok:

```shell
curl http://localhost:31100/core/currency
```

You should get bulky array of json objects describing existing currencies, looking like this:

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
  },
  {
    "currencyID": "5e555f23-fa28-44fc-b561-4c55705d5205",
    "currencyCode": 12,
    "currencyCodeNumChar": "012",
    "currencyCodeAlfaChar": "DZD",
    "unit": 1,
    "convertible": 0,
    "currencyNameSerCyrl": "Алжирски динар",
    "currencyNameSerLat": "Alžirski dinar",
    "currencyNameEng": "Algerian Dinar",
    "indicator": 1,
    "countryID": "ce6c6696-7882-45bc-ad0b-5c74c533e0ed",
    "countryCode": 12,
    "countryCodeNumChar": "012",
    "countryCodeAlfaChar3": "DZA",
    "countryCodeAlfaChar2": "DZ",
    "countryNameSerCyrl": "Алжир",
    "countryNameSerLat": "Alžir",
    "countryNameEng": "Algeria"
  },
  ...
]
```