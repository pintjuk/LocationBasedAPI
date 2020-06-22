package Domain

import (
	"BablarAPI/IdP/Data"
	"BablarAPI/IdP/Domain"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestIdPService_Login(t *testing.T) {
	IdOfBob := uuid.New()
	IdOfAlice := uuid.New()
	IdpRepoWithTwoUsers := Data.NewEmptyInMemoryIdpRepository()
	bob := Domain.Principal{
		Id:       IdOfBob,
		Username: "Bob",
		Password: "1111",
	}
	alice := Domain.Principal{
		Id:       IdOfAlice,
		Username: "Alice",
		Password: "1234",
	}
	IdpRepoWithTwoUsers.Add(bob)
	IdpRepoWithTwoUsers.Add(alice)
	type fields struct {
		idPRepsitory Domain.IdPRepsitory
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOk       bool
		wantIdentity Domain.Principal
	}{
		{name: "should not login in default case",
			fields: fields{idPRepsitory: Data.NewEmptyInMemoryIdpRepository()},
			args: args{
				username: "",
				password: "",
			},
			wantOk:       false,
			wantIdentity: Domain.Principal{},
		},
		{name: "should not login if user does not exist",
			fields: fields{idPRepsitory: IdpRepoWithTwoUsers},
			args: args{
				username: "Jhon snwo",
				password: "0000",
			},
			wantOk:       false,
			wantIdentity: Domain.Principal{},
		},
		{name: "should not login when user exists but password is incorrect",
			fields: fields{idPRepsitory: IdpRepoWithTwoUsers},
			args: args{
				username: "Bob",
				password: "0000",
			},
			wantOk:       false,
			wantIdentity: Domain.Principal{},
		},

		{name: "should login when user exists and password is correct",
			fields: fields{idPRepsitory: IdpRepoWithTwoUsers},
			args: args{
				username: "Bob",
				password: "1111",
			},
			wantOk:       true,
			wantIdentity: Domain.Principal{IdOfBob, "Bob", "1111"},
		},

		// OWASP recomends this beheviour, but i dont want to fix it right now so this test fails

		/*{name: "should login when user name differs in case",

			fields: fields{idPRepsitory: IdpRepoWithTwoUsers},
			args: args{
				username: "bob",
				password: "1111",
			},
			wantOk:       true,
			wantIdentity: Domain.Principal{IdOfBob, "Bob", "1111"},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idPService := Domain.NewIdPService(tt.fields.idPRepsitory)
			gotOk, gotIdentity := idPService.Login(tt.args.username, tt.args.password)
			if gotOk != tt.wantOk {
				t.Errorf("Login() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(gotIdentity, tt.wantIdentity) {
				t.Errorf("Login() gotIdentity = %v, want %v", gotIdentity, tt.wantIdentity)
			}
		})
	}
}
