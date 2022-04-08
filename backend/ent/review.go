// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"imdbv2/ent/movie"
	"imdbv2/ent/review"
	"imdbv2/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Review is the model entity for the Review schema.
type Review struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank int `json:"rank,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewQuery when eager-loading is set.
	Edges        ReviewEdges `json:"edges"`
	review_movie *int
	user_reviews *int
}

// ReviewEdges holds the relations/edges for other nodes in the graph.
type ReviewEdges struct {
	// Movie holds the value of the movie edge.
	Movie *Movie `json:"movie,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MovieOrErr returns the Movie value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) MovieOrErr() (*Movie, error) {
	if e.loadedTypes[0] {
		if e.Movie == nil {
			// The edge movie was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: movie.Label}
		}
		return e.Movie, nil
	}
	return nil, &NotLoadedError{edge: "movie"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Review) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case review.FieldID, review.FieldRank:
			values[i] = new(sql.NullInt64)
		case review.FieldText:
			values[i] = new(sql.NullString)
		case review.ForeignKeys[0]: // review_movie
			values[i] = new(sql.NullInt64)
		case review.ForeignKeys[1]: // user_reviews
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Review", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Review fields.
func (r *Review) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case review.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				r.Text = value.String
			}
		case review.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				r.Rank = int(value.Int64)
			}
		case review.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field review_movie", value)
			} else if value.Valid {
				r.review_movie = new(int)
				*r.review_movie = int(value.Int64)
			}
		case review.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_reviews", value)
			} else if value.Valid {
				r.user_reviews = new(int)
				*r.user_reviews = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMovie queries the "movie" edge of the Review entity.
func (r *Review) QueryMovie() *MovieQuery {
	return (&ReviewClient{config: r.config}).QueryMovie(r)
}

// QueryUser queries the "user" edge of the Review entity.
func (r *Review) QueryUser() *UserQuery {
	return (&ReviewClient{config: r.config}).QueryUser(r)
}

// Update returns a builder for updating this Review.
// Note that you need to call Review.Unwrap() before calling this method if this Review
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Review) Update() *ReviewUpdateOne {
	return (&ReviewClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Review entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Review) Unwrap() *Review {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Review is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Review) String() string {
	var builder strings.Builder
	builder.WriteString("Review(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", text=")
	builder.WriteString(r.Text)
	builder.WriteString(", rank=")
	builder.WriteString(fmt.Sprintf("%v", r.Rank))
	builder.WriteByte(')')
	return builder.String()
}

// Reviews is a parsable slice of Review.
type Reviews []*Review

func (r Reviews) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
