
### /v1/address

#### GET
##### Summary

Get accounts

##### Description

Submit public key to receive accounts

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | return 200 with the array of accounts. | [ [ [models.Wallet](#modelswallet) ] ] |

#### POST
##### Summary

Create an address

##### Description

create address use public key

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| publicKey | body | public key | Yes | string |
| signatureAlgorithm | body | sign algorithm | Yes | string |
| hashAlgorithm | body | hash algorithm | Yes | string |
| weight | body | weight of the key | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | return 200 with the transaction id. | [controllers.WalletReturn](#controllerswalletreturn) |

### /v1/address/testnet

#### GET
##### Summary

Get accounts

##### Description

Submit public key to receive accounts

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | return 200 with the array of accounts from testnet. | [ [ [models.Wallet](#modelswallet) ] ] |

#### POST
##### Summary

Create an address

##### Description

use public key to create address

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| publicKey | body | public key | Yes | string |
| signatureAlgorithm | body | sign algorithm | Yes | string |
| hashAlgorithm | body | hash algorithm | Yes | string |
| weight | body | weight of the key | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | return 200 with the transaction id. | [controllers.WalletReturn](#controllerswalletreturn) |

### Models

#### controllers.WalletReturn

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| txId | string |  | No |

#### models.Wallet

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string |  | No |
| created_at | string |  | No |
| hash_algo | integer |  | No |
| hash_algo_string | string |  | No |
| id | integer |  | No |
| public_key | string |  | No |
| sign_algo | integer |  | No |
| sign_algo_string | string |  | No |
| updated_at | string |  | No |
| weight | integer |  | No |
