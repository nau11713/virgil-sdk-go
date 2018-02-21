# Virgil Security Go SDK

[![Build Status](https://travis-ci.org/go-virgil/virgil.png?branch=v5)](https://travis-ci.org/go-virgil/virgil)
[![GitHub license](https://img.shields.io/badge/license-BSD%203--Clause-blue.svg)](https://github.com/VirgilSecurity/virgil/blob/master/LICENSE)


[Introduction](#installation) | [SDK Features](#sdk-features) | [Installation](#installation) | [Usage Examples](#usage-examples) | [Docs](#docs) | [Support](#support)



## Introduction

[Virgil Security](https://virgilsecurity.com) provides a set of APIs for adding security to any application. In a few simple steps you can encrypt communication, securely store data, provide passwordless login, and ensure data integrity.

The Virgil SDK allows developers to get up and running with Virgil API quickly and add full end-to-end security to their existing digital solutions to become HIPAA and GDPR compliant and more.

## SDK Features
- communicate with [Virgil Cards Service][_cards_service]
- manage users' Public Keys
- store private keys in secure local storage
- use Virgil [Crypto library][_virgil_crypto]
- use your own Crypto


## Installation

The Virgil Go SDK is provided as a package named virgil. The package is distributed via github. Also in this guide, you find one more package called Virgil Crypto (Virgil Crypto Library) that is used by the SDK to perform cryptographic operations.

The package is available for Go 1.10 or newer.

Installing the package:

- go get -u gopkg.in/virgil.v5


### Crypto library notice

The built in crypto library supports only following primitives:

- ED25519 keys
- SHA512 hashes
- AES256_GCM encrypting

and is not recommended for production use.

On linux and macOS consider using external crypto library written in c++

### Using external crypto library (c++)

```bash
go get -u gopkg.in/virgilsecurity/virgil-crypto-go.v5
cd $GOPATH/src/gopkg.in/virgilsecurity/virgil-crypto-go.v5
make
```

in your source code use crypto objects from this library as follows:

```go

var (
	crypto      = virgil_crypto_go.NewCrypto()
	cardCrypto  = virgil_crypto_go.NewCardCrypto()
	tokenSigner = virgil_crypto_go.NewVirgilAccessTokenSigner()
)
```


## Usage Examples

#### Generate and publish user's Cards with Public Keys inside on Cards Service
Use the following lines of code to create and publish a user's Card with Public Key inside on Virgil Cards Service:

```go
import (
	"gopkg.in/virgil.v5/cryptoimpl"
	"gopkg.in/virgil.v5/sdk"
)
// generate a new Key
crypto := cryptoimpl.NewVirgilCrypto()
keypair, err := crypto.GenerateKeypair()

// save the Key into the storage
err = privateKeyStorage.Store(keypair.PrivateKey(), "Alice", nil)
if err != nil{
		//handle error
}
// create a Card
card, err := cardManager.PublishCard(&sdk.CardParams{
	PublicKey:  keypair.PublicKey(),
	PrivateKey: keypair.PrivateKey(),
	Identity:   "Alice",
}

if err != nil{
	//handle error
}
```

#### Sign then encrypt data

Virgil SDK lets you use a user's Private key and his or her Cards to sign, then encrypt any kind of data.

In the following example, we load a Private Key from a customized Key Storage and get recipient's Card from the Virgil Cards Services. Recipient's Card contains a Public Key on which we will encrypt the data and verify a signature.

```go

import (
	"gopkg.in/virgil.v5/cryptoimpl"
	"gopkg.in/virgil.v5/sdk"
)


messageToEncrypt := []byte("Hello, Bob!")

// load a Key from a device storage
alicePrivateKey, err := privateKeyStorage.Load("Alice")
if err != nil{
	//handle error
}

// using cardManager search for user's cards on Cards Service
cards, err := cardManager.SearchCards("Bob")

if err != nil{
	//handle error
}

//encrypt message for all valid Bob cards
encryptedMessage, err := crypto.SignThenEncrypt(messageToEncrypt, alicePrivateKey, cards.ExtractPublicKeys()...)

if err != nil{
	//handle error
}

```

#### Decrypt then verify data
Once the Users receive the signed and encrypted message, they can decrypt it with their own Private Key and verify signature with a Sender's Card:

```go
import (
	"gopkg.in/virgil.v5/cryptoimpl"
	"gopkg.in/virgil.v5/sdk"
)

// load a Key from a device storage
bobPrivateKey, err := privateKeyStorage.Load("Bob")
if err != nil{
	//handle error
}

// using cardManager search for user's cards on Cards Service
aliceCards, err := cardManager.SearchCards("Alice")

if err != nil{
	//handle error
}

//decrypt message and verify with one of Alice's valid
decryptedMessage, err := crypto.DecryptThenVerify(encryptedMessage, bobPrivateKey, cards.ExtractPublicKeys()...)

if err != nil{
	//handle error
}

```

## Docs
Virgil Security has a powerful set of APIs, and the documentation below can get you started today.

In order to use the Virgil SDK with your application, you will need to first configure your application. By default, the SDK will attempt to look for Virgil-specific settings in your application but you can change it during SDK configuration.

* [Configure the SDK][_configure_sdk] documentation
  * [Setup authentication][_setup_authentication] to make API calls to Virgil Services
  * [Setup Card Manager][_card_manager] to manage user's Public Keys
  * [Setup Card Verifier][_card_verifier] to verify signatures inside of user's Card
  * [Setup Key storage][_key_storage] to store Private Keys
  * [Setup your own Crypto library][_own_crypto] inside of the SDK
* [More usage examples][_more_examples]
  * [Create & publish a Card][_create_card] that has a Public Key on Virgil Cards Service
  * [Search user's Card by user's identity][_search_card]
  * [Get user's Card by its ID][_get_card]
  * [Use Card for crypto operations][_use_card]
* [Reference API][_reference_api]


## License

This library is released under the [3-clause BSD License](LICENSE).

## Support
Our developer support team is here to help you.

You can find us on [Twitter](https://twitter.com/VirgilSecurity) or send us email support@VirgilSecurity.com.

Also, get extra help from our support team on [Slack](https://join.slack.com/t/VirgilSecurity/shared_invite/enQtMjg4MDE4ODM3ODA4LTc2OWQwOTQ3YjNhNTQ0ZjJiZDc2NjkzYjYxNTI0YzhmNTY2ZDliMGJjYWQ5YmZiOGU5ZWEzNmJiMWZhYWVmYTM).

[_virgil_crypto]: https://github.com/VirgilSecurity/virgil-crypto-go/tree/master
[_cards_service]: https://developer.virgilsecurity.com/docs/api-reference/card-service/v5
[_use_card]: https://developer.virgilsecurity.com/docs/go/how-to/public-key-management/use-card-for-crypto-operation
[_get_card]: https://developer.virgilsecurity.com/docs/go/how-to/public-key-management/get-card
[_search_card]: https://developer.virgilsecurity.com/docs/go/how-to/public-key-management/search-card
[_create_card]: https://developer.virgilsecurity.com/docs/go/how-to/public-key-management/create-card
[_own_crypto]: https://developer.virgilsecurity.com/docs/go/how-to/setup/setup-own-crypto-library
[_key_storage]: https://developer.virgilsecurity.com/docs/go/how-to/setup/setup-key-storage
[_card_verifier]: https://developer.virgilsecurity.com/docs/go/how-to/setup/setup-card-verifier
[_card_manager]: https://developer.virgilsecurity.com/docs/go/how-to/setup/setup-card-manager
[_setup_authentication]: https://developer.virgilsecurity.com/docs/go/how-to/setup/setup-authentication
[_reference_api]: https://developer.virgilsecurity.com/docs/api-reference
[_configure_sdk]: https://developer.virgilsecurity.com/docs/how-to#sdk-configuration
[_more_examples]: https://developer.virgilsecurity.com/docs/how-to#public-key-management
