---
title: POKTRollSDK
sidebar_position: 1
---

# POKTRollSDK <!-- omit in toc -->

- [What is the POKTRollSDK?](#what-is-the-poktrollsdk)
- [Target audience](#target-audience)
  - [Applications](#applications)
  - [Gateways](#gateways)
- [POKTRollSDK API](#poktrollsdk-api)
  - [GetSessionSupplierEndpoints](#getsessionsupplierendpoints)
  - [SendRelay](#sendrelay)
  - [NewPOKTRollSDK](#newpoktrollsdk)
  - [POKTRollSDKConfig](#poktrollsdkconfig)
- [POKTRollSDK usage](#poktrollsdk-usage)
  - [Example usage](#example-usage)
- [POKTRollSDK sequence diagram](#poktrollsdk-sequence-diagram)
- [How to contribute](#how-to-contribute)
- [Where to get help](#where-to-get-help)

## What is the POKTRollSDK?

`POKTRollSDK` is a package that provides the needed functionality for `Gateway`s
and `Application`s to interact with the `Supplier`s of the Pocket Network, allowing
them to do so in a way that complies with the Pocket Network's protocol.

It takes care of providing a list of `Supplier`s that are allowed to serve the
`Application` based on the service requested and the current block height. It
also takes care of signing the relay requests, forwarding them to the selected
`Supplier` and verifying the response signature.

It lets `Application` and `Gateway` developers to easily integrate the
Pocket Network into their workflow while leaving room for customization of the
different aspects of the relay request and response lifecycle.

## Target audience

`POKTRollSDK` is intended to be used by `Gateway`s and `Application`s that want
to interact with the Pocket Network in a way that complies with the Pocket Network's
protocol.

### Applications

`Application`s that want to use the Pocket Network to query the services provided
by the Pokt Network's `Supplier`s without resorting to a `Gateway` can integrate the
`POKTRollSDK` into their code in order to permissionlessly interact with the
available `Supplier`s, Given that they staked the required amount of POKT for
the services they want to use.

The `POKTRollSDK` takes care of the following:

* Providing them with a list of `Supplier`s that are allowed to serve the `Application`.
It is up to the `Application` to implement the desired strategy of selecting the
`Supplier` that will serve the request.
* It also handles the proper signing of the relay requests
* Forwards them to the selected `Supplier`
* Verifies the response signature.

The following diagram shows the different components involved in the case of an
`Application` integrating the `POKTRollSDK` into their workflow.

```mermaid
flowchart RL
	SDK[POKTRollSDK]
  A[Application logic]
  S[Supplier]
  Node[PocketNode]

  subgraph Application runtime
      A -- Suppliers list --> SDK
      A <-- Relay Req/Res --> SDK
  end
  Node <-. websocket subscription .-> SDK
  SDK -- Session --> Node
  SDK -- RelayRequest --> S
  S -- RelayResponse --> SDK
```

### Gateways

`Gateway`s are services that forward relay requests to `Supplier`s on behalf of
`Application`s. End-users that do not want or cannot run their own `Application`
logic such as wallets, can use `Gateway`s to interact with the Pocket Network.

By integrating the `POKTRollSDK` into their workflow, Operators can permissionlessly
setup a `Gateway` that complies with the Pocket Network's protocol.
Given that they staked the required amount of POKT and have `Application`s delegating
POKT tokens to them, `Gateway`s can customize:

* The way they authenticate their users.
* How they charge them for the service provided
* How they select the `Supplier` that will serve the request by running QoS tests.

While having the `POKTRollSDK` takeing care of:

* Providing a list of `Supplier`s that are allowed to serve the `Application`.
* Handling the proper signing of the `RelayRequests`.
* Forwarding them to the selected `Supplier`
* Verifying the `RelayResponse` signature.

The following diagram shows the different components involved in the case of a
`Gateway` integrating the `POKTRollSDK` into their workflow.

```mermaid
flowchart RL
	SDK[POKTRollSDK]
  G[Gateway]
  S[Supplier]
  Network[PocketNetwork]
  Node[PocketNode]

  subgraph Gateway infra
    subgraph Gateway runtime
        G -- Session --> SDK
        G <-- Relay Req/Res --> SDK
    end
    Node <-. websocket subscription .-> SDK
    SDK -- Session --> Node
  end
  Network <-. websocket subscription .-> Node
  SDK -- RelayRequest --> S
  S -- RelayResponse --> SDK
```

## POKTRollSDK API

`POKTRollSDK` consists of an interface that exposes the following methods:

```go
type POKTRollSDK interface {
	GetSessionSupplierEndpoints(
		ctx context.Context,
		appAddress string,
		serviceId string,
	) (session *SessionSuppliers, err error)

	SendRelay(
		ctx context.Context,
		sessionSupplierEndpoint *SingleSupplierEndpoint,
		request *http.Request,
	) (response *servicetypes.RelayResponse, err error)
}
```

The interface, its implementation and the relevant types involved could be found
at the following directory [pkg/sdk/](https://github.com/pokt-network/poktroll/blob/main/pkg/sdk)

```
pkg/sdk/
├── deps_builder.go   // Logic to build the dependencies of the SDK
├── errors.go         // Errors returned by the SDK
├── interface.go      // POKTRollSDK interface
├── relay_verifier.go // Logic to verify the relay response
├── sdk.go            // POKTRollSDK implementation
├── send_relay.go     // Logic to send the relay request
├── session.go        // Logic to handle the session and retrieve the suppliers list
└── urls.go           // Helpers to parse the urls used by the SDK to interact with the Pocket Network
```

### GetSessionSupplierEndpoints

`GetSessionSupplierEndpoints` returns a `SessionSuppliers` struct that contains
the fully-hydrated session corresponding to the `appAddress` and `serviceId`
provided and a list of `SingleSupplierEndpoint`s, where each `SingleSupplierEndpoint`
contains the necessary information to send a relay request to the `Supplier` it represents.

The `POKTRollSDK` consumer is free to run any strategy to select a `Supplier` from
the list returned by the `GetSessionSupplierEndpoints` method.

### SendRelay

Once the `Supplier` is selected, the `POKTRollSDK` consumer can pass the corresponding
`SingleSupplierEndpoint` to the `SendRelay` method which will take care of constructing
the `RelayRequest`, sending it to the `Supplier` and verifying the response signature.

Once the `RelayRequest` is sent, the `SendRelay` method will wait for the `RelayResponse`
and verify its signature before returning it to the `POKTRollSDK` consumer.

:::warning

The error returned by the `SendRelay` only indicates that an error occurred during
the process of sending the `RelayRequest` or verifying the `RelayResponse` signature
and does not indicate errors returned by the `Supplier` or the requested service.
These ones should be available in the `RelayResponse` returned by the `SendRelay`
and the `POKTRollSDK` considers them as valid responses.

:::

### NewPOKTRollSDK

`NewPOKTRollSDK` is an initializing function that returns a fully functional
`POKTRollSDK` implementation provided its `POKTRollSDKConfig` struct which contains
the necessary information to build the dependencies of the `POKTRollSDK` implementation.

```go
func NewPOKTRollSDK(
  ctx context.Context,
  config *POKTRollSDKConfig,
) (POKTRollSDK, error)
```

### POKTRollSDKConfig

`POKTRollSDKConfig` is a struct that contains the necessary information to build
the dependencies needed by the `POKTRollSDK` implementation.

```go
type POKTRollSDKConfig struct {
	QueryNodeGRPCUrl *url.URL
	QueryNodeUrl     *url.URL
	PrivateKey       cryptotypes.PrivKey
	Deps             depinject.Config
}
```

It consists of the following fields:

* `QueryNodeGRPCUrl` is the url of the Pocket Node's gRPC endpoint, used to query
the Pocket Network's state for sessions, account information, delegations, etc.
* `QueryNodeUrl` is the url of the Pocket Node's HTTP endpoint, used to subscribe
to the Pocket Network's new block events needed to keep the `POKTRollSDK` session
information up to date.
* `PrivateKey` is the private key used to sign the relay requests. It could be
either the `Gateway` or the `Application` private key depending on the use case.
* `Deps` is a `depinject.Config` struct that contains the dependencies needed by
the `POKTRollSDK` implementation. It is used to inject the dependencies into the
`POKTRollSDK` implementation. This field is optional and if not provided, the
`POKTRollSDK` implementation will use the default dependencies.

## POKTRollSDK usage

In order to use the `POKTRollSDK` the consumer needs to:

1. Import the `POKTRollSDK` package
2. Initialize a new `POKTRollSDK` instance by calling the `NewPOKTRollSDK` function
and providing the necessary `POKTRollSDKConfig` struct.
3. Call the `GetSessionSupplierEndpoints` method to get the `SessionSuppliers`.
4. Select a `Supplier` from the list of `SingleSupplierEndpoint`s returned.
5. Call the `SendRelay` method providing the selected `SingleSupplierEndpoint` and
the request to send.

### Example usage

:::warning

The code below is for illustrative purposes only. It shows how `GetSessionSupplierEndpoints`
and `SendRelay` could be used in a simple scenario. It does not show how to handle
errors or how to select a `Supplier` from the list of `SingleSupplierEndpoint`s

:::

```go
package main

import sdk "github.com/pokt-network/poktroll/pkg/sdk"

func main() {

  // Build the POKTRollSDKConfig struct
  sdkConfig := sdk.POKTRollSDKConfig{
      QueryNodeGRPCUrl grpcURL,
      QueryNodeUrl     rpcURL
      PrivateKey:      privateKey,
      // Deps are omitted and the default dependencies will be used
  }

  // Initialize a new POKTRollSDK instance
  poktrollSDK, err := sdk.NewPOKTRollSDK(ctx, &sdkConfig)

  // Get the session and the corresponding list of suppliers
  sessionSupplier, err := poktrollSDK.GetSessionSupplierEndpoints(
    ctx,
    appAddress,
    serviceId,
  )

  // Naively select the first supplier from the list of SingleSupplierEndpoints
  selectedSupplier := sessionSupplier.SuppliersEndpoints[0]

  // Send the request to the selected supplier and wait for the response
  response, err := poktrollSDK.SendRelay(ctx, selectedSupplier, httpRequest)
}
```

## POKTRollSDK sequence diagram

The following diagram shows the sequence of events that take place when an `Application`
or a `Gateway` uses the `POKTRollSDK` to interact with the Pocket Network.

```mermaid
sequenceDiagram
    participant G as "Gateway/Application"
    participant SDK as POKTRollSDK
    participant Node as PocketNode
    participant Network as PocketNetwork
    participant S as Supplier

		loop async network communication
	    Network -->> Node: Blocks subscription
	    Node -->> SDK: Blocks subscription
		end
		%% Session Retrieval
		alt only when a new block is commited
			SDK -->> G: Block
	    G ->> +SDK: GetSession
	    SDK ->> +Node: GetSession
	    Node ->> -SDK: Session
			SDK ->> -G: Session
		end
		%% Relay Propagation
    G ->> +SDK: SendRelay
    SDK ->> +S: RelayRequest
    S ->> -SDK: RelayResponse
		SDK ->> -G: RelayResponse
```

## How to contribute

If you want to contribute to the `POKTRollSDK` you can find the current issues
and the project's roadmap in the [Issues](https://github.com/pokt-network/poktroll/labels/sdk)
section of the repository.

Feel free to open a new issue and add the `sdk` label if you find a bug or if
you have a feature request.

You can also open a PR if you want to contribute with a new feature or a bug fix.

## Where to get help

If you want to discuss the `POKTRollSDK`, you can join the [Pocket Network Discord](https://discord.gg/build-with-grove) and join the `#protocol-public` channel.

:::note

The `POKTRollSDK` is still in its early stages and it is subject to change. We
will try to keep the changes to a minimum and to keep the community informed of
any changes that could affect the `POKTRollSDK` consumers.

:::