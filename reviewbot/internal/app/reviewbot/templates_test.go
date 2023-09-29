package reviewbot_test

import (
	"context"
	"reflect"
	"testing"

	rb "github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot"
	"github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot/pb"
)

func TestTemplates_Create(t *testing.T) {
	testDB := &rb.Templates{TemplatesMap: map[string]*pb.MessageTemplate{}}
	msgTempl := &pb.MessageTemplate{Name: "new-template", Content: "test"}

	tests := []struct {
		name    string
		want    *pb.TemplateResponse
		wantErr bool
	}{
		{
			name:    "Success",
			want:    &pb.TemplateResponse{Templates: []*pb.MessageTemplate{msgTempl}},
			wantErr: false,
		},
		{
			name:    "AlreadyExists",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.Create(context.Background(), msgTempl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Templates.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Templates.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplates_List(t *testing.T) {
	templ1 := &pb.MessageTemplate{Name: "test", Content: "test"}
	testDB := &rb.Templates{TemplatesMap: map[string]*pb.MessageTemplate{"test": templ1}}
	want := &pb.TemplateResponse{Templates: []*pb.MessageTemplate{templ1}}

	got, err := testDB.List(context.Background(), nil)
	if err != nil {
		t.Errorf("Templates.List() error = %v", err)
		return
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Templates.List() = %v, want %v", got, want)
	}
}

func TestTemplates_Update(t *testing.T) {
	msgTempl := &pb.MessageTemplate{Name: "test", Content: "test"}
	testDB := &rb.Templates{TemplatesMap: map[string]*pb.MessageTemplate{"test": msgTempl}}

	tests := []struct {
		name    string
		mtempl  *pb.MessageTemplate
		want    *pb.TemplateResponse
		wantErr bool
	}{
		{
			name:    "Success",
			mtempl:  &pb.MessageTemplate{Name: "test", Content: "updated content"},
			want:    &pb.TemplateResponse{Templates: []*pb.MessageTemplate{{Name: "test", Content: "updated content"}}},
			wantErr: false,
		},
		{
			name:    "EmptyTemplate",
			mtempl:  nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "EmptyName",
			mtempl:  &pb.MessageTemplate{Name: "", Content: "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "EmptyContent",
			mtempl:  &pb.MessageTemplate{Name: "test", Content: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.Update(context.Background(), tt.mtempl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Templates.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Templates.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplates_Delete(t *testing.T) {
	msgTempl1 := &pb.MessageTemplate{Name: "test", Content: "test"}
	msgTempl2 := &pb.MessageTemplate{Name: "test2", Content: "test2"}
	testDB := &rb.Templates{TemplatesMap: map[string]*pb.MessageTemplate{"test": msgTempl1, "test2": msgTempl2}}

	tests := []struct {
		name    string
		want    *pb.TemplateResponse
		wantErr bool
	}{
		{
			name:    "Success",
			want:    &pb.TemplateResponse{Templates: []*pb.MessageTemplate{msgTempl2}},
			wantErr: false,
		},
		{
			name:    "NotFound",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.Delete(context.Background(), &pb.DeleteTemplateRequest{Name: "test"})
			if (err != nil) != tt.wantErr {
				t.Errorf("Templates.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Templates.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplates_GenerateMessage(t *testing.T) {
	msgTempl := &pb.MessageTemplate{Name: "test", Content: "Hello {{.FirstName}} {{.LastName}}!"}
	testDB := &rb.Templates{TemplatesMap: map[string]*pb.MessageTemplate{"test": msgTempl}}

	type args struct {
		templateName string
		firstName    string
		lastName     string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{templateName: "test", firstName: "John", lastName: "Doe"},
			want:    "Hello John Doe!",
			wantErr: false,
		},
		{
			name:    "TemplateNotFound",
			args:    args{templateName: "not-found", firstName: "John", lastName: "Doe"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.GenerateMessage(tt.args.templateName, tt.args.firstName, tt.args.lastName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Templates.GenerateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Templates.GenerateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
