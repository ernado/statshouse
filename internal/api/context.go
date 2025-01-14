// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/vkcom/statshouse/internal/util"
)

type contextKey int

const (
	debugQueriesContextKey contextKey = iota
	accessInfoContextKey
	endpointStatContextKey
)

func debugQueriesContext(ctx context.Context, queries *[]string) context.Context {
	return context.WithValue(ctx, debugQueriesContextKey, queries)
}

func saveDebugQuery(ctx context.Context, query string) {
	p, ok := ctx.Value(debugQueriesContextKey).(*[]string)
	if ok {
		query = strings.TrimSpace(strings.ReplaceAll(query, "\n", " "))
		*p = append(*p, query)
	}
}

func withAccessInfo(ctx context.Context, ai *accessInfo) context.Context {
	return context.WithValue(ctx, accessInfoContextKey, ai)
}

func getAccessInfo(ctx context.Context) *accessInfo {
	if ai, ok := ctx.Value(accessInfoContextKey).(*accessInfo); ok {
		return ai
	}
	return nil
}

func withHTTPEndpointStat(ctx context.Context, es *endpointStat) context.Context {
	return context.WithValue(ctx, endpointStatContextKey, es)
}

func withRPCEndpointStat(ctx context.Context, ms *rpcMethodStat) context.Context {
	return context.WithValue(ctx, endpointStatContextKey, ms)
}

func endpointStatSetQueryKind(ctx context.Context, isFast, isLight bool) {
	switch s := ctx.Value(endpointStatContextKey).(type) {
	case *endpointStat:
		s.lane = strconv.Itoa(util.QueryKind(isFast, isLight))
	case *rpcMethodStat:
		s.lane = strconv.Itoa(util.QueryKind(isFast, isLight))
	}
}
