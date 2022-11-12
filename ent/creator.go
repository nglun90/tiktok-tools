// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/4meepo/tiktok-tools/ent/creator"
)

// Creator is the model entity for the Creator schema.
type Creator struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 先知网ID
	Xzid string `json:"xzid,omitempty"`
	// 唯一ID
	UniqueID string `json:"unique_id,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty"`
	// 地区
	Region string `json:"region,omitempty"`
	// 粉丝数
	FollowerNum uint32 `json:"follower_num,omitempty"`
	// tiktok creator_id
	CreatorID string `json:"creator_id,omitempty"`
	// tiktok shop creator oecuid
	CreatorOecuid string `json:"creator_oecuid,omitempty"`
	// 先知侧商品类目
	Cate1NameCn []string `json:"cate1_name_cn,omitempty"`
	// tiktok侧商品类目
	TiktokCategory []string `json:"tiktok_category,omitempty"`
	// 先知网爬的email
	Email string `json:"email,omitempty"`
	// 先知网爬的whatsapp
	Whatsapp string `json:"whatsapp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Creator) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case creator.FieldCate1NameCn, creator.FieldTiktokCategory:
			values[i] = new([]byte)
		case creator.FieldID, creator.FieldFollowerNum:
			values[i] = new(sql.NullInt64)
		case creator.FieldXzid, creator.FieldUniqueID, creator.FieldNickName, creator.FieldRegion, creator.FieldCreatorID, creator.FieldCreatorOecuid, creator.FieldEmail, creator.FieldWhatsapp:
			values[i] = new(sql.NullString)
		case creator.FieldCreateTime, creator.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Creator", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Creator fields.
func (c *Creator) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case creator.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case creator.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case creator.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case creator.FieldXzid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field xzid", values[i])
			} else if value.Valid {
				c.Xzid = value.String
			}
		case creator.FieldUniqueID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unique_id", values[i])
			} else if value.Valid {
				c.UniqueID = value.String
			}
		case creator.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				c.NickName = value.String
			}
		case creator.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				c.Region = value.String
			}
		case creator.FieldFollowerNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follower_num", values[i])
			} else if value.Valid {
				c.FollowerNum = uint32(value.Int64)
			}
		case creator.FieldCreatorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				c.CreatorID = value.String
			}
		case creator.FieldCreatorOecuid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_oecuid", values[i])
			} else if value.Valid {
				c.CreatorOecuid = value.String
			}
		case creator.FieldCate1NameCn:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cate1_name_cn", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Cate1NameCn); err != nil {
					return fmt.Errorf("unmarshal field cate1_name_cn: %w", err)
				}
			}
		case creator.FieldTiktokCategory:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tiktok_category", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.TiktokCategory); err != nil {
					return fmt.Errorf("unmarshal field tiktok_category: %w", err)
				}
			}
		case creator.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case creator.FieldWhatsapp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field whatsapp", values[i])
			} else if value.Valid {
				c.Whatsapp = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Creator.
// Note that you need to call Creator.Unwrap() before calling this method if this Creator
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Creator) Update() *CreatorUpdateOne {
	return (&CreatorClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Creator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Creator) Unwrap() *Creator {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Creator is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Creator) String() string {
	var builder strings.Builder
	builder.WriteString("Creator(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("xzid=")
	builder.WriteString(c.Xzid)
	builder.WriteString(", ")
	builder.WriteString("unique_id=")
	builder.WriteString(c.UniqueID)
	builder.WriteString(", ")
	builder.WriteString("nick_name=")
	builder.WriteString(c.NickName)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(c.Region)
	builder.WriteString(", ")
	builder.WriteString("follower_num=")
	builder.WriteString(fmt.Sprintf("%v", c.FollowerNum))
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(c.CreatorID)
	builder.WriteString(", ")
	builder.WriteString("creator_oecuid=")
	builder.WriteString(c.CreatorOecuid)
	builder.WriteString(", ")
	builder.WriteString("cate1_name_cn=")
	builder.WriteString(fmt.Sprintf("%v", c.Cate1NameCn))
	builder.WriteString(", ")
	builder.WriteString("tiktok_category=")
	builder.WriteString(fmt.Sprintf("%v", c.TiktokCategory))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("whatsapp=")
	builder.WriteString(c.Whatsapp)
	builder.WriteByte(')')
	return builder.String()
}

// Creators is a parsable slice of Creator.
type Creators []*Creator

func (c Creators) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
