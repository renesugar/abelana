// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/data_proto/bigtable_data.proto
// DO NOT EDIT!

/*
Package google_bigtable_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/data_proto/bigtable_data.proto

It has these top-level messages:
	Row
	Family
	Column
	Cell
	RowRange
	ColumnRange
	TimestampRange
	ValueRange
	RowFilter
	Mutation
	ReadModifyWriteRule
*/
package google_bigtable_v1

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Specifies the complete (requested) contents of a single row of a table.
// Rows which exceed 256MiB in size cannot be read in full.
type Row struct {
	// The unique key which identifies this row within its table. This is the same
	// key that's used to identify the row in, for example, a MutateRowRequest.
	// May contain any non-empty byte string up to 16KiB in length.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// May be empty, but only if the entire row is empty.
	// The mutual ordering of column families is not specified.
	Families []*Family `protobuf:"bytes,2,rep,name=families" json:"families,omitempty"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}

func (m *Row) GetFamilies() []*Family {
	if m != nil {
		return m.Families
	}
	return nil
}

// Specifies (some of) the contents of a single row/column family of a table.
type Family struct {
	// The unique key which identifies this family within its row. This is the
	// same key that's used to identify the family in, for example, a RowFilter
	// which sets its "family_name_regex_filter" field.
	// Must match [-_.a-zA-Z0-9]+, except that AggregatingRowProcessors may
	// produce cells in a sentinel family with an empty name.
	// Must be no greater than 64 characters in length.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Must not be empty. Sorted in order of increasing "qualifier".
	Columns []*Column `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
}

func (m *Family) Reset()         { *m = Family{} }
func (m *Family) String() string { return proto.CompactTextString(m) }
func (*Family) ProtoMessage()    {}

func (m *Family) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Specifies (some of) the contents of a single row/column of a table.
type Column struct {
	// The unique key which identifies this column within its family. This is the
	// same key that's used to identify the column in, for example, a RowFilter
	// which sets its "column_qualifier_regex_filter" field.
	// May contain any byte string, including the empty string, up to 16kiB in
	// length.
	Qualifier []byte `protobuf:"bytes,1,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Must not be empty. Sorted in order of decreasing "timestamp_micros".
	Cells []*Cell `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}

func (m *Column) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

// Specifies (some of) the contents of a single row/column/timestamp of a table.
type Cell struct {
	// The cell's stored timestamp, which also uniquely identifies it within
	// its column.
	// Values are always expressed in microseconds, but individual tables may set
	// a coarser "granularity" to further restrict the allowed values. For
	// example, a table which specifies millisecond granularity will only allow
	// values of "timestamp_micros" which are multiples of 1000.
	TimestampMicros int64 `protobuf:"varint,1,opt,name=timestamp_micros" json:"timestamp_micros,omitempty"`
	// The value stored in the cell.
	// May contain any byte string, including the empty string, up to 100MiB in
	// length.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}

// Specifies a contiguous range of rows.
type RowRange struct {
	// Inclusive lower bound. If left empty, interpreted as the empty string.
	StartKey []byte `protobuf:"bytes,2,opt,name=start_key,proto3" json:"start_key,omitempty"`
	// Exclusive upper bound. If left empty, interpreted as infinity.
	EndKey []byte `protobuf:"bytes,3,opt,name=end_key,proto3" json:"end_key,omitempty"`
}

func (m *RowRange) Reset()         { *m = RowRange{} }
func (m *RowRange) String() string { return proto.CompactTextString(m) }
func (*RowRange) ProtoMessage()    {}

// Specifies a contiguous range of columns within a single column family.
// The range spans from <column_family>:<start_qualifier> to
// <column_family>:<end_qualifier>, where both bounds can be either inclusive or
// exclusive.
type ColumnRange struct {
	// The name of the column family within which this range falls.
	FamilyName string `protobuf:"bytes,1,opt,name=family_name" json:"family_name,omitempty"`
	// Used when giving an inclusive lower bound for the range.
	StartQualifierInclusive []byte `protobuf:"bytes,2,opt,name=start_qualifier_inclusive,proto3" json:"start_qualifier_inclusive,omitempty"`
	// Used when giving an exclusive lower bound for the range.
	StartQualifierExclusive []byte `protobuf:"bytes,3,opt,name=start_qualifier_exclusive,proto3" json:"start_qualifier_exclusive,omitempty"`
	// Used when giving an inclusive upper bound for the range.
	EndQualifierInclusive []byte `protobuf:"bytes,4,opt,name=end_qualifier_inclusive,proto3" json:"end_qualifier_inclusive,omitempty"`
	// Used when giving an exclusive upper bound for the range.
	EndQualifierExclusive []byte `protobuf:"bytes,5,opt,name=end_qualifier_exclusive,proto3" json:"end_qualifier_exclusive,omitempty"`
}

func (m *ColumnRange) Reset()         { *m = ColumnRange{} }
func (m *ColumnRange) String() string { return proto.CompactTextString(m) }
func (*ColumnRange) ProtoMessage()    {}

// Specified a contiguous range of microsecond timestamps.
type TimestampRange struct {
	// Inclusive lower bound. If left empty, interpreted as 0.
	StartTimestampMicros int64 `protobuf:"varint,1,opt,name=start_timestamp_micros" json:"start_timestamp_micros,omitempty"`
	// Exclusive upper bound. If left empty, interpreted as infinity.
	EndTimestampMicros int64 `protobuf:"varint,2,opt,name=end_timestamp_micros" json:"end_timestamp_micros,omitempty"`
}

func (m *TimestampRange) Reset()         { *m = TimestampRange{} }
func (m *TimestampRange) String() string { return proto.CompactTextString(m) }
func (*TimestampRange) ProtoMessage()    {}

// Specifies a contiguous range of raw byte values.
type ValueRange struct {
	// Used when giving an inclusive lower bound for the range.
	StartValueInclusive []byte `protobuf:"bytes,1,opt,name=start_value_inclusive,proto3" json:"start_value_inclusive,omitempty"`
	// Used when giving an exclusive lower bound for the range.
	StartValueExclusive []byte `protobuf:"bytes,2,opt,name=start_value_exclusive,proto3" json:"start_value_exclusive,omitempty"`
	// Used when giving an inclusive upper bound for the range.
	EndValueInclusive []byte `protobuf:"bytes,3,opt,name=end_value_inclusive,proto3" json:"end_value_inclusive,omitempty"`
	// Used when giving an exclusive upper bound for the range.
	EndValueExclusive []byte `protobuf:"bytes,4,opt,name=end_value_exclusive,proto3" json:"end_value_exclusive,omitempty"`
}

func (m *ValueRange) Reset()         { *m = ValueRange{} }
func (m *ValueRange) String() string { return proto.CompactTextString(m) }
func (*ValueRange) ProtoMessage()    {}

// Takes a row as input and produces an alternate view of the row based on
// specified rules. For example, a RowFilter might trim down a row to include
// just the cells from columns matching a given regular expression, or might
// return all the cells of a row but not their values. More complicated filters
// can be composed out of these components to express requests such as, "within
// every column of a particular family, give just the two most recent cells
// which are older than timestamp X."
//
// There are two broad categories of RowFilters (true filters and transformers),
// as well as two ways to compose simple filters into more complex ones
// (chains and interleaves). They work as follows:
//
// * True filters alter the input row by excluding some of its cells wholesale
// from the output row. An example of a true filter is the "value_regex_filter",
// which excludes cells whose values don't match the specified pattern. All
// regex true filters use RE2 syntax (https://github.com/google/re2/wiki/Syntax)
// in raw byte mode (RE2::Latin1), and are evaluated as full matches. An
// important point to keep in mind is that RE2(.) is equivalent by default to
// RE2([^\n]), meaning that it does not match newlines. When attempting to match
// an arbitrary byte, you should therefore use the escape sequence '\C', which
// may need to be further escaped as '\\C' in your client language.
//
// * Transformers alter the input row by changing the values of some of its
// cells in the output, without excluding them completely. Currently, the only
// supported transformer is the "strip_value_transformer", which replaces every
// cell's value with the empty string.
//
// * Chains and interleaves are described in more detail in the
// RowFilter.Chain and RowFilter.Interleave documentation.
//
// The total serialized size of a RowFilter message must not
// exceed 4096 bytes, and RowFilters may not be nested within each other
// (in Chains or Interleaves) to a depth of more than 20.
type RowFilter struct {
	// Applies several RowFilters to the data in sequence, progressively
	// narrowing the results.
	Chain *RowFilter_Chain `protobuf:"bytes,1,opt,name=chain" json:"chain,omitempty"`
	// Applies several RowFilters to the data in parallel and combines the
	// results.
	Interleave *RowFilter_Interleave `protobuf:"bytes,2,opt,name=interleave" json:"interleave,omitempty"`
	// Applies one of two possible RowFilters to the data based on the output of
	// a predicate RowFilter.
	Condition *RowFilter_Condition `protobuf:"bytes,3,opt,name=condition" json:"condition,omitempty"`
	// Matches only cells from rows whose keys satisfy the given RE2 regex. In
	// other words, passes through the entire row when the key matches, and
	// otherwise produces an empty row.
	// Note that, since row keys can contain arbitrary bytes, the '\C' escape
	// sequence must be used if a true wildcard is desired. The '.' character
	// will not match the new line character '\n', which may be present in a
	// binary key.
	RowKeyRegexFilter []byte `protobuf:"bytes,4,opt,name=row_key_regex_filter,proto3" json:"row_key_regex_filter,omitempty"`
	// Matches all cells from a row with probability p, and matches no cells
	// from the row with probability 1-p.
	RowSampleFilter float64 `protobuf:"fixed64,14,opt,name=row_sample_filter" json:"row_sample_filter,omitempty"`
	// Matches only cells from columns whose families satisfy the given RE2
	// regex. For technical reasons, the regex must not contain the ':'
	// character, even if it is not being used as a literal.
	// Note that, since column families cannot contain the new line character
	// '\n', it is sufficient to use '.' as a full wildcard when matching
	// column family names.
	FamilyNameRegexFilter string `protobuf:"bytes,5,opt,name=family_name_regex_filter" json:"family_name_regex_filter,omitempty"`
	// Matches only cells from columns whose qualifiers satisfy the given RE2
	// regex.
	// Note that, since column qualifiers can contain arbitrary bytes, the '\C'
	// escape sequence must be used if a true wildcard is desired. The '.'
	// character will not match the new line character '\n', which may be
	// present in a binary qualifier.
	ColumnQualifierRegexFilter []byte `protobuf:"bytes,6,opt,name=column_qualifier_regex_filter,proto3" json:"column_qualifier_regex_filter,omitempty"`
	// Matches only cells from columns within the given range.
	ColumnRangeFilter *ColumnRange `protobuf:"bytes,7,opt,name=column_range_filter" json:"column_range_filter,omitempty"`
	// Matches only cells with timestamps within the given range.
	TimestampRangeFilter *TimestampRange `protobuf:"bytes,8,opt,name=timestamp_range_filter" json:"timestamp_range_filter,omitempty"`
	// Matches only cells with values that satisfy the given regular expression.
	// Note that, since cell values can contain arbitrary bytes, the '\C' escape
	// sequence must be used if a true wildcard is desired. The '.' character
	// will not match the new line character '\n', which may be present in a
	// binary value.
	ValueRegexFilter []byte `protobuf:"bytes,9,opt,name=value_regex_filter,proto3" json:"value_regex_filter,omitempty"`
	// Matches only cells with values that fall within the given range.
	ValueRangeFilter *ValueRange `protobuf:"bytes,15,opt,name=value_range_filter" json:"value_range_filter,omitempty"`
	// Skips the first N cells of each row, matching all subsequent cells.
	CellsPerRowOffsetFilter int32 `protobuf:"varint,10,opt,name=cells_per_row_offset_filter" json:"cells_per_row_offset_filter,omitempty"`
	// Matches only the first N cells of each row.
	CellsPerRowLimitFilter int32 `protobuf:"varint,11,opt,name=cells_per_row_limit_filter" json:"cells_per_row_limit_filter,omitempty"`
	// Matches only the most recent N cells within each column. For example,
	// if N=2, this filter would match column "foo:bar" at timestamps 10 and 9,
	// skip all earlier cells in "foo:bar", and then begin matching again in
	// column "foo:bar2".
	CellsPerColumnLimitFilter int32 `protobuf:"varint,12,opt,name=cells_per_column_limit_filter" json:"cells_per_column_limit_filter,omitempty"`
	// Replaces each cell's value with the empty string.
	StripValueTransformer bool `protobuf:"varint,13,opt,name=strip_value_transformer" json:"strip_value_transformer,omitempty"`
}

func (m *RowFilter) Reset()         { *m = RowFilter{} }
func (m *RowFilter) String() string { return proto.CompactTextString(m) }
func (*RowFilter) ProtoMessage()    {}

func (m *RowFilter) GetChain() *RowFilter_Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *RowFilter) GetInterleave() *RowFilter_Interleave {
	if m != nil {
		return m.Interleave
	}
	return nil
}

func (m *RowFilter) GetCondition() *RowFilter_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *RowFilter) GetColumnRangeFilter() *ColumnRange {
	if m != nil {
		return m.ColumnRangeFilter
	}
	return nil
}

func (m *RowFilter) GetTimestampRangeFilter() *TimestampRange {
	if m != nil {
		return m.TimestampRangeFilter
	}
	return nil
}

func (m *RowFilter) GetValueRangeFilter() *ValueRange {
	if m != nil {
		return m.ValueRangeFilter
	}
	return nil
}

// A RowFilter which sends rows through several RowFilters in sequence.
type RowFilter_Chain struct {
	// The elements of "filters" are chained together to process the input row:
	// in row -> f(0) -> intermediate row -> f(1) -> ... -> f(N) -> out row
	// The full chain is executed atomically.
	Filters []*RowFilter `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (m *RowFilter_Chain) Reset()         { *m = RowFilter_Chain{} }
func (m *RowFilter_Chain) String() string { return proto.CompactTextString(m) }
func (*RowFilter_Chain) ProtoMessage()    {}

func (m *RowFilter_Chain) GetFilters() []*RowFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// A RowFilter which sends each row to each of several component
// RowFilters and interleaves the results.
type RowFilter_Interleave struct {
	// The elements of "filters" all process a copy of the input row, and the
	// results are pooled, sorted, and combined into a single output row.
	// If multiple cells are produced with the same column and timestamp,
	// they will all appear in the output row in an unspecified mutual order.
	// Consider the following example, with three filters:
	//
	//                              input row
	//                                  |
	//        -----------------------------------------------------
	//        |                         |                         |
	//       f(0)                      f(1)                      f(2)
	//        |                         |                         |
	// 1: foo,bar,10,x             foo,bar,10,z              far,bar,7,a
	// 2: foo,blah,11,z            far,blah,5,x              far,blah,5,x
	//        |                         |                         |
	//        -----------------------------------------------------
	//                                  |
	// 1:                        foo,bar,10,z     // could have switched with #2
	// 2:                        foo,bar,10,x     // could have switched with #1
	// 3:                        foo,blah,11,z
	// 4:                        far,bar,7,a
	// 5:                        far,blah,5,x     // identical to #6
	// 6:                        far,blah,5,x     // identical to #5
	// All interleaved filters are executed atomically.
	Filters []*RowFilter `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (m *RowFilter_Interleave) Reset()         { *m = RowFilter_Interleave{} }
func (m *RowFilter_Interleave) String() string { return proto.CompactTextString(m) }
func (*RowFilter_Interleave) ProtoMessage()    {}

func (m *RowFilter_Interleave) GetFilters() []*RowFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// A RowFilter which evaluates one of two possible RowFilters, depending on
// whether or not a predicate RowFilter outputs any cells from the input row.
//
// IMPORTANT NOTE: The predicate filter does not execute atomically with the
// true and false filters, which may lead to inconsistent or unexpected
// results. Additionally, Condition filters have poor performance, especially
// when filters are set for the false condition.
type RowFilter_Condition struct {
	// If "predicate_filter" outputs any cells, then "true_filter" will be
	// evaluated on the input row. Otherwise, "false_filter" will be evaluated.
	PredicateFilter *RowFilter `protobuf:"bytes,1,opt,name=predicate_filter" json:"predicate_filter,omitempty"`
	// The filter to apply to the input row if "predicate_filter" returns any
	// results. If not provided, no results will be returned in the true case.
	TrueFilter *RowFilter `protobuf:"bytes,2,opt,name=true_filter" json:"true_filter,omitempty"`
	// The filter to apply to the input row if "predicate_filter" does not
	// return any results. If not provided, no results will be returned in the
	// false case.
	FalseFilter *RowFilter `protobuf:"bytes,3,opt,name=false_filter" json:"false_filter,omitempty"`
}

func (m *RowFilter_Condition) Reset()         { *m = RowFilter_Condition{} }
func (m *RowFilter_Condition) String() string { return proto.CompactTextString(m) }
func (*RowFilter_Condition) ProtoMessage()    {}

func (m *RowFilter_Condition) GetPredicateFilter() *RowFilter {
	if m != nil {
		return m.PredicateFilter
	}
	return nil
}

func (m *RowFilter_Condition) GetTrueFilter() *RowFilter {
	if m != nil {
		return m.TrueFilter
	}
	return nil
}

func (m *RowFilter_Condition) GetFalseFilter() *RowFilter {
	if m != nil {
		return m.FalseFilter
	}
	return nil
}

// Specifies a particular change to be made to the contents of a row.
type Mutation struct {
	// Set a cell's value.
	SetCell *Mutation_SetCell `protobuf:"bytes,1,opt,name=set_cell" json:"set_cell,omitempty"`
	// Deletes cells from a column.
	DeleteFromColumn *Mutation_DeleteFromColumn `protobuf:"bytes,2,opt,name=delete_from_column" json:"delete_from_column,omitempty"`
	// Deletes cells from a column family.
	DeleteFromFamily *Mutation_DeleteFromFamily `protobuf:"bytes,3,opt,name=delete_from_family" json:"delete_from_family,omitempty"`
	// Deletes cells from the entire row.
	DeleteFromRow *Mutation_DeleteFromRow `protobuf:"bytes,4,opt,name=delete_from_row" json:"delete_from_row,omitempty"`
}

func (m *Mutation) Reset()         { *m = Mutation{} }
func (m *Mutation) String() string { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()    {}

func (m *Mutation) GetSetCell() *Mutation_SetCell {
	if m != nil {
		return m.SetCell
	}
	return nil
}

func (m *Mutation) GetDeleteFromColumn() *Mutation_DeleteFromColumn {
	if m != nil {
		return m.DeleteFromColumn
	}
	return nil
}

func (m *Mutation) GetDeleteFromFamily() *Mutation_DeleteFromFamily {
	if m != nil {
		return m.DeleteFromFamily
	}
	return nil
}

func (m *Mutation) GetDeleteFromRow() *Mutation_DeleteFromRow {
	if m != nil {
		return m.DeleteFromRow
	}
	return nil
}

// A Mutation which sets the value of the specified cell.
type Mutation_SetCell struct {
	// The name of the family into which new data should be written.
	// Must match [-_.a-zA-Z0-9]+
	FamilyName string `protobuf:"bytes,1,opt,name=family_name" json:"family_name,omitempty"`
	// The qualifier of the column into which new data should be written.
	// Can be any byte string, including the empty string.
	ColumnQualifier []byte `protobuf:"bytes,2,opt,name=column_qualifier,proto3" json:"column_qualifier,omitempty"`
	// The timestamp of the cell into which new data should be written.
	// Use -1 for current Bigtable server time.
	// Otherwise, the client should set this value itself, noting that the
	// default value is a timestamp of zero if the field is left unspecified.
	// Values must match the "granularity" of the table (e.g. micros, millis).
	TimestampMicros int64 `protobuf:"varint,3,opt,name=timestamp_micros" json:"timestamp_micros,omitempty"`
	// The value to be written into the specified cell.
	Value []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Mutation_SetCell) Reset()         { *m = Mutation_SetCell{} }
func (m *Mutation_SetCell) String() string { return proto.CompactTextString(m) }
func (*Mutation_SetCell) ProtoMessage()    {}

// A Mutation which deletes cells from the specified column, optionally
// restricting the deletions to a given timestamp range.
type Mutation_DeleteFromColumn struct {
	// The name of the family from which cells should be deleted.
	// Must match [-_.a-zA-Z0-9]+
	FamilyName string `protobuf:"bytes,1,opt,name=family_name" json:"family_name,omitempty"`
	// The qualifier of the column from which cells should be deleted.
	// Can be any byte string, including the empty string.
	ColumnQualifier []byte `protobuf:"bytes,2,opt,name=column_qualifier,proto3" json:"column_qualifier,omitempty"`
	// The range of timestamps within which cells should be deleted.
	TimeRange *TimestampRange `protobuf:"bytes,3,opt,name=time_range" json:"time_range,omitempty"`
}

func (m *Mutation_DeleteFromColumn) Reset()         { *m = Mutation_DeleteFromColumn{} }
func (m *Mutation_DeleteFromColumn) String() string { return proto.CompactTextString(m) }
func (*Mutation_DeleteFromColumn) ProtoMessage()    {}

func (m *Mutation_DeleteFromColumn) GetTimeRange() *TimestampRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

// A Mutation which deletes all cells from the specified column family.
type Mutation_DeleteFromFamily struct {
	// The name of the family from which cells should be deleted.
	// Must match [-_.a-zA-Z0-9]+
	FamilyName string `protobuf:"bytes,1,opt,name=family_name" json:"family_name,omitempty"`
}

func (m *Mutation_DeleteFromFamily) Reset()         { *m = Mutation_DeleteFromFamily{} }
func (m *Mutation_DeleteFromFamily) String() string { return proto.CompactTextString(m) }
func (*Mutation_DeleteFromFamily) ProtoMessage()    {}

// A Mutation which deletes all cells from the containing row.
type Mutation_DeleteFromRow struct {
}

func (m *Mutation_DeleteFromRow) Reset()         { *m = Mutation_DeleteFromRow{} }
func (m *Mutation_DeleteFromRow) String() string { return proto.CompactTextString(m) }
func (*Mutation_DeleteFromRow) ProtoMessage()    {}

// Specifies an atomic read/modify/write operation on the latest value of the
// specified column.
type ReadModifyWriteRule struct {
	// The name of the family to which the read/modify/write should be applied.
	// Must match [-_.a-zA-Z0-9]+
	FamilyName string `protobuf:"bytes,1,opt,name=family_name" json:"family_name,omitempty"`
	// The qualifier of the column to which the read/modify/write should be
	// applied.
	// Can be any byte string, including the empty string.
	ColumnQualifier []byte `protobuf:"bytes,2,opt,name=column_qualifier,proto3" json:"column_qualifier,omitempty"`
	// Rule specifying that "append_value" be appended to the existing value.
	// If the targeted cell is unset, it will be treated as containing the
	// empty string.
	AppendValue []byte `protobuf:"bytes,3,opt,name=append_value,proto3" json:"append_value,omitempty"`
	// Rule specifying that "increment_amount" be added to the existing value.
	// If the targeted cell is unset, it will be treated as containing a zero.
	// Otherwise, the targeted cell must contain an 8-byte value (interpreted
	// as a 64-bit big-endian signed integer), or the entire request will fail.
	IncrementAmount int64 `protobuf:"varint,4,opt,name=increment_amount" json:"increment_amount,omitempty"`
}

func (m *ReadModifyWriteRule) Reset()         { *m = ReadModifyWriteRule{} }
func (m *ReadModifyWriteRule) String() string { return proto.CompactTextString(m) }
func (*ReadModifyWriteRule) ProtoMessage()    {}

func init() {
}