package common

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func SafeBasicAuthLoad(filename string, userOptions ...basicauth.UserAuthOption) (handler context.Handler, err error) {
	defer func() {
		result := recover()
		if result != nil {
			if panicErr, ok := result.(error); ok {
				err = fmt.Errorf("failed to create auth: %w", panicErr)
			} else {
				panic(result)
			}
		}
	}()

	opts := basicauth.Options{
		Realm: basicauth.DefaultRealm,
		Allow: basicauth.AllowUsersFile(filename, userOptions...),
	}
	handler = basicauth.New(opts)
	return handler, err
}

type CachedAuth struct {
	expirationDuration time.Duration
	savedCredentials   map[string]time.Time
}

func NewCachedAuth(expirationDuration time.Duration) *CachedAuth {
	return &CachedAuth{
		expirationDuration: expirationDuration,
		savedCredentials:   map[string]time.Time{},
	}
}

func (a *CachedAuth) Option(options *basicauth.UserAuthOptions) {
	cmpFunc := options.ComparePassword
	if cmpFunc == nil {
		return
	}

	options.ComparePassword = func(stored, userPassword string) bool {
		if stored == "" || userPassword == "" {
			return cmpFunc(stored, userPassword)
		}

		key := stored + ":" + userPassword
		loginTime, ok := a.savedCredentials[key]
		if !ok || time.Since(loginTime) >= a.expirationDuration {
			if cmpFunc(stored, userPassword) {
				a.savedCredentials[key] = time.Now()
				return true
			}

			delete(a.savedCredentials, key)
			return false
		}

		return true
	}
}
