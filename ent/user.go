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
	"github.com/witchs-lounge_backend/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	// Global custom UUID ID
	ID uuid.UUID `json:"id,omitempty"`
	// Created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Updated time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 플랫폼 타입
	PlatformType user.PlatformType `json:"platform_type,omitempty"`
	// 플랫폼에서의 유저 ID
	PlatformUserID string `json:"platform_user_id,omitempty"`
	// 플랫폼 이메일
	PlatformEmail string `json:"platform_email,omitempty"`
	// 플랫폼 프로필 이미지 URL
	PlatformAvatarURL string `json:"platform_avatar_url,omitempty"`
	// 플랫폼에서 표시되는 이름
	PlatformDisplayName string `json:"platform_display_name,omitempty"`
	// 선호 언어
	Language string `json:"language,omitempty"`
	// 플랫폼별 추가 데이터
	PlatformData map[string]interface{} `json:"platform_data,omitempty"`
	// 플랫폼 인증 여부
	IsVerified bool `json:"is_verified,omitempty"`
	// 게임 내 닉네임
	Nickname string `json:"nickname,omitempty"`
	// 게임 내 표시 이름
	DisplayName string `json:"display_name,omitempty"`
	// 마지막 로그인 시간
	LastLoginAt time.Time `json:"last_login_at,omitempty"`
	// 유저 레벨
	Level int `json:"level,omitempty"`
	// 경험치
	Exp int `json:"exp,omitempty"`
	// 게임 내 코인
	Coin int `json:"coin,omitempty"`
	// 프리미엄 재화
	Gem int `json:"gem,omitempty"`
	// 게임 설정
	Settings map[string]interface{} `json:"settings,omitempty"`
	// 커스터마이징 데이터
	CustomizeData map[string]interface{} `json:"customize_data,omitempty"`
	// 게임 저장 데이터
	SaveData map[string]interface{} `json:"save_data,omitempty"`
	// 밴 여부
	IsBanned bool `json:"is_banned,omitempty"`
	// 밴 해제 시간
	BannedUntil *time.Time `json:"banned_until,omitempty"`
	// 밴 사유
	BanReason string `json:"ban_reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// PurchasedProducts holds the value of the purchased_products edge.
	PurchasedProducts []*Product `json:"purchased_products,omitempty"`
	// Records holds the value of the records edge.
	Records []*Record `json:"records,omitempty"`
	// UserAchievements holds the value of the user_achievements edge.
	UserAchievements []*UserAchievement `json:"user_achievements,omitempty"`
	// UserPurchases holds the value of the user_purchases edge.
	UserPurchases []*UserPurchase `json:"user_purchases,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PurchasedProductsOrErr returns the PurchasedProducts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PurchasedProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.PurchasedProducts, nil
	}
	return nil, &NotLoadedError{edge: "purchased_products"}
}

// RecordsOrErr returns the Records value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RecordsOrErr() ([]*Record, error) {
	if e.loadedTypes[1] {
		return e.Records, nil
	}
	return nil, &NotLoadedError{edge: "records"}
}

// UserAchievementsOrErr returns the UserAchievements value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserAchievementsOrErr() ([]*UserAchievement, error) {
	if e.loadedTypes[2] {
		return e.UserAchievements, nil
	}
	return nil, &NotLoadedError{edge: "user_achievements"}
}

// UserPurchasesOrErr returns the UserPurchases value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserPurchasesOrErr() ([]*UserPurchase, error) {
	if e.loadedTypes[3] {
		return e.UserPurchases, nil
	}
	return nil, &NotLoadedError{edge: "user_purchases"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldPlatformData, user.FieldSettings, user.FieldCustomizeData, user.FieldSaveData:
			values[i] = new([]byte)
		case user.FieldIsVerified, user.FieldIsBanned:
			values[i] = new(sql.NullBool)
		case user.FieldLevel, user.FieldExp, user.FieldCoin, user.FieldGem:
			values[i] = new(sql.NullInt64)
		case user.FieldPlatformType, user.FieldPlatformUserID, user.FieldPlatformEmail, user.FieldPlatformAvatarURL, user.FieldPlatformDisplayName, user.FieldLanguage, user.FieldNickname, user.FieldDisplayName, user.FieldBanReason:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldLastLoginAt, user.FieldBannedUntil:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldPlatformType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform_type", values[i])
			} else if value.Valid {
				u.PlatformType = user.PlatformType(value.String)
			}
		case user.FieldPlatformUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform_user_id", values[i])
			} else if value.Valid {
				u.PlatformUserID = value.String
			}
		case user.FieldPlatformEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform_email", values[i])
			} else if value.Valid {
				u.PlatformEmail = value.String
			}
		case user.FieldPlatformAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform_avatar_url", values[i])
			} else if value.Valid {
				u.PlatformAvatarURL = value.String
			}
		case user.FieldPlatformDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform_display_name", values[i])
			} else if value.Valid {
				u.PlatformDisplayName = value.String
			}
		case user.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				u.Language = value.String
			}
		case user.FieldPlatformData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field platform_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.PlatformData); err != nil {
					return fmt.Errorf("unmarshal field platform_data: %w", err)
				}
			}
		case user.FieldIsVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_verified", values[i])
			} else if value.Valid {
				u.IsVerified = value.Bool
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldLastLoginAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_at", values[i])
			} else if value.Valid {
				u.LastLoginAt = value.Time
			}
		case user.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				u.Level = int(value.Int64)
			}
		case user.FieldExp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exp", values[i])
			} else if value.Valid {
				u.Exp = int(value.Int64)
			}
		case user.FieldCoin:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				u.Coin = int(value.Int64)
			}
		case user.FieldGem:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gem", values[i])
			} else if value.Valid {
				u.Gem = int(value.Int64)
			}
		case user.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		case user.FieldCustomizeData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field customize_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.CustomizeData); err != nil {
					return fmt.Errorf("unmarshal field customize_data: %w", err)
				}
			}
		case user.FieldSaveData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field save_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.SaveData); err != nil {
					return fmt.Errorf("unmarshal field save_data: %w", err)
				}
			}
		case user.FieldIsBanned:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_banned", values[i])
			} else if value.Valid {
				u.IsBanned = value.Bool
			}
		case user.FieldBannedUntil:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field banned_until", values[i])
			} else if value.Valid {
				u.BannedUntil = new(time.Time)
				*u.BannedUntil = value.Time
			}
		case user.FieldBanReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ban_reason", values[i])
			} else if value.Valid {
				u.BanReason = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryPurchasedProducts queries the "purchased_products" edge of the User entity.
func (u *User) QueryPurchasedProducts() *ProductQuery {
	return NewUserClient(u.config).QueryPurchasedProducts(u)
}

// QueryRecords queries the "records" edge of the User entity.
func (u *User) QueryRecords() *RecordQuery {
	return NewUserClient(u.config).QueryRecords(u)
}

// QueryUserAchievements queries the "user_achievements" edge of the User entity.
func (u *User) QueryUserAchievements() *UserAchievementQuery {
	return NewUserClient(u.config).QueryUserAchievements(u)
}

// QueryUserPurchases queries the "user_purchases" edge of the User entity.
func (u *User) QueryUserPurchases() *UserPurchaseQuery {
	return NewUserClient(u.config).QueryUserPurchases(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("platform_type=")
	builder.WriteString(fmt.Sprintf("%v", u.PlatformType))
	builder.WriteString(", ")
	builder.WriteString("platform_user_id=")
	builder.WriteString(u.PlatformUserID)
	builder.WriteString(", ")
	builder.WriteString("platform_email=")
	builder.WriteString(u.PlatformEmail)
	builder.WriteString(", ")
	builder.WriteString("platform_avatar_url=")
	builder.WriteString(u.PlatformAvatarURL)
	builder.WriteString(", ")
	builder.WriteString("platform_display_name=")
	builder.WriteString(u.PlatformDisplayName)
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(u.Language)
	builder.WriteString(", ")
	builder.WriteString("platform_data=")
	builder.WriteString(fmt.Sprintf("%v", u.PlatformData))
	builder.WriteString(", ")
	builder.WriteString("is_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.IsVerified))
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("last_login_at=")
	builder.WriteString(u.LastLoginAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", u.Level))
	builder.WriteString(", ")
	builder.WriteString("exp=")
	builder.WriteString(fmt.Sprintf("%v", u.Exp))
	builder.WriteString(", ")
	builder.WriteString("coin=")
	builder.WriteString(fmt.Sprintf("%v", u.Coin))
	builder.WriteString(", ")
	builder.WriteString("gem=")
	builder.WriteString(fmt.Sprintf("%v", u.Gem))
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(fmt.Sprintf("%v", u.Settings))
	builder.WriteString(", ")
	builder.WriteString("customize_data=")
	builder.WriteString(fmt.Sprintf("%v", u.CustomizeData))
	builder.WriteString(", ")
	builder.WriteString("save_data=")
	builder.WriteString(fmt.Sprintf("%v", u.SaveData))
	builder.WriteString(", ")
	builder.WriteString("is_banned=")
	builder.WriteString(fmt.Sprintf("%v", u.IsBanned))
	builder.WriteString(", ")
	if v := u.BannedUntil; v != nil {
		builder.WriteString("banned_until=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("ban_reason=")
	builder.WriteString(u.BanReason)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
