package main

import(
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/fatih/color"
	"encoding/json"
	"log"
	"fmt"
)

type golevelInterface interface {
	NewDatabase(path string)(*golevelDatabase ,error)
	Set(key string , value Request) error                 //insert the key(string) <--> value([]byte) pair into the database 
	Get(key []byte)(*golevelDatabase , error)				 //fetch the value with  the key from the database
	GetallInCsv()(error)									 //using for the debugging :) creates the csv file contains all the key value pairs in the database
}
type tuple struct {
	Key string `json:"key"`
	Val string	`json:"val"`
}

//struct defining the database
type golevelDatabase struct {
	db *leveldb.DB
}


//creates the database
//Parameters:
//	 -path(string) at what place we want to locate / previously located.
func Create_Database(path string)(*golevelDatabase) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil{
		fmt.Println("Error in Creating Database ** err --> Create_Database")
	}
	color.Green("database created")
	return &golevelDatabase{db:db}

}


//Get the value from the database
//Parameters:
//	-key(string) 
func (b *golevelDatabase) Get(key string) (Request , error) {
	fetched_byte_stream , err:= b.db.Get([]byte(key) , nil)

	var reqData Request
	err = json.Unmarshal(fetched_byte_stream , &reqData)
	return reqData, err
}


//insert or update the data with the key
// Parameters
// 	-key(string) with which key we want to insert into the database
// 	-value(leveldbVal (struct)) with which value we want to insert into the database
func (b *golevelDatabase) Set(key string , value Request)error{

	jsonStr , err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	return b.db.Put([]byte(key) , []byte(jsonStr) , nil)
}

func (b *golevelDatabase) Delete(key string) error{
	err := b.db.Delete([]byte(key), nil)
	if(err != nil){
		return err
	}
	return nil
}


func (b *golevelDatabase) Getall() string {
	iter := b.db.NewIterator(nil, nil)

	str := "["
	if iter.Next() {
		
		record := tuple{
			Key: string(iter.Key()),
			Val: string(iter.Value()),
		}
		//fmt.Println(record.key , )
		app , err := json.Marshal(record)
		if err != nil {
			fmt.Println("Marshal error in Get all function")
		}
		str += string(app)
	}

	for iter.Next() {
		

		record := tuple{
			Key: string(iter.Key()),
			Val: string(iter.Value()),
		}
		str += ","
		app , _ := json.Marshal(record)
		str += string(app)
	}
	iter.Release()
	str += "]"
	return str
}

