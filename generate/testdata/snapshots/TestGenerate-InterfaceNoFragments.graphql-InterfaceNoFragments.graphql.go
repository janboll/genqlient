// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNoFragmentsQueryRandomItemArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryRandomItemArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemArticle) GetTypename() string { return v.Typename }

// GetId returns InterfaceNoFragmentsQueryRandomItemArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemArticle) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRandomItemArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemArticle) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryRandomItemContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNoFragmentsQueryRandomItemContent is implemented by the following types:
// InterfaceNoFragmentsQueryRandomItemArticle
// InterfaceNoFragmentsQueryRandomItemVideo
// InterfaceNoFragmentsQueryRandomItemTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryRandomItemContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *InterfaceNoFragmentsQueryRandomItemArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}
func (v *InterfaceNoFragmentsQueryRandomItemVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}
func (v *InterfaceNoFragmentsQueryRandomItemTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}

func __unmarshalInterfaceNoFragmentsQueryRandomItemContent(b []byte, v *InterfaceNoFragmentsQueryRandomItemContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNoFragmentsQueryRandomItemArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryRandomItemVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryRandomItemTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceNoFragmentsQueryRandomItemContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceNoFragmentsQueryRandomItemContent(v *InterfaceNoFragmentsQueryRandomItemContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceNoFragmentsQueryRandomItemArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryRandomItemArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNoFragmentsQueryRandomItemVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryRandomItemVideo
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNoFragmentsQueryRandomItemTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryRandomItemTopic
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceNoFragmentsQueryRandomItemContent: "%T"`, v)
	}
}

// InterfaceNoFragmentsQueryRandomItemTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryRandomItemTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemTopic) GetTypename() string { return v.Typename }

// GetId returns InterfaceNoFragmentsQueryRandomItemTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRandomItemTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemTopic) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryRandomItemVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryRandomItemVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemVideo) GetTypename() string { return v.Typename }

// GetId returns InterfaceNoFragmentsQueryRandomItemVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemVideo) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRandomItemVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemVideo) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryRandomItemWithTypeNameContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNoFragmentsQueryRandomItemWithTypeNameContent is implemented by the following types:
// InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle
// InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo
// InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent() {
}
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent() {
}
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent() {
}

func __unmarshalInterfaceNoFragmentsQueryRandomItemWithTypeNameContent(b []byte, v *InterfaceNoFragmentsQueryRandomItemWithTypeNameContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceNoFragmentsQueryRandomItemWithTypeNameContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceNoFragmentsQueryRandomItemWithTypeNameContent(v *InterfaceNoFragmentsQueryRandomItemWithTypeNameContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceNoFragmentsQueryRandomItemWithTypeNameContent: "%T"`, v)
	}
}

// InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryResponse is returned by InterfaceNoFragmentsQuery on success.
type InterfaceNoFragmentsQueryResponse struct {
	Root                   InterfaceNoFragmentsQueryRootTopic                     `json:"root"`
	RandomItem             InterfaceNoFragmentsQueryRandomItemContent             `json:"-"`
	RandomItemWithTypeName InterfaceNoFragmentsQueryRandomItemWithTypeNameContent `json:"-"`
	WithPointer            *InterfaceNoFragmentsQueryWithPointerContent           `json:"-"`
}

// GetRoot returns InterfaceNoFragmentsQueryResponse.Root, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryResponse) GetRoot() InterfaceNoFragmentsQueryRootTopic {
	return v.Root
}

// GetRandomItem returns InterfaceNoFragmentsQueryResponse.RandomItem, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryResponse) GetRandomItem() InterfaceNoFragmentsQueryRandomItemContent {
	return v.RandomItem
}

// GetRandomItemWithTypeName returns InterfaceNoFragmentsQueryResponse.RandomItemWithTypeName, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryResponse) GetRandomItemWithTypeName() InterfaceNoFragmentsQueryRandomItemWithTypeNameContent {
	return v.RandomItemWithTypeName
}

// GetWithPointer returns InterfaceNoFragmentsQueryResponse.WithPointer, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryResponse) GetWithPointer() *InterfaceNoFragmentsQueryWithPointerContent {
	return v.WithPointer
}

func (v *InterfaceNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InterfaceNoFragmentsQueryResponse
		RandomItem             json.RawMessage `json:"randomItem"`
		RandomItemWithTypeName json.RawMessage `json:"randomItemWithTypeName"`
		WithPointer            json.RawMessage `json:"withPointer"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceNoFragmentsQueryResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomItem
		src := firstPass.RandomItem
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalInterfaceNoFragmentsQueryRandomItemContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal InterfaceNoFragmentsQueryResponse.RandomItem: %w", err)
			}
		}
	}

	{
		dst := &v.RandomItemWithTypeName
		src := firstPass.RandomItemWithTypeName
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalInterfaceNoFragmentsQueryRandomItemWithTypeNameContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal InterfaceNoFragmentsQueryResponse.RandomItemWithTypeName: %w", err)
			}
		}
	}

	{
		dst := &v.WithPointer
		src := firstPass.WithPointer
		if len(src) != 0 && string(src) != "null" {
			*dst = new(InterfaceNoFragmentsQueryWithPointerContent)
			err = __unmarshalInterfaceNoFragmentsQueryWithPointerContent(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal InterfaceNoFragmentsQueryResponse.WithPointer: %w", err)
			}
		}
	}
	return nil
}

type __premarshalInterfaceNoFragmentsQueryResponse struct {
	Root InterfaceNoFragmentsQueryRootTopic `json:"root"`

	RandomItem json.RawMessage `json:"randomItem"`

	RandomItemWithTypeName json.RawMessage `json:"randomItemWithTypeName"`

	WithPointer json.RawMessage `json:"withPointer"`
}

func (v *InterfaceNoFragmentsQueryResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InterfaceNoFragmentsQueryResponse) __premarshalJSON() (*__premarshalInterfaceNoFragmentsQueryResponse, error) {
	var retval __premarshalInterfaceNoFragmentsQueryResponse

	retval.Root = v.Root
	{

		dst := &retval.RandomItem
		src := v.RandomItem
		var err error
		*dst, err = __marshalInterfaceNoFragmentsQueryRandomItemContent(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal InterfaceNoFragmentsQueryResponse.RandomItem: %w", err)
		}
	}
	{

		dst := &retval.RandomItemWithTypeName
		src := v.RandomItemWithTypeName
		var err error
		*dst, err = __marshalInterfaceNoFragmentsQueryRandomItemWithTypeNameContent(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal InterfaceNoFragmentsQueryResponse.RandomItemWithTypeName: %w", err)
		}
	}
	{

		dst := &retval.WithPointer
		src := v.WithPointer
		if src != nil {
			var err error
			*dst, err = __marshalInterfaceNoFragmentsQueryWithPointerContent(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal InterfaceNoFragmentsQueryResponse.WithPointer: %w", err)
			}
		}
	}
	return &retval, nil
}

// InterfaceNoFragmentsQueryRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRootTopic struct {
	// ID is documented in the Content interface.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns InterfaceNoFragmentsQueryRootTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRootTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryRootTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryRootTopic) GetName() string { return v.Name }

// InterfaceNoFragmentsQueryWithPointerArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryWithPointerArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryWithPointerArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerArticle) GetTypename() string { return v.Typename }

// GetId returns InterfaceNoFragmentsQueryWithPointerArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerArticle) GetId() *testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryWithPointerArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerArticle) GetName() *string { return v.Name }

// InterfaceNoFragmentsQueryWithPointerContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNoFragmentsQueryWithPointerContent is implemented by the following types:
// InterfaceNoFragmentsQueryWithPointerArticle
// InterfaceNoFragmentsQueryWithPointerVideo
// InterfaceNoFragmentsQueryWithPointerTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryWithPointerContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() *testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() *string
}

func (v *InterfaceNoFragmentsQueryWithPointerArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}
func (v *InterfaceNoFragmentsQueryWithPointerVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}
func (v *InterfaceNoFragmentsQueryWithPointerTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}

func __unmarshalInterfaceNoFragmentsQueryWithPointerContent(b []byte, v *InterfaceNoFragmentsQueryWithPointerContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNoFragmentsQueryWithPointerArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryWithPointerVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryWithPointerTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceNoFragmentsQueryWithPointerContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceNoFragmentsQueryWithPointerContent(v *InterfaceNoFragmentsQueryWithPointerContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceNoFragmentsQueryWithPointerArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryWithPointerArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNoFragmentsQueryWithPointerVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryWithPointerVideo
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNoFragmentsQueryWithPointerTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNoFragmentsQueryWithPointerTopic
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceNoFragmentsQueryWithPointerContent: "%T"`, v)
	}
}

// InterfaceNoFragmentsQueryWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryWithPointerTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryWithPointerTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerTopic) GetTypename() string { return v.Typename }

// GetId returns InterfaceNoFragmentsQueryWithPointerTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerTopic) GetId() *testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryWithPointerTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerTopic) GetName() *string { return v.Name }

// InterfaceNoFragmentsQueryWithPointerVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryWithPointerVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// GetTypename returns InterfaceNoFragmentsQueryWithPointerVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerVideo) GetTypename() string { return v.Typename }

// GetId returns InterfaceNoFragmentsQueryWithPointerVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerVideo) GetId() *testutil.ID { return v.Id }

// GetName returns InterfaceNoFragmentsQueryWithPointerVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceNoFragmentsQueryWithPointerVideo) GetName() *string { return v.Name }

func InterfaceNoFragmentsQuery(
	client graphql.Client,
) (*InterfaceNoFragmentsQueryResponse, error) {
	req := &graphql.Payload{
		OpName: "InterfaceNoFragmentsQuery",
		Query: `
query InterfaceNoFragmentsQuery {
	root {
		id
		name
	}
	randomItem {
		__typename
		id
		name
	}
	randomItemWithTypeName: randomItem {
		__typename
		id
		name
	}
	withPointer: randomItem {
		__typename
		id
		name
	}
}
`,
	}
	var err error

	resp := &graphql.Response{
		Data: &InterfaceNoFragmentsQueryResponse{},
	}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	retval := resp.Data.(*InterfaceNoFragmentsQueryResponse)
	return retval, err
}

