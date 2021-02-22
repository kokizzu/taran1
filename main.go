package main

import (
	"github.com/francoispqt/onelog"
	"github.com/kokizzu/gotro/D/Tr"
	"github.com/kokizzu/gotro/L"
	"github.com/tarantool/go-tarantool"
	"os"
	"taran1/pkg/config"
)
var log *onelog.Logger

func init() {
	log = onelog.New(
		os.Stdout,
		onelog.ALL,
	)
}

const DummySpace = `dummy`
const Id = `id`
const Title = `title`
const Created = `created`

var dummyTable = &Tr.TableProp{
	SpaceName: DummySpace,
	Droppable: true,
	Fields: []Tr.Field{
		{Name: Id, Type: Tr.Unsigned},
		{Name: Title, Type: Tr.String},
		{Name: Created, Type: Tr.Unsigned},
	},
	Unique:        Id,
	AutoIncrement: true,
}


func main() {
	config.LoadEnv()
		
	// tarantool connect
	taranHost := os.Getenv(config.TaranHost)
	taranUser := os.Getenv(config.TaranUser)
	opts := tarantool.Opts{User: taranUser}
	taranConn, err := tarantool.Connect(taranHost, opts)
	L.PanicIf(err,`cannot connect to tarantool`)
	
	// truncate tables
	taran := Tr.Taran{taranConn,log}
	taran.MigrateTarantool(dummyTable)
	
	// TODO: generate 1m records vinyl vs memtx
	
	// TODO: benchmark update/select vinyl vs memtx
	
	// TODO: reformat/alter table
	
	// TODO: backup table
	
	// TODO: restore table
	
	// TODO: test replication, test sync on readonly replica
}
