// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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

func GetContextCaller(ctx context.Context) any {
	return ctx.Value(contextMethodKey{})
}

func IsContextCallerSet(ctx context.Context) bool {
	return GetContextCaller(ctx) != nil
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
