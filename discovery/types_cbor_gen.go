// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package discovery

import (
	"fmt"
	"io"
	"math"
	"sort"

	retrievalmarket "github.com/filecoin-project/go-fil-markets/retrievalmarket"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *RetrievalPeers) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{161}); err != nil {
		return err
	}

	// t.Peers ([]retrievalmarket.RetrievalPeer) (slice)
	if len("Peers") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Peers\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Peers"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Peers")); err != nil {
		return err
	}

	if len(t.Peers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Peers was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Peers))); err != nil {
		return err
	}
	for _, v := range t.Peers {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *RetrievalPeers) UnmarshalCBOR(r io.Reader) (err error) {
	*t = RetrievalPeers{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("RetrievalPeers: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Peers ([]retrievalmarket.RetrievalPeer) (slice)
		case "Peers":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Peers: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Peers = make([]retrievalmarket.RetrievalPeer, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v retrievalmarket.RetrievalPeer
				if err := v.UnmarshalCBOR(cr); err != nil {
					return err
				}

				t.Peers[i] = v
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}