// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/fluent/role"
	"github.com/usalko/fluent/flc/integration/fluent/schema"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/field_type"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	mutation *FieldTypeMutation
	hooks    []Hook
}

// SetInt sets the "int" field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.mutation.SetInt(i)
	return ftc
}

// SetInt8 sets the "int8" field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetInt8(i)
	return ftc
}

// SetInt16 sets the "int16" field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetInt16(i)
	return ftc
}

// SetInt32 sets the "int32" field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetInt32(i)
	return ftc
}

// SetInt64 sets the "int64" field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetInt64(i)
	return ftc
}

// SetOptionalInt sets the "optional_int" field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt(i)
	return ftc
}

// SetNillableOptionalInt sets the "optional_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the "optional_int8" field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt8(i)
	return ftc
}

// SetNillableOptionalInt8 sets the "optional_int8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the "optional_int16" field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt16(i)
	return ftc
}

// SetNillableOptionalInt16 sets the "optional_int16" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the "optional_int32" field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt32(i)
	return ftc
}

// SetNillableOptionalInt32 sets the "optional_int32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the "optional_int64" field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt64(i)
	return ftc
}

// SetNillableOptionalInt64 sets the "optional_int64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNillableInt sets the "nillable_int" field.
func (ftc *FieldTypeCreate) SetNillableInt(i int) *FieldTypeCreate {
	ftc.mutation.SetNillableInt(i)
	return ftc
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt(*i)
	}
	return ftc
}

// SetNillableInt8 sets the "nillable_int8" field.
func (ftc *FieldTypeCreate) SetNillableInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetNillableInt8(i)
	return ftc
}

// SetNillableNillableInt8 sets the "nillable_int8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt8(*i)
	}
	return ftc
}

// SetNillableInt16 sets the "nillable_int16" field.
func (ftc *FieldTypeCreate) SetNillableInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetNillableInt16(i)
	return ftc
}

// SetNillableNillableInt16 sets the "nillable_int16" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt16(*i)
	}
	return ftc
}

// SetNillableInt32 sets the "nillable_int32" field.
func (ftc *FieldTypeCreate) SetNillableInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetNillableInt32(i)
	return ftc
}

// SetNillableNillableInt32 sets the "nillable_int32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt32(*i)
	}
	return ftc
}

// SetNillableInt64 sets the "nillable_int64" field.
func (ftc *FieldTypeCreate) SetNillableInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetNillableInt64(i)
	return ftc
}

// SetNillableNillableInt64 sets the "nillable_int64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt64(*i)
	}
	return ftc
}

// SetValidateOptionalInt32 sets the "validate_optional_int32" field.
func (ftc *FieldTypeCreate) SetValidateOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetValidateOptionalInt32(i)
	return ftc
}

// SetNillableValidateOptionalInt32 sets the "validate_optional_int32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetValidateOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalUint sets the "optional_uint" field.
func (ftc *FieldTypeCreate) SetOptionalUint(u uint) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint(u)
	return ftc
}

// SetNillableOptionalUint sets the "optional_uint" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint(u *uint) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint(*u)
	}
	return ftc
}

// SetOptionalUint8 sets the "optional_uint8" field.
func (ftc *FieldTypeCreate) SetOptionalUint8(u uint8) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint8(u)
	return ftc
}

// SetNillableOptionalUint8 sets the "optional_uint8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint8(u *uint8) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint8(*u)
	}
	return ftc
}

// SetOptionalUint16 sets the "optional_uint16" field.
func (ftc *FieldTypeCreate) SetOptionalUint16(u uint16) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint16(u)
	return ftc
}

// SetNillableOptionalUint16 sets the "optional_uint16" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint16(u *uint16) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint16(*u)
	}
	return ftc
}

// SetOptionalUint32 sets the "optional_uint32" field.
func (ftc *FieldTypeCreate) SetOptionalUint32(u uint32) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint32(u)
	return ftc
}

// SetNillableOptionalUint32 sets the "optional_uint32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint32(u *uint32) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint32(*u)
	}
	return ftc
}

// SetOptionalUint64 sets the "optional_uint64" field.
func (ftc *FieldTypeCreate) SetOptionalUint64(u uint64) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint64(u)
	return ftc
}

// SetNillableOptionalUint64 sets the "optional_uint64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint64(u *uint64) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint64(*u)
	}
	return ftc
}

// SetState sets the "state" field.
func (ftc *FieldTypeCreate) SetState(field_type_state field_type.State) *FieldTypeCreate {
	ftc.mutation.SetState(field_type_state)
	return ftc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableState(field_type_state *field_type.State) *FieldTypeCreate {
	if field_type_state != nil {
		ftc.SetState(*field_type_state)
	}
	return ftc
}

// SetOptionalFloat sets the "optional_float" field.
func (ftc *FieldTypeCreate) SetOptionalFloat(f float64) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat(f)
	return ftc
}

// SetNillableOptionalFloat sets the "optional_float" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat(*f)
	}
	return ftc
}

// SetOptionalFloat32 sets the "optional_float32" field.
func (ftc *FieldTypeCreate) SetOptionalFloat32(f float32) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat32(f)
	return ftc
}

// SetNillableOptionalFloat32 sets the "optional_float32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat32(f *float32) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat32(*f)
	}
	return ftc
}

// SetText sets the "text" field.
func (ftc *FieldTypeCreate) SetText(s string) *FieldTypeCreate {
	ftc.mutation.SetText(s)
	return ftc
}

// SetNillableText sets the "text" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableText(s *string) *FieldTypeCreate {
	if s != nil {
		ftc.SetText(*s)
	}
	return ftc
}

// SetDatetime sets the "datetime" field.
func (ftc *FieldTypeCreate) SetDatetime(time_time time.Time) *FieldTypeCreate {
	ftc.mutation.SetDatetime(time_time)
	return ftc
}

// SetNillableDatetime sets the "datetime" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDatetime(time_time *time.Time) *FieldTypeCreate {
	if time_time != nil {
		ftc.SetDatetime(*time_time)
	}
	return ftc
}

// SetDecimal sets the "decimal" field.
func (ftc *FieldTypeCreate) SetDecimal(f float64) *FieldTypeCreate {
	ftc.mutation.SetDecimal(f)
	return ftc
}

// SetNillableDecimal sets the "decimal" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDecimal(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetDecimal(*f)
	}
	return ftc
}

// SetLinkOther sets the "link_other" field.
func (ftc *FieldTypeCreate) SetLinkOther(schema_link *schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLinkOther(schema_link)
	return ftc
}

// SetLinkOtherFunc sets the "link_other_func" field.
func (ftc *FieldTypeCreate) SetLinkOtherFunc(schema_link *schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLinkOtherFunc(schema_link)
	return ftc
}

// SetMAC sets the "mac" field.
func (ftc *FieldTypeCreate) SetMAC(schema_mac schema.MAC) *FieldTypeCreate {
	ftc.mutation.SetMAC(schema_mac)
	return ftc
}

// SetNillableMAC sets the "mac" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableMAC(schema_mac *schema.MAC) *FieldTypeCreate {
	if schema_mac != nil {
		ftc.SetMAC(*schema_mac)
	}
	return ftc
}

// SetStringArray sets the "string_array" field.
func (ftc *FieldTypeCreate) SetStringArray(schema_strings schema.Strings) *FieldTypeCreate {
	ftc.mutation.SetStringArray(schema_strings)
	return ftc
}

// SetPassword sets the "password" field.
func (ftc *FieldTypeCreate) SetPassword(s string) *FieldTypeCreate {
	ftc.mutation.SetPassword(s)
	return ftc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillablePassword(s *string) *FieldTypeCreate {
	if s != nil {
		ftc.SetPassword(*s)
	}
	return ftc
}

// SetStringScanner sets the "string_scanner" field.
func (ftc *FieldTypeCreate) SetStringScanner(schema_stringscanner schema.StringScanner) *FieldTypeCreate {
	ftc.mutation.SetStringScanner(schema_stringscanner)
	return ftc
}

// SetNillableStringScanner sets the "string_scanner" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableStringScanner(schema_stringscanner *schema.StringScanner) *FieldTypeCreate {
	if schema_stringscanner != nil {
		ftc.SetStringScanner(*schema_stringscanner)
	}
	return ftc
}

// SetDuration sets the "duration" field.
func (ftc *FieldTypeCreate) SetDuration(time_duration time.Duration) *FieldTypeCreate {
	ftc.mutation.SetDuration(time_duration)
	return ftc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDuration(time_duration *time.Duration) *FieldTypeCreate {
	if time_duration != nil {
		ftc.SetDuration(*time_duration)
	}
	return ftc
}

// SetDir sets the "dir" field.
func (ftc *FieldTypeCreate) SetDir(http_dir http.Dir) *FieldTypeCreate {
	ftc.mutation.SetDir(http_dir)
	return ftc
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDir(http_dir *http.Dir) *FieldTypeCreate {
	if http_dir != nil {
		ftc.SetDir(*http_dir)
	}
	return ftc
}

// SetNdir sets the "ndir" field.
func (ftc *FieldTypeCreate) SetNdir(http_dir http.Dir) *FieldTypeCreate {
	ftc.mutation.SetNdir(http_dir)
	return ftc
}

// SetNillableNdir sets the "ndir" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNdir(http_dir *http.Dir) *FieldTypeCreate {
	if http_dir != nil {
		ftc.SetNdir(*http_dir)
	}
	return ftc
}

// SetStr sets the "str" field.
func (ftc *FieldTypeCreate) SetStr(sql_nullstring sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetStr(sql_nullstring)
	return ftc
}

// SetNillableStr sets the "str" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableStr(sql_nullstring *sql.NullString) *FieldTypeCreate {
	if sql_nullstring != nil {
		ftc.SetStr(*sql_nullstring)
	}
	return ftc
}

// SetNullStr sets the "null_str" field.
func (ftc *FieldTypeCreate) SetNullStr(sql_nullstring *sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetNullStr(sql_nullstring)
	return ftc
}

// SetLink sets the "link" field.
func (ftc *FieldTypeCreate) SetLink(schema_link schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLink(schema_link)
	return ftc
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableLink(schema_link *schema.Link) *FieldTypeCreate {
	if schema_link != nil {
		ftc.SetLink(*schema_link)
	}
	return ftc
}

// SetNullLink sets the "null_link" field.
func (ftc *FieldTypeCreate) SetNullLink(schema_link *schema.Link) *FieldTypeCreate {
	ftc.mutation.SetNullLink(schema_link)
	return ftc
}

// SetActive sets the "active" field.
func (ftc *FieldTypeCreate) SetActive(schema_status schema.Status) *FieldTypeCreate {
	ftc.mutation.SetActive(schema_status)
	return ftc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableActive(schema_status *schema.Status) *FieldTypeCreate {
	if schema_status != nil {
		ftc.SetActive(*schema_status)
	}
	return ftc
}

// SetNullActive sets the "null_active" field.
func (ftc *FieldTypeCreate) SetNullActive(schema_status schema.Status) *FieldTypeCreate {
	ftc.mutation.SetNullActive(schema_status)
	return ftc
}

// SetNillableNullActive sets the "null_active" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullActive(schema_status *schema.Status) *FieldTypeCreate {
	if schema_status != nil {
		ftc.SetNullActive(*schema_status)
	}
	return ftc
}

// SetDeleted sets the "deleted" field.
func (ftc *FieldTypeCreate) SetDeleted(sql_nullbool *sql.NullBool) *FieldTypeCreate {
	ftc.mutation.SetDeleted(sql_nullbool)
	return ftc
}

// SetDeletedAt sets the "deleted_at" field.
func (ftc *FieldTypeCreate) SetDeletedAt(sql_nulltime *sql.NullTime) *FieldTypeCreate {
	ftc.mutation.SetDeletedAt(sql_nulltime)
	return ftc
}

// SetRawData sets the "raw_data" field.
func (ftc *FieldTypeCreate) SetRawData(b []byte) *FieldTypeCreate {
	ftc.mutation.SetRawData(b)
	return ftc
}

// SetSensitive sets the "sensitive" field.
func (ftc *FieldTypeCreate) SetSensitive(b []byte) *FieldTypeCreate {
	ftc.mutation.SetSensitive(b)
	return ftc
}

// SetIP sets the "ip" field.
func (ftc *FieldTypeCreate) SetIP(net_ip net.IP) *FieldTypeCreate {
	ftc.mutation.SetIP(net_ip)
	return ftc
}

// SetNullInt64 sets the "null_int64" field.
func (ftc *FieldTypeCreate) SetNullInt64(sql_nullint *sql.NullInt64) *FieldTypeCreate {
	ftc.mutation.SetNullInt64(sql_nullint)
	return ftc
}

// SetSchemaInt sets the "schema_int" field.
func (ftc *FieldTypeCreate) SetSchemaInt(schema_int schema.Int) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt(schema_int)
	return ftc
}

// SetNillableSchemaInt sets the "schema_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt(schema_int *schema.Int) *FieldTypeCreate {
	if schema_int != nil {
		ftc.SetSchemaInt(*schema_int)
	}
	return ftc
}

// SetSchemaInt8 sets the "schema_int8" field.
func (ftc *FieldTypeCreate) SetSchemaInt8(schema_int schema.Int8) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt8(schema_int)
	return ftc
}

// SetNillableSchemaInt8 sets the "schema_int8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt8(schema_int *schema.Int8) *FieldTypeCreate {
	if schema_int != nil {
		ftc.SetSchemaInt8(*schema_int)
	}
	return ftc
}

// SetSchemaInt64 sets the "schema_int64" field.
func (ftc *FieldTypeCreate) SetSchemaInt64(schema_int schema.Int64) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt64(schema_int)
	return ftc
}

// SetNillableSchemaInt64 sets the "schema_int64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt64(schema_int *schema.Int64) *FieldTypeCreate {
	if schema_int != nil {
		ftc.SetSchemaInt64(*schema_int)
	}
	return ftc
}

// SetSchemaFloat sets the "schema_float" field.
func (ftc *FieldTypeCreate) SetSchemaFloat(schema_float schema.Float64) *FieldTypeCreate {
	ftc.mutation.SetSchemaFloat(schema_float)
	return ftc
}

// SetNillableSchemaFloat sets the "schema_float" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaFloat(schema_float *schema.Float64) *FieldTypeCreate {
	if schema_float != nil {
		ftc.SetSchemaFloat(*schema_float)
	}
	return ftc
}

// SetSchemaFloat32 sets the "schema_float32" field.
func (ftc *FieldTypeCreate) SetSchemaFloat32(schema_float schema.Float32) *FieldTypeCreate {
	ftc.mutation.SetSchemaFloat32(schema_float)
	return ftc
}

// SetNillableSchemaFloat32 sets the "schema_float32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaFloat32(schema_float *schema.Float32) *FieldTypeCreate {
	if schema_float != nil {
		ftc.SetSchemaFloat32(*schema_float)
	}
	return ftc
}

// SetNullFloat sets the "null_float" field.
func (ftc *FieldTypeCreate) SetNullFloat(sql_nullfloat *sql.NullFloat64) *FieldTypeCreate {
	ftc.mutation.SetNullFloat(sql_nullfloat)
	return ftc
}

// SetRole sets the "role" field.
func (ftc *FieldTypeCreate) SetRole(role_role role.Role) *FieldTypeCreate {
	ftc.mutation.SetRole(role_role)
	return ftc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableRole(role_role *role.Role) *FieldTypeCreate {
	if role_role != nil {
		ftc.SetRole(*role_role)
	}
	return ftc
}

// SetPriority sets the "priority" field.
func (ftc *FieldTypeCreate) SetPriority(role_priority role.Priority) *FieldTypeCreate {
	ftc.mutation.SetPriority(role_priority)
	return ftc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillablePriority(role_priority *role.Priority) *FieldTypeCreate {
	if role_priority != nil {
		ftc.SetPriority(*role_priority)
	}
	return ftc
}

// SetOptionalUUID sets the "optional_uuid" field.
func (ftc *FieldTypeCreate) SetOptionalUUID(uuid_uuid uuid.UUID) *FieldTypeCreate {
	ftc.mutation.SetOptionalUUID(uuid_uuid)
	return ftc
}

// SetNillableOptionalUUID sets the "optional_uuid" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUUID(uuid_uuid *uuid.UUID) *FieldTypeCreate {
	if uuid_uuid != nil {
		ftc.SetOptionalUUID(*uuid_uuid)
	}
	return ftc
}

// SetNillableUUID sets the "nillable_uuid" field.
func (ftc *FieldTypeCreate) SetNillableUUID(uuid_uuid uuid.UUID) *FieldTypeCreate {
	ftc.mutation.SetNillableUUID(uuid_uuid)
	return ftc
}

// SetNillableNillableUUID sets the "nillable_uuid" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableUUID(uuid_uuid *uuid.UUID) *FieldTypeCreate {
	if uuid_uuid != nil {
		ftc.SetNillableUUID(*uuid_uuid)
	}
	return ftc
}

// SetStrings sets the "strings" field.
func (ftc *FieldTypeCreate) SetStrings(s []string) *FieldTypeCreate {
	ftc.mutation.SetStrings(s)
	return ftc
}

// SetPair sets the "pair" field.
func (ftc *FieldTypeCreate) SetPair(schema_pair schema.Pair) *FieldTypeCreate {
	ftc.mutation.SetPair(schema_pair)
	return ftc
}

// SetNillablePair sets the "pair" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillablePair(schema_pair *schema.Pair) *FieldTypeCreate {
	if schema_pair != nil {
		ftc.SetPair(*schema_pair)
	}
	return ftc
}

// SetNilPair sets the "nil_pair" field.
func (ftc *FieldTypeCreate) SetNilPair(schema_pair *schema.Pair) *FieldTypeCreate {
	ftc.mutation.SetNilPair(schema_pair)
	return ftc
}

// SetVstring sets the "vstring" field.
func (ftc *FieldTypeCreate) SetVstring(schema_vstring schema.VString) *FieldTypeCreate {
	ftc.mutation.SetVstring(schema_vstring)
	return ftc
}

// SetNillableVstring sets the "vstring" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableVstring(schema_vstring *schema.VString) *FieldTypeCreate {
	if schema_vstring != nil {
		ftc.SetVstring(*schema_vstring)
	}
	return ftc
}

// SetTriple sets the "triple" field.
func (ftc *FieldTypeCreate) SetTriple(schema_triple schema.Triple) *FieldTypeCreate {
	ftc.mutation.SetTriple(schema_triple)
	return ftc
}

// SetNillableTriple sets the "triple" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableTriple(schema_triple *schema.Triple) *FieldTypeCreate {
	if schema_triple != nil {
		ftc.SetTriple(*schema_triple)
	}
	return ftc
}

// SetBigInt sets the "big_int" field.
func (ftc *FieldTypeCreate) SetBigInt(schema_bigint schema.BigInt) *FieldTypeCreate {
	ftc.mutation.SetBigInt(schema_bigint)
	return ftc
}

// SetNillableBigInt sets the "big_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableBigInt(schema_bigint *schema.BigInt) *FieldTypeCreate {
	if schema_bigint != nil {
		ftc.SetBigInt(*schema_bigint)
	}
	return ftc
}

// SetPasswordOther sets the "password_other" field.
func (ftc *FieldTypeCreate) SetPasswordOther(schema_password schema.Password) *FieldTypeCreate {
	ftc.mutation.SetPasswordOther(schema_password)
	return ftc
}

// SetNillablePasswordOther sets the "password_other" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillablePasswordOther(schema_password *schema.Password) *FieldTypeCreate {
	if schema_password != nil {
		ftc.SetPasswordOther(*schema_password)
	}
	return ftc
}

// Mutation returns the FieldTypeMutation object of the builder.
func (ftc *FieldTypeCreate) Mutation() *FieldTypeMutation {
	return ftc.mutation
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	ftc.defaults()
	return withHooks(ctx, ftc.gremlinSave, ftc.mutation, ftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FieldTypeCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FieldTypeCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftc *FieldTypeCreate) defaults() {
	if _, ok := ftc.mutation.LinkOther(); !ok {
		v := field_type.DefaultLinkOther
		ftc.mutation.SetLinkOther(v)
	}
	if _, ok := ftc.mutation.LinkOtherFunc(); !ok {
		v := field_type.DefaultLinkOtherFunc()
		ftc.mutation.SetLinkOtherFunc(v)
	}
	if _, ok := ftc.mutation.Dir(); !ok {
		v := field_type.DefaultDir()
		ftc.mutation.SetDir(v)
	}
	if _, ok := ftc.mutation.Str(); !ok {
		v := field_type.DefaultStr()
		ftc.mutation.SetStr(v)
	}
	if _, ok := ftc.mutation.NullStr(); !ok {
		v := field_type.DefaultNullStr()
		ftc.mutation.SetNullStr(v)
	}
	if _, ok := ftc.mutation.DeletedAt(); !ok {
		v := field_type.DefaultDeletedAt()
		ftc.mutation.SetDeletedAt(v)
	}
	if _, ok := ftc.mutation.IP(); !ok {
		v := field_type.DefaultIP()
		ftc.mutation.SetIP(v)
	}
	if _, ok := ftc.mutation.Role(); !ok {
		v := field_type.DefaultRole
		ftc.mutation.SetRole(v)
	}
	if _, ok := ftc.mutation.Pair(); !ok {
		v := field_type.DefaultPair()
		ftc.mutation.SetPair(v)
	}
	if _, ok := ftc.mutation.Vstring(); !ok {
		v := field_type.DefaultVstring()
		ftc.mutation.SetVstring(v)
	}
	if _, ok := ftc.mutation.Triple(); !ok {
		v := field_type.DefaultTriple()
		ftc.mutation.SetTriple(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FieldTypeCreate) check() error {
	if _, ok := ftc.mutation.Int(); !ok {
		return &ValidationError{Name: "int", err: errors.New(`fluent: missing required field "FieldType.int"`)}
	}
	if _, ok := ftc.mutation.Int8(); !ok {
		return &ValidationError{Name: "int8", err: errors.New(`fluent: missing required field "FieldType.int8"`)}
	}
	if _, ok := ftc.mutation.Int16(); !ok {
		return &ValidationError{Name: "int16", err: errors.New(`fluent: missing required field "FieldType.int16"`)}
	}
	if _, ok := ftc.mutation.Int32(); !ok {
		return &ValidationError{Name: "int32", err: errors.New(`fluent: missing required field "FieldType.int32"`)}
	}
	if _, ok := ftc.mutation.Int64(); !ok {
		return &ValidationError{Name: "int64", err: errors.New(`fluent: missing required field "FieldType.int64"`)}
	}
	if v, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		if err := field_type.ValidateOptionalInt32Validator(v); err != nil {
			return &ValidationError{Name: "validate_optional_int32", err: fmt.Errorf(`fluent: validator failed for field "FieldType.validate_optional_int32": %w`, err)}
		}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := field_type.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`fluent: validator failed for field "FieldType.state": %w`, err)}
		}
	}
	if v, ok := ftc.mutation.MAC(); ok {
		if err := field_type.MACValidator(v.String()); err != nil {
			return &ValidationError{Name: "mac", err: fmt.Errorf(`fluent: validator failed for field "FieldType.mac": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.Dir(); !ok {
		return &ValidationError{Name: "dir", err: errors.New(`fluent: missing required field "FieldType.dir"`)}
	}
	if v, ok := ftc.mutation.Ndir(); ok {
		if err := field_type.NdirValidator(string(v)); err != nil {
			return &ValidationError{Name: "ndir", err: fmt.Errorf(`fluent: validator failed for field "FieldType.ndir": %w`, err)}
		}
	}
	if v, ok := ftc.mutation.Link(); ok {
		if err := field_type.LinkValidator(v.String()); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`fluent: validator failed for field "FieldType.link": %w`, err)}
		}
	}
	if v, ok := ftc.mutation.RawData(); ok {
		if err := field_type.RawDataValidator(v); err != nil {
			return &ValidationError{Name: "raw_data", err: fmt.Errorf(`fluent: validator failed for field "FieldType.raw_data": %w`, err)}
		}
	}
	if v, ok := ftc.mutation.IP(); ok {
		if err := field_type.IPValidator([]byte(v)); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`fluent: validator failed for field "FieldType.ip": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`fluent: missing required field "FieldType.role"`)}
	}
	if v, ok := ftc.mutation.Role(); ok {
		if err := field_type.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`fluent: validator failed for field "FieldType.role": %w`, err)}
		}
	}
	if v, ok := ftc.mutation.Priority(); ok {
		if err := field_type.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`fluent: validator failed for field "FieldType.priority": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.Pair(); !ok {
		return &ValidationError{Name: "pair", err: errors.New(`fluent: missing required field "FieldType.pair"`)}
	}
	if _, ok := ftc.mutation.Vstring(); !ok {
		return &ValidationError{Name: "vstring", err: errors.New(`fluent: missing required field "FieldType.vstring"`)}
	}
	if _, ok := ftc.mutation.Triple(); !ok {
		return &ValidationError{Name: "triple", err: errors.New(`fluent: missing required field "FieldType.triple"`)}
	}
	return nil
}

func (ftc *FieldTypeCreate) gremlinSave(ctx context.Context) (*FieldType, error) {
	if err := ftc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := ftc.gremlin().Query()
	if err := ftc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &FieldType{config: ftc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	ftc.mutation.id = &rnode.ID
	ftc.mutation.done = true
	return rnode, nil
}

func (ftc *FieldTypeCreate) gremlin() *dsl.Traversal {
	v := g.AddV(field_type.Label)
	if value, ok := ftc.mutation.Int(); ok {
		v.Property(dsl.Single, field_type.FieldInt, value)
	}
	if value, ok := ftc.mutation.Int8(); ok {
		v.Property(dsl.Single, field_type.FieldInt8, value)
	}
	if value, ok := ftc.mutation.Int16(); ok {
		v.Property(dsl.Single, field_type.FieldInt16, value)
	}
	if value, ok := ftc.mutation.Int32(); ok {
		v.Property(dsl.Single, field_type.FieldInt32, value)
	}
	if value, ok := ftc.mutation.Int64(); ok {
		v.Property(dsl.Single, field_type.FieldInt64, value)
	}
	if value, ok := ftc.mutation.OptionalInt(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalInt, value)
	}
	if value, ok := ftc.mutation.OptionalInt8(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalInt8, value)
	}
	if value, ok := ftc.mutation.OptionalInt16(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalInt16, value)
	}
	if value, ok := ftc.mutation.OptionalInt32(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalInt32, value)
	}
	if value, ok := ftc.mutation.OptionalInt64(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalInt64, value)
	}
	if value, ok := ftc.mutation.NillableInt(); ok {
		v.Property(dsl.Single, field_type.FieldNillableInt, value)
	}
	if value, ok := ftc.mutation.NillableInt8(); ok {
		v.Property(dsl.Single, field_type.FieldNillableInt8, value)
	}
	if value, ok := ftc.mutation.NillableInt16(); ok {
		v.Property(dsl.Single, field_type.FieldNillableInt16, value)
	}
	if value, ok := ftc.mutation.NillableInt32(); ok {
		v.Property(dsl.Single, field_type.FieldNillableInt32, value)
	}
	if value, ok := ftc.mutation.NillableInt64(); ok {
		v.Property(dsl.Single, field_type.FieldNillableInt64, value)
	}
	if value, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		v.Property(dsl.Single, field_type.FieldValidateOptionalInt32, value)
	}
	if value, ok := ftc.mutation.OptionalUint(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalUint, value)
	}
	if value, ok := ftc.mutation.OptionalUint8(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalUint8, value)
	}
	if value, ok := ftc.mutation.OptionalUint16(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalUint16, value)
	}
	if value, ok := ftc.mutation.OptionalUint32(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalUint32, value)
	}
	if value, ok := ftc.mutation.OptionalUint64(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalUint64, value)
	}
	if value, ok := ftc.mutation.State(); ok {
		v.Property(dsl.Single, field_type.FieldState, value)
	}
	if value, ok := ftc.mutation.OptionalFloat(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalFloat, value)
	}
	if value, ok := ftc.mutation.OptionalFloat32(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalFloat32, value)
	}
	if value, ok := ftc.mutation.Text(); ok {
		v.Property(dsl.Single, field_type.FieldText, value)
	}
	if value, ok := ftc.mutation.Datetime(); ok {
		v.Property(dsl.Single, field_type.FieldDatetime, value)
	}
	if value, ok := ftc.mutation.Decimal(); ok {
		v.Property(dsl.Single, field_type.FieldDecimal, value)
	}
	if value, ok := ftc.mutation.LinkOther(); ok {
		v.Property(dsl.Single, field_type.FieldLinkOther, value)
	}
	if value, ok := ftc.mutation.LinkOtherFunc(); ok {
		v.Property(dsl.Single, field_type.FieldLinkOtherFunc, value)
	}
	if value, ok := ftc.mutation.MAC(); ok {
		v.Property(dsl.Single, field_type.FieldMAC, value)
	}
	if value, ok := ftc.mutation.StringArray(); ok {
		v.Property(dsl.Single, field_type.FieldStringArray, value)
	}
	if value, ok := ftc.mutation.Password(); ok {
		v.Property(dsl.Single, field_type.FieldPassword, value)
	}
	if value, ok := ftc.mutation.StringScanner(); ok {
		v.Property(dsl.Single, field_type.FieldStringScanner, value)
	}
	if value, ok := ftc.mutation.Duration(); ok {
		v.Property(dsl.Single, field_type.FieldDuration, value)
	}
	if value, ok := ftc.mutation.Dir(); ok {
		v.Property(dsl.Single, field_type.FieldDir, value)
	}
	if value, ok := ftc.mutation.Ndir(); ok {
		v.Property(dsl.Single, field_type.FieldNdir, value)
	}
	if value, ok := ftc.mutation.Str(); ok {
		v.Property(dsl.Single, field_type.FieldStr, value)
	}
	if value, ok := ftc.mutation.NullStr(); ok {
		v.Property(dsl.Single, field_type.FieldNullStr, value)
	}
	if value, ok := ftc.mutation.Link(); ok {
		v.Property(dsl.Single, field_type.FieldLink, value)
	}
	if value, ok := ftc.mutation.NullLink(); ok {
		v.Property(dsl.Single, field_type.FieldNullLink, value)
	}
	if value, ok := ftc.mutation.Active(); ok {
		v.Property(dsl.Single, field_type.FieldActive, value)
	}
	if value, ok := ftc.mutation.NullActive(); ok {
		v.Property(dsl.Single, field_type.FieldNullActive, value)
	}
	if value, ok := ftc.mutation.Deleted(); ok {
		v.Property(dsl.Single, field_type.FieldDeleted, value)
	}
	if value, ok := ftc.mutation.DeletedAt(); ok {
		v.Property(dsl.Single, field_type.FieldDeletedAt, value)
	}
	if value, ok := ftc.mutation.RawData(); ok {
		v.Property(dsl.Single, field_type.FieldRawData, value)
	}
	if value, ok := ftc.mutation.Sensitive(); ok {
		v.Property(dsl.Single, field_type.FieldSensitive, value)
	}
	if value, ok := ftc.mutation.IP(); ok {
		v.Property(dsl.Single, field_type.FieldIP, value)
	}
	if value, ok := ftc.mutation.NullInt64(); ok {
		v.Property(dsl.Single, field_type.FieldNullInt64, value)
	}
	if value, ok := ftc.mutation.SchemaInt(); ok {
		v.Property(dsl.Single, field_type.FieldSchemaInt, value)
	}
	if value, ok := ftc.mutation.SchemaInt8(); ok {
		v.Property(dsl.Single, field_type.FieldSchemaInt8, value)
	}
	if value, ok := ftc.mutation.SchemaInt64(); ok {
		v.Property(dsl.Single, field_type.FieldSchemaInt64, value)
	}
	if value, ok := ftc.mutation.SchemaFloat(); ok {
		v.Property(dsl.Single, field_type.FieldSchemaFloat, value)
	}
	if value, ok := ftc.mutation.SchemaFloat32(); ok {
		v.Property(dsl.Single, field_type.FieldSchemaFloat32, value)
	}
	if value, ok := ftc.mutation.NullFloat(); ok {
		v.Property(dsl.Single, field_type.FieldNullFloat, value)
	}
	if value, ok := ftc.mutation.Role(); ok {
		v.Property(dsl.Single, field_type.FieldRole, value)
	}
	if value, ok := ftc.mutation.Priority(); ok {
		v.Property(dsl.Single, field_type.FieldPriority, value)
	}
	if value, ok := ftc.mutation.OptionalUUID(); ok {
		v.Property(dsl.Single, field_type.FieldOptionalUUID, value)
	}
	if value, ok := ftc.mutation.NillableUUID(); ok {
		v.Property(dsl.Single, field_type.FieldNillableUUID, value)
	}
	if value, ok := ftc.mutation.Strings(); ok {
		v.Property(dsl.Single, field_type.FieldStrings, value)
	}
	if value, ok := ftc.mutation.Pair(); ok {
		v.Property(dsl.Single, field_type.FieldPair, value)
	}
	if value, ok := ftc.mutation.NilPair(); ok {
		v.Property(dsl.Single, field_type.FieldNilPair, value)
	}
	if value, ok := ftc.mutation.Vstring(); ok {
		v.Property(dsl.Single, field_type.FieldVstring, value)
	}
	if value, ok := ftc.mutation.Triple(); ok {
		v.Property(dsl.Single, field_type.FieldTriple, value)
	}
	if value, ok := ftc.mutation.BigInt(); ok {
		v.Property(dsl.Single, field_type.FieldBigInt, value)
	}
	if value, ok := ftc.mutation.PasswordOther(); ok {
		v.Property(dsl.Single, field_type.FieldPasswordOther, value)
	}
	return v.ValueMap(true)
}

// FieldTypeCreateBulk is the builder for creating many FieldType entities in bulk.
type FieldTypeCreateBulk struct {
	config
	err      error
	builders []*FieldTypeCreate
}
