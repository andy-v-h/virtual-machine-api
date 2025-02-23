// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package virtualmachine

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldName, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldOwnerID, v))
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldLocationID, v))
}

// Userdata applies equality check predicate on the "userdata" field. It's identical to UserdataEQ.
func Userdata(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldUserdata, v))
}

// VMCPUConfigID applies equality check predicate on the "vm_cpu_config_id" field. It's identical to VMCPUConfigIDEQ.
func VMCPUConfigID(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldVMCPUConfigID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldName, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContains(FieldOwnerID, vc))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldOwnerID, vc))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldOwnerID, vc))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldOwnerID, vc))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldOwnerID, vc))
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldLocationID, v))
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldLocationID, v))
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldLocationID, vs...))
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldLocationID, vs...))
}

// LocationIDGT applies the GT predicate on the "location_id" field.
func LocationIDGT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldLocationID, v))
}

// LocationIDGTE applies the GTE predicate on the "location_id" field.
func LocationIDGTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldLocationID, v))
}

// LocationIDLT applies the LT predicate on the "location_id" field.
func LocationIDLT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldLocationID, v))
}

// LocationIDLTE applies the LTE predicate on the "location_id" field.
func LocationIDLTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldLocationID, v))
}

// LocationIDContains applies the Contains predicate on the "location_id" field.
func LocationIDContains(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContains(FieldLocationID, vc))
}

// LocationIDHasPrefix applies the HasPrefix predicate on the "location_id" field.
func LocationIDHasPrefix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldLocationID, vc))
}

// LocationIDHasSuffix applies the HasSuffix predicate on the "location_id" field.
func LocationIDHasSuffix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldLocationID, vc))
}

// LocationIDEqualFold applies the EqualFold predicate on the "location_id" field.
func LocationIDEqualFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldLocationID, vc))
}

// LocationIDContainsFold applies the ContainsFold predicate on the "location_id" field.
func LocationIDContainsFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldLocationID, vc))
}

// UserdataEQ applies the EQ predicate on the "userdata" field.
func UserdataEQ(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldUserdata, v))
}

// UserdataNEQ applies the NEQ predicate on the "userdata" field.
func UserdataNEQ(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldUserdata, v))
}

// UserdataIn applies the In predicate on the "userdata" field.
func UserdataIn(vs ...string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldUserdata, vs...))
}

// UserdataNotIn applies the NotIn predicate on the "userdata" field.
func UserdataNotIn(vs ...string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldUserdata, vs...))
}

// UserdataGT applies the GT predicate on the "userdata" field.
func UserdataGT(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldUserdata, v))
}

// UserdataGTE applies the GTE predicate on the "userdata" field.
func UserdataGTE(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldUserdata, v))
}

// UserdataLT applies the LT predicate on the "userdata" field.
func UserdataLT(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldUserdata, v))
}

// UserdataLTE applies the LTE predicate on the "userdata" field.
func UserdataLTE(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldUserdata, v))
}

// UserdataContains applies the Contains predicate on the "userdata" field.
func UserdataContains(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldContains(FieldUserdata, v))
}

// UserdataHasPrefix applies the HasPrefix predicate on the "userdata" field.
func UserdataHasPrefix(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldUserdata, v))
}

// UserdataHasSuffix applies the HasSuffix predicate on the "userdata" field.
func UserdataHasSuffix(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldUserdata, v))
}

// UserdataIsNil applies the IsNil predicate on the "userdata" field.
func UserdataIsNil() predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIsNull(FieldUserdata))
}

// UserdataNotNil applies the NotNil predicate on the "userdata" field.
func UserdataNotNil() predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotNull(FieldUserdata))
}

// UserdataEqualFold applies the EqualFold predicate on the "userdata" field.
func UserdataEqualFold(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldUserdata, v))
}

// UserdataContainsFold applies the ContainsFold predicate on the "userdata" field.
func UserdataContainsFold(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldUserdata, v))
}

// VMCPUConfigIDEQ applies the EQ predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldVMCPUConfigID, v))
}

// VMCPUConfigIDNEQ applies the NEQ predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDNEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldVMCPUConfigID, v))
}

// VMCPUConfigIDIn applies the In predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldVMCPUConfigID, vs...))
}

// VMCPUConfigIDNotIn applies the NotIn predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDNotIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldVMCPUConfigID, vs...))
}

// VMCPUConfigIDGT applies the GT predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDGT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldVMCPUConfigID, v))
}

// VMCPUConfigIDGTE applies the GTE predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDGTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldVMCPUConfigID, v))
}

// VMCPUConfigIDLT applies the LT predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDLT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldVMCPUConfigID, v))
}

// VMCPUConfigIDLTE applies the LTE predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDLTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldVMCPUConfigID, v))
}

// VMCPUConfigIDContains applies the Contains predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDContains(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContains(FieldVMCPUConfigID, vc))
}

// VMCPUConfigIDHasPrefix applies the HasPrefix predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDHasPrefix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldVMCPUConfigID, vc))
}

// VMCPUConfigIDHasSuffix applies the HasSuffix predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDHasSuffix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldVMCPUConfigID, vc))
}

// VMCPUConfigIDEqualFold applies the EqualFold predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDEqualFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldVMCPUConfigID, vc))
}

// VMCPUConfigIDContainsFold applies the ContainsFold predicate on the "vm_cpu_config_id" field.
func VMCPUConfigIDContainsFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldVMCPUConfigID, vc))
}

// HasVirtualMachineCPUConfig applies the HasEdge predicate on the "virtual_machine_cpu_config" edge.
func HasVirtualMachineCPUConfig() predicate.VirtualMachine {
	return predicate.VirtualMachine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, VirtualMachineCPUConfigTable, VirtualMachineCPUConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVirtualMachineCPUConfigWith applies the HasEdge predicate on the "virtual_machine_cpu_config" edge with a given conditions (other predicates).
func HasVirtualMachineCPUConfigWith(preds ...predicate.VirtualMachineCPUConfig) predicate.VirtualMachine {
	return predicate.VirtualMachine(func(s *sql.Selector) {
		step := newVirtualMachineCPUConfigStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VirtualMachine) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VirtualMachine) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VirtualMachine) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.NotPredicates(p))
}
