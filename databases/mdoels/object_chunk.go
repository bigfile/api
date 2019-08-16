//   Copyright 2019 The bigfile Authors. All rights reserved.
//   Use of this source code is governed by a MIT-style
//   license that can be found in the LICENSE file.

package models

import "time"

// ObjectChunk is a middle table, that associates chunk and object
type ObjectChunk struct {
	ID        uint64    `gorm:"type:BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT;primary_key"`
	ObjectID  uint64    `gorm:"type:BIGINT(20) UNSIGNED NOT NULL;column:objectId"`
	ChunkID   uint64    `gorm:"type:BIGINT(20) UNSIGNED NOT NULL;column:chunkId"`
	Number    int       `gorm:"type:int;column:number"`
	HashState *string   `gorm:"type:CHAR(64) NOT NULL;UNIQUE;column:hashState"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP(6) NOT NULL;DEFAULT:CURRENT_TIMESTAMP(6);column:createdAt"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP(6) NOT NULL;DEFAULT:CURRENT_TIMESTAMP(6);column:updatedAt"`
}

// TableName represent the db table name
func (oc ObjectChunk) TableName() string {
	return "object_chunk"
}
