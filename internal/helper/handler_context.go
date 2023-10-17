/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import "context"

type (
	Caller           string
	contextMethodKey struct{}
)

const (
	DataRead     Caller = "DataRead"
	RefreshState Caller = "RefreshState"
	CreateState  Caller = "CreateState"
	UpdateState  Caller = "UpdateState"
	DeleteState  Caller = "DeleteState"
)

func GetContextWithCaller(ctx context.Context, caller Caller) context.Context {
	return context.WithValue(ctx, contextMethodKey{}, caller)
}

func IsContextCallerSet(ctx context.Context) bool {
	return ctx.Value(contextMethodKey{}).(Caller) != ""
}

func IsDataRead(ctx context.Context) bool {
	return ctx.Value(contextMethodKey{}) == DataRead
}

func IsRefreshState(ctx context.Context) bool {
	return ctx.Value(contextMethodKey{}) == RefreshState
}

func IsCreateState(ctx context.Context) bool {
	return ctx.Value(contextMethodKey{}) == CreateState
}

func IsUpdateState(ctx context.Context) bool {
	return ctx.Value(contextMethodKey{}) == UpdateState
}

func IsDeleteState(ctx context.Context) bool {
	return ctx.Value(contextMethodKey{}) == DeleteState
}
