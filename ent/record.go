// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/character"
	"github.com/witchs-lounge_backend/ent/music"
	"github.com/witchs-lounge_backend/ent/record"
	"github.com/witchs-lounge_backend/ent/stage"
	"github.com/witchs-lounge_backend/ent/user"
)

// Record is the model entity for the Record schema.
type Record struct {
	config `json:"-"`
	// ID of the ent.
	// Global custom UUID ID
	ID uuid.UUID `json:"id,omitempty"`
	// Created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Updated time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 유저 ID
	UserID uuid.UUID `json:"user_id,omitempty"`
	// 음악 ID
	MusicID uuid.UUID `json:"music_id,omitempty"`
	// 스테이지 ID
	StageID uuid.UUID `json:"stage_id,omitempty"`
	// 캐릭터 ID
	CharacterID uuid.UUID `json:"character_id,omitempty"`
	// 점수
	Score int `json:"score,omitempty"`
	// Perfect 개수
	PerfectCount int `json:"perfect_count,omitempty"`
	// Good 개수
	GoodCount int `json:"good_count,omitempty"`
	// Bad 개수
	BadCount int `json:"bad_count,omitempty"`
	// Miss 개수
	MissCount int `json:"miss_count,omitempty"`
	// 최대 콤보
	MaxCombo int `json:"max_combo,omitempty"`
	// 정확도 (%)
	Accuracy float64 `json:"accuracy,omitempty"`
	// 랭크
	Rank record.Rank `json:"rank,omitempty"`
	// 풀콤보 여부
	IsFullCombo bool `json:"is_full_combo,omitempty"`
	// 퍼펙트 플레이 여부
	IsPerfectPlay bool `json:"is_perfect_play,omitempty"`
	// 플레이 시간
	PlayedAt time.Time `json:"played_at,omitempty"`
	// 플레이 소요시간(초)
	PlayDuration int `json:"play_duration,omitempty"`
	// 추가 정보
	AdditionalInfo map[string]interface{} `json:"additional_info,omitempty"`
	// 유효한 기록 여부
	IsValid bool `json:"is_valid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecordQuery when eager-loading is set.
	Edges        RecordEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RecordEdges holds the relations/edges for other nodes in the graph.
type RecordEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Music holds the value of the music edge.
	Music *Music `json:"music,omitempty"`
	// Stage holds the value of the stage edge.
	Stage *Stage `json:"stage,omitempty"`
	// Character holds the value of the character edge.
	Character *Character `json:"character,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MusicOrErr returns the Music value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) MusicOrErr() (*Music, error) {
	if e.Music != nil {
		return e.Music, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: music.Label}
	}
	return nil, &NotLoadedError{edge: "music"}
}

// StageOrErr returns the Stage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) StageOrErr() (*Stage, error) {
	if e.Stage != nil {
		return e.Stage, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: stage.Label}
	}
	return nil, &NotLoadedError{edge: "stage"}
}

// CharacterOrErr returns the Character value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) CharacterOrErr() (*Character, error) {
	if e.Character != nil {
		return e.Character, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: character.Label}
	}
	return nil, &NotLoadedError{edge: "character"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Record) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case record.FieldAdditionalInfo:
			values[i] = new([]byte)
		case record.FieldIsFullCombo, record.FieldIsPerfectPlay, record.FieldIsValid:
			values[i] = new(sql.NullBool)
		case record.FieldAccuracy:
			values[i] = new(sql.NullFloat64)
		case record.FieldScore, record.FieldPerfectCount, record.FieldGoodCount, record.FieldBadCount, record.FieldMissCount, record.FieldMaxCombo, record.FieldPlayDuration:
			values[i] = new(sql.NullInt64)
		case record.FieldRank:
			values[i] = new(sql.NullString)
		case record.FieldCreatedAt, record.FieldUpdatedAt, record.FieldPlayedAt:
			values[i] = new(sql.NullTime)
		case record.FieldID, record.FieldUserID, record.FieldMusicID, record.FieldStageID, record.FieldCharacterID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Record fields.
func (r *Record) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case record.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case record.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case record.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				r.UserID = *value
			}
		case record.FieldMusicID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field music_id", values[i])
			} else if value != nil {
				r.MusicID = *value
			}
		case record.FieldStageID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field stage_id", values[i])
			} else if value != nil {
				r.StageID = *value
			}
		case record.FieldCharacterID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field character_id", values[i])
			} else if value != nil {
				r.CharacterID = *value
			}
		case record.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				r.Score = int(value.Int64)
			}
		case record.FieldPerfectCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field perfect_count", values[i])
			} else if value.Valid {
				r.PerfectCount = int(value.Int64)
			}
		case record.FieldGoodCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field good_count", values[i])
			} else if value.Valid {
				r.GoodCount = int(value.Int64)
			}
		case record.FieldBadCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bad_count", values[i])
			} else if value.Valid {
				r.BadCount = int(value.Int64)
			}
		case record.FieldMissCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field miss_count", values[i])
			} else if value.Valid {
				r.MissCount = int(value.Int64)
			}
		case record.FieldMaxCombo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_combo", values[i])
			} else if value.Valid {
				r.MaxCombo = int(value.Int64)
			}
		case record.FieldAccuracy:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field accuracy", values[i])
			} else if value.Valid {
				r.Accuracy = value.Float64
			}
		case record.FieldRank:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				r.Rank = record.Rank(value.String)
			}
		case record.FieldIsFullCombo:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_full_combo", values[i])
			} else if value.Valid {
				r.IsFullCombo = value.Bool
			}
		case record.FieldIsPerfectPlay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_perfect_play", values[i])
			} else if value.Valid {
				r.IsPerfectPlay = value.Bool
			}
		case record.FieldPlayedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field played_at", values[i])
			} else if value.Valid {
				r.PlayedAt = value.Time
			}
		case record.FieldPlayDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field play_duration", values[i])
			} else if value.Valid {
				r.PlayDuration = int(value.Int64)
			}
		case record.FieldAdditionalInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.AdditionalInfo); err != nil {
					return fmt.Errorf("unmarshal field additional_info: %w", err)
				}
			}
		case record.FieldIsValid:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_valid", values[i])
			} else if value.Valid {
				r.IsValid = value.Bool
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Record.
// This includes values selected through modifiers, order, etc.
func (r *Record) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Record entity.
func (r *Record) QueryUser() *UserQuery {
	return NewRecordClient(r.config).QueryUser(r)
}

// QueryMusic queries the "music" edge of the Record entity.
func (r *Record) QueryMusic() *MusicQuery {
	return NewRecordClient(r.config).QueryMusic(r)
}

// QueryStage queries the "stage" edge of the Record entity.
func (r *Record) QueryStage() *StageQuery {
	return NewRecordClient(r.config).QueryStage(r)
}

// QueryCharacter queries the "character" edge of the Record entity.
func (r *Record) QueryCharacter() *CharacterQuery {
	return NewRecordClient(r.config).QueryCharacter(r)
}

// Update returns a builder for updating this Record.
// Note that you need to call Record.Unwrap() before calling this method if this Record
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Record) Update() *RecordUpdateOne {
	return NewRecordClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Record entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Record) Unwrap() *Record {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Record is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Record) String() string {
	var builder strings.Builder
	builder.WriteString("Record(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("music_id=")
	builder.WriteString(fmt.Sprintf("%v", r.MusicID))
	builder.WriteString(", ")
	builder.WriteString("stage_id=")
	builder.WriteString(fmt.Sprintf("%v", r.StageID))
	builder.WriteString(", ")
	builder.WriteString("character_id=")
	builder.WriteString(fmt.Sprintf("%v", r.CharacterID))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", r.Score))
	builder.WriteString(", ")
	builder.WriteString("perfect_count=")
	builder.WriteString(fmt.Sprintf("%v", r.PerfectCount))
	builder.WriteString(", ")
	builder.WriteString("good_count=")
	builder.WriteString(fmt.Sprintf("%v", r.GoodCount))
	builder.WriteString(", ")
	builder.WriteString("bad_count=")
	builder.WriteString(fmt.Sprintf("%v", r.BadCount))
	builder.WriteString(", ")
	builder.WriteString("miss_count=")
	builder.WriteString(fmt.Sprintf("%v", r.MissCount))
	builder.WriteString(", ")
	builder.WriteString("max_combo=")
	builder.WriteString(fmt.Sprintf("%v", r.MaxCombo))
	builder.WriteString(", ")
	builder.WriteString("accuracy=")
	builder.WriteString(fmt.Sprintf("%v", r.Accuracy))
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", r.Rank))
	builder.WriteString(", ")
	builder.WriteString("is_full_combo=")
	builder.WriteString(fmt.Sprintf("%v", r.IsFullCombo))
	builder.WriteString(", ")
	builder.WriteString("is_perfect_play=")
	builder.WriteString(fmt.Sprintf("%v", r.IsPerfectPlay))
	builder.WriteString(", ")
	builder.WriteString("played_at=")
	builder.WriteString(r.PlayedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("play_duration=")
	builder.WriteString(fmt.Sprintf("%v", r.PlayDuration))
	builder.WriteString(", ")
	builder.WriteString("additional_info=")
	builder.WriteString(fmt.Sprintf("%v", r.AdditionalInfo))
	builder.WriteString(", ")
	builder.WriteString("is_valid=")
	builder.WriteString(fmt.Sprintf("%v", r.IsValid))
	builder.WriteByte(')')
	return builder.String()
}

// Records is a parsable slice of Record.
type Records []*Record
