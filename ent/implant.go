// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
)

// Implant is the model entity for the Implant schema.
type Implant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// MachineID holds the value of the "machine_id" field.
	MachineID string `json:"machine_id,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// LastSeenAt holds the value of the "last_seen_at" field.
	LastSeenAt time.Time `json:"last_seen_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImplantQuery when eager-loading is set.
	Edges ImplantEdges `json:"edges"`
}

// ImplantEdges holds the relations/edges for other nodes in the graph.
type ImplantEdges struct {
	// Heartbeats holds the value of the heartbeats edge.
	Heartbeats []*Heartbeat `json:"heartbeats,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// HeartbeatsOrErr returns the Heartbeats value or an error if the edge
// was not loaded in eager-loading.
func (e ImplantEdges) HeartbeatsOrErr() ([]*Heartbeat, error) {
	if e.loadedTypes[0] {
		return e.Heartbeats, nil
	}
	return nil, &NotLoadedError{edge: "heartbeats"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ImplantEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Implant) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case implant.FieldID:
			values[i] = new(sql.NullInt64)
		case implant.FieldUUID, implant.FieldMachineID, implant.FieldHostname, implant.FieldIP:
			values[i] = new(sql.NullString)
		case implant.FieldLastSeenAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Implant", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Implant fields.
func (i *Implant) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case implant.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case implant.FieldUUID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[j])
			} else if value.Valid {
				i.UUID = value.String
			}
		case implant.FieldMachineID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field machine_id", values[j])
			} else if value.Valid {
				i.MachineID = value.String
			}
		case implant.FieldHostname:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[j])
			} else if value.Valid {
				i.Hostname = value.String
			}
		case implant.FieldIP:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[j])
			} else if value.Valid {
				i.IP = value.String
			}
		case implant.FieldLastSeenAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_seen_at", values[j])
			} else if value.Valid {
				i.LastSeenAt = value.Time
			}
		}
	}
	return nil
}

// QueryHeartbeats queries the "heartbeats" edge of the Implant entity.
func (i *Implant) QueryHeartbeats() *HeartbeatQuery {
	return (&ImplantClient{config: i.config}).QueryHeartbeats(i)
}

// QueryTasks queries the "tasks" edge of the Implant entity.
func (i *Implant) QueryTasks() *TaskQuery {
	return (&ImplantClient{config: i.config}).QueryTasks(i)
}

// Update returns a builder for updating this Implant.
// Note that you need to call Implant.Unwrap() before calling this method if this Implant
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Implant) Update() *ImplantUpdateOne {
	return (&ImplantClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Implant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Implant) Unwrap() *Implant {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Implant is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Implant) String() string {
	var builder strings.Builder
	builder.WriteString("Implant(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(i.UUID)
	builder.WriteString(", machine_id=")
	builder.WriteString(i.MachineID)
	builder.WriteString(", hostname=")
	builder.WriteString(i.Hostname)
	builder.WriteString(", ip=")
	builder.WriteString(i.IP)
	builder.WriteString(", last_seen_at=")
	builder.WriteString(i.LastSeenAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Implants is a parsable slice of Implant.
type Implants []*Implant

func (i Implants) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
