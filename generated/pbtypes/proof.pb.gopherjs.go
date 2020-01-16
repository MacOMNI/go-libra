// +build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: proof.proto

package pbtypes

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type AccumulatorProof struct {
	// The siblings. The ones near the leaf are at the beginning of the list. The
	// placeholder nodes are represented by empty byte arrays, other nodes should
	// be exactly 32-bytes long.
	Siblings [][]byte
}

// GetSiblings gets the Siblings of the AccumulatorProof.
func (m *AccumulatorProof) GetSiblings() (x [][]byte) {
	if m == nil {
		return x
	}
	return m.Siblings
}

// MarshalToWriter marshals AccumulatorProof to the provided writer.
func (m *AccumulatorProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, val := range m.Siblings {
		writer.WriteBytes(1, val)
	}

	return
}

// Marshal marshals AccumulatorProof to a slice of bytes.
func (m *AccumulatorProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccumulatorProof from the provided reader.
func (m *AccumulatorProof) UnmarshalFromReader(reader jspb.Reader) *AccumulatorProof {
	for reader.Next() {
		if m == nil {
			m = &AccumulatorProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Siblings = append(m.Siblings, reader.ReadBytes())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccumulatorProof from a slice of bytes.
func (m *AccumulatorProof) Unmarshal(rawBytes []byte) (*AccumulatorProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type SparseMerkleProof struct {
	// This proof can be used to authenticate whether a given leaf exists in the
	// tree or not. In Rust:
	//   - If this is `Some(HashValue, HashValue)`
	//     - If the first `HashValue` equals requested key, this is an inclusion
	//       proof and the second `HashValue` equals the hash of the
	//       corresponding account blob.
	//     - Otherwise this is a non-inclusion proof. The first `HashValue` is
	//       the only key that exists in the subtree and the second `HashValue`
	//       equals the hash of the corresponding account blob.
	//   - If this is `None`, this is also a non-inclusion proof which indicates
	//     the subtree is empty.
	//
	// In protobuf, this leaf field should either be
	//   - empty, which corresponds to None in the Rust structure.
	//   - exactly 64 bytes, which corresponds to Some<(HashValue, HashValue)>
	//     in the Rust structure.
	Leaf []byte
	// The siblings. The ones near the leaf are at the beginning of the list. The
	// placeholder nodes are represented by empty byte arrays, other nodes should
	// be exactly 32-bytes long.
	Siblings [][]byte
}

// GetLeaf gets the Leaf of the SparseMerkleProof.
func (m *SparseMerkleProof) GetLeaf() (x []byte) {
	if m == nil {
		return x
	}
	return m.Leaf
}

// GetSiblings gets the Siblings of the SparseMerkleProof.
func (m *SparseMerkleProof) GetSiblings() (x [][]byte) {
	if m == nil {
		return x
	}
	return m.Siblings
}

// MarshalToWriter marshals SparseMerkleProof to the provided writer.
func (m *SparseMerkleProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Leaf) > 0 {
		writer.WriteBytes(1, m.Leaf)
	}

	for _, val := range m.Siblings {
		writer.WriteBytes(2, val)
	}

	return
}

// Marshal marshals SparseMerkleProof to a slice of bytes.
func (m *SparseMerkleProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SparseMerkleProof from the provided reader.
func (m *SparseMerkleProof) UnmarshalFromReader(reader jspb.Reader) *SparseMerkleProof {
	for reader.Next() {
		if m == nil {
			m = &SparseMerkleProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Leaf = reader.ReadBytes()
		case 2:
			m.Siblings = append(m.Siblings, reader.ReadBytes())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SparseMerkleProof from a slice of bytes.
func (m *SparseMerkleProof) Unmarshal(rawBytes []byte) (*SparseMerkleProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type AccumulatorConsistencyProof struct {
	// The root hashes of the subtrees that represent new leaves. Note that none
	// of these hashes should be default hash.
	Subtrees [][]byte
}

// GetSubtrees gets the Subtrees of the AccumulatorConsistencyProof.
func (m *AccumulatorConsistencyProof) GetSubtrees() (x [][]byte) {
	if m == nil {
		return x
	}
	return m.Subtrees
}

// MarshalToWriter marshals AccumulatorConsistencyProof to the provided writer.
func (m *AccumulatorConsistencyProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, val := range m.Subtrees {
		writer.WriteBytes(1, val)
	}

	return
}

// Marshal marshals AccumulatorConsistencyProof to a slice of bytes.
func (m *AccumulatorConsistencyProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccumulatorConsistencyProof from the provided reader.
func (m *AccumulatorConsistencyProof) UnmarshalFromReader(reader jspb.Reader) *AccumulatorConsistencyProof {
	for reader.Next() {
		if m == nil {
			m = &AccumulatorConsistencyProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Subtrees = append(m.Subtrees, reader.ReadBytes())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccumulatorConsistencyProof from a slice of bytes.
func (m *AccumulatorConsistencyProof) Unmarshal(rawBytes []byte) (*AccumulatorConsistencyProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type AccumulatorRangeProof struct {
	// The siblings on the left of the path from root to the first leaf. The ones
	// near the leaf are at the beginning of the list. The placeholder nodes are
	// represented by empty byte arrays, other nodes should be exactly 32-bytes
	// long.
	LeftSiblings [][]byte
	// The siblings on the right of the path from root to the last leaf. The ones
	// near the leaf are at the beginning of the list. The placeholder nodes are
	// represented by empty byte arrays, other nodes should be exactly 32-bytes
	// long.
	RightSiblings [][]byte
}

// GetLeftSiblings gets the LeftSiblings of the AccumulatorRangeProof.
func (m *AccumulatorRangeProof) GetLeftSiblings() (x [][]byte) {
	if m == nil {
		return x
	}
	return m.LeftSiblings
}

// GetRightSiblings gets the RightSiblings of the AccumulatorRangeProof.
func (m *AccumulatorRangeProof) GetRightSiblings() (x [][]byte) {
	if m == nil {
		return x
	}
	return m.RightSiblings
}

// MarshalToWriter marshals AccumulatorRangeProof to the provided writer.
func (m *AccumulatorRangeProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, val := range m.LeftSiblings {
		writer.WriteBytes(1, val)
	}

	for _, val := range m.RightSiblings {
		writer.WriteBytes(2, val)
	}

	return
}

// Marshal marshals AccumulatorRangeProof to a slice of bytes.
func (m *AccumulatorRangeProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccumulatorRangeProof from the provided reader.
func (m *AccumulatorRangeProof) UnmarshalFromReader(reader jspb.Reader) *AccumulatorRangeProof {
	for reader.Next() {
		if m == nil {
			m = &AccumulatorRangeProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.LeftSiblings = append(m.LeftSiblings, reader.ReadBytes())
		case 2:
			m.RightSiblings = append(m.RightSiblings, reader.ReadBytes())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccumulatorRangeProof from a slice of bytes.
func (m *AccumulatorRangeProof) Unmarshal(rawBytes []byte) (*AccumulatorRangeProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type SparseMerkleRangeProof struct {
	// The siblings on the right of the path from root to the last leaf. The ones
	// near the leaf are at the beginning of the list. The placeholder nodes are
	// represented by empty byte arrays, other nodes should be exactly 32-bytes
	// long.
	RightSiblings [][]byte
}

// GetRightSiblings gets the RightSiblings of the SparseMerkleRangeProof.
func (m *SparseMerkleRangeProof) GetRightSiblings() (x [][]byte) {
	if m == nil {
		return x
	}
	return m.RightSiblings
}

// MarshalToWriter marshals SparseMerkleRangeProof to the provided writer.
func (m *SparseMerkleRangeProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, val := range m.RightSiblings {
		writer.WriteBytes(1, val)
	}

	return
}

// Marshal marshals SparseMerkleRangeProof to a slice of bytes.
func (m *SparseMerkleRangeProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SparseMerkleRangeProof from the provided reader.
func (m *SparseMerkleRangeProof) UnmarshalFromReader(reader jspb.Reader) *SparseMerkleRangeProof {
	for reader.Next() {
		if m == nil {
			m = &SparseMerkleRangeProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.RightSiblings = append(m.RightSiblings, reader.ReadBytes())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SparseMerkleRangeProof from a slice of bytes.
func (m *SparseMerkleRangeProof) Unmarshal(rawBytes []byte) (*SparseMerkleRangeProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// The complete proof used to authenticate a transaction.
type TransactionProof struct {
	LedgerInfoToTransactionInfoProof *AccumulatorProof
	TransactionInfo                  *TransactionInfo
}

// GetLedgerInfoToTransactionInfoProof gets the LedgerInfoToTransactionInfoProof of the TransactionProof.
func (m *TransactionProof) GetLedgerInfoToTransactionInfoProof() (x *AccumulatorProof) {
	if m == nil {
		return x
	}
	return m.LedgerInfoToTransactionInfoProof
}

// GetTransactionInfo gets the TransactionInfo of the TransactionProof.
func (m *TransactionProof) GetTransactionInfo() (x *TransactionInfo) {
	if m == nil {
		return x
	}
	return m.TransactionInfo
}

// MarshalToWriter marshals TransactionProof to the provided writer.
func (m *TransactionProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.LedgerInfoToTransactionInfoProof != nil {
		writer.WriteMessage(1, func() {
			m.LedgerInfoToTransactionInfoProof.MarshalToWriter(writer)
		})
	}

	if m.TransactionInfo != nil {
		writer.WriteMessage(2, func() {
			m.TransactionInfo.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals TransactionProof to a slice of bytes.
func (m *TransactionProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a TransactionProof from the provided reader.
func (m *TransactionProof) UnmarshalFromReader(reader jspb.Reader) *TransactionProof {
	for reader.Next() {
		if m == nil {
			m = &TransactionProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.LedgerInfoToTransactionInfoProof = m.LedgerInfoToTransactionInfoProof.UnmarshalFromReader(reader)
			})
		case 2:
			reader.ReadMessage(func() {
				m.TransactionInfo = m.TransactionInfo.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a TransactionProof from a slice of bytes.
func (m *TransactionProof) Unmarshal(rawBytes []byte) (*TransactionProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// The complete proof used to authenticate an account state.
type AccountStateProof struct {
	LedgerInfoToTransactionInfoProof *AccumulatorProof
	TransactionInfo                  *TransactionInfo
	TransactionInfoToAccountProof    *SparseMerkleProof
}

// GetLedgerInfoToTransactionInfoProof gets the LedgerInfoToTransactionInfoProof of the AccountStateProof.
func (m *AccountStateProof) GetLedgerInfoToTransactionInfoProof() (x *AccumulatorProof) {
	if m == nil {
		return x
	}
	return m.LedgerInfoToTransactionInfoProof
}

// GetTransactionInfo gets the TransactionInfo of the AccountStateProof.
func (m *AccountStateProof) GetTransactionInfo() (x *TransactionInfo) {
	if m == nil {
		return x
	}
	return m.TransactionInfo
}

// GetTransactionInfoToAccountProof gets the TransactionInfoToAccountProof of the AccountStateProof.
func (m *AccountStateProof) GetTransactionInfoToAccountProof() (x *SparseMerkleProof) {
	if m == nil {
		return x
	}
	return m.TransactionInfoToAccountProof
}

// MarshalToWriter marshals AccountStateProof to the provided writer.
func (m *AccountStateProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.LedgerInfoToTransactionInfoProof != nil {
		writer.WriteMessage(1, func() {
			m.LedgerInfoToTransactionInfoProof.MarshalToWriter(writer)
		})
	}

	if m.TransactionInfo != nil {
		writer.WriteMessage(2, func() {
			m.TransactionInfo.MarshalToWriter(writer)
		})
	}

	if m.TransactionInfoToAccountProof != nil {
		writer.WriteMessage(3, func() {
			m.TransactionInfoToAccountProof.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals AccountStateProof to a slice of bytes.
func (m *AccountStateProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccountStateProof from the provided reader.
func (m *AccountStateProof) UnmarshalFromReader(reader jspb.Reader) *AccountStateProof {
	for reader.Next() {
		if m == nil {
			m = &AccountStateProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.LedgerInfoToTransactionInfoProof = m.LedgerInfoToTransactionInfoProof.UnmarshalFromReader(reader)
			})
		case 2:
			reader.ReadMessage(func() {
				m.TransactionInfo = m.TransactionInfo.UnmarshalFromReader(reader)
			})
		case 3:
			reader.ReadMessage(func() {
				m.TransactionInfoToAccountProof = m.TransactionInfoToAccountProof.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccountStateProof from a slice of bytes.
func (m *AccountStateProof) Unmarshal(rawBytes []byte) (*AccountStateProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// The complete proof used to authenticate an event.
type EventProof struct {
	LedgerInfoToTransactionInfoProof *AccumulatorProof
	TransactionInfo                  *TransactionInfo
	TransactionInfoToEventProof      *AccumulatorProof
}

// GetLedgerInfoToTransactionInfoProof gets the LedgerInfoToTransactionInfoProof of the EventProof.
func (m *EventProof) GetLedgerInfoToTransactionInfoProof() (x *AccumulatorProof) {
	if m == nil {
		return x
	}
	return m.LedgerInfoToTransactionInfoProof
}

// GetTransactionInfo gets the TransactionInfo of the EventProof.
func (m *EventProof) GetTransactionInfo() (x *TransactionInfo) {
	if m == nil {
		return x
	}
	return m.TransactionInfo
}

// GetTransactionInfoToEventProof gets the TransactionInfoToEventProof of the EventProof.
func (m *EventProof) GetTransactionInfoToEventProof() (x *AccumulatorProof) {
	if m == nil {
		return x
	}
	return m.TransactionInfoToEventProof
}

// MarshalToWriter marshals EventProof to the provided writer.
func (m *EventProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.LedgerInfoToTransactionInfoProof != nil {
		writer.WriteMessage(1, func() {
			m.LedgerInfoToTransactionInfoProof.MarshalToWriter(writer)
		})
	}

	if m.TransactionInfo != nil {
		writer.WriteMessage(2, func() {
			m.TransactionInfo.MarshalToWriter(writer)
		})
	}

	if m.TransactionInfoToEventProof != nil {
		writer.WriteMessage(3, func() {
			m.TransactionInfoToEventProof.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals EventProof to a slice of bytes.
func (m *EventProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a EventProof from the provided reader.
func (m *EventProof) UnmarshalFromReader(reader jspb.Reader) *EventProof {
	for reader.Next() {
		if m == nil {
			m = &EventProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.LedgerInfoToTransactionInfoProof = m.LedgerInfoToTransactionInfoProof.UnmarshalFromReader(reader)
			})
		case 2:
			reader.ReadMessage(func() {
				m.TransactionInfo = m.TransactionInfo.UnmarshalFromReader(reader)
			})
		case 3:
			reader.ReadMessage(func() {
				m.TransactionInfoToEventProof = m.TransactionInfoToEventProof.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a EventProof from a slice of bytes.
func (m *EventProof) Unmarshal(rawBytes []byte) (*EventProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// The complete proof used to authenticate a list of transactions.
type TransactionListProof struct {
	LedgerInfoToTransactionInfosProof *AccumulatorRangeProof
	TransactionInfos                  []*TransactionInfo
}

// GetLedgerInfoToTransactionInfosProof gets the LedgerInfoToTransactionInfosProof of the TransactionListProof.
func (m *TransactionListProof) GetLedgerInfoToTransactionInfosProof() (x *AccumulatorRangeProof) {
	if m == nil {
		return x
	}
	return m.LedgerInfoToTransactionInfosProof
}

// GetTransactionInfos gets the TransactionInfos of the TransactionListProof.
func (m *TransactionListProof) GetTransactionInfos() (x []*TransactionInfo) {
	if m == nil {
		return x
	}
	return m.TransactionInfos
}

// MarshalToWriter marshals TransactionListProof to the provided writer.
func (m *TransactionListProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.LedgerInfoToTransactionInfosProof != nil {
		writer.WriteMessage(1, func() {
			m.LedgerInfoToTransactionInfosProof.MarshalToWriter(writer)
		})
	}

	for _, msg := range m.TransactionInfos {
		writer.WriteMessage(2, func() {
			msg.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals TransactionListProof to a slice of bytes.
func (m *TransactionListProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a TransactionListProof from the provided reader.
func (m *TransactionListProof) UnmarshalFromReader(reader jspb.Reader) *TransactionListProof {
	for reader.Next() {
		if m == nil {
			m = &TransactionListProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.LedgerInfoToTransactionInfosProof = m.LedgerInfoToTransactionInfosProof.UnmarshalFromReader(reader)
			})
		case 2:
			reader.ReadMessage(func() {
				m.TransactionInfos = append(m.TransactionInfos, new(TransactionInfo).UnmarshalFromReader(reader))
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a TransactionListProof from a slice of bytes.
func (m *TransactionListProof) Unmarshal(rawBytes []byte) (*TransactionListProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
