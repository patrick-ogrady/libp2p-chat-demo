# LibP2P Chat Demo

This demo combines the `routed-echo` and `chat-with-rendezvous` demo found in [go-libp2p-examples](https://github.com/libp2p/go-libp2p-examples) and borrows shamelessly from that repo.

## Build

```
> make
```

## Usage
On one computer (or terminal) run the command `./chat -p 1000` where `1000` is the port the libp2p node will listen on. It should look something like this:
```
> ./chat -p 1000
2019/12/18 21:00:49 bootstrapped with QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM
2019/12/18 21:00:49 bootstrapped with QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd
2019/12/18 21:00:50 bootstrapped with QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu
2019/12/18 21:00:54 NODE ADDRESSES:
2019/12/18 21:00:54 /ip4/127.0.0.1/tcp/1001/p2p/QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ
2019/12/18 21:00:54 /ip4/127.94.0.1/tcp/1001/p2p/QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ
2019/12/18 21:00:54 /ip4/192.168.1.77/tcp/1001/p2p/QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ
2019/12/18 21:00:54 Run "./chat -p 1001 -d QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ" on a different terminal or computer to start a chat
2019/12/18 21:00:54 listening for streams....
```

On another computer (or terminal) run the command `./chat -p 1001 -d QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ` where `1001` is the port where the other libp2p node with listen and `QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ` is the identify of the first node started. If you are running these nodes on the same computer, make sure that the ports you use are different. It should look something like this:
```
2019/12/18 21:08:59 bootstrapped with QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM
2019/12/18 21:08:59 bootstrapped with QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd
2019/12/18 21:09:00 bootstrapped with QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu
2019/12/18 21:09:09 NODE ADDRESSES:
2019/12/18 21:09:09 /ip4/127.0.0.1/tcp/1002/p2p/QmRMBN32fxhdRs1QvfJsXiqUcAnPfx3ABwkd4fdCqLUbD2
2019/12/18 21:09:09 /ip4/127.94.0.1/tcp/1002/p2p/QmRMBN32fxhdRs1QvfJsXiqUcAnPfx3ABwkd4fdCqLUbD2
2019/12/18 21:09:09 /ip4/192.168.1.77/tcp/1002/p2p/QmRMBN32fxhdRs1QvfJsXiqUcAnPfx3ABwkd4fdCqLUbD2
2019/12/18 21:09:09 /ip4/71.145.210.138/tcp/1002/p2p/QmRMBN32fxhdRs1QvfJsXiqUcAnPfx3ABwkd4fdCqLUbD2
2019/12/18 21:09:09 stream opened with QmQQm1Zfi6rx88JCy5TteHqAQ9CNQYsSbkH6nJuNho7hCJ
> 
```
You can now message between peers.

# TBD
This example is intended to show how easy it is to create a p2p chat app using libp2p the ipfs distributed hash table to lookup peers.

Functionally, this example works similarly to the echo example, however setup of the host includes wrapping it with a Kademila hash table, so it can find peers using only their IDs. 

We'll also enable NAT port mapping to illustrate the setup, although it isn't guaranteed to actually be used to make the connections.  Additionally, this example uses the newer `libp2p.New` constructor.

## Build

From `go-libp2p-examples` base folder:

```
> make deps
> go build ./routed-echo
```

## Usage


```
> ./routed-echo -l 10000
2018/02/19 12:22:32 I can be reached at:
2018/02/19 12:22:32 /ip4/127.0.0.1/tcp/10000/ipfs/QmfRY4vuKpU2tApACrbmYFn9xoeNzMQhLXg7nKnyvnzHeL
2018/02/19 12:22:32 /ip4/192.168.1.203/tcp/10000/ipfs/QmfRY4vuKpU2tApACrbmYFn9xoeNzMQhLXg7nKnyvnzHeL
2018/02/19 12:22:32 Now run "./routed-echo -l 10001 -d QmfRY4vuKpU2tApACrbmYFn9xoeNzMQhLXg7nKnyvnzHeL" on a different terminal
2018/02/19 12:22:32 listening for connections
```

The listener libp2p host will print its randomly generated Base58 encoded ID string, which combined with the ipfs DHT, can be used to reach the host, despite lacking other connection details.  By default, this example will bootstrap off your local IPFS peer (assuming one is running). If you'd rather bootstrap off the same peers go-ipfs uses, pass the `-global` flag in both terminals.

Now, launch another node that talks to the listener:

```
> ./routed-echo -l 10001 -d QmfRY4vuKpU2tApACrbmYFn9xoeNzMQhLXg7nKnyvnzHeL
```

As in other examples, the new node will send the message `"Hello, world!"` to the listener, which will in turn echo it over the stream and close it. The listener logs the message, and the sender logs the response.

## Details

The `makeRoutedHost()` function creates a [go-libp2p routedhost](https://godoc.org/github.com/libp2p/go-libp2p/p2p/host/routed) object. `routedhost` objects wrap [go-libp2p basichost](https://godoc.org/github.com/libp2p/go-libp2p/p2p/host/basic) and add the ability to lookup a peers address using the ipfs distributed hash table as implemented by [go-libp2p-kad-dht](https://godoc.org/github.com/libp2p/go-libp2p-kad-dht).

In order to create the routed host, the example needs:

- A [go-libp2p basichost](https://godoc.org/github.com/libp2p/go-libp2p/p2p/host/basic) as in other examples.
- A [go-libp2p-kad-dht](https://godoc.org/github.com/libp2p/go-libp2p-kad-dht) which provides the ability to lookup peers by ID.  Wrapping takes place via `routedHost := rhost.Wrap(basicHost, dht)`

A `routedhost` can now open streams (bi-directional channel between to peers) using [NewStream](https://godoc.org/github.com/libp2p/go-libp2p/p2p/host/basic#BasicHost.NewStream) and use them to send and receive data tagged with a `Protocol.ID` (a string). The host can also listen for incoming connections for a given
`Protocol` with [`SetStreamHandle()`](https://godoc.org/github.com/libp2p/go-libp2p/p2p/host/basic#BasicHost.SetStreamHandler).  The advantage of the routed host is that only the Peer ID is required to make the connection, not the underlying address details, since they are provided by the DHT.

The example makes use of all of this to enable communication between a listener and a sender using protocol `/echo/1.0.0` (which could be any other thing).
