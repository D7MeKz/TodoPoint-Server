package mysqlS_test

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"modules/common/testutils"
	"modules/d7mysql/ent"
	"modules/d7mysql/ent/enttest"
	"testing"
	"todopoint/auth/v2/data"
	"todopoint/auth/v2/spi/mysqlS"
)

func TestUserStore_IsExist(t *testing.T) {
	ctx := testutils.GetTestGinContext()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			t.Log(err)
		}
	}(client)

	// test
	tests := []struct {
		description string
		input       interface{}
		isExist     bool
	}{
		{
			description: "Success : User Exist",
			input: data.Credential{
				Email:    "Test",
				Password: "Test",
			},
			isExist: true,
		},
		{
			description: "Fail#1 : Invalid Email",
			input: data.Credential{
				Email:    "Test22",
				Password: "Test",
			},
			isExist: false,
		},
		{
			description: "Fail#2 : Invalid Password",
			input: data.Credential{
				Email:    "Test",
				Password: "Test22",
			},
			isExist: false,
		},
		{
			description: "Fail#3 : Empty value",
			input: data.Credential{
				Email:    "",
				Password: "Test",
			},
			isExist: false,
		},
	}

	// given
	testD := data.Credential{
		Email:    "Test",
		Password: "Test",
	}
	_, err := client.User.Create().SetEmail(testD.Email).SetPassword(testD.Password).Save(ctx)
	if err != nil {
		t.Log(err)
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			ok, _ := mysqlS.NewUserStore(client).IsExist(ctx, tc.input)
			assert.Equal(t, tc.isExist, ok)
		})
	}

}

func TestUserStore_Create(t *testing.T) {
	// Setup test
	ctx := testutils.GetTestGinContext()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			t.Log(err)
		}
	}(client)

	tests := []struct {
		description string
		input       interface{}
		isError     bool
	}{
		{
			description: "Success",
			input: data.RegisterRequest{
				Email:    "Test",
				Password: "Test",
			},
			isError: false,
		},
		{
			description: "Fail#1 Empty value",
			input: data.RegisterRequest{
				Email:    "",
				Password: "",
			},
			isError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			err := mysqlS.NewUserStore(client).Create(ctx, tc.input)
			if tc.isError {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestUserStore_GetId(t *testing.T) {
	// Setup test
	ctx := testutils.GetTestGinContext()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			t.Log(err)
		}
	}(client)

	// test
	tests := []struct {
		description string
		input       interface{}
		expected    int
		isError     bool
	}{
		{
			description: "Success",
			input: data.Credential{
				Email:    "Test",
				Password: "Test",
			},
			expected: 1,
			isError:  false,
		},
		{
			description: "Fail#1 : Invalid Email",
			input: data.Credential{
				Email:    "Test22",
				Password: "Test",
			},
			expected: 0,
			isError:  true,
		},
		{
			description: "Fail#2 : Invalid Password",
			input: data.Credential{
				Email:    "Test",
				Password: "Test22",
			},
			expected: 0,
			isError:  true,
		},
		{
			description: "Fail#3 : Empty value",
			input: data.Credential{
				Email:    "",
				Password: "Test",
			},
			expected: 0,
			isError:  true,
		},
	}

	// given
	testD := data.Credential{
		Email:    "Test",
		Password: "Test",
	}
	_, err := client.User.Create().SetEmail(testD.Email).SetPassword(testD.Password).Save(ctx)
	if err != nil {
		t.Log(err)
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			id, err := mysqlS.NewUserStore(client).GetId(ctx, tc.input)
			if tc.isError {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tc.expected, id)
			}
		})
	}
}
