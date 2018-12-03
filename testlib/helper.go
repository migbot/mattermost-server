// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package testlib

import (
	"os"
	"testing"

	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/store"
	"github.com/mattermost/mattermost-server/store/sqlstore"
	"github.com/mattermost/mattermost-server/store/storetest"
	"github.com/mattermost/mattermost-server/utils"
)

type MainHelper struct {
	Store            store.Store
	SqlSupplier      *sqlstore.SqlSupplier
	ClusterInterface *FakeClusterInterface

	status int
}

func NewMainHelper() *MainHelper {
	// Setup a global logger to catch tests logging outside of app context
	// The global logger will be stomped by apps initalizing but that's fine for testing.
	// Ideally this won't happen.
	mlog.InitGlobalLogger(mlog.NewLogger(&mlog.LoggerConfiguration{
		EnableConsole: true,
		ConsoleJson:   true,
		ConsoleLevel:  "error",
		EnableFile:    false,
	}))

	utils.TranslationsPreInit()

	settings := storetest.MySQLSettings()
	testClusterInterface := &FakeClusterInterface{}
	testStoreSqlSupplier := sqlstore.NewSqlSupplier(*settings, nil)
	testStore := &TestStore{store.NewLayeredStore(testStoreSqlSupplier, nil, testClusterInterface)}

	return &MainHelper{
		Store:            testStore,
		SqlSupplier:      testStoreSqlSupplier,
		ClusterInterface: testClusterInterface,
	}
}

func (h *MainHelper) Main(m *testing.M) {
	h.status = m.Run()
}

func (h *MainHelper) Close() error {
	os.Exit(h.status)

	return nil
}
