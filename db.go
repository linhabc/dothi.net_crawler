package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/syndtr/goleveldb/leveldb"
)

func createOrOpenDb(path string) *leveldb.DB {
	db, _ := leveldb.OpenFile(path, nil)
	return db
}

func getData(db *leveldb.DB, key string) string {
	data, _ := db.Get([]byte(key), nil)
	return bytesToString(data)
}

func putData(db *leveldb.DB, key string, data string) error {
	err := db.Put([]byte(key), []byte(data), nil)
	return err
}

func bytesToString(data []byte) string {
	return string(data[:])
}

func exportDatabaseToJson(url string) {
	db := createOrOpenDb("./db/" + url)
	users := NewUsers()
	var count int
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := bytesToString(iter.Key())
		value := bytesToString(iter.Value())
		count++
		user := User{
			ID:          key,
			PhoneNumber: value,
		}

		users.List = append(users.List, user)

		fmt.Printf("Key: %s, Value: %s", key, value)
		println("")
	}

	users.TotalUsers = count

	// convert User sang JSON
	userJSON, err := json.Marshal(users)
	checkError(err)

	// Ghi dữ liệu vào file JSON
	err = ioutil.WriteFile("./output/"+url+".json", userJSON, 0644)
	checkError(err)
}

// func mainTest() {

// 	db1 := createOrOpenDb("./db/cat1")
// 	db2 := createOrOpenDb("./db/cat2")

// 	defer db1.Close()
// 	defer db2.Close()

// 	// _ = putData(db1, "id1", "sdt1")
// 	// _ = putData(db1, "id2", "sdt2")

// 	// _ = putData(db2, "id3", "sdt3")
// 	// _ = putData(db2, "id4", "sdt4")

// 	data1, _ := db1.Get([]byte("id1"), nil)
// 	data := getData(db1, "id2")

// 	fmt.Printf("db1 value: %s\n", data1)
// 	fmt.Printf("db1 value: %s\n", data)

// 	data1, _ = db2.Get([]byte("id3"), nil)
// 	data = getData(db2, "id5")

// 	fmt.Printf("db2 value: %s\n", data1)
// 	fmt.Printf("db2 value: %s\n", data)
// }

// func main() {

// 	var count int
// 	db := createOrOpenDb("./db/Ô tô")

// 	defer db.Close()

// 	iter := db.NewIterator(nil, nil)
// 	for iter.Next() {
// 		count++
// 		// key := iter.Key()
// 		// value := iter.Value()

// 		// fmt.Printf("Key: %s, Value: %s", key, value)
// 		// println("")
// 	}
// 	println(count)
// 	iter.Release()
// 	_ = iter.Error()
// }

// func main() {
// 	exportDatabaseToJson("Bat_dong_san")
// }

//  level db ko co muti thread
//  nhan vao id -> co hay ko
//  2 routine, 1 read + 1 write

// store key -> id
//       value -> phone number

// err = db.Delete([]byte("key"), nil)
