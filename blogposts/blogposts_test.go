
// TODO : modify this chapter in forked repo to make it more beginner friendly
package blogposts

import (
	"log"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("reading correct number of files or not", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md":  {Data: []byte("hi")},
			"hello_world2.md": {Data: []byte("hola")},
		}
		posts, err := NewPostsFromFS(fs)
		CheckError(t, err)
		//_, err = NewPostsFromFS(StubFailingFS{})
		//CheckError(t ,err)
		if len(posts) != len(fs) {
			t.Errorf("got %d posts , wanted %d posts", len(posts), len(fs))
		}
	})
	t.Run("test  blog post schema , title field", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := NewPostsFromFS(fs)
		CheckError(t,err)
		// rest of test code cut for brevity
		got := posts[0]
		want := Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

}

func CheckError(t testing.TB, e error) {
	t.Helper()
	if e != nil {
		log.Fatal(e)
	}
}
