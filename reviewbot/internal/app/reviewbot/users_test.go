package reviewbot_test

import (
	"reflect"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	rb "github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot"
)

func TestCustomers_AddCustomer(t *testing.T) {
	custDB := &rb.Customers{CustomersMap: make(map[int64]*rb.Customer)}

	tests := []struct {
		name    string
		tgmsg   *tgbotapi.Message
		want    *rb.Customer
		wantErr bool
	}{
		{
			name: "Success",
			tgmsg: &tgbotapi.Message{
				From: &tgbotapi.User{FirstName: "John", LastName: "Doe", ID: 1},
				Chat: &tgbotapi.Chat{ID: 12345},
			},
			want:    &rb.Customer{FirstName: "John", LastName: "Doe", ChatID: 12345},
			wantErr: false,
		},
		{
			name: "CustomerAlreadyExists",
			tgmsg: &tgbotapi.Message{
				From: &tgbotapi.User{FirstName: "John", LastName: "Doe", ID: 1},
				Chat: &tgbotapi.Chat{ID: 12345},
			},
			want:    &rb.Customer{FirstName: "John", LastName: "Doe", ChatID: 12345},
			wantErr: false,
		},
		{
			name: "EmptyFirstName",
			tgmsg: &tgbotapi.Message{
				From: &tgbotapi.User{FirstName: "", LastName: "Doe", ID: 2},
				Chat: &tgbotapi.Chat{ID: 12345},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "EmptyLastName",
			tgmsg: &tgbotapi.Message{
				From: &tgbotapi.User{FirstName: "Jane", LastName: "", ID: 2},
				Chat: &tgbotapi.Chat{ID: 12345},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NoUserID",
			tgmsg: &tgbotapi.Message{
				From: &tgbotapi.User{FirstName: "Jane", LastName: "Doe"},
				Chat: &tgbotapi.Chat{ID: 12345},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NoChatID",
			tgmsg: &tgbotapi.Message{
				From: &tgbotapi.User{FirstName: "Jane", LastName: "Doe", ID: 2},
				Chat: &tgbotapi.Chat{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := custDB.AddCustomer(tt.tgmsg); (err != nil) != tt.wantErr {
				t.Errorf("Customers.AddCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(custDB.CustomersMap[tt.tgmsg.From.ID], tt.want) {
				t.Errorf("Customers.AddCustomer() = %v, want %v", custDB.CustomersMap[tt.tgmsg.Chat.ID], tt.want)
			}
		})
	}
}

func TestCustomers_GetCustomerByName(t *testing.T) {
	custDB := &rb.Customers{CustomersMap: make(map[int64]*rb.Customer)}
	customer := &rb.Customer{FirstName: "John", LastName: "Doe", ChatID: 12345}
	custDB.CustomersMap[1] = customer

	type args struct {
		firstName string
		lastName  string
	}

	tests := []struct {
		name   string
		args   args
		want   rb.Customer
	}{
		{
			name: "Success",
			args: args{firstName: "John", lastName: "Doe"},
			want: *customer,
		},
		{
			name: "CustomerDoesNotExist",
			args: args{firstName: "Jane", lastName: "Doe"},
			want: rb.Customer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := custDB.GetCustomerByName(tt.args.firstName, tt.args.lastName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Customers.GetCustomerByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
