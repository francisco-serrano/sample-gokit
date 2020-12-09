// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrations)
	t.Run("Messages", testMessages)
}

func TestDelete(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsDelete)
	t.Run("Messages", testMessagesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
	t.Run("Messages", testMessagesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
	t.Run("Messages", testMessagesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsExists)
	t.Run("Messages", testMessagesExists)
}

func TestFind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsFind)
	t.Run("Messages", testMessagesFind)
}

func TestBind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsBind)
	t.Run("Messages", testMessagesBind)
}

func TestOne(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsOne)
	t.Run("Messages", testMessagesOne)
}

func TestAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsAll)
	t.Run("Messages", testMessagesAll)
}

func TestCount(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsCount)
	t.Run("Messages", testMessagesCount)
}

func TestHooks(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsHooks)
	t.Run("Messages", testMessagesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
	t.Run("Messages", testMessagesInsert)
	t.Run("Messages", testMessagesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReload)
	t.Run("Messages", testMessagesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
	t.Run("Messages", testMessagesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSelect)
	t.Run("Messages", testMessagesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
	t.Run("Messages", testMessagesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
	t.Run("Messages", testMessagesSliceUpdateAll)
}
