// Code generated by ent, DO NOT EDIT.

package music

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldName, v))
}

// Artist applies equality check predicate on the "artist" field. It's identical to ArtistEQ.
func Artist(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldArtist, v))
}

// Composer applies equality check predicate on the "composer" field. It's identical to ComposerEQ.
func Composer(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldComposer, v))
}

// MusicSource applies equality check predicate on the "music_source" field. It's identical to MusicSourceEQ.
func MusicSource(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldMusicSource, v))
}

// JacketSource applies equality check predicate on the "jacket_source" field. It's identical to JacketSourceEQ.
func JacketSource(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldJacketSource, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v float64) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldDuration, v))
}

// Bpm applies equality check predicate on the "bpm" field. It's identical to BpmEQ.
func Bpm(v float64) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldBpm, v))
}

// Genre applies equality check predicate on the "genre" field. It's identical to GenreEQ.
func Genre(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldGenre, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldDescription, v))
}

// IsFeatured applies equality check predicate on the "is_featured" field. It's identical to IsFeaturedEQ.
func IsFeatured(v bool) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldIsFeatured, v))
}

// IsFree applies equality check predicate on the "is_free" field. It's identical to IsFreeEQ.
func IsFree(v bool) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldIsFree, v))
}

// UnlockLevel applies equality check predicate on the "unlock_level" field. It's identical to UnlockLevelEQ.
func UnlockLevel(v int) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldUnlockLevel, v))
}

// ReleaseDate applies equality check predicate on the "release_date" field. It's identical to ReleaseDateEQ.
func ReleaseDate(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldReleaseDate, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldIsActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldName, v))
}

// ArtistEQ applies the EQ predicate on the "artist" field.
func ArtistEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldArtist, v))
}

// ArtistNEQ applies the NEQ predicate on the "artist" field.
func ArtistNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldArtist, v))
}

// ArtistIn applies the In predicate on the "artist" field.
func ArtistIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldArtist, vs...))
}

// ArtistNotIn applies the NotIn predicate on the "artist" field.
func ArtistNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldArtist, vs...))
}

// ArtistGT applies the GT predicate on the "artist" field.
func ArtistGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldArtist, v))
}

// ArtistGTE applies the GTE predicate on the "artist" field.
func ArtistGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldArtist, v))
}

// ArtistLT applies the LT predicate on the "artist" field.
func ArtistLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldArtist, v))
}

// ArtistLTE applies the LTE predicate on the "artist" field.
func ArtistLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldArtist, v))
}

// ArtistContains applies the Contains predicate on the "artist" field.
func ArtistContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldArtist, v))
}

// ArtistHasPrefix applies the HasPrefix predicate on the "artist" field.
func ArtistHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldArtist, v))
}

// ArtistHasSuffix applies the HasSuffix predicate on the "artist" field.
func ArtistHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldArtist, v))
}

// ArtistEqualFold applies the EqualFold predicate on the "artist" field.
func ArtistEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldArtist, v))
}

// ArtistContainsFold applies the ContainsFold predicate on the "artist" field.
func ArtistContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldArtist, v))
}

// ComposerEQ applies the EQ predicate on the "composer" field.
func ComposerEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldComposer, v))
}

// ComposerNEQ applies the NEQ predicate on the "composer" field.
func ComposerNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldComposer, v))
}

// ComposerIn applies the In predicate on the "composer" field.
func ComposerIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldComposer, vs...))
}

// ComposerNotIn applies the NotIn predicate on the "composer" field.
func ComposerNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldComposer, vs...))
}

// ComposerGT applies the GT predicate on the "composer" field.
func ComposerGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldComposer, v))
}

// ComposerGTE applies the GTE predicate on the "composer" field.
func ComposerGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldComposer, v))
}

// ComposerLT applies the LT predicate on the "composer" field.
func ComposerLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldComposer, v))
}

// ComposerLTE applies the LTE predicate on the "composer" field.
func ComposerLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldComposer, v))
}

// ComposerContains applies the Contains predicate on the "composer" field.
func ComposerContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldComposer, v))
}

// ComposerHasPrefix applies the HasPrefix predicate on the "composer" field.
func ComposerHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldComposer, v))
}

// ComposerHasSuffix applies the HasSuffix predicate on the "composer" field.
func ComposerHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldComposer, v))
}

// ComposerIsNil applies the IsNil predicate on the "composer" field.
func ComposerIsNil() predicate.Music {
	return predicate.Music(sql.FieldIsNull(FieldComposer))
}

// ComposerNotNil applies the NotNil predicate on the "composer" field.
func ComposerNotNil() predicate.Music {
	return predicate.Music(sql.FieldNotNull(FieldComposer))
}

// ComposerEqualFold applies the EqualFold predicate on the "composer" field.
func ComposerEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldComposer, v))
}

// ComposerContainsFold applies the ContainsFold predicate on the "composer" field.
func ComposerContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldComposer, v))
}

// MusicSourceEQ applies the EQ predicate on the "music_source" field.
func MusicSourceEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldMusicSource, v))
}

// MusicSourceNEQ applies the NEQ predicate on the "music_source" field.
func MusicSourceNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldMusicSource, v))
}

// MusicSourceIn applies the In predicate on the "music_source" field.
func MusicSourceIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldMusicSource, vs...))
}

// MusicSourceNotIn applies the NotIn predicate on the "music_source" field.
func MusicSourceNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldMusicSource, vs...))
}

// MusicSourceGT applies the GT predicate on the "music_source" field.
func MusicSourceGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldMusicSource, v))
}

// MusicSourceGTE applies the GTE predicate on the "music_source" field.
func MusicSourceGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldMusicSource, v))
}

// MusicSourceLT applies the LT predicate on the "music_source" field.
func MusicSourceLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldMusicSource, v))
}

// MusicSourceLTE applies the LTE predicate on the "music_source" field.
func MusicSourceLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldMusicSource, v))
}

// MusicSourceContains applies the Contains predicate on the "music_source" field.
func MusicSourceContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldMusicSource, v))
}

// MusicSourceHasPrefix applies the HasPrefix predicate on the "music_source" field.
func MusicSourceHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldMusicSource, v))
}

// MusicSourceHasSuffix applies the HasSuffix predicate on the "music_source" field.
func MusicSourceHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldMusicSource, v))
}

// MusicSourceEqualFold applies the EqualFold predicate on the "music_source" field.
func MusicSourceEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldMusicSource, v))
}

// MusicSourceContainsFold applies the ContainsFold predicate on the "music_source" field.
func MusicSourceContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldMusicSource, v))
}

// JacketSourceEQ applies the EQ predicate on the "jacket_source" field.
func JacketSourceEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldJacketSource, v))
}

// JacketSourceNEQ applies the NEQ predicate on the "jacket_source" field.
func JacketSourceNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldJacketSource, v))
}

// JacketSourceIn applies the In predicate on the "jacket_source" field.
func JacketSourceIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldJacketSource, vs...))
}

// JacketSourceNotIn applies the NotIn predicate on the "jacket_source" field.
func JacketSourceNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldJacketSource, vs...))
}

// JacketSourceGT applies the GT predicate on the "jacket_source" field.
func JacketSourceGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldJacketSource, v))
}

// JacketSourceGTE applies the GTE predicate on the "jacket_source" field.
func JacketSourceGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldJacketSource, v))
}

// JacketSourceLT applies the LT predicate on the "jacket_source" field.
func JacketSourceLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldJacketSource, v))
}

// JacketSourceLTE applies the LTE predicate on the "jacket_source" field.
func JacketSourceLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldJacketSource, v))
}

// JacketSourceContains applies the Contains predicate on the "jacket_source" field.
func JacketSourceContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldJacketSource, v))
}

// JacketSourceHasPrefix applies the HasPrefix predicate on the "jacket_source" field.
func JacketSourceHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldJacketSource, v))
}

// JacketSourceHasSuffix applies the HasSuffix predicate on the "jacket_source" field.
func JacketSourceHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldJacketSource, v))
}

// JacketSourceEqualFold applies the EqualFold predicate on the "jacket_source" field.
func JacketSourceEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldJacketSource, v))
}

// JacketSourceContainsFold applies the ContainsFold predicate on the "jacket_source" field.
func JacketSourceContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldJacketSource, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v float64) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v float64) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...float64) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...float64) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v float64) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v float64) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v float64) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v float64) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldDuration, v))
}

// BpmEQ applies the EQ predicate on the "bpm" field.
func BpmEQ(v float64) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldBpm, v))
}

// BpmNEQ applies the NEQ predicate on the "bpm" field.
func BpmNEQ(v float64) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldBpm, v))
}

// BpmIn applies the In predicate on the "bpm" field.
func BpmIn(vs ...float64) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldBpm, vs...))
}

// BpmNotIn applies the NotIn predicate on the "bpm" field.
func BpmNotIn(vs ...float64) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldBpm, vs...))
}

// BpmGT applies the GT predicate on the "bpm" field.
func BpmGT(v float64) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldBpm, v))
}

// BpmGTE applies the GTE predicate on the "bpm" field.
func BpmGTE(v float64) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldBpm, v))
}

// BpmLT applies the LT predicate on the "bpm" field.
func BpmLT(v float64) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldBpm, v))
}

// BpmLTE applies the LTE predicate on the "bpm" field.
func BpmLTE(v float64) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldBpm, v))
}

// GenreEQ applies the EQ predicate on the "genre" field.
func GenreEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldGenre, v))
}

// GenreNEQ applies the NEQ predicate on the "genre" field.
func GenreNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldGenre, v))
}

// GenreIn applies the In predicate on the "genre" field.
func GenreIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldGenre, vs...))
}

// GenreNotIn applies the NotIn predicate on the "genre" field.
func GenreNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldGenre, vs...))
}

// GenreGT applies the GT predicate on the "genre" field.
func GenreGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldGenre, v))
}

// GenreGTE applies the GTE predicate on the "genre" field.
func GenreGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldGenre, v))
}

// GenreLT applies the LT predicate on the "genre" field.
func GenreLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldGenre, v))
}

// GenreLTE applies the LTE predicate on the "genre" field.
func GenreLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldGenre, v))
}

// GenreContains applies the Contains predicate on the "genre" field.
func GenreContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldGenre, v))
}

// GenreHasPrefix applies the HasPrefix predicate on the "genre" field.
func GenreHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldGenre, v))
}

// GenreHasSuffix applies the HasSuffix predicate on the "genre" field.
func GenreHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldGenre, v))
}

// GenreIsNil applies the IsNil predicate on the "genre" field.
func GenreIsNil() predicate.Music {
	return predicate.Music(sql.FieldIsNull(FieldGenre))
}

// GenreNotNil applies the NotNil predicate on the "genre" field.
func GenreNotNil() predicate.Music {
	return predicate.Music(sql.FieldNotNull(FieldGenre))
}

// GenreEqualFold applies the EqualFold predicate on the "genre" field.
func GenreEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldGenre, v))
}

// GenreContainsFold applies the ContainsFold predicate on the "genre" field.
func GenreContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldGenre, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Music {
	return predicate.Music(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Music {
	return predicate.Music(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Music {
	return predicate.Music(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Music {
	return predicate.Music(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Music {
	return predicate.Music(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Music {
	return predicate.Music(sql.FieldContainsFold(FieldDescription, v))
}

// IsFeaturedEQ applies the EQ predicate on the "is_featured" field.
func IsFeaturedEQ(v bool) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldIsFeatured, v))
}

// IsFeaturedNEQ applies the NEQ predicate on the "is_featured" field.
func IsFeaturedNEQ(v bool) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldIsFeatured, v))
}

// IsFreeEQ applies the EQ predicate on the "is_free" field.
func IsFreeEQ(v bool) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldIsFree, v))
}

// IsFreeNEQ applies the NEQ predicate on the "is_free" field.
func IsFreeNEQ(v bool) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldIsFree, v))
}

// UnlockLevelEQ applies the EQ predicate on the "unlock_level" field.
func UnlockLevelEQ(v int) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldUnlockLevel, v))
}

// UnlockLevelNEQ applies the NEQ predicate on the "unlock_level" field.
func UnlockLevelNEQ(v int) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldUnlockLevel, v))
}

// UnlockLevelIn applies the In predicate on the "unlock_level" field.
func UnlockLevelIn(vs ...int) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldUnlockLevel, vs...))
}

// UnlockLevelNotIn applies the NotIn predicate on the "unlock_level" field.
func UnlockLevelNotIn(vs ...int) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldUnlockLevel, vs...))
}

// UnlockLevelGT applies the GT predicate on the "unlock_level" field.
func UnlockLevelGT(v int) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldUnlockLevel, v))
}

// UnlockLevelGTE applies the GTE predicate on the "unlock_level" field.
func UnlockLevelGTE(v int) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldUnlockLevel, v))
}

// UnlockLevelLT applies the LT predicate on the "unlock_level" field.
func UnlockLevelLT(v int) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldUnlockLevel, v))
}

// UnlockLevelLTE applies the LTE predicate on the "unlock_level" field.
func UnlockLevelLTE(v int) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldUnlockLevel, v))
}

// ReleaseDateEQ applies the EQ predicate on the "release_date" field.
func ReleaseDateEQ(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldReleaseDate, v))
}

// ReleaseDateNEQ applies the NEQ predicate on the "release_date" field.
func ReleaseDateNEQ(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldReleaseDate, v))
}

// ReleaseDateIn applies the In predicate on the "release_date" field.
func ReleaseDateIn(vs ...time.Time) predicate.Music {
	return predicate.Music(sql.FieldIn(FieldReleaseDate, vs...))
}

// ReleaseDateNotIn applies the NotIn predicate on the "release_date" field.
func ReleaseDateNotIn(vs ...time.Time) predicate.Music {
	return predicate.Music(sql.FieldNotIn(FieldReleaseDate, vs...))
}

// ReleaseDateGT applies the GT predicate on the "release_date" field.
func ReleaseDateGT(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldGT(FieldReleaseDate, v))
}

// ReleaseDateGTE applies the GTE predicate on the "release_date" field.
func ReleaseDateGTE(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldGTE(FieldReleaseDate, v))
}

// ReleaseDateLT applies the LT predicate on the "release_date" field.
func ReleaseDateLT(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldLT(FieldReleaseDate, v))
}

// ReleaseDateLTE applies the LTE predicate on the "release_date" field.
func ReleaseDateLTE(v time.Time) predicate.Music {
	return predicate.Music(sql.FieldLTE(FieldReleaseDate, v))
}

// ReleaseDateIsNil applies the IsNil predicate on the "release_date" field.
func ReleaseDateIsNil() predicate.Music {
	return predicate.Music(sql.FieldIsNull(FieldReleaseDate))
}

// ReleaseDateNotNil applies the NotNil predicate on the "release_date" field.
func ReleaseDateNotNil() predicate.Music {
	return predicate.Music(sql.FieldNotNull(FieldReleaseDate))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Music {
	return predicate.Music(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Music {
	return predicate.Music(sql.FieldNEQ(FieldIsActive, v))
}

// HasStages applies the HasEdge predicate on the "stages" edge.
func HasStages() predicate.Music {
	return predicate.Music(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StagesTable, StagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStagesWith applies the HasEdge predicate on the "stages" edge with a given conditions (other predicates).
func HasStagesWith(preds ...predicate.Stage) predicate.Music {
	return predicate.Music(func(s *sql.Selector) {
		step := newStagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRecords applies the HasEdge predicate on the "records" edge.
func HasRecords() predicate.Music {
	return predicate.Music(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RecordsTable, RecordsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordsWith applies the HasEdge predicate on the "records" edge with a given conditions (other predicates).
func HasRecordsWith(preds ...predicate.Record) predicate.Music {
	return predicate.Music(func(s *sql.Selector) {
		step := newRecordsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Music) predicate.Music {
	return predicate.Music(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Music) predicate.Music {
	return predicate.Music(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Music) predicate.Music {
	return predicate.Music(sql.NotPredicates(p))
}
