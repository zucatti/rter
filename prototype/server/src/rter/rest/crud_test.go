package rest

import (
	"encoding/json"
	"github.com/drewolson/testflight"
	"net/http"
	"rter/data"
	"rter/storage"
	"strconv"
	"testing"
	"time"
)

var (
	role      *data.Role
	user      *data.User
	direction *data.UserDirection

	item    *data.Item
	comment *data.ItemComment

	term    *data.Term
	ranking *data.TermRanking
)

func TestOpenStorage(t *testing.T) {
	err := storage.OpenStorage("rter", "j2pREch8", "tcp", "localhost:3306", "rter")

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateRole(t *testing.T) {
	role = new(data.Role)
	role.Title = "TestRole"
	role.Permissions = 1

	testCreate(t, "/roles", role)
}

func TestUpdateRole(t *testing.T) {
	role.Permissions = 5

	testUpdate(t, "/roles/"+role.Title, role)

	structJSONCompare(t, role.Permissions, 5)
}

func TestReadRole(t *testing.T) {
	readrole := new(data.Role)

	testRead(t, "/roles/"+role.Title, readrole)

	structJSONCompare(t, readrole, role)
}

func TestReadAllRole(t *testing.T) {
	v := make([]*data.Role, 0)

	testRead(t, "/roles", &v)

	if len(v) < 1 {
		t.Error("ReadAll Failed")
	}
}

func TestCreateUser(t *testing.T) {
	user = new(data.User)
	user.Username = "TestUser"
	user.Password = "passwordhash"
	user.Salt = "serioussalt"
	user.Role = role.Title
	user.TrustLevel = 1
	user.CreateTime = time.Now()

	testCreate(t, "/users", user)
}

func TestUpdateUser(t *testing.T) {
	user.TrustLevel = 5

	testUpdate(t, "/users/"+user.Username, user)

	structJSONCompare(t, user.TrustLevel, 5)
}

func TestReadUser(t *testing.T) {
	readUser := new(data.User)

	testRead(t, "/users/"+user.Username, readUser)

	readUser.CreateTime = user.CreateTime

	structJSONCompare(t, user, readUser)
}

func TestReadAllUser(t *testing.T) {
	v := make([]*data.User, 0)

	testRead(t, "/users", &v)

	if len(v) < 1 {
		t.Error("ReadAll Failed")
	}
}

func TestUpdateUserDirection(t *testing.T) {
	direction = new(data.UserDirection)
	direction.Username = user.Username
	direction.Command = "none"
	direction.Heading = 12.123
	direction.Lat = 123.234
	direction.Lng = -74.234
	direction.UpdateTime = time.Now()

	testUpdate(t, "/users/"+user.Username+"/direction", direction)
}

func TestReadUserDirection(t *testing.T) {
	readDirection := new(data.UserDirection)

	testRead(t, "/users/"+user.Username+"/direction", readDirection)

	readDirection.UpdateTime = direction.UpdateTime // hack

	structJSONCompare(t, direction, readDirection)
}

func TestCreateItem(t *testing.T) {
	item = new(data.Item)
	item.Type = "generic"
	item.Author = user.Username
	item.ThumbnailURI = "http://fun.com/thumb.jpg"
	item.ContentURI = "http://fun.com"
	item.UploadURI = "http://fun.com/upload"
	item.HasGeo = false
	item.Heading = -40.3
	item.Lat = 47.123
	item.Lng = -123.123
	item.StartTime = time.Now()

	testCreate(t, "/items", item)
}

func TestReadItem1(t *testing.T) {
	readItem := new(data.Item)

	testRead(t, "/items/"+strconv.FormatInt(item.ID, 10), readItem)

	readItem.StartTime = item.StartTime // hack
	readItem.StopTime = item.StopTime   // hack

	structJSONCompare(t, item, readItem)
}

func TestUpdateItem(t *testing.T) {
	item.Type = "different"

	testUpdate(t, "/items/"+strconv.FormatInt(item.ID, 10), item)

	structJSONCompare(t, item.Type, "different")
}

func TestReadItem2(t *testing.T) {
	readItem := new(data.Item)

	testRead(t, "/items/"+strconv.FormatInt(item.ID, 10), readItem)

	readItem.StartTime = item.StartTime // hack
	readItem.StopTime = item.StopTime   // hack

	structJSONCompare(t, item, readItem)
}

func TestReadAllItems(t *testing.T) {
	items := make([]*data.Item, 0)

	testRead(t, "/items", &items)

	if len(items) < 1 {
		t.Error("ReadAll Failed")
	}
}

func TestCreateItemComment(t *testing.T) {
	comment = new(data.ItemComment)
	comment.ItemID = item.ID
	comment.Author = user.Username
	comment.Body = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	comment.UpdateTime = time.Now()

	testCreate(t, "/items/"+strconv.FormatInt(item.ID, 10)+"/comments", comment)
}

func TestUpdateItemComment(t *testing.T) {
	comment.Body = "[deleted]"

	testUpdate(t, "/items/"+strconv.FormatInt(item.ID, 10)+"/comments/"+strconv.FormatInt(comment.ID, 10), comment)

	structJSONCompare(t, comment.Body, "[deleted]")
}

func TestReadItemComment(t *testing.T) {
	readComment := new(data.ItemComment)

	testRead(t, "/items/"+strconv.FormatInt(item.ID, 10)+"/comments/"+strconv.FormatInt(comment.ID, 10), readComment)

	readComment.UpdateTime = comment.UpdateTime

	structJSONCompare(t, comment, readComment)
}

func TestReadAllComment(t *testing.T) {
	comments := make([]*data.ItemComment, 0)

	testRead(t, "/items/"+strconv.FormatInt(item.ID, 10)+"/comments", &comments)

	if len(comments) < 1 {
		t.Error("ReadAll Failed")
	}
}

func TestCreateTerm(t *testing.T) {
	term = new(data.Term)
	term.Term = "testterm"
	term.Automated = false
	term.Author = user.Username
	term.UpdateTime = time.Now()

	testCreate(t, "/taxonomy", term)
}

func TestReadTerm(t *testing.T) {
	readTerm := new(data.Term)

	testRead(t, "/taxonomy/"+term.Term, readTerm)

	readTerm.UpdateTime = term.UpdateTime

	structJSONCompare(t, term, readTerm)
}

func TestAddTermToItem(t *testing.T) {
	item.Terms = make([]*data.Term, 1)
	item.Terms[0] = term

	testUpdate(t, "/items/"+strconv.FormatInt(item.ID, 10), item)
}

func TestReadItem3(t *testing.T) {
	readItem := new(data.Item)

	testRead(t, "/items/"+strconv.FormatInt(item.ID, 10), readItem)

	readItem.StartTime = item.StartTime // hack
	readItem.StopTime = item.StopTime   // hack

	if len(readItem.Terms) < 1 {
		t.Error("There should be a term here")
	}

	readItem.Terms[0].UpdateTime = term.UpdateTime

	structJSONCompare(t, item, readItem)
}

func TestReadAllTerm(t *testing.T) {
	terms := make([]*data.Term, 0)

	testRead(t, "/taxonomy", &terms)

	if len(terms) < 1 {
		t.Error("ReadAll Failed")
	}
}

func TestUpdateTermRanking(t *testing.T) {
	ranking = new(data.TermRanking)
	ranking.Term = term.Term
	ranking.Ranking = "1,2,3,4,5"
	ranking.UpdateTime = time.Now()

	testUpdate(t, "/taxonomy/"+term.Term+"/ranking", ranking)
}

func TestReadTermRanking(t *testing.T) {
	readRanking := new(data.TermRanking)

	testRead(t, "/taxonomy/"+term.Term+"/ranking", readRanking)

	readRanking.UpdateTime = ranking.UpdateTime
	structJSONCompare(t, ranking, readRanking)
}

func TestDeleteTerm(t *testing.T) {
	testDelete(t, "/taxonomy/"+term.Term)
}

func TestDeleteItemComment(t *testing.T) {
	testDelete(t, "/items/"+strconv.FormatInt(item.ID, 10)+"/comments/"+strconv.FormatInt(comment.ID, 10))
}

func TestDeleteItem(t *testing.T) {
	testDelete(t, "/items/"+strconv.FormatInt(item.ID, 10))
}

func TestDeleteUser(t *testing.T) {
	testDelete(t, "/users/"+user.Username)
}

func TestDeleteRole(t *testing.T) {
	testDelete(t, "/roles/"+role.Title)
}

func testCreate(t *testing.T, url string, v interface{}) {
	enc, err := json.Marshal(v)

	if err != nil {
		t.Error(err)
	}

	testflight.WithServer(
		CRUDRouter(),
		func(r *testflight.Requester) {
			response := r.Post(url, testflight.JSON, string(enc))

			structJSONCompare(t, 201, response.StatusCode)

			err = json.Unmarshal([]byte(response.Body), v)

			if err != nil {
				t.Error(err)
			}
		},
	)
}

func testUpdate(t *testing.T, url string, v interface{}) {
	enc, err := json.Marshal(v)

	if err != nil {
		t.Error(err)
	}

	testflight.WithServer(
		CRUDRouter(),
		func(r *testflight.Requester) {
			response := r.Put(url, testflight.JSON, string(enc))

			structJSONCompare(t, http.StatusOK, response.StatusCode)

			err = json.Unmarshal([]byte(response.Body), v)

			if err != nil {
				t.Error(err)
			}
		},
	)
}

func testRead(t *testing.T, url string, v interface{}) {
	testflight.WithServer(
		CRUDRouter(),
		func(r *testflight.Requester) {
			response := r.Get(url)

			structJSONCompare(t, http.StatusOK, response.StatusCode)

			err := json.Unmarshal([]byte(response.Body), v)

			if err != nil {
				t.Error(err)
			}
		},
	)
}

func testDelete(t *testing.T, url string) {
	testflight.WithServer(
		CRUDRouter(),
		func(r *testflight.Requester) {
			response := r.Delete(url, testflight.JSON, "")

			structJSONCompare(t, http.StatusNoContent, response.StatusCode)
		},
	)
}

func structJSONCompare(t *testing.T, a interface{}, b interface{}) {
	j1, _ := json.Marshal(a)
	j2, _ := json.Marshal(b)

	// t.Log(string(j1), string(j2))

	if string(j1) != string(j2) {
		t.Error("Objects are not equal:")
		t.Error(string(j1))
		t.Error(string(j2))
	}
}
