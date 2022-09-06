package controller

import (
	"testing"
	"net/http/httptest"
	"encoding/json"

	"github.com/sou1991/learn_golan_g/controller/dto"
	"github.com/sou1991/learn_golan_g/entity"
	"github.com/sou1991/learn_golan_g/tests"

)
//httptestは関数名をTestプレフィックスが必要
func TestNotFoundUser(t *testing.T){
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/usres/", nil)
	
	mock := tests.UserRepositoryMock{}
	uc := NewUsersController{mock}

	uc.Find(w, r)

	body := make([]byte, w.body.Len())
	w.body.Read(body)

	var res dto.User

	json.Unmarshal(body, &res)
	
	if len(res.User) != 0 {
		t.Errorf("Response is %v", res.User)
	}

	//WriteHeaderが明示的に呼ばれない場合、200(http.StatusOK)がトリガーされるっぽい
}
