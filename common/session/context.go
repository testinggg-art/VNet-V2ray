package session

import (
	"context"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/net"
)

type sessionKey int

const (
	idSessionKey sessionKey = iota
	connectionSessionKey
	inboundSessionKey
	destinationSessionKey
	outboundSessionKey
	contentSessionKey
	muxPreferedSessionKey
	proxyLimiterSessionKey
)

type Connection interface {
	net.Conn
}

// ContextWithID returns a new context with the given ID.
func ContextWithID(ctx context.Context, id ID) context.Context {
	return context.WithValue(ctx, idSessionKey, id)
}

// IDFromContext returns ID in this context, or 0 if not contained.
func IDFromContext(ctx context.Context) ID {
	if id, ok := ctx.Value(idSessionKey).(ID); ok {
		return id
	}
	return 0
}

func ContextWithInbound(ctx context.Context, inbound *Inbound) context.Context {
	return context.WithValue(ctx, inboundSessionKey, inbound)
}

func InboundFromContext(ctx context.Context) *Inbound {
	if inbound, ok := ctx.Value(inboundSessionKey).(*Inbound); ok {
		return inbound
	}
	return nil
}

func ContextWithOutbound(ctx context.Context, outbound *Outbound) context.Context {
	return context.WithValue(ctx, outboundSessionKey, outbound)
}

func OutboundFromContext(ctx context.Context) *Outbound {
	if outbound, ok := ctx.Value(outboundSessionKey).(*Outbound); ok {
		return outbound
	}
	return nil
}

func ContextWithContent(ctx context.Context, content *Content) context.Context {
	return context.WithValue(ctx, contentSessionKey, content)
}

func ContentFromContext(ctx context.Context) *Content {
	if content, ok := ctx.Value(contentSessionKey).(*Content); ok {
		return content
	}
	return nil
}

// ContextWithMuxPrefered returns a new context with the given bool
func ContextWithMuxPrefered(ctx context.Context, forced bool) context.Context {
	return context.WithValue(ctx, muxPreferedSessionKey, forced)
}

// MuxPreferedFromContext returns value in this context, or false if not contained.
func MuxPreferedFromContext(ctx context.Context) bool {
	if val, ok := ctx.Value(muxPreferedSessionKey).(bool); ok {
		return val
	}
	return false
}

// ContextWithMuxPrefered returns a new context with the given bool
func ContextWithProxyLimiter(ctx context.Context, limter *buf.ProxyLimiter) context.Context {
	return context.WithValue(ctx, proxyLimiterSessionKey, limter)
}

// MuxPreferedFromContext returns value in this context, or false if not contained.
func ProxyLimiterFromContext(ctx context.Context) *buf.ProxyLimiter {
	if val, ok := ctx.Value(proxyLimiterSessionKey).(*buf.ProxyLimiter); ok {
		return val
	}
	return nil
}
