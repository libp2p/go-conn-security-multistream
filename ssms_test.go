package csms

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/libp2p/go-libp2p-core/sec/insecure"
	tnet "github.com/libp2p/go-libp2p-testing/net"
	sst "github.com/libp2p/go-libp2p-testing/suites/sec"
)

func TestCommonProto(t *testing.T) {
	idA := tnet.RandIdentityOrFatal(t)
	idB := tnet.RandIdentityOrFatal(t)

	var at, bt SSMuxer

	atInsecure := insecure.NewWithIdentity(idA.ID(), idA.PrivateKey())
	btInsecure := insecure.NewWithIdentity(idB.ID(), idB.PrivateKey())
	at.AddTransport("/plaintext/1.0.0", atInsecure)
	bt.AddTransport("/plaintext/1.1.0", btInsecure)
	bt.AddTransport("/plaintext/1.0.0", btInsecure)
	sst.SubtestRW(t, &at, &bt, idA.ID(), idB.ID())
}

func TestNoCommonProto(t *testing.T) {
	idA := tnet.RandIdentityOrFatal(t)
	idB := tnet.RandIdentityOrFatal(t)

	var at, bt SSMuxer
	atInsecure := insecure.NewWithIdentity(idA.ID(), idA.PrivateKey())
	btInsecure := insecure.NewWithIdentity(idB.ID(), idB.PrivateKey())

	at.AddTransport("/plaintext/1.0.0", atInsecure)
	bt.AddTransport("/plaintext/1.1.0", btInsecure)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	a, b := net.Pipe()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer a.Close()
		_, err := at.SecureInbound(ctx, a)
		if err == nil {
			t.Fatal("conection should have failed")
		}
	}()

	go func() {
		defer wg.Done()
		defer b.Close()
		_, err := bt.SecureOutbound(ctx, b, idA.ID())
		if err == nil {
			t.Fatal("connection should have failed")
		}
	}()
	wg.Wait()
}
