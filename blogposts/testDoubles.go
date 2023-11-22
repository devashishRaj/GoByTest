package blogposts

import (
	"errors"
	"io/fs"
)

type StubFailingFS struct{

}

func ( s StubFailingFS) Open(name string) (fs.File , error){
	return nil , errors.New("I always fail")
}