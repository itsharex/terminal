// Code generated by ent, DO NOT EDIT.

package keys

import (
	"terminal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Keys {
	return predicate.Keys(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Keys {
	return predicate.Keys(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Keys {
	return predicate.Keys(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Keys {
	return predicate.Keys(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Keys {
	return predicate.Keys(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Keys {
	return predicate.Keys(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Keys {
	return predicate.Keys(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Keys {
	return predicate.Keys(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Keys {
	return predicate.Keys(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldEQ(FieldContent, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.Keys {
	return predicate.Keys(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.Keys {
	return predicate.Keys(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.Keys {
	return predicate.Keys(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.Keys {
	return predicate.Keys(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.Keys {
	return predicate.Keys(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.Keys {
	return predicate.Keys(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.Keys {
	return predicate.Keys(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.Keys {
	return predicate.Keys(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.Keys {
	return predicate.Keys(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.Keys {
	return predicate.Keys(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.Keys {
	return predicate.Keys(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.Keys {
	return predicate.Keys(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.Keys {
	return predicate.Keys(sql.FieldContainsFold(FieldLabel, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...[]byte) predicate.Keys {
	return predicate.Keys(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...[]byte) predicate.Keys {
	return predicate.Keys(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v []byte) predicate.Keys {
	return predicate.Keys(sql.FieldLTE(FieldContent, v))
}

// HasHost applies the HasEdge predicate on the "host" edge.
func HasHost() predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HostTable, HostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostWith applies the HasEdge predicate on the "host" edge with a given conditions (other predicates).
func HasHostWith(preds ...predicate.Hosts) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		step := newHostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Keys) predicate.Keys {
	return predicate.Keys(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Keys) predicate.Keys {
	return predicate.Keys(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Keys) predicate.Keys {
	return predicate.Keys(sql.NotPredicates(p))
}
