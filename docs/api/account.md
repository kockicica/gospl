---
title: account  
parent: API  
permalink: "/api/account"
---

## Account REST endpoints

### Get company account

Request:

```http
  GET /account
```

Request params:

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | company name (min 3 chars) |
| `national-id` | `int` | national identification number |
| `tax-id` | `string` | tax identification number |
| `bank-code` | `int` | bank code |
| `account-number` | `int` | account number |
| `control-number` | `int` | account control number |
| `city` | `string` | city |

There are some simple rules concerning request params:

- at least one of the ```name```, ```national-id``` or ```tax-id``` should be present
- if ```city``` is present, param ```name``` must also be set
- params ```bank-code``` ```account-number``` and ```control-number``` must be set together

Response:

```json
[
  {
    "account": "xxx-xxxxxxxx-xx",
    "bankCode": 111,
    "accountNumber": 11111111,
    "controlNumber": 11,
    "companyName": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "nationalIdentificationNumber": 99999999,
    "taxIdentificationNumber": "000000000    ",
    "address": "xxxxxxxxxxxxxxx  xx.x",
    "city": "xxxxxxx",
    "municipalityCode": 11111,
    "activityCode": 2222,
    "municipalityName": "xxxxxxxxxxxxxxxx",
    "activityName": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "bankName": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "companyAccountStatusID": 0,
    "companyAccountBlockadeStatusID": 0,
    "companyAccountTypeID": 1,
    "legalUserTypeID": 2,
    "initializationDate": "2018-09-10T00:00:00+02:00",
    "changeDate": "2018-09-25T00:00:00+02:00",
    "updateDate": "2018-09-25T19:30:08.013+02:00",
    "bankID": ""
  }
]
```

### Get company account status

Request:

```html
    GET /account/status
```

Response:

```json
[
  {
    "companyAccountStatusID": 0,
    "name": "xxxxxxx",
    "description": ""
  }
]
```

### Get company account type

Request:

```html
    GET /account/type
```

Response:

```json
[
  {
    "companyAccountTypeID": 1,
    "name": "xxxxxxxxxxxxxx",
    "description": ""
  }
]
```
