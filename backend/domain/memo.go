package domain

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	ShortHash string // 短縮URL用ハッシュ
	Title     string // タイトル
	Content   string // 内容
	Pinned    bool   // ピン留め
	ViewCount uint   // 閲覧回数
	ParentID  uint   // 親メモのID
	Parent    *Memo  `gorm:"foreignKey:ParentID"` // 親メモ
	Children  []Memo `gorm:"foreignKey:ParentID"` // 子メモ
}

type Memos []Memo

func (m *Memo) NewMemo(title, content string) *Memo {
	return &Memo{
		Title:   title,
		Content: content,
	}
}

// validator for memo
func (m Memo) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.ShortHash, validation.Required),
	)
}

func (m *Memo) GenerateShortHash() {
	hash := sha256.Sum256([]byte(uuid.New().String()))
	m.ShortHash = hex.EncodeToString(hash[:])[:10]
}
