// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package testlib

import (
	"github.com/mattermost/mattermost-server/store"
	"github.com/mattermost/mattermost-server/store/sqlstore"
	"github.com/mattermost/mattermost-server/store/storetest"
)

type TestStore struct {
	store.Store
}

func (*TestStore) Close() {
}

func NewTestStore() *TestStore {
	settings := storetest.MySQLSettings()
	return &TestStore{store.NewLayeredStore(sqlstore.NewSqlSupplier(*settings, nil), nil, nil)}
}
