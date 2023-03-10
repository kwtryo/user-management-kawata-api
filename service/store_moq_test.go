// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"sync"
)

// DBConnectionMock is a mock implementation of store.DBConnection.
//
//	func TestSomethingThatUsesDBConnection(t *testing.T) {
//
//		// make and configure a mocked store.DBConnection
//		mockedDBConnection := &DBConnectionMock{
//			ExecContextFunc: func(ctx context.Context, query string, args ...any) (sql.Result, error) {
//				panic("mock out the ExecContext method")
//			},
//			GetContextFunc: func(ctx context.Context, dest interface{}, query string, args ...any) error {
//				panic("mock out the GetContext method")
//			},
//			NamedExecContextFunc: func(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
//				panic("mock out the NamedExecContext method")
//			},
//			PreparexContextFunc: func(ctx context.Context, query string) (*sqlx.Stmt, error) {
//				panic("mock out the PreparexContext method")
//			},
//			QueryRowContextFunc: func(ctx context.Context, query string, args ...any) *sql.Row {
//				panic("mock out the QueryRowContext method")
//			},
//			QueryxContextFunc: func(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
//				panic("mock out the QueryxContext method")
//			},
//			SelectContextFunc: func(ctx context.Context, dest interface{}, query string, args ...any) error {
//				panic("mock out the SelectContext method")
//			},
//		}
//
//		// use mockedDBConnection in code that requires store.DBConnection
//		// and then make assertions.
//
//	}
type DBConnectionMock struct {
	// ExecContextFunc mocks the ExecContext method.
	ExecContextFunc func(ctx context.Context, query string, args ...any) (sql.Result, error)

	// GetContextFunc mocks the GetContext method.
	GetContextFunc func(ctx context.Context, dest interface{}, query string, args ...any) error

	// NamedExecContextFunc mocks the NamedExecContext method.
	NamedExecContextFunc func(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	// PreparexContextFunc mocks the PreparexContext method.
	PreparexContextFunc func(ctx context.Context, query string) (*sqlx.Stmt, error)

	// QueryRowContextFunc mocks the QueryRowContext method.
	QueryRowContextFunc func(ctx context.Context, query string, args ...any) *sql.Row

	// QueryxContextFunc mocks the QueryxContext method.
	QueryxContextFunc func(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)

	// SelectContextFunc mocks the SelectContext method.
	SelectContextFunc func(ctx context.Context, dest interface{}, query string, args ...any) error

	// calls tracks calls to the methods.
	calls struct {
		// ExecContext holds details about calls to the ExecContext method.
		ExecContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// GetContext holds details about calls to the GetContext method.
		GetContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Dest is the dest argument value.
			Dest interface{}
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// NamedExecContext holds details about calls to the NamedExecContext method.
		NamedExecContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Arg is the arg argument value.
			Arg interface{}
		}
		// PreparexContext holds details about calls to the PreparexContext method.
		PreparexContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
		}
		// QueryRowContext holds details about calls to the QueryRowContext method.
		QueryRowContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// QueryxContext holds details about calls to the QueryxContext method.
		QueryxContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// SelectContext holds details about calls to the SelectContext method.
		SelectContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Dest is the dest argument value.
			Dest interface{}
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
	}
	lockExecContext      sync.RWMutex
	lockGetContext       sync.RWMutex
	lockNamedExecContext sync.RWMutex
	lockPreparexContext  sync.RWMutex
	lockQueryRowContext  sync.RWMutex
	lockQueryxContext    sync.RWMutex
	lockSelectContext    sync.RWMutex
}

// ExecContext calls ExecContextFunc.
func (mock *DBConnectionMock) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if mock.ExecContextFunc == nil {
		panic("DBConnectionMock.ExecContextFunc: method is nil but DBConnection.ExecContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	mock.lockExecContext.Lock()
	mock.calls.ExecContext = append(mock.calls.ExecContext, callInfo)
	mock.lockExecContext.Unlock()
	return mock.ExecContextFunc(ctx, query, args...)
}

// ExecContextCalls gets all the calls that were made to ExecContext.
// Check the length with:
//
//	len(mockedDBConnection.ExecContextCalls())
func (mock *DBConnectionMock) ExecContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []any
	}
	mock.lockExecContext.RLock()
	calls = mock.calls.ExecContext
	mock.lockExecContext.RUnlock()
	return calls
}

// GetContext calls GetContextFunc.
func (mock *DBConnectionMock) GetContext(ctx context.Context, dest interface{}, query string, args ...any) error {
	if mock.GetContextFunc == nil {
		panic("DBConnectionMock.GetContextFunc: method is nil but DBConnection.GetContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Dest:  dest,
		Query: query,
		Args:  args,
	}
	mock.lockGetContext.Lock()
	mock.calls.GetContext = append(mock.calls.GetContext, callInfo)
	mock.lockGetContext.Unlock()
	return mock.GetContextFunc(ctx, dest, query, args...)
}

// GetContextCalls gets all the calls that were made to GetContext.
// Check the length with:
//
//	len(mockedDBConnection.GetContextCalls())
func (mock *DBConnectionMock) GetContextCalls() []struct {
	Ctx   context.Context
	Dest  interface{}
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}
	mock.lockGetContext.RLock()
	calls = mock.calls.GetContext
	mock.lockGetContext.RUnlock()
	return calls
}

// NamedExecContext calls NamedExecContextFunc.
func (mock *DBConnectionMock) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	if mock.NamedExecContextFunc == nil {
		panic("DBConnectionMock.NamedExecContextFunc: method is nil but DBConnection.NamedExecContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Arg   interface{}
	}{
		Ctx:   ctx,
		Query: query,
		Arg:   arg,
	}
	mock.lockNamedExecContext.Lock()
	mock.calls.NamedExecContext = append(mock.calls.NamedExecContext, callInfo)
	mock.lockNamedExecContext.Unlock()
	return mock.NamedExecContextFunc(ctx, query, arg)
}

// NamedExecContextCalls gets all the calls that were made to NamedExecContext.
// Check the length with:
//
//	len(mockedDBConnection.NamedExecContextCalls())
func (mock *DBConnectionMock) NamedExecContextCalls() []struct {
	Ctx   context.Context
	Query string
	Arg   interface{}
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Arg   interface{}
	}
	mock.lockNamedExecContext.RLock()
	calls = mock.calls.NamedExecContext
	mock.lockNamedExecContext.RUnlock()
	return calls
}

// PreparexContext calls PreparexContextFunc.
func (mock *DBConnectionMock) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	if mock.PreparexContextFunc == nil {
		panic("DBConnectionMock.PreparexContextFunc: method is nil but DBConnection.PreparexContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
	}{
		Ctx:   ctx,
		Query: query,
	}
	mock.lockPreparexContext.Lock()
	mock.calls.PreparexContext = append(mock.calls.PreparexContext, callInfo)
	mock.lockPreparexContext.Unlock()
	return mock.PreparexContextFunc(ctx, query)
}

// PreparexContextCalls gets all the calls that were made to PreparexContext.
// Check the length with:
//
//	len(mockedDBConnection.PreparexContextCalls())
func (mock *DBConnectionMock) PreparexContextCalls() []struct {
	Ctx   context.Context
	Query string
} {
	var calls []struct {
		Ctx   context.Context
		Query string
	}
	mock.lockPreparexContext.RLock()
	calls = mock.calls.PreparexContext
	mock.lockPreparexContext.RUnlock()
	return calls
}

// QueryRowContext calls QueryRowContextFunc.
func (mock *DBConnectionMock) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if mock.QueryRowContextFunc == nil {
		panic("DBConnectionMock.QueryRowContextFunc: method is nil but DBConnection.QueryRowContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	mock.lockQueryRowContext.Lock()
	mock.calls.QueryRowContext = append(mock.calls.QueryRowContext, callInfo)
	mock.lockQueryRowContext.Unlock()
	return mock.QueryRowContextFunc(ctx, query, args...)
}

// QueryRowContextCalls gets all the calls that were made to QueryRowContext.
// Check the length with:
//
//	len(mockedDBConnection.QueryRowContextCalls())
func (mock *DBConnectionMock) QueryRowContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []any
	}
	mock.lockQueryRowContext.RLock()
	calls = mock.calls.QueryRowContext
	mock.lockQueryRowContext.RUnlock()
	return calls
}

// QueryxContext calls QueryxContextFunc.
func (mock *DBConnectionMock) QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	if mock.QueryxContextFunc == nil {
		panic("DBConnectionMock.QueryxContextFunc: method is nil but DBConnection.QueryxContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	mock.lockQueryxContext.Lock()
	mock.calls.QueryxContext = append(mock.calls.QueryxContext, callInfo)
	mock.lockQueryxContext.Unlock()
	return mock.QueryxContextFunc(ctx, query, args...)
}

// QueryxContextCalls gets all the calls that were made to QueryxContext.
// Check the length with:
//
//	len(mockedDBConnection.QueryxContextCalls())
func (mock *DBConnectionMock) QueryxContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []any
	}
	mock.lockQueryxContext.RLock()
	calls = mock.calls.QueryxContext
	mock.lockQueryxContext.RUnlock()
	return calls
}

// SelectContext calls SelectContextFunc.
func (mock *DBConnectionMock) SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error {
	if mock.SelectContextFunc == nil {
		panic("DBConnectionMock.SelectContextFunc: method is nil but DBConnection.SelectContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Dest:  dest,
		Query: query,
		Args:  args,
	}
	mock.lockSelectContext.Lock()
	mock.calls.SelectContext = append(mock.calls.SelectContext, callInfo)
	mock.lockSelectContext.Unlock()
	return mock.SelectContextFunc(ctx, dest, query, args...)
}

// SelectContextCalls gets all the calls that were made to SelectContext.
// Check the length with:
//
//	len(mockedDBConnection.SelectContextCalls())
func (mock *DBConnectionMock) SelectContextCalls() []struct {
	Ctx   context.Context
	Dest  interface{}
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}
	mock.lockSelectContext.RLock()
	calls = mock.calls.SelectContext
	mock.lockSelectContext.RUnlock()
	return calls
}
