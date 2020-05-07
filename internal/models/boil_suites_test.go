// Code generated by SQLBoiler 4.1.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("AccessTokens", testAccessTokens)
	t.Run("AppUserProfiles", testAppUserProfiles)
	t.Run("PasswordResetTokens", testPasswordResetTokens)
	t.Run("RefreshTokens", testRefreshTokens)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensDelete)
	t.Run("AppUserProfiles", testAppUserProfilesDelete)
	t.Run("PasswordResetTokens", testPasswordResetTokensDelete)
	t.Run("RefreshTokens", testRefreshTokensDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensQueryDeleteAll)
	t.Run("AppUserProfiles", testAppUserProfilesQueryDeleteAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensQueryDeleteAll)
	t.Run("RefreshTokens", testRefreshTokensQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensSliceDeleteAll)
	t.Run("AppUserProfiles", testAppUserProfilesSliceDeleteAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensSliceDeleteAll)
	t.Run("RefreshTokens", testRefreshTokensSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensExists)
	t.Run("AppUserProfiles", testAppUserProfilesExists)
	t.Run("PasswordResetTokens", testPasswordResetTokensExists)
	t.Run("RefreshTokens", testRefreshTokensExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensFind)
	t.Run("AppUserProfiles", testAppUserProfilesFind)
	t.Run("PasswordResetTokens", testPasswordResetTokensFind)
	t.Run("RefreshTokens", testRefreshTokensFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensBind)
	t.Run("AppUserProfiles", testAppUserProfilesBind)
	t.Run("PasswordResetTokens", testPasswordResetTokensBind)
	t.Run("RefreshTokens", testRefreshTokensBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensOne)
	t.Run("AppUserProfiles", testAppUserProfilesOne)
	t.Run("PasswordResetTokens", testPasswordResetTokensOne)
	t.Run("RefreshTokens", testRefreshTokensOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensAll)
	t.Run("AppUserProfiles", testAppUserProfilesAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensAll)
	t.Run("RefreshTokens", testRefreshTokensAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensCount)
	t.Run("AppUserProfiles", testAppUserProfilesCount)
	t.Run("PasswordResetTokens", testPasswordResetTokensCount)
	t.Run("RefreshTokens", testRefreshTokensCount)
	t.Run("Users", testUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensInsert)
	t.Run("AccessTokens", testAccessTokensInsertWhitelist)
	t.Run("AppUserProfiles", testAppUserProfilesInsert)
	t.Run("AppUserProfiles", testAppUserProfilesInsertWhitelist)
	t.Run("PasswordResetTokens", testPasswordResetTokensInsert)
	t.Run("PasswordResetTokens", testPasswordResetTokensInsertWhitelist)
	t.Run("RefreshTokens", testRefreshTokensInsert)
	t.Run("RefreshTokens", testRefreshTokensInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AccessTokenToUserUsingUser", testAccessTokenToOneUserUsingUser)
	t.Run("AppUserProfileToUserUsingUser", testAppUserProfileToOneUserUsingUser)
	t.Run("PasswordResetTokenToUserUsingUser", testPasswordResetTokenToOneUserUsingUser)
	t.Run("RefreshTokenToUserUsingUser", testRefreshTokenToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("UserToAppUserProfileUsingAppUserProfile", testUserOneToOneAppUserProfileUsingAppUserProfile)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UserToAccessTokens", testUserToManyAccessTokens)
	t.Run("UserToPasswordResetTokens", testUserToManyPasswordResetTokens)
	t.Run("UserToRefreshTokens", testUserToManyRefreshTokens)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AccessTokenToUserUsingAccessTokens", testAccessTokenToOneSetOpUserUsingUser)
	t.Run("AppUserProfileToUserUsingAppUserProfile", testAppUserProfileToOneSetOpUserUsingUser)
	t.Run("PasswordResetTokenToUserUsingPasswordResetTokens", testPasswordResetTokenToOneSetOpUserUsingUser)
	t.Run("RefreshTokenToUserUsingRefreshTokens", testRefreshTokenToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("UserToAppUserProfileUsingAppUserProfile", testUserOneToOneSetOpAppUserProfileUsingAppUserProfile)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UserToAccessTokens", testUserToManyAddOpAccessTokens)
	t.Run("UserToPasswordResetTokens", testUserToManyAddOpPasswordResetTokens)
	t.Run("UserToRefreshTokens", testUserToManyAddOpRefreshTokens)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensReload)
	t.Run("AppUserProfiles", testAppUserProfilesReload)
	t.Run("PasswordResetTokens", testPasswordResetTokensReload)
	t.Run("RefreshTokens", testRefreshTokensReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensReloadAll)
	t.Run("AppUserProfiles", testAppUserProfilesReloadAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensReloadAll)
	t.Run("RefreshTokens", testRefreshTokensReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensSelect)
	t.Run("AppUserProfiles", testAppUserProfilesSelect)
	t.Run("PasswordResetTokens", testPasswordResetTokensSelect)
	t.Run("RefreshTokens", testRefreshTokensSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensUpdate)
	t.Run("AppUserProfiles", testAppUserProfilesUpdate)
	t.Run("PasswordResetTokens", testPasswordResetTokensUpdate)
	t.Run("RefreshTokens", testRefreshTokensUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensSliceUpdateAll)
	t.Run("AppUserProfiles", testAppUserProfilesSliceUpdateAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensSliceUpdateAll)
	t.Run("RefreshTokens", testRefreshTokensSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
