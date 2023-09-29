package reviewbot

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Templates is a collection of templates.
type Templates struct {
	pb.UnimplementedTemplatesServer
	// map of template names to text/template.Template objects
	TemplatesMap map[string]*pb.MessageTemplate
}

// Create creates a new template and adds it to the Templates object.
func (t *Templates) Create(ctx context.Context, in *pb.MessageTemplate) (*pb.TemplateResponse, error) {
	// When calling Create, the template should not already exist.
	if _, ok := t.TemplatesMap[in.Name]; ok {
		return nil, status.Error(
			codes.AlreadyExists, fmt.Sprintf("template with name '%s' already exists", in.Name))
	}

	// Aside from overriding existing templates, Create is the same as Update.
	return t.Update(ctx, in)
}

// ListTemplates returns all templates in the Templates object.
func (t *Templates) List(ctx context.Context, _ *pb.ListTemplatesRequest) (*pb.TemplateResponse, error) {
	var templates []*pb.MessageTemplate
	for _, v := range t.TemplatesMap {
		templates = append(templates, v)
	}

	return &pb.TemplateResponse{Templates: templates}, nil
}

// Update creates or updates a template in the Templates object.
func (t *Templates) Update(ctx context.Context, in *pb.MessageTemplate) (*pb.TemplateResponse, error) {
	var e string
	switch {
	case in == nil:
		e = "MessageTemplate is nil"
	case in.Name == "":
		e = "MessageTemplate.Name is empty"
	case in.Content == "":
		e = "MessageTemplate.Content is empty"
	}

	if e != "" {
		return nil, status.Error(codes.InvalidArgument, e)
	}

	// Update template in the TemplatesMap.
	t.TemplatesMap[in.Name] = in

	// Return a list of all templates.
	return t.List(ctx, &pb.ListTemplatesRequest{})
}

// Delete deletes a template from the Templates object.
func (t *Templates) Delete(ctx context.Context, in *pb.DeleteTemplateRequest) (*pb.TemplateResponse, error) {
	// When calling Delete, the template should already exist.
	if _, ok := t.TemplatesMap[in.Name]; !ok {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("template with name '%s' does not exist", in.Name))
	}

	// Delete template from the TemplatesMap.
	delete(t.TemplatesMap, in.Name)

	// Return a list of all remaining templates.
	return t.List(ctx, &pb.ListTemplatesRequest{})
}

// GenerateMessage generates a message from a template.
func (t *Templates) GenerateMessage(templateName, firstName, lastName string) (string, error) {
	tT, ok := t.TemplatesMap[templateName]
	if !ok {
		return "", fmt.Errorf("template %s does not exist", templateName)
	}

	// create a template.Template object from the MessageTemplate.Content
	templ, err := template.New(templateName).Parse(tT.Content)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %s", err.Error())
	}
	// the template.Execute method requires an io.Writer object
	// so we have to write to a buffer to get the desired string
	buf := new(bytes.Buffer)
	err = templ.Execute(buf, &pb.SendMessageRequest{FirstName: firstName, LastName: lastName})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
