// Code generated by entc, DO NOT EDIT.

package review

import (
	"imdbv2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...int) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRank), v...))
	})
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...int) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRank), v...))
	})
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// HasMovie applies the HasEdge predicate on the "movie" edge.
func HasMovie() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MovieTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MovieTable, MovieColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMovieWith applies the HasEdge predicate on the "movie" edge with a given conditions (other predicates).
func HasMovieWith(preds ...predicate.Movie) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MovieInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MovieTable, MovieColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		p(s.Not())
	})
}
