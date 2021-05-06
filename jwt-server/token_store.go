package main

import (
	"context"
	"fmt"

	"github.com/go-oauth2/oauth2/v4"
)

func NewDummyTokenStore() (oauth2.TokenStore, error) {
	return &TokenStore{}, nil
}

// TokenStore token storage based on buntdb(https://github.com/tidwall/buntdb)
type TokenStore struct {
}

// Create create and store the new token information
func (ts *TokenStore) Create(ctx context.Context, info oauth2.TokenInfo) error {
	fmt.Println("Create", info)
	return nil
}

// remove key
func (ts *TokenStore) remove(key string) error {
	fmt.Println("remove", key)
	return nil
}

// RemoveByCode use the authorization code to delete the token information
func (ts *TokenStore) RemoveByCode(ctx context.Context, code string) error {
	return ts.remove(code)
}

// RemoveByAccess use the access token to delete the token information
func (ts *TokenStore) RemoveByAccess(ctx context.Context, access string) error {
	return ts.remove(access)
}

// RemoveByRefresh use the refresh token to delete the token information
func (ts *TokenStore) RemoveByRefresh(ctx context.Context, refresh string) error {
	return ts.remove(refresh)
}

func (ts *TokenStore) getData(key string) (oauth2.TokenInfo, error) {
	fmt.Println("getData", key)
	var ti oauth2.TokenInfo
	return ti, nil
}

func (ts *TokenStore) getBasicID(key string) (string, error) {
	var basicID string
	return basicID, nil
}

// GetByCode use the authorization code for token information data
func (ts *TokenStore) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	return ts.getData(code)
}

// GetByAccess use the access token for token information data
func (ts *TokenStore) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	basicID, err := ts.getBasicID(access)
	if err != nil {
		return nil, err
	}
	return ts.getData(basicID)
}

// GetByRefresh use the refresh token for token information data
func (ts *TokenStore) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	basicID, err := ts.getBasicID(refresh)
	if err != nil {
		return nil, err
	}
	return ts.getData(basicID)
}
