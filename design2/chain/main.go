// 责任链模式 是一种行为型设计模式，
// 其核心思想是 将多个处理对象连成一条链，请求沿链传递直到被处理或到达链尾。
// 每个处理对象决定是否处理请求，或将其传递给下一个处理者。
// 该模式解耦了请求的发送者和接收者，允许多个对象动态处理同一请求

// 责任链模式的使用场景
// 多级审批流程（如费用报销需多级领导审批）。
// 请求过滤（如 HTTP 中间件链：认证 → 日志 → 限流）。
// 异常处理链（如不同异常类型由不同处理器处理）。

package main

import (
	"fmt"
	"net/http"
)

// --------------------- 处理者接口 ---------------------
type Middleware interface {
	Handle(next http.HandlerFunc) http.HandlerFunc
}

// --------------------- 具体处理者 ---------------------
// 认证中间件
type AuthMiddleware struct{}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "认证失败")
			return
		}
		fmt.Println("认证通过")
		next(w, r) // 调用下一个中间件
	}
}

// 日志中间件
type LogMiddleware struct{}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("请求路径: %s\n", r.URL.Path)
		next(w, r)
	}
}

// 限流中间件
type RateLimitMiddleware struct{}

func (m *RateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 模拟限流检查
		if r.URL.Path == "/api/limited" {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, "请求频率过高")
			return
		}
		next(w, r)
	}
}

// --------------------- 客户端：组装处理链 ---------------------
func Chain(middlewares ...Middleware) http.HandlerFunc {
	// 最终的处理器（实际业务逻辑）
	finalHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "业务处理完成")
	}

	// 反向组装中间件链：限流 → 日志 → 认证 → 业务逻辑
	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i].Handle(finalHandler)
	}

	return finalHandler
}

func main() {
	// 定义中间件链
	handler := Chain(
		&AuthMiddleware{},
		&LogMiddleware{},
		&RateLimitMiddleware{},
	)

	// 模拟 HTTP 请求
	req1 := &http.Request{
		Header: http.Header{"Authorization": []string{"valid-token"}},
		// URL:    &http.URL{Path: "/api/data"},
	}
	req2 := &http.Request{
		Header: http.Header{"Authorization": []string{"invalid-token"}},
		// URL:    &http.URL{Path: "/api/limited"},
	}

	// 测试请求1（正常流程）
	fmt.Println("请求1结果:")
	handler(nil, req1)

	// 测试请求2（认证失败）
	fmt.Println("\n请求2结果:")
	handler(nil, req2)
}
