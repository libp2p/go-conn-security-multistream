// Deprecated: This package has moved into go-libp2p as a sub-package: github.com/libp2p/go-libp2p/p2p/net/conn-security-multistream.
package csms

import csms "github.com/libp2p/go-libp2p/p2p/net/conn-security-multistream"

// SSMuxer is safe to use without initialization. However, it's not safe to move
// after use.
//
// Deprecated: use github.com/libp2p/go-libp2p/p2p/net/conn-security-multistream.SSMuxer instead.
type SSMuxer = csms.SSMuxer
