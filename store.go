package mdbx

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	DataFileName = "mdbx.dat"
	LockFileName = "mdbx.lck"
)

var (
	storeMu sync.Mutex
	stores  = make(map[*Store]struct{})
	syncer  = make(chan struct {
		s     *Store
		delay time.Duration
	}, 10)
)

type Store struct {
	env              *Env
	path             string
	recovery         string
	recoveryDuration time.Duration
	closed           int64
	updates          uint64
	synced           uint64
	syncQueued       uint64
	syncPeriod       time.Duration
	writeMu          sync.Mutex
	syncMu           sync.Mutex
	mu               sync.Mutex
}

func Open(
	path string,
	flags EnvFlags,
	mode os.FileMode,
	initEnv func(env *Env, create bool) error,
	init func(store *Store, create bool) error,
) (*Store, error) {
	store := &Store{
		path: path,
	}

	env, e := NewEnv()
	if e != ErrSuccess {
		return nil, e
	}
	store.env = env

	var err error
	var stat os.FileInfo
	stat, err = os.Stat(filepath.Join(path, DataFileName))
	var create bool
	if err != nil {
		create = true
	} else {
		// Init
		create = false
	}
	_ = stat

	if initEnv != nil {
		if err = initEnv(env, create); err != nil && err != ErrSuccess {
			store.Close()
			return nil, err
		}
	}

	if mode == 0 {
		mode = 0664
	}

	if err = store.env.Open(path, flags, mode); err != ErrSuccess {
		if err == ErrWannaRecovery {
			start := time.Now()
			_, output, err := Chk("-v", "-w", store.path)
			if err != ErrSuccess {
				_ = store.Close()
				return nil, err
			}
			store.recovery = string(output)
			store.recoveryDuration = time.Now().Sub(start)

			if err = store.env.Open(path, flags, mode); err != ErrSuccess {
				_ = store.Close()
				return nil, err
			}
			//_ = store.Close()
			//return nil, err
		} else {
			_ = store.Close()
			return nil, err
		}
	}

	if _, err = store.env.ReaderCheck(); err != nil && err != ErrSuccess {
		_ = store.Close()
		return nil, err
	}

	if init != nil {
		if err = init(store, create); err != nil && err != ErrSuccess {
			_ = store.Close()
			return nil, err
		}
	}

	if flags&EnvSafeNoSync != 0 || flags&EnvNoMetaSync != 0 {
		//syncBytes, _ := env.GetSyncBytes()
		syncPeriod, _ := env.GetSyncPeriod()

		if syncPeriod > 0 {

		}
	}

	storeMu.Lock()
	stores[store] = struct{}{}
	storeMu.Unlock()
	return store, nil
}

func (s *Store) Env() *Env {
	return s.env
}

func (s *Store) IsClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed > 0
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	if s.closed > 0 {
		return os.ErrClosed
	}
	s.closed = time.Now().UnixNano()
	if s.env != nil {
		_ = s.env.Close(false)
		s.env = nil
	}
	storeMu.Lock()
	delete(stores, s)
	storeMu.Unlock()
	return nil
}

func (s *Store) Update(fn func(tx *Tx) error) error {
	return s.UpdateLock(true, fn)
}

func (s *Store) UpdateLock(lockThread bool, fn func(tx *Tx) error) (err error) {
	if lockThread {
		// Write transactions must be bound to a single thread.
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// Get exclusive write lock.
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	tx := Tx{}
	defer func() {
		// Abort if panic
		if !tx.IsCommitted() && !tx.IsAborted() && tx.txn != nil {
			if e := tx.Abort(); e != ErrSuccess {
				// Ignore
			}
		}

		// Propagate
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()

	if err = s.env.Begin(&tx, TxReadWrite); err != ErrSuccess {
		return err
	}
	if err = fn(&tx); err != nil && err != ErrSuccess {
		// Abort if necessary
		if !tx.IsAborted() && !tx.IsCommitted() {
			if e := tx.Abort(); e != ErrSuccess {
				// Ignore
			}
		}
		return err
	} else {
		if !tx.IsCommitted() {
			if err = tx.Commit(); err != ErrSuccess {
				return err
			}
		}
		atomic.AddUint64(&s.updates, 1)
		return nil
	}
}

func (s *Store) View(fn func(tx *Tx) error) (err error) {
	tx := Tx{}
	defer func() {
		if !tx.IsAborted() {
			_ = tx.Abort()
		}
		// recovery
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	if err = s.env.Begin(&tx, TxReadOnly); err != ErrSuccess {
		return err
	}
	return fn(&tx)
}

func (s *Store) ViewRenew(tx *Tx, fn func(tx *Tx) error) (err error) {
	if tx == nil {
		return s.View(fn)
	}
	defer func() {
		if !tx.IsReset() {
			err = tx.Reset()
		}
		// recovery
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	if err = tx.Renew(); err != ErrSuccess {
		return err
	}
	return fn(tx)
}

func (s *Store) Sync() error {
	update := atomic.LoadUint64(&s.updates)
	if err := s.env.Sync(true, false); err != ErrSuccess && err != ErrResultTrue {
		return err
	}
	atomic.StoreUint64(&s.synced, update)
	return nil
}

func (s *Store) Chk() (result int32, output []byte, err error) {
	return Chk("-v", filepath.Join(s.path, DataFileName))
}

func (s *Store) ChkRecover() (result int32, output []byte, err error) {
	return Chk("-v", "-w", filepath.Join(s.path, DataFileName))
}
