package postings

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
	"github.com/prometheus/prometheus/tsdb/encoding"
	"github.com/prometheus/prometheus/tsdb/index"
)

const (
	RoaringEncoder = "roaringbitmap"
	RawEncoder     = "raw"
)

var _ index.PostingsEncoder = EncodePostingsRoaring

func EncodePostingsRoaring(e *encoding.Encbuf, in []uint32) error {
	bm := roaring.NewBitmap()

	bm.AddMany(in)
	bm.RunOptimize()

	out, err := bm.MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshaling roaring bitmap: %w", err)
	}

	e.PutBytes(out)
	return nil
}
