package utils

import (
	"os"
	"strings"
)

// HostEnv 运行的环境标识，通过 env.Current 获取当前的环境标识。
//
// ⚠️ 请勿将普通 string 显式或隐式地转换到此类型上，
// 应总是通过 env.Parse 方法对标识进行解释，
// 或通过 env.Current 获取当前环境标识。
type HostEnv string

// 已知环境标识
const (
	_UNKNOWN HostEnv = ""

	LOCAL       HostEnv = "local"       // 本地环境
	DEVELOP     HostEnv = "develop"     // 开发环境
	TEST        HostEnv = "test"        // 测试环境
	SANDBOX     HostEnv = "sandbox"     // 沙盒环境
	INTEGRATION HostEnv = "integration" // 集成环境
	PRESSURE    HostEnv = "pressure"    // 压测环境
	PREVIEW     HostEnv = "preview"     // 预发环境
	PRODUCT     HostEnv = "product"     // 生产环境
)

// Kind 运行环境类型。
type Kind string

// 已知环境类型
const (
	_KIND_UNKNOWN Kind = ""
	KIND_LOCAL    Kind = "local"  // 本地环境
	KIND_SERVER   Kind = "server" // 服务器环境
)

var _SERVER_ENV_MAPPING = map[HostEnv]string{
	DEVELOP:     "dev",
	TEST:        "test",
	SANDBOX:     "sandbox",
	INTEGRATION: "integrate",
	PRESSURE:    "press",
	PREVIEW:     "pre",
	PRODUCT:     "prod",
}

var _CURRENT_ENV = func() HostEnv {
	val, exist := os.LookupEnv("MY_ENV_NAME")
	if !exist {
		return LOCAL
	}
	for k, v := range _SERVER_ENV_MAPPING {
		if v == val {
			return k
		}
	}
	return _UNKNOWN
}()

// Kind 环境类型。
func (this HostEnv) Kind() Kind {
	if this == LOCAL {
		return KIND_LOCAL
	} else if _, ok := _SERVER_ENV_MAPPING[this]; ok {
		return KIND_SERVER
	} else {
		return _KIND_UNKNOWN
	}
}

// IsLocal 是否本地环境。
func (this HostEnv) IsLocal() bool {
	return this == LOCAL
}

// IsServer 是否服务器环境。
func (this HostEnv) IsServer() bool {
	_, ok := _SERVER_ENV_MAPPING[this]
	return ok
}

// IsValid 是否一个有效的环境标识。
func (this HostEnv) IsValid() bool {
	return this.IsLocal() || this.IsServer()
}

// MustValid 保证当前标识必须是一个有效的环境标识，否则 panic。
func (this HostEnv) MustValid() HostEnv {
	if !this.IsValid() {
		panic("invalid env: " + this)
	}
	return this
}

// StdName 返回环境标准名称
func (this HostEnv) StdName() string {
	if this == LOCAL {
		return string(this)
	} else if _, ok := _SERVER_ENV_MAPPING[this]; ok {
		return string(this)
	} else {
		return string(_UNKNOWN)
	}
}

// Parse 环境名称，如 env 不是一个有效的名称时返回的 HostEnv 的 IsUnknown == true。
func Parse(env string) HostEnv {
	env = strings.ToLower(env)
	if env == string(LOCAL) {
		return LOCAL
	}
	for k, v := range _SERVER_ENV_MAPPING {
		if string(k) == env || v == env {
			return k
		}
	}
	return _UNKNOWN
}

// Current 获取当前环境。
func Current() HostEnv {
	return _CURRENT_ENV
}

// IsLocal 判断当前环境 Current 的类型是否本地环境 KIND_LOCAL。
func IsLocal() bool {
	return Current().IsLocal()
}

// IsServer 判断当前环境 Current 的类型是否服务器环境 KIND_SERVER。
func IsServer() bool {
	return Current().IsServer()
}

// IsValid 判断当前环境 Current 是否有效。
func IsValid() bool {
	return Current().IsValid()
}

// MustValid 保证当前环境 Current 必须是一个有效的环境标识，否则 panic。
func MustValid() HostEnv {
	if current := Current(); current.IsValid() {
		return current
	} else {
		panic("invalid current env: " + Current())
	}
}

// AllKnownEnvs 所有已知的环境标识
func AllKnownEnvs() []HostEnv {
	return []HostEnv{
		LOCAL,
		DEVELOP,
		TEST,
		SANDBOX,
		INTEGRATION,
		PRESSURE,
		PREVIEW,
		PRODUCT,
	}
}

// AllKnownEnvNames 所有已知的环境标识标准名称
func AllKnownEnvNames() []string {
	return []string{
		LOCAL.StdName(),
		DEVELOP.StdName(),
		TEST.StdName(),
		SANDBOX.StdName(),
		INTEGRATION.StdName(),
		PRESSURE.StdName(),
		PREVIEW.StdName(),
		PRODUCT.StdName(),
	}
}
