package mdbx

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"testing"
	"unsafe"
)

func stringify(v interface{}) string {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(d)
}

func TestSysRamInfo(t *testing.T) {
	info, err := SysRamInfo()
	if err != ErrSuccess {
		t.Fatal(err)
	}
	t.Log("SysRamInfo:", info)
	fmt.Println(stringify(info))
}

func TestIsReadAheadReasonable(t *testing.T) {
	fmt.Println("IsReadAheadReasonable:", IsReadAheadReasonable(1024*1024*1024*4, 0))
}

func TestOpen(t *testing.T) {
	env, err := NewEnv()
	if err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetGeometry(Geometry{
		SizeLower:       1024 * 512,
		SizeNow:         1024 * 512,
		SizeUpper:       1024 * 1024 * 1024 * 16,
		GrowthStep:      1024 * 1024 * 20,
		ShrinkThreshold: 0,
		PageSize:        4096,
	}); err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetMaxDBS(1); err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetMaxReaders(1); err != ErrSuccess {
		t.Fatal(err)
	}
	os.Remove("testdata/db.dat")
	os.Remove("testdata/db.dat-lck")
	os.MkdirAll("testdata", 0755)
	err = env.Open(
		"testdata/db.dat",
		EnvNoTLS|EnvNoReadAhead|EnvCoalesce|EnvLIFOReclaim|EnvSafeNoSync|EnvWriteMap|EnvNoSubDir,
		0664,
	)
	if err != ErrSuccess {
		t.Fatal(err)
	}

	var txn Tx
	if err = env.Begin(&txn, TxReadWrite); err != ErrSuccess {
		t.Fatal(err)
	}

	if err = env.Warmup(&txn, WarmupOOMSafe, 0); err != ErrSuccess {
		t.Fatal(err)
	}

	var dbi DBI
	if dbi, err = txn.OpenDBI("m", DBIntegerKey|DBCreate); err != ErrSuccess {
		t.Fatal(err)
	}

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, uint64(101))
	value := []byte("hello")

	keyVal := Bytes(&key)
	valueVal := Bytes(&value)

	if err = txn.Put(dbi, &keyVal, &valueVal, PutUpsert); err != ErrSuccess {
		t.Fatal(err)
	}

	var info EnvInfo
	info, err = env.Info(&txn)
	if err != ErrSuccess {
		t.Fatal(err)
	}
	_ = info
	//fmt.Println(info)

	// var latency CommitLatency
	if err = txn.Commit(); err != ErrSuccess {
		t.Fatal(err)
	}

	err = env.Close(false)
	if err != ErrSuccess {
		t.Fatal(err)
	}
}

func TestGC(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	env, err := NewEnv()
	if err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetGeometry(Geometry{
		SizeLower:       1024 * 512,
		SizeNow:         1024 * 512,
		SizeUpper:       1024 * 1024 * 2,
		GrowthStep:      1024 * 512,
		ShrinkThreshold: 0,
		PageSize:        4096,
	}); err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetMaxDBS(4); err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetMaxReaders(64); err != ErrSuccess {
		t.Fatal(err)
	}
	//os.Remove("testdata/db.dat")
	//os.Remove("testdata/db.dat-lck")
	os.MkdirAll("testdata", 0755)
	err = env.Open(
		"testdata/db.dat",
		EnvNoTLS|EnvNoReadAhead|EnvCoalesce|EnvLIFOReclaim|EnvWriteMap|EnvSyncDurable|EnvNoSubDir,
		0664,
	)
	if err != ErrSuccess {
		t.Fatal(err)
	}

	var tx Tx
	if err = env.Begin(&tx, TxReadWrite); err != ErrSuccess {
		t.Fatal(err)
	}

	if err = env.Warmup(&tx, WarmupOOMSafe, 0); err != ErrSuccess {
		t.Fatal(err)
	}

	var (
		dbi     DBI
		dbiLogs DBI
	)
	if dbi, err = tx.OpenDBI("m", DBIntegerKey|DBCreate); err != ErrSuccess {
		t.Fatal(err)
	}
	if dbiLogs, err = tx.OpenDBI("l", DBIntegerKey|DBCreate); err != ErrSuccess {
		t.Fatal(err)
	}

	if err = tx.Commit(); err != ErrSuccess {
		t.Fatal(err)
	}

	data := make([]byte, 32)
	var id uint64 = 0

	for i := 0; i < 5000; i++ {
		if err = env.Begin(&tx, TxReadWrite); err != ErrSuccess {
			t.Fatal(err)
		}

		var (
			key_ = uint64(1)
			key  = U64(&key_)
			val  = Bytes(&data)
			old  Val
		)

		if err = tx.Replace(dbi, &key, &val, &old, 0); err != ErrSuccess {
			t.Fatal(err)
		}

		var cursor Cursor
		cursor, err = tx.OpenCursor(dbiLogs)
		if err != ErrSuccess {
			t.Fatal(err)
		}

		if err = cursor.Get(&key, &val, CursorLast); err != ErrSuccess {
			if err != ErrENODATA && err != ErrNotFound {
				t.Fatal(err)
			}
		}

		if err != ErrENODATA && err != ErrNotFound {
			id = key.U64()
		}

		for x := 0; x < 50; x++ {
			id++
			key_ = id
			key = U64(&key_)
			if err = cursor.Put(&key, &val, PutAppend); err != ErrSuccess {
				//if err = tx.Put(dbiLogs, &key, &val, PutAppend); err != ErrSuccess {
				t.Fatal(err)
			}
		}

		if err = cursor.Close(); err != ErrSuccess {
			t.Fatal(err)
		}

		if err = tx.Commit(); err != ErrSuccess {
			t.Fatal(err)
		}

		//chkPrint("-v", "-w", "testdata/db.dat")

		if i >= 0 {
			tx = Tx{}
			if err = env.Begin(&tx, TxReadWrite); err != ErrSuccess {
				t.Fatal(err)
			}

			first, last, count, err := tx.DeleteIntegerRange(dbiLogs, 0, math.MaxUint64, 100)
			if err != ErrSuccess {
				if err != ErrENODATA && err != ErrNotFound {
					t.Fatal(err)
				}
				_ = tx.Abort()
				goto NEXT
			}

			fmt.Println("Delete Range", "[", first, ",", last, "]  count:", count)

			//if cursor, err = tx.OpenCursor(dbiLogs); err != ErrSuccess {
			//	t.Fatal(err)
			//}
			//
			//key_ = 0
			//if err = cursor.Get(&key, &val, CursorFirst); err != ErrSuccess {
			//	if err != ErrENODATA && err != ErrNotFound {
			//		t.Fatal(err)
			//	}
			//	_ = cursor.Close()
			//	_ = tx.Abort()
			//	goto NEXT
			//}
			//
			//fmt.Println(key.U64())
			//key_ = key.U64()
			//
			//for x := 0; x < 100; x++ {
			//	if err = cursor.Get(&key, &val, CursorGetCurrent); err != ErrSuccess {
			//		if err != ErrENODATA && err != ErrNotFound {
			//			t.Fatal(err)
			//		}
			//		break
			//	}
			//
			//	//fmt.Println("delete key:", key.U64())
			//	if err = cursor.Delete(PutCurrent); err != ErrSuccess {
			//		if err == ErrENODATA || err == ErrNotFound {
			//			break
			//		}
			//		t.Fatal(err)
			//	}
			//
			//}
			//
			//if err = cursor.Close(); err != ErrSuccess {
			//	t.Fatal(err)
			//}

			if err = tx.Commit(); err != ErrSuccess {
				t.Fatal(err)
			}
		}

	NEXT:
		tx = Tx{}
		if err = env.Begin(&tx, TxReadWrite); err != ErrSuccess {
			t.Fatal(err)
		}
		if cursor, err = tx.OpenCursor(dbiLogs); err != ErrSuccess {
			t.Fatal(err)
		}
		if err = cursor.Get(&key, &val, CursorLast); err != ErrSuccess {
			if err != ErrENODATA && err != ErrNotFound {
				t.Fatal(err)
			}
		}

		if err != ErrENODATA && err != ErrNotFound {
			id = key.U64()
		}
		for x := 0; x < 50; x++ {
			id++
			key_ = id
			key = U64(&key_)
			if err = cursor.Put(&key, &val, PutAppend); err != ErrSuccess {
				//if err = tx.Put(dbiLogs, &key, &val, PutAppend); err != ErrSuccess {
				t.Fatal(err)
			}
		}
		if err = cursor.Close(); err != ErrSuccess {
			t.Fatal(err)
		}
		if err = tx.Commit(); err != ErrSuccess {
			t.Fatal(err)
		}
	}

	if err = env.Close(false); err != ErrSuccess {
		fmt.Println(err)
	}
}

type Engine struct {
	env    *Env
	rootDB DBI
	write  Tx
	rd     Tx
}

func (engine *Engine) BeginWrite() (*Tx, Error) {
	engine.write.ptr = nil
	return &engine.write, engine.env.Begin(&engine.write, TxReadWrite)
}

func (engine *Engine) BeginRead() (*Tx, Error) {
	return &engine.rd, engine.rd.Renew()
}

func initDB(path string, flags EnvFlags) (*Engine, Error) {
	_ = os.RemoveAll(path)
	os.MkdirAll(path, 0755)
	engine := &Engine{}
	env, err := NewEnv()
	if err != ErrSuccess {
		return nil, err
	}
	engine.env = env
	if err = env.SetGeometry(Geometry{
		SizeLower:       1024 * 1024 * 64,
		SizeNow:         1024 * 1024 * 64,
		SizeUpper:       1024 * 1024 * 1024 * 16,
		GrowthStep:      1024 * 1024 * 64,
		ShrinkThreshold: 0,
		PageSize:        65536,
	}); err != ErrSuccess {
		return nil, err
	}
	if err = env.SetMaxDBS(1); err != ErrSuccess {
		return nil, err
	}

	env.SetMaxReaders(2)
	if err != ErrSuccess {
		return nil, err
	}

	err = env.Open(
		path,
		//EnvNoMemInit|EnvCoalesce|EnvLIFOReclaim|EnvSyncDurable,
		// EnvNoMemInit|EnvCoalesce|EnvLIFOReclaim|EnvSafeNoSync|EnvWriteMap,
		EnvNoTLS|EnvNoMemInit|flags|EnvWriteMap|EnvUtterlyNoSync,
		0664,
	)

	if err = env.Begin(&engine.write, TxReadWrite); err != ErrSuccess {
		return nil, err
	}

	if engine.rootDB, err = engine.write.OpenDBI("m", DBIntegerKey|DBCreate); err != ErrSuccess {
		return nil, err
	}
	//if engine.rootDB, err = engine.write.OpenDBIEx("m", DBCreate, CmpU64, nil); err != ErrSuccess {
	//	return nil, err
	//}

	if err = engine.write.Commit(); err != ErrSuccess {
		return nil, err
	}

	//if err = env.Begin(&engine.rd, TxReadOnly); err != ErrSuccess {
	//	return nil, err
	//}
	//if err = engine.rd.Reset(); err != ErrSuccess {
	//	return nil, err
	//}

	return engine, ErrSuccess
}

func BenchmarkTxn_Put(b *testing.B) {
	defer os.RemoveAll("testdata/db")
	engine, err := initDB("testdata/db/"+strconv.Itoa(b.N), EnvSafeNoSync)
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(true)

	key := make([]byte, 8)
	data := []byte("hello")

	keyVal := Bytes(&key)
	dataVal := Bytes(&data)

	b.ResetTimer()
	b.ReportAllocs()

	txn, err := engine.BeginWrite()
	if err != ErrSuccess {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		//binary.BigEndian.PutUint64(key, uint64(20))
		//binary.LittleEndian.PutUint64(key, uint64(i))
		*(*uint64)(unsafe.Pointer(keyVal.Base)) = uint64(i)
		//keyVal = U64(uint64(i))
		if err = txn.Put(engine.rootDB, &keyVal, &dataVal, PutAppend); err != ErrSuccess {
			txn.Abort()
			b.Fatal(err)
		}
	}

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	if err = txn.Commit(); err != ErrSuccess {
		b.Fatal(err)
	}
	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}

func BenchmarkTxn_PutCursor(b *testing.B) {
	//defer os.RemoveAll("testdata/db")
	engine, err := initDB("testdata/db/"+strconv.Itoa(b.N), EnvSafeNoSync)
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(false)

	key := uint64(0)
	data := []byte("hello")
	keyVal := U64(&key)
	dataVal := Bytes(&data)
	b.ReportAllocs()
	b.ResetTimer()

	{
		insert := func(low, high uint64) {
			txn, err := engine.BeginWrite()
			if err != ErrSuccess {
				b.Fatal(err)
			}

			cursor, err := txn.OpenCursor(engine.rootDB)
			if err != ErrSuccess {
				b.Fatal(err)
			}

			for i := low; i < high; i++ {
				key = i
				if err = cursor.Put(&keyVal, &dataVal, PutAppend); err != ErrSuccess {
					cursor.Close()
					txn.Abort()
					b.Fatal(err)
				}
			}

			if err = cursor.Close(); err != ErrSuccess {
				b.Fatal(err)
			}
			if err = txn.Commit(); err != ErrSuccess {
				b.Fatal(err)
			}
		}

		const batchSize = 10000000
		for i := 0; i < b.N; i += batchSize {
			end := i + batchSize
			if end > b.N {
				end = b.N
			}
			insert(uint64(i), uint64(end))
		}
	}
	b.StopTimer()
}

func BenchmarkTxn_Get(b *testing.B) {
	defer os.RemoveAll("testdata/db")
	engine, err := initDB("testdata/db/"+strconv.Itoa(b.N), EnvSafeNoSync)
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(true)

	key := uint64(0)
	data := []byte("hello")

	keyVal := U64(&key)
	dataVal := Bytes(&data)

	{
		insert := func(low, high uint64) {
			txn, err := engine.BeginWrite()
			if err != ErrSuccess {
				b.Fatal(err)
			}

			cursor, err := txn.OpenCursor(engine.rootDB)
			if err != ErrSuccess {
				b.Fatal(err)
			}

			for i := low; i < high; i++ {
				key = i
				if err = cursor.Put(&keyVal, &dataVal, PutAppend); err != ErrSuccess {
					cursor.Close()
					txn.Abort()
					b.Fatal(err)
				}
			}

			if err = cursor.Close(); err != ErrSuccess {
				b.Fatal(err)
			}
			if err = txn.Commit(); err != ErrSuccess {
				b.Fatal(err)
			}
		}

		const batchSize = 10000000
		for i := 0; i < b.N; i += batchSize {
			end := i + batchSize
			if end > b.N {
				end = b.N
			}
			insert(uint64(i), uint64(end))
		}
	}

	txn := &Tx{}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)

	if err = engine.env.Begin(txn, TxReadOnly); err != ErrSuccess {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()

	//fmt.Println(dataVal.String())

	//binary.LittleEndian.PutUint64(key, 0)

	count := 0

	for i := 1; i < b.N; i++ {
		key = uint64(i)
		keyVal = U64(&key)
		//binary.BigEndian.PutUint64(key, uint64(20))
		//binary.BigEndian.PutUint64(key[8:], uint64(i))
		if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess && err != ErrNotFound {
			txn.Reset()
			b.Fatal(err)
		}
		count++
	}

	if err = txn.Reset(); err != ErrSuccess {
		b.Fatal(err)
	}

	b.StopTimer()

	fmt.Println("count", count)

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}

func BenchmarkTxn_GetCursor(b *testing.B) {
	defer os.RemoveAll("testdata/db")
	engine, err := initDB("testdata/db/"+strconv.Itoa(b.N), EnvSafeNoSync)
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(true)

	key := uint64(0)
	data := []byte("hello")

	keyVal := U64(&key)
	dataVal := Bytes(&data)

	{
		insert := func(low, high uint64) {
			txn, err := engine.BeginWrite()
			if err != ErrSuccess {
				b.Fatal(err)
			}

			cursor, err := txn.OpenCursor(engine.rootDB)
			if err != ErrSuccess {
				b.Fatal(err)
			}

			for i := low; i < high; i++ {
				key = i
				if err = cursor.Put(&keyVal, &dataVal, PutAppend); err != ErrSuccess {
					cursor.Close()
					txn.Abort()
					b.Fatal(err)
				}
			}

			if err = cursor.Close(); err != ErrSuccess {
				b.Fatal(err)
			}
			if err = txn.Commit(); err != ErrSuccess {
				b.Fatal(err)
			}
		}

		const batchSize = 1000000
		for i := 0; i < b.N; i += batchSize {
			end := i + batchSize
			if end > b.N {
				end = b.N
			}
			insert(uint64(i), uint64(end))
		}
	}

	txn := Tx{}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)

	if err = engine.env.Begin(&txn, TxReadOnly); err != ErrSuccess {
		b.Fatal(err)
	}
	//txn, err = engine.BeginRead()
	//if err != ErrSuccess {
	//	b.Fatal(err)
	//}

	b.ResetTimer()
	b.ReportAllocs()

	cursor, err := txn.OpenCursor(engine.rootDB)
	if err != ErrSuccess {
		b.Fatal(err)
	}

	//binary.LittleEndian.PutUint64(key, uint64(b.N))

	//if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//keyInt := binary.LittleEndian.Uint64(key)

	//if err = cursor.Get(&keyVal, &dataVal, CursorSet); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	dataVal = Val{}
	keyVal = Val{}

	//fmt.Println(dataVal.String())

	//binary.LittleEndian.PutUint64(key, 0)

	count := 0
	//
	for {
		if err = cursor.Get(&keyVal, &dataVal, CursorNextNoDup); err != ErrSuccess {
			break
		}
		//if keyVal.Base == nil {
		//	break
		//}
		count++
		//keyInt = binary.LittleEndian.Uint64(key)
		//_ = keyInt

		//keyVal = Val{}
		//dataVal = Val{}

		//if cursor.EOF() != 0 {
		//	break
		//}
	}

	//if count == 1000000 {
	//	println("1m")
	//}

	//for i := 0; i < b.N; i++ {
	//	*(*uint64)(unsafe.Pointer(&key[0])) = uint64(i)
	//	//binary.BigEndian.PutUint64(key, uint64(20))
	//	//binary.BigEndian.PutUint64(key[8:], uint64(i))
	//	//keyVal = U64(uint64(i))
	//	if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess && err != ErrNotFound {
	//		txn.Reset()
	//		b.Fatal(err)
	//	}
	//}

	if err = cursor.Close(); err != ErrSuccess {
		b.Fatal(err)
	}
	if err = txn.Reset(); err != ErrSuccess {
		b.Fatal(err)
	}

	b.StopTimer()

	fmt.Println("count", count)

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}

func TestTxn_Cursor(b *testing.T) {
	// defer func() {
	// 	_ = Delete("testdata/db", DeleteModeJustDelete)
	// 	os.RemoveAll("testdata/db")
	// }()
	_ = Delete("testdata/db", DeleteModeJustDelete)
	os.RemoveAll("testdata/db")
	iterations := 100
	engine, err := initDB("testdata/db/"+strconv.Itoa(iterations), EnvSafeNoSync)
	if err != ErrSuccess {
		b.Fatal(err)
	}

	key := make([]byte, 8)
	data := []byte("hello")

	keyVal := Bytes(&key)
	dataVal := Bytes(&data)

	txn, err := engine.BeginWrite()
	if err != ErrSuccess {
		b.Fatal(err)
	}

	for i := 1; i <= iterations; i++ {
		//binary.BigEndian.PutUint64(key, uint64(20))
		//*(*uint64)(unsafe.Pointer(&key[0])) = uint64(i)
		binary.LittleEndian.PutUint64(key, uint64(i))
		//keyVal = U64(uint64(i))
		if err = txn.Put(engine.rootDB, &keyVal, &dataVal, 0); err != ErrSuccess {
			txn.Abort()
			b.Fatal(err)
		}
	}

	//*(*uint64)(unsafe.Pointer(&key[0])) = 0

	if err = txn.Commit(); err != ErrSuccess {
		b.Fatal(err)
	}

	txn = &Tx{}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)

	if err = engine.env.Begin(txn, TxReadOnly); err != ErrSuccess {
		b.Fatal(err)
	}
	//txn, err = engine.BeginRead()
	//if err != ErrSuccess {
	//	b.Fatal(err)
	//}

	cursor, err := txn.OpenCursor(engine.rootDB)
	if err != ErrSuccess {
		b.Fatal(err)
	}

	dataVal = Val{}
	keyVal = Val{}

	binary.LittleEndian.PutUint64(key, 0)

	if err = cursor.Get(&keyVal, &dataVal, CursorFirst); err != ErrSuccess {
		b.Fatal(err)
	}
	println("first key", keyVal.U64())
	count := 1
	//
	for {
		if err = cursor.Get(&keyVal, &dataVal, CursorNextNoDup); err != ErrSuccess {
			break
		}
		//if keyVal.Base == nil {
		//	break
		//}
		count++
		keyInt := keyVal.U64()
		println("key", keyInt)
		_ = keyInt

		//keyVal = Val{}
		//dataVal = Val{}

		//if cursor.EOF() != 0 {
		//	break
		//}
	}

	//for i := 0; i < b.N; i++ {
	//	*(*uint64)(unsafe.Pointer(&key[0])) = uint64(i)
	//	//binary.BigEndian.PutUint64(key, uint64(20))
	//	//binary.BigEndian.PutUint64(key[8:], uint64(i))
	//	//keyVal = U64(uint64(i))
	//	if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess && err != ErrNotFound {
	//		txn.Reset()
	//		b.Fatal(err)
	//	}
	//}

	if err = cursor.Close(); err != ErrSuccess {
		b.Fatal(err)
	}
	if err = txn.Reset(); err != ErrSuccess {
		b.Fatal(err)
	}

	fmt.Println("count", count)

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}

func BenchmarkWrite(b *testing.B) {
	const runPebble = false
	const all = 10000000000

	runMDBXAppendBatched := func(batchSize int, name string, flags EnvFlags) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		batchSizeString := "ALL"
		if batchSize < all {
			batchSizeString = strconv.Itoa(batchSize)
		}
		b.Run("MDBX("+name+") Append "+batchSizeString, func(b *testing.B) {
			defer func() {
				//if err := Delete("testdata/db", DeleteModeWaitForUnused); err != ErrSuccess {
				//	b.Fatal(err)
				//}
				defer os.RemoveAll("testdata/db")
			}()
			engine, err := initDB("testdata/db", flags)
			if err != ErrSuccess {
				b.Fatal(err)
			}
			key := uint64(0)
			data := []byte("hello")
			keyVal := U64(&key)
			dataVal := Bytes(&data)
			b.ReportAllocs()
			b.ResetTimer()
			{
				insert := func(low, high uint64) {
					txn, err := engine.BeginWrite()
					if err != ErrSuccess {
						b.Fatal(err)
					}

					cursor, err := txn.OpenCursor(engine.rootDB)
					if err != ErrSuccess {
						b.Fatal(err)
					}

					for i := low; i < high; i++ {
						key = i
						if err = cursor.Put(&keyVal, &dataVal, PutAppend); err != ErrSuccess {
							cursor.Close()
							txn.Abort()
							b.Fatal(err)
						}
					}

					if err = cursor.Close(); err != ErrSuccess {
						b.Fatal(err)
					}
					if err = txn.Commit(); err != ErrSuccess {
						b.Fatal(err)
					}

					//if flags != EnvSyncDurable {
					//	if err = engine.env.Sync(false, false); err != ErrSuccess {
					//		b.Fatal(err)
					//	}
					//}
				}

				for i := 0; i < b.N; i += batchSize {
					end := i + batchSize
					if end > b.N {
						end = b.N
					}
					insert(uint64(i), uint64(end))
				}

				if flags != EnvSyncDurable {
					if err = engine.env.Sync(true, false); err != ErrSuccess {
						b.Fatal(err)
					}
				}
			}
			b.StopTimer()

			if err = engine.env.Close(false); err != ErrSuccess {
				b.Fatal(err)
			}
		})
	}

	//runMDBXAppendBatched(all, "SyncDurable", EnvSyncDurable)
	//runMDBXAppendBatched(100000, "SyncDurable", EnvSyncDurable)
	//runMDBXAppendBatched(10000, "SyncDurable", EnvSyncDurable)
	//runMDBXAppendBatched(1000, "SyncDurable", EnvSyncDurable)
	//runMDBXAppendBatched(100, "SyncDurable", EnvSyncDurable)
	//
	//runMDBXAppendBatched(all, "SafeNoSync", EnvSafeNoSync)
	//runMDBXAppendBatched(100000, "SafeNoSync", EnvSafeNoSync)
	//runMDBXAppendBatched(10000, "SafeNoSync", EnvSafeNoSync)
	//runMDBXAppendBatched(1000, "SafeNoSync", EnvSafeNoSync)
	//runMDBXAppendBatched(100, "SafeNoSync", EnvSafeNoSync)

	runMDBXAppendBatched(all, "NoSync", EnvUtterlyNoSync)
	runMDBXAppendBatched(100000, "NoSync", EnvUtterlyNoSync)
	runMDBXAppendBatched(10000, "NoSync", EnvUtterlyNoSync)
	runMDBXAppendBatched(1000, "NoSync", EnvUtterlyNoSync)
	runMDBXAppendBatched(100, "NoSync", EnvUtterlyNoSync)
}
