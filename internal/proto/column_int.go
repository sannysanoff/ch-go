package proto

import "github.com/go-faster/errors"

type ColumnUInt8 []uint8

func (ColumnUInt8) Type() ColumnType { return ColumnTypeUInt8 }
func (c ColumnUInt8) Rows() int      { return len(c) }
func (c *ColumnUInt8) Reset()        { *c = (*c)[:0] }

func (c ColumnUInt8) EncodeColumn(b *Buffer) {
	for _, v := range c {
		b.PutUInt8(v)
	}
}

func (c *ColumnUInt8) DecodeColumn(r *Reader, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := r.UInt8()
		if err != nil {
			return errors.Wrapf(err, "[%d]: read", i)
		}
		*c = append(*c, v)
	}
	return nil
}

type ColumnUInt32 []uint32

func (ColumnUInt32) Type() ColumnType { return ColumnTypeUInt32 }
func (c ColumnUInt32) Rows() int      { return len(c) }
func (c *ColumnUInt32) Reset()        { *c = (*c)[:0] }

func (c ColumnUInt32) EncodeColumn(b *Buffer) {
	for _, v := range c {
		b.PutUInt32(v)
	}
}

func (c *ColumnUInt32) DecodeColumn(r *Reader, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := r.UInt32()
		if err != nil {
			return errors.Wrapf(err, "[%d]: read", i)
		}
		*c = append(*c, v)
	}
	return nil
}
