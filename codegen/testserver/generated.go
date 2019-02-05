package testserver

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"

	introspection1 "github.com/99designs/gqlgen/codegen/testserver/introspection"
	"github.com/99designs/gqlgen/codegen/testserver/invalid-packagename"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// region    ************************** generated!.gotpl **************************

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	ForcedResolver() ForcedResolverResolver
	ModelMethods() ModelMethodsResolver
	Query() QueryResolver
	Subscription() SubscriptionResolver
	User() UserResolver
}

type DirectiveRoot struct {
	Length func(ctx context.Context, obj interface{}, next graphql.Resolver, min int, max *int, message string) (res interface{}, err error)

	Range func(ctx context.Context, obj interface{}, next graphql.Resolver, min *int, max *int, message *string) (res interface{}, err error)
}

type ComplexityRoot struct {
	Circle struct {
		Radius func(childComplexity int) int
		Area   func(childComplexity int) int
	}

	EmbeddedPointer struct {
		Id    func(childComplexity int) int
		Title func(childComplexity int) int
	}

	Error struct {
		Id                      func(childComplexity int) int
		ErrorOnNonRequiredField func(childComplexity int) int
		ErrorOnRequiredField    func(childComplexity int) int
		NilOnRequiredField      func(childComplexity int) int
	}

	ForcedResolver struct {
		Field func(childComplexity int) int
	}

	InnerObject struct {
		Id func(childComplexity int) int
	}

	InvalidIdentifier struct {
		Id func(childComplexity int) int
	}

	It struct {
		Id func(childComplexity int) int
	}

	ModelMethods struct {
		ResolverField func(childComplexity int) int
		NoContext     func(childComplexity int) int
		WithContext   func(childComplexity int) int
	}

	OuterObject struct {
		Inner func(childComplexity int) int
	}

	Query struct {
		InvalidIdentifier      func(childComplexity int) int
		Collision              func(childComplexity int) int
		MapInput               func(childComplexity int, input map[string]interface{}) int
		Recursive              func(childComplexity int, input *RecursiveInputSlice) int
		NestedInputs           func(childComplexity int, input [][]*OuterInput) int
		NestedOutputs          func(childComplexity int) int
		Keywords               func(childComplexity int, input *Keywords) int
		Shapes                 func(childComplexity int) int
		ErrorBubble            func(childComplexity int) int
		ModelMethods           func(childComplexity int) int
		Valid                  func(childComplexity int) int
		User                   func(childComplexity int, id int) int
		NullableArg            func(childComplexity int, arg *int) int
		DirectiveArg           func(childComplexity int, arg string) int
		DirectiveNullableArg   func(childComplexity int, arg *int, arg2 *int) int
		DirectiveInputNullable func(childComplexity int, arg *InputDirectives) int
		DirectiveInput         func(childComplexity int, arg InputDirectives) int
		KeywordArgs            func(childComplexity int, breakArg string, defaultArg string, funcArg string, interfaceArg string, selectArg string, caseArg string, deferArg string, goArg string, mapArg string, structArg string, chanArg string, elseArg string, gotoArg string, packageArg string, switchArg string, constArg string, fallthroughArg string, ifArg string, rangeArg string, typeArg string, continueArg string, forArg string, importArg string, returnArg string, varArg string) int
	}

	Rectangle struct {
		Length func(childComplexity int) int
		Width  func(childComplexity int) int
		Area   func(childComplexity int) int
	}

	Subscription struct {
		Updated     func(childComplexity int) int
		InitPayload func(childComplexity int) int
	}

	User struct {
		Id      func(childComplexity int) int
		Friends func(childComplexity int) int
	}
}

type ForcedResolverResolver interface {
	Field(ctx context.Context, obj *ForcedResolver) (*Circle, error)
}
type ModelMethodsResolver interface {
	ResolverField(ctx context.Context, obj *ModelMethods) (bool, error)
}
type QueryResolver interface {
	InvalidIdentifier(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error)
	Collision(ctx context.Context) (*introspection1.It, error)
	MapInput(ctx context.Context, input map[string]interface{}) (*bool, error)
	Recursive(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
	NestedInputs(ctx context.Context, input [][]*OuterInput) (*bool, error)
	NestedOutputs(ctx context.Context) ([][]*OuterObject, error)
	Keywords(ctx context.Context, input *Keywords) (bool, error)
	Shapes(ctx context.Context) ([]Shape, error)
	ErrorBubble(ctx context.Context) (*Error, error)
	ModelMethods(ctx context.Context) (*ModelMethods, error)
	Valid(ctx context.Context) (string, error)
	User(ctx context.Context, id int) (User, error)
	NullableArg(ctx context.Context, arg *int) (*string, error)
	DirectiveArg(ctx context.Context, arg string) (*string, error)
	DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int) (*string, error)
	DirectiveInputNullable(ctx context.Context, arg *InputDirectives) (*string, error)
	DirectiveInput(ctx context.Context, arg InputDirectives) (*string, error)
	KeywordArgs(ctx context.Context, breakArg string, defaultArg string, funcArg string, interfaceArg string, selectArg string, caseArg string, deferArg string, goArg string, mapArg string, structArg string, chanArg string, elseArg string, gotoArg string, packageArg string, switchArg string, constArg string, fallthroughArg string, ifArg string, rangeArg string, typeArg string, continueArg string, forArg string, importArg string, returnArg string, varArg string) (bool, error)
}
type SubscriptionResolver interface {
	Updated(ctx context.Context) (<-chan string, error)
	InitPayload(ctx context.Context) (<-chan string, error)
}
type UserResolver interface {
	Friends(ctx context.Context, obj *User) ([]User, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	switch typeName + "." + field {

	case "Circle.radius":
		if e.complexity.Circle.Radius == nil {
			break
		}

		return e.complexity.Circle.Radius(childComplexity), true

	case "Circle.area":
		if e.complexity.Circle.Area == nil {
			break
		}

		return e.complexity.Circle.Area(childComplexity), true

	case "EmbeddedPointer.ID":
		if e.complexity.EmbeddedPointer.Id == nil {
			break
		}

		return e.complexity.EmbeddedPointer.Id(childComplexity), true

	case "EmbeddedPointer.Title":
		if e.complexity.EmbeddedPointer.Title == nil {
			break
		}

		return e.complexity.EmbeddedPointer.Title(childComplexity), true

	case "Error.id":
		if e.complexity.Error.Id == nil {
			break
		}

		return e.complexity.Error.Id(childComplexity), true

	case "Error.errorOnNonRequiredField":
		if e.complexity.Error.ErrorOnNonRequiredField == nil {
			break
		}

		return e.complexity.Error.ErrorOnNonRequiredField(childComplexity), true

	case "Error.errorOnRequiredField":
		if e.complexity.Error.ErrorOnRequiredField == nil {
			break
		}

		return e.complexity.Error.ErrorOnRequiredField(childComplexity), true

	case "Error.nilOnRequiredField":
		if e.complexity.Error.NilOnRequiredField == nil {
			break
		}

		return e.complexity.Error.NilOnRequiredField(childComplexity), true

	case "ForcedResolver.field":
		if e.complexity.ForcedResolver.Field == nil {
			break
		}

		return e.complexity.ForcedResolver.Field(childComplexity), true

	case "InnerObject.id":
		if e.complexity.InnerObject.Id == nil {
			break
		}

		return e.complexity.InnerObject.Id(childComplexity), true

	case "InvalidIdentifier.id":
		if e.complexity.InvalidIdentifier.Id == nil {
			break
		}

		return e.complexity.InvalidIdentifier.Id(childComplexity), true

	case "It.id":
		if e.complexity.It.Id == nil {
			break
		}

		return e.complexity.It.Id(childComplexity), true

	case "ModelMethods.resolverField":
		if e.complexity.ModelMethods.ResolverField == nil {
			break
		}

		return e.complexity.ModelMethods.ResolverField(childComplexity), true

	case "ModelMethods.noContext":
		if e.complexity.ModelMethods.NoContext == nil {
			break
		}

		return e.complexity.ModelMethods.NoContext(childComplexity), true

	case "ModelMethods.withContext":
		if e.complexity.ModelMethods.WithContext == nil {
			break
		}

		return e.complexity.ModelMethods.WithContext(childComplexity), true

	case "OuterObject.inner":
		if e.complexity.OuterObject.Inner == nil {
			break
		}

		return e.complexity.OuterObject.Inner(childComplexity), true

	case "Query.invalidIdentifier":
		if e.complexity.Query.InvalidIdentifier == nil {
			break
		}

		return e.complexity.Query.InvalidIdentifier(childComplexity), true

	case "Query.collision":
		if e.complexity.Query.Collision == nil {
			break
		}

		return e.complexity.Query.Collision(childComplexity), true

	case "Query.mapInput":
		if e.complexity.Query.MapInput == nil {
			break
		}

		args, err := e.field_Query_mapInput_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.MapInput(childComplexity, args["input"].(map[string]interface{})), true

	case "Query.recursive":
		if e.complexity.Query.Recursive == nil {
			break
		}

		args, err := e.field_Query_recursive_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Recursive(childComplexity, args["input"].(*RecursiveInputSlice)), true

	case "Query.nestedInputs":
		if e.complexity.Query.NestedInputs == nil {
			break
		}

		args, err := e.field_Query_nestedInputs_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.NestedInputs(childComplexity, args["input"].([][]*OuterInput)), true

	case "Query.nestedOutputs":
		if e.complexity.Query.NestedOutputs == nil {
			break
		}

		return e.complexity.Query.NestedOutputs(childComplexity), true

	case "Query.keywords":
		if e.complexity.Query.Keywords == nil {
			break
		}

		args, err := e.field_Query_keywords_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Keywords(childComplexity, args["input"].(*Keywords)), true

	case "Query.shapes":
		if e.complexity.Query.Shapes == nil {
			break
		}

		return e.complexity.Query.Shapes(childComplexity), true

	case "Query.errorBubble":
		if e.complexity.Query.ErrorBubble == nil {
			break
		}

		return e.complexity.Query.ErrorBubble(childComplexity), true

	case "Query.modelMethods":
		if e.complexity.Query.ModelMethods == nil {
			break
		}

		return e.complexity.Query.ModelMethods(childComplexity), true

	case "Query.valid":
		if e.complexity.Query.Valid == nil {
			break
		}

		return e.complexity.Query.Valid(childComplexity), true

	case "Query.user":
		if e.complexity.Query.User == nil {
			break
		}

		args, err := e.field_Query_user_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.User(childComplexity, args["id"].(int)), true

	case "Query.nullableArg":
		if e.complexity.Query.NullableArg == nil {
			break
		}

		args, err := e.field_Query_nullableArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.NullableArg(childComplexity, args["arg"].(*int)), true

	case "Query.directiveArg":
		if e.complexity.Query.DirectiveArg == nil {
			break
		}

		args, err := e.field_Query_directiveArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveArg(childComplexity, args["arg"].(string)), true

	case "Query.directiveNullableArg":
		if e.complexity.Query.DirectiveNullableArg == nil {
			break
		}

		args, err := e.field_Query_directiveNullableArg_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveNullableArg(childComplexity, args["arg"].(*int), args["arg2"].(*int)), true

	case "Query.directiveInputNullable":
		if e.complexity.Query.DirectiveInputNullable == nil {
			break
		}

		args, err := e.field_Query_directiveInputNullable_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveInputNullable(childComplexity, args["arg"].(*InputDirectives)), true

	case "Query.directiveInput":
		if e.complexity.Query.DirectiveInput == nil {
			break
		}

		args, err := e.field_Query_directiveInput_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.DirectiveInput(childComplexity, args["arg"].(InputDirectives)), true

	case "Query.keywordArgs":
		if e.complexity.Query.KeywordArgs == nil {
			break
		}

		args, err := e.field_Query_keywordArgs_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.KeywordArgs(childComplexity, args["break"].(string), args["default"].(string), args["func"].(string), args["interface"].(string), args["select"].(string), args["case"].(string), args["defer"].(string), args["go"].(string), args["map"].(string), args["struct"].(string), args["chan"].(string), args["else"].(string), args["goto"].(string), args["package"].(string), args["switch"].(string), args["const"].(string), args["fallthrough"].(string), args["if"].(string), args["range"].(string), args["type"].(string), args["continue"].(string), args["for"].(string), args["import"].(string), args["return"].(string), args["var"].(string)), true

	case "Rectangle.length":
		if e.complexity.Rectangle.Length == nil {
			break
		}

		return e.complexity.Rectangle.Length(childComplexity), true

	case "Rectangle.width":
		if e.complexity.Rectangle.Width == nil {
			break
		}

		return e.complexity.Rectangle.Width(childComplexity), true

	case "Rectangle.area":
		if e.complexity.Rectangle.Area == nil {
			break
		}

		return e.complexity.Rectangle.Area(childComplexity), true

	case "Subscription.updated":
		if e.complexity.Subscription.Updated == nil {
			break
		}

		return e.complexity.Subscription.Updated(childComplexity), true

	case "Subscription.initPayload":
		if e.complexity.Subscription.InitPayload == nil {
			break
		}

		return e.complexity.Subscription.InitPayload(childComplexity), true

	case "User.id":
		if e.complexity.User.Id == nil {
			break
		}

		return e.complexity.User.Id(childComplexity), true

	case "User.friends":
		if e.complexity.User.Friends == nil {
			break
		}

		return e.complexity.User.Friends(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	return graphql.ErrorResponse(ctx, "mutations are not supported")
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	next := ec._Subscription(ctx, op.SelectionSet)
	if ec.Errors != nil {
		return graphql.OneShot(&graphql.Response{Data: []byte("null"), Errors: ec.Errors})
	}

	var buf bytes.Buffer
	return func() *graphql.Response {
		buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
			buf.Reset()
			data := next()

			if data == nil {
				return nil
			}
			data.MarshalGQL(&buf)
			return buf.Bytes()
		})

		if buf == nil {
			return nil
		}

		return &graphql.Response{
			Data:       buf,
			Errors:     ec.Errors,
			Extensions: ec.Extensions,
		}
	}
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

func (ec *executionContext) FieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	rctx := graphql.GetResolverContext(ctx)
	for _, d := range rctx.Field.Definition.Directives {
		switch d.Name {
		case "length":
			if ec.directives.Length != nil {
				rawArgs := d.ArgumentMap(ec.Variables)
				args, err := ec.dir_length_args(ctx, rawArgs)
				if err != nil {
					ec.Error(ctx, err)
					return nil
				}
				n := next
				next = func(ctx context.Context) (interface{}, error) {
					return ec.directives.Length(ctx, obj, n, args["min"].(int), args["max"].(*int), args["message"].(string))
				}
			}
		case "range":
			if ec.directives.Range != nil {
				rawArgs := d.ArgumentMap(ec.Variables)
				args, err := ec.dir_range_args(ctx, rawArgs)
				if err != nil {
					ec.Error(ctx, err)
					return nil
				}
				n := next
				next = func(ctx context.Context) (interface{}, error) {
					return ec.directives.Range(ctx, obj, n, args["min"].(*int), args["max"].(*int), args["message"].(*string))
				}
			}
		}
	}
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `type Query {
    invalidIdentifier: InvalidIdentifier
    collision: It
    mapInput(input: Changes): Boolean
    recursive(input: RecursiveInputSlice): Boolean
    nestedInputs(input: [[OuterInput]] = [[{inner: {id: 1}}]]): Boolean
    nestedOutputs: [[OuterObject]]
    keywords(input: Keywords): Boolean!
    shapes: [Shape]
    errorBubble: Error
    modelMethods: ModelMethods
    valid: String!
    user(id: Int!): User!
    nullableArg(arg: Int = 123): String
    directiveArg(arg: String! @length(min:1, max: 255, message: "invalid length")): String
    directiveNullableArg(arg: Int @range(min:0), arg2: Int @range): String
    directiveInputNullable(arg: InputDirectives): String
    directiveInput(arg: InputDirectives!): String
}

type Subscription {
    updated: String!
    initPayload: String!
}

type User {
    id: Int!
    friends: [User!]!
}

type Error {
    id: ID!
    errorOnNonRequiredField: String
    errorOnRequiredField: String!
    nilOnRequiredField: String!
}

type ModelMethods {
    resolverField: Boolean!
    noContext: Boolean!
    withContext: Boolean!
}

type InvalidIdentifier {
    id: Int!
}

type It {
    id: ID!
}

input Changes {
    a: Int
    b: Int
}

input RecursiveInputSlice {
    self: [RecursiveInputSlice!]
}

input InnerInput {
    id:Int!
}

input OuterInput {
    inner: InnerInput!
}

input InputDirectives {
    text: String! @length(min: 0, max: 7, message: "not valid")
    inner: InnerDirectives!
    innerNullable: InnerDirectives
}

input InnerDirectives {
    message: String! @length(min: 1, message: "not valid")
}

type OuterObject {
    inner: InnerObject!
}

type InnerObject {
    id: Int!
}

input Keywords {
    break:       String!
    default:     String!
    func:        String!
    interface:   String!
    select:      String!
    case:        String!
    defer:       String!
    go:          String!
    map:         String!
    struct:      String!
    chan:        String!
    else:        String!
    goto:        String!
    package:     String!
    switch:      String!
    const:       String!
    fallthrough: String!
    if:          String!
    range:       String!
    type:        String!
    continue:    String!
    for:         String!
    import:      String!
    return:      String!
    var:         String!
}

extend type Query {
    keywordArgs(
        break:       String!,
        default:     String!,
        func:        String!,
        interface:   String!,
        select:      String!,
        case:        String!,
        defer:       String!,
        go:          String!,
        map:         String!,
        struct:      String!,
        chan:        String!,
        else:        String!,
        goto:        String!,
        package:     String!,
        switch:      String!,
        const:       String!,
        fallthrough: String!,
        if:          String!,
        range:       String!,
        type:        String!,
        continue:    String!,
        for:         String!,
        import:      String!,
        return:      String!,
        var:         String!,
    ): Boolean!
}

interface Shape {
    area: Float
}
type Circle implements Shape {
    radius: Float
    area: Float
}
type Rectangle implements Shape {
    length: Float
    width: Float
    area: Float
}
union ShapeUnion = Circle | Rectangle

type ForcedResolver {
    field: Circle
}

type EmbeddedPointer {
    ID: String
    Title: String
}

directive @length(min: Int!, max: Int, message: String!) on ARGUMENT_DEFINITION | INPUT_FIELD_DEFINITION
directive @range(min: Int = 0, max: Int, message: String) on ARGUMENT_DEFINITION

enum Status {
    OK
    ERROR
}
`},
)

// ChainFieldMiddleware add chain by FieldMiddleware
// nolint: deadcode
func chainFieldMiddleware(handleFunc ...graphql.FieldMiddleware) graphql.FieldMiddleware {
	n := len(handleFunc)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
			var (
				chainHandler graphql.Resolver
				curI         int
			)
			chainHandler = func(currentCtx context.Context) (interface{}, error) {
				if curI == lastI {
					return next(currentCtx)
				}
				curI++
				res, err := handleFunc[curI](currentCtx, chainHandler)
				curI--
				return res, err

			}
			return handleFunc[0](ctx, chainHandler)
		}
	}

	if n == 1 {
		return handleFunc[0]
	}

	return func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
		return next(ctx)
	}
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (e *executableSchema) dir_length_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["min"]; ok {
		arg0, err = e.unmarshalInt2int(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["min"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["max"]; ok {
		arg1, err = e.unmarshalInt2ᚖint(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["max"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["message"]; ok {
		arg2, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["message"] = arg2
	return args, nil
}

func (e *executableSchema) dir_range_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["min"]; ok {
		arg0, err = e.unmarshalInt2ᚖint(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["min"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["max"]; ok {
		arg1, err = e.unmarshalInt2ᚖint(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["max"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["message"]; ok {
		arg2, err = e.unmarshalString2ᚖstring(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["message"] = arg2
	return args, nil
}

func (e *executableSchema) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		arg0, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_directiveArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["arg"]; ok {
		getArg0 := func(ctx context.Context) (interface{}, error) { return e.unmarshalString2string(tmp) }
		getArg1 := func(ctx context.Context) (res interface{}, err error) {
			max := 255
			n := getArg0
			return e.directives.Length(ctx, tmp, n, 1, &max, "invalid length")
		}

		tmp, err = getArg1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(string); ok {
			arg0 = data
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_directiveInputNullable_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *InputDirectives
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = e.unmarshalInputDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_directiveInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 InputDirectives
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = e.unmarshalInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_directiveNullableArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["arg"]; ok {
		getArg0 := func(ctx context.Context) (interface{}, error) { return e.unmarshalInt2ᚖint(tmp) }
		getArg1 := func(ctx context.Context) (res interface{}, err error) {
			min := 0
			n := getArg0
			return e.directives.Range(ctx, tmp, n, &min, nil, nil)
		}

		tmp, err = getArg1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*int); ok {
			arg0 = data
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *int`, tmp)
		}
	}
	args["arg"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["arg2"]; ok {
		getArg0 := func(ctx context.Context) (interface{}, error) { return e.unmarshalInt2ᚖint(tmp) }
		getArg1 := func(ctx context.Context) (res interface{}, err error) {
			min := 0
			n := getArg0
			return e.directives.Range(ctx, tmp, n, &min, nil, nil)
		}

		tmp, err = getArg1(ctx)
		if err != nil {
			return nil, err
		}
		if data, ok := tmp.(*int); ok {
			arg1 = data
		} else {
			return nil, fmt.Errorf(`unexpected type %T from directive, should be *int`, tmp)
		}
	}
	args["arg2"] = arg1
	return args, nil
}

func (e *executableSchema) field_Query_keywordArgs_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["break"]; ok {
		arg0, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["break"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["default"]; ok {
		arg1, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["default"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["func"]; ok {
		arg2, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["func"] = arg2
	var arg3 string
	if tmp, ok := rawArgs["interface"]; ok {
		arg3, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["interface"] = arg3
	var arg4 string
	if tmp, ok := rawArgs["select"]; ok {
		arg4, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["select"] = arg4
	var arg5 string
	if tmp, ok := rawArgs["case"]; ok {
		arg5, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["case"] = arg5
	var arg6 string
	if tmp, ok := rawArgs["defer"]; ok {
		arg6, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["defer"] = arg6
	var arg7 string
	if tmp, ok := rawArgs["go"]; ok {
		arg7, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["go"] = arg7
	var arg8 string
	if tmp, ok := rawArgs["map"]; ok {
		arg8, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["map"] = arg8
	var arg9 string
	if tmp, ok := rawArgs["struct"]; ok {
		arg9, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["struct"] = arg9
	var arg10 string
	if tmp, ok := rawArgs["chan"]; ok {
		arg10, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["chan"] = arg10
	var arg11 string
	if tmp, ok := rawArgs["else"]; ok {
		arg11, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["else"] = arg11
	var arg12 string
	if tmp, ok := rawArgs["goto"]; ok {
		arg12, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["goto"] = arg12
	var arg13 string
	if tmp, ok := rawArgs["package"]; ok {
		arg13, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["package"] = arg13
	var arg14 string
	if tmp, ok := rawArgs["switch"]; ok {
		arg14, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["switch"] = arg14
	var arg15 string
	if tmp, ok := rawArgs["const"]; ok {
		arg15, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["const"] = arg15
	var arg16 string
	if tmp, ok := rawArgs["fallthrough"]; ok {
		arg16, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["fallthrough"] = arg16
	var arg17 string
	if tmp, ok := rawArgs["if"]; ok {
		arg17, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["if"] = arg17
	var arg18 string
	if tmp, ok := rawArgs["range"]; ok {
		arg18, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["range"] = arg18
	var arg19 string
	if tmp, ok := rawArgs["type"]; ok {
		arg19, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["type"] = arg19
	var arg20 string
	if tmp, ok := rawArgs["continue"]; ok {
		arg20, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["continue"] = arg20
	var arg21 string
	if tmp, ok := rawArgs["for"]; ok {
		arg21, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["for"] = arg21
	var arg22 string
	if tmp, ok := rawArgs["import"]; ok {
		arg22, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["import"] = arg22
	var arg23 string
	if tmp, ok := rawArgs["return"]; ok {
		arg23, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["return"] = arg23
	var arg24 string
	if tmp, ok := rawArgs["var"]; ok {
		arg24, err = e.unmarshalString2string(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["var"] = arg24
	return args, nil
}

func (e *executableSchema) field_Query_keywords_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *Keywords
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = e.unmarshalKeywords2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐKeywords(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_mapInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 map[string]interface{}
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = e.unmarshalChanges2map(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_nestedInputs_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 [][]*OuterInput
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = e.unmarshalOuterInput2ᚕᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_nullableArg_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["arg"]; ok {
		arg0, err = e.unmarshalInt2ᚖint(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["arg"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_recursive_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *RecursiveInputSlice
	if tmp, ok := rawArgs["input"]; ok {
		arg0, err = e.unmarshalRecursiveInputSlice2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (e *executableSchema) field_Query_user_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = e.unmarshalInt2int(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (e *executableSchema) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = e.unmarshalBoolean2bool(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (e *executableSchema) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = e.unmarshalBoolean2bool(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    **************************** field.gotpl *****************************

// nolint: vetshadow
func (ec *executionContext) _Circle_radius(ctx context.Context, field graphql.CollectedField, obj *Circle) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Circle",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Radius, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalFloat(res)
}

// nolint: vetshadow
func (ec *executionContext) _Circle_area(ctx context.Context, field graphql.CollectedField, obj *Circle) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Circle",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Area(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalFloat(res)
}

// nolint: vetshadow
func (ec *executionContext) _EmbeddedPointer_ID(ctx context.Context, field graphql.CollectedField, obj *EmbeddedPointerModel) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "EmbeddedPointer",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _EmbeddedPointer_Title(ctx context.Context, field graphql.CollectedField, obj *EmbeddedPointerModel) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "EmbeddedPointer",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Title, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Error_id(ctx context.Context, field graphql.CollectedField, obj *Error) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Error",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalID(res)
}

// nolint: vetshadow
func (ec *executionContext) _Error_errorOnNonRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Error",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorOnNonRequiredField()
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Error_errorOnRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Error",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorOnRequiredField()
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Error_nilOnRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Error",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NilOnRequiredField(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _ForcedResolver_field(ctx context.Context, field graphql.CollectedField, obj *ForcedResolver) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ForcedResolver",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.ForcedResolver().Field(rctx, obj)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Circle)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Circle(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _InnerObject_id(ctx context.Context, field graphql.CollectedField, obj *InnerObject) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "InnerObject",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalInt(res)
}

// nolint: vetshadow
func (ec *executionContext) _InvalidIdentifier_id(ctx context.Context, field graphql.CollectedField, obj *invalid_packagename.InvalidIdentifier) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "InvalidIdentifier",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalInt(res)
}

// nolint: vetshadow
func (ec *executionContext) _It_id(ctx context.Context, field graphql.CollectedField, obj *introspection1.It) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "It",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalID(res)
}

// nolint: vetshadow
func (ec *executionContext) _ModelMethods_resolverField(ctx context.Context, field graphql.CollectedField, obj *ModelMethods) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ModelMethods",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.ModelMethods().ResolverField(rctx, obj)
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _ModelMethods_noContext(ctx context.Context, field graphql.CollectedField, obj *ModelMethods) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ModelMethods",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NoContext(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _ModelMethods_withContext(ctx context.Context, field graphql.CollectedField, obj *ModelMethods) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "ModelMethods",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.WithContext(ctx), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _OuterObject_inner(ctx context.Context, field graphql.CollectedField, obj *OuterObject) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "OuterObject",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Inner, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(InnerObject)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	return ec._InnerObject(ctx, field.Selections, &res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_invalidIdentifier(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().InvalidIdentifier(rctx)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*invalid_packagename.InvalidIdentifier)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._InvalidIdentifier(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_collision(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Collision(rctx)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection1.It)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._It(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_mapInput(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_mapInput_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().MapInput(rctx, args["input"].(map[string]interface{}))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalBoolean(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_recursive(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_recursive_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Recursive(rctx, args["input"].(*RecursiveInputSlice))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalBoolean(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_nestedInputs(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_nestedInputs_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().NestedInputs(rctx, args["input"].([][]*OuterInput))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalBoolean(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_nestedOutputs(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().NestedOutputs(rctx)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([][]*OuterObject)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				arr2 := make(graphql.Array, len(res[idx1]))

				isLen1 := len(res[idx1]) == 1
				if !isLen1 {
					wg.Add(len(res[idx1]))
				}

				for idx2 := range res[idx1] {
					idx2 := idx2
					rctx := &graphql.ResolverContext{
						Index:  &idx2,
						Result: res[idx1][idx2],
					}
					ctx := graphql.WithResolverContext(ctx, rctx)
					f := func(idx2 int) {
						if !isLen1 {
							defer wg.Done()
						}
						arr2[idx2] = func() graphql.Marshaler {

							if res[idx1][idx2] == nil {
								return graphql.Null
							}

							return ec._OuterObject(ctx, field.Selections, res[idx1][idx2])
						}()
					}
					if isLen1 {
						f(idx2)
					} else {
						go f(idx2)
					}

				}

				return arr2
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) _Query_keywords(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_keywords_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Keywords(rctx, args["input"].(*Keywords))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_shapes(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Shapes(rctx)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]Shape)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec._Shape(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) _Query_errorBubble(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ErrorBubble(rctx)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Error)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._Error(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_modelMethods(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().ModelMethods(rctx)
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ModelMethods)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec._ModelMethods(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_valid(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Valid(rctx)
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_user(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_user_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().User(rctx, args["id"].(int))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(User)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	return ec._User(ctx, field.Selections, &res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_nullableArg(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_nullableArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().NullableArg(rctx, args["arg"].(*int))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_directiveArg(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_directiveArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveArg(rctx, args["arg"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_directiveNullableArg(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_directiveNullableArg_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveNullableArg(rctx, args["arg"].(*int), args["arg2"].(*int))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_directiveInputNullable(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_directiveInputNullable_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveInputNullable(rctx, args["arg"].(*InputDirectives))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_directiveInput(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_directiveInput_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveInput(rctx, args["arg"].(InputDirectives))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) _Query_keywordArgs(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query_keywordArgs_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().KeywordArgs(rctx, args["break"].(string), args["default"].(string), args["func"].(string), args["interface"].(string), args["select"].(string), args["case"].(string), args["defer"].(string), args["go"].(string), args["map"].(string), args["struct"].(string), args["chan"].(string), args["else"].(string), args["goto"].(string), args["package"].(string), args["switch"].(string), args["const"].(string), args["fallthrough"].(string), args["if"].(string), args["range"].(string), args["type"].(string), args["continue"].(string), args["for"].(string), args["import"].(string), args["return"].(string), args["var"].(string))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query___type_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Schema(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Rectangle_length(ctx context.Context, field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Rectangle",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Length, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalFloat(res)
}

// nolint: vetshadow
func (ec *executionContext) _Rectangle_width(ctx context.Context, field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Rectangle",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Width, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalFloat(res)
}

// nolint: vetshadow
func (ec *executionContext) _Rectangle_area(ctx context.Context, field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "Rectangle",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Area(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Subscription_updated(ctx context.Context, field graphql.CollectedField) func() graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Field: field,
		Args:  nil,
	})
	// FIXME: subscriptions are missing request middleware stack https://github.com/99designs/gqlgen/issues/259
	//          and Tracer stack
	rctx := ctx
	results, err := ec.resolvers.Subscription().Updated(rctx)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-results
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			func() graphql.Marshaler {
				return graphql.MarshalString(res)
			}().MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

func (ec *executionContext) _Subscription_initPayload(ctx context.Context, field graphql.CollectedField) func() graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Field: field,
		Args:  nil,
	})
	// FIXME: subscriptions are missing request middleware stack https://github.com/99designs/gqlgen/issues/259
	//          and Tracer stack
	rctx := ctx
	results, err := ec.resolvers.Subscription().InitPayload(rctx)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return func() graphql.Marshaler {
		res, ok := <-results
		if !ok {
			return nil
		}
		return graphql.WriterFunc(func(w io.Writer) {
			w.Write([]byte{'{'})
			graphql.MarshalString(field.Alias).MarshalGQL(w)
			w.Write([]byte{':'})
			func() graphql.Marshaler {
				return graphql.MarshalString(res)
			}().MarshalGQL(w)
			w.Write([]byte{'}'})
		})
	}
}

// nolint: vetshadow
func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *User) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "User",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalInt(res)
}

// nolint: vetshadow
func (ec *executionContext) _User_friends(ctx context.Context, field graphql.CollectedField, obj *User) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "User",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.User().Friends(rctx, obj)
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]User)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec._User(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))

	for idx1 := range res {
		arr1[idx1] = func() graphql.Marshaler {
			return graphql.MarshalString(res[idx1])
		}()
	}

	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Directive(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Field(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx.Args = args
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___EnumValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	ctx = ec.Tracer.StartFieldExecution(ctx, field)
	defer func() { ec.Tracer.EndFieldExecution(ctx) }()
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Field:  field,
		Args:   nil,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	ctx = ec.Tracer.StartFieldResolverExecution(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res
	ctx = ec.Tracer.StartFieldChildExecution(ctx)

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (e *executableSchema) unmarshalInputInnerDirectives(v interface{}) (InnerDirectives, error) {
	var it InnerDirectives
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "message":
			var err error
			it.Message, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (e *executableSchema) unmarshalInputInnerInput(v interface{}) (InnerInput, error) {
	var it InnerInput
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "id":
			var err error
			it.ID, err = graphql.UnmarshalInt(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (e *executableSchema) unmarshalInputInputDirectives(v interface{}) (InputDirectives, error) {
	var it InputDirectives
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "text":
			var err error
			it.Text, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "inner":
			var err error
			it.Inner, err = e.unmarshalInputInnerDirectives(v)
			if err != nil {
				return it, err
			}
		case "innerNullable":
			var err error
			var ptr1 InnerDirectives
			if v != nil {
				ptr1, err = e.unmarshalInputInnerDirectives(v)
				it.InnerNullable = &ptr1
			}

			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (e *executableSchema) unmarshalInputKeywords(v interface{}) (Keywords, error) {
	var it Keywords
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "break":
			var err error
			it.Break, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "default":
			var err error
			it.Default, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "func":
			var err error
			it.Func, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "interface":
			var err error
			it.Interface, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "select":
			var err error
			it.Select, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "case":
			var err error
			it.Case, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "defer":
			var err error
			it.Defer, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "go":
			var err error
			it.Go, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "map":
			var err error
			it.Map, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "struct":
			var err error
			it.Struct, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "chan":
			var err error
			it.Chan, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "else":
			var err error
			it.Else, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "goto":
			var err error
			it.Goto, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "package":
			var err error
			it.Package, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "switch":
			var err error
			it.Switch, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "const":
			var err error
			it.Const, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "fallthrough":
			var err error
			it.Fallthrough, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "if":
			var err error
			it.If, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "range":
			var err error
			it.Range, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "type":
			var err error
			it.Type, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "continue":
			var err error
			it.Continue, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "for":
			var err error
			it.For, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "import":
			var err error
			it.Import, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "return":
			var err error
			it.Return, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "var":
			var err error
			it.Var, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (e *executableSchema) unmarshalInputOuterInput(v interface{}) (OuterInput, error) {
	var it OuterInput
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "inner":
			var err error
			it.Inner, err = e.unmarshalInputInnerInput(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (e *executableSchema) unmarshalInputRecursiveInputSlice(v interface{}) (RecursiveInputSlice, error) {
	var it RecursiveInputSlice
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "self":
			var err error
			var rawIf1 []interface{}
			if v != nil {
				if tmp1, ok := v.([]interface{}); ok {
					rawIf1 = tmp1
				} else {
					rawIf1 = []interface{}{v}
				}
			}
			it.Self = make([]RecursiveInputSlice, len(rawIf1))
			for idx1 := range rawIf1 {
				it.Self[idx1], err = e.unmarshalInputRecursiveInputSlice(rawIf1[idx1])
			}
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Shape(ctx context.Context, sel ast.SelectionSet, obj *Shape) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _ShapeUnion(ctx context.Context, sel ast.SelectionSet, obj *ShapeUnion) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var circleImplementors = []string{"Circle", "Shape"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Circle(ctx context.Context, sel ast.SelectionSet, obj *Circle) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, circleImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Circle")
		case "radius":
			out.Values[i] = ec._Circle_radius(ctx, field, obj)
		case "area":
			out.Values[i] = ec._Circle_area(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var embeddedPointerImplementors = []string{"EmbeddedPointer"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _EmbeddedPointer(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedPointerModel) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, embeddedPointerImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedPointer")
		case "ID":
			out.Values[i] = ec._EmbeddedPointer_ID(ctx, field, obj)
		case "Title":
			out.Values[i] = ec._EmbeddedPointer_Title(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var errorImplementors = []string{"Error"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Error(ctx context.Context, sel ast.SelectionSet, obj *Error) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, errorImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Error")
		case "id":
			out.Values[i] = ec._Error_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "errorOnNonRequiredField":
			out.Values[i] = ec._Error_errorOnNonRequiredField(ctx, field, obj)
		case "errorOnRequiredField":
			out.Values[i] = ec._Error_errorOnRequiredField(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "nilOnRequiredField":
			out.Values[i] = ec._Error_nilOnRequiredField(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var forcedResolverImplementors = []string{"ForcedResolver"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _ForcedResolver(ctx context.Context, sel ast.SelectionSet, obj *ForcedResolver) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, forcedResolverImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ForcedResolver")
		case "field":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._ForcedResolver_field(ctx, field, obj)
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var innerObjectImplementors = []string{"InnerObject"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _InnerObject(ctx context.Context, sel ast.SelectionSet, obj *InnerObject) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, innerObjectImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InnerObject")
		case "id":
			out.Values[i] = ec._InnerObject_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var invalidIdentifierImplementors = []string{"InvalidIdentifier"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _InvalidIdentifier(ctx context.Context, sel ast.SelectionSet, obj *invalid_packagename.InvalidIdentifier) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, invalidIdentifierImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InvalidIdentifier")
		case "id":
			out.Values[i] = ec._InvalidIdentifier_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var itImplementors = []string{"It"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _It(ctx context.Context, sel ast.SelectionSet, obj *introspection1.It) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, itImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("It")
		case "id":
			out.Values[i] = ec._It_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var modelMethodsImplementors = []string{"ModelMethods"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _ModelMethods(ctx context.Context, sel ast.SelectionSet, obj *ModelMethods) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, modelMethodsImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ModelMethods")
		case "resolverField":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._ModelMethods_resolverField(ctx, field, obj)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "noContext":
			out.Values[i] = ec._ModelMethods_noContext(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "withContext":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._ModelMethods_withContext(ctx, field, obj)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var outerObjectImplementors = []string{"OuterObject"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _OuterObject(ctx context.Context, sel ast.SelectionSet, obj *OuterObject) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, outerObjectImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OuterObject")
		case "inner":
			out.Values[i] = ec._OuterObject_inner(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "invalidIdentifier":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_invalidIdentifier(ctx, field)
				return res
			})
		case "collision":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_collision(ctx, field)
				return res
			})
		case "mapInput":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_mapInput(ctx, field)
				return res
			})
		case "recursive":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_recursive(ctx, field)
				return res
			})
		case "nestedInputs":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_nestedInputs(ctx, field)
				return res
			})
		case "nestedOutputs":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_nestedOutputs(ctx, field)
				return res
			})
		case "keywords":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_keywords(ctx, field)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "shapes":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_shapes(ctx, field)
				return res
			})
		case "errorBubble":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_errorBubble(ctx, field)
				return res
			})
		case "modelMethods":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_modelMethods(ctx, field)
				return res
			})
		case "valid":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_valid(ctx, field)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "user":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_user(ctx, field)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "nullableArg":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_nullableArg(ctx, field)
				return res
			})
		case "directiveArg":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_directiveArg(ctx, field)
				return res
			})
		case "directiveNullableArg":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_directiveNullableArg(ctx, field)
				return res
			})
		case "directiveInputNullable":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_directiveInputNullable(ctx, field)
				return res
			})
		case "directiveInput":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_directiveInput(ctx, field)
				return res
			})
		case "keywordArgs":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._Query_keywordArgs(ctx, field)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var rectangleImplementors = []string{"Rectangle", "Shape"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Rectangle(ctx context.Context, sel ast.SelectionSet, obj *Rectangle) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, rectangleImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Rectangle")
		case "length":
			out.Values[i] = ec._Rectangle_length(ctx, field, obj)
		case "width":
			out.Values[i] = ec._Rectangle_width(ctx, field, obj)
		case "area":
			out.Values[i] = ec._Rectangle_area(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var subscriptionImplementors = []string{"Subscription"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Subscription(ctx context.Context, sel ast.SelectionSet) func() graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, subscriptionImplementors)
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Subscription",
	})
	if len(fields) != 1 {
		ec.Errorf(ctx, "must subscribe to exactly one stream")
		return nil
	}

	switch fields[0].Name {
	case "updated":
		return ec._Subscription_updated(ctx, fields[0])
	case "initPayload":
		return ec._Subscription_initPayload(ctx, fields[0])
	default:
		panic("unknown field " + strconv.Quote(fields[0].Name))
	}
}

var userImplementors = []string{"User"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *User) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, userImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "id":
			out.Values[i] = ec._User_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "friends":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				res = ec._User_friends(ctx, field, obj)
				if res == graphql.Null {
					invalid = true
				}
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	invalid := false
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalid {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (e *executableSchema) unmarshalInputDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(v interface{}) (*InputDirectives, error) {
	if v == nil {
		return nil, nil
	}
	res, err := e.unmarshalInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(v)
	return &res, err
}
func (e *executableSchema) unmarshalKeywords2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐKeywords(v interface{}) (*Keywords, error) {
	if v == nil {
		return nil, nil
	}
	res, err := e.unmarshalKeywords2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐKeywords(v)
	return &res, err
}
func (e *executableSchema) unmarshalOuterInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(v interface{}) (*OuterInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := e.unmarshalOuterInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(v)
	return &res, err
}
func (e *executableSchema) unmarshalRecursiveInputSlice2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(v interface{}) (*RecursiveInputSlice, error) {
	if v == nil {
		return nil, nil
	}
	res, err := e.unmarshalRecursiveInputSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(v)
	return &res, err
}
func (e *executableSchema) unmarshalInt2ᚖint(v interface{}) (*int, error) {
	if v == nil {
		return nil, nil
	}
	res, err := e.unmarshalInt2int(v)
	return &res, err
}
func (e *executableSchema) unmarshalString2ᚖstring(v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := e.unmarshalString2string(v)
	return &res, err
}
func (e *executableSchema) unmarshalOuterInput2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(v interface{}) ([]*OuterInput, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]*OuterInput, len(vSlice))
	for i := range vSlice {
		res[i], err = e.unmarshalOuterInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
func (e *executableSchema) unmarshalOuterInput2ᚕᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(v interface{}) ([][]*OuterInput, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([][]*OuterInput, len(vSlice))
	for i := range vSlice {
		res[i], err = e.unmarshalOuterInput2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
func (e *executableSchema) unmarshalBoolean2bool(v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}
func (e *executableSchema) unmarshalInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐInputDirectives(v interface{}) (InputDirectives, error) {
	return e.unmarshalInputInputDirectives(v)
}
func (e *executableSchema) unmarshalKeywords2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐKeywords(v interface{}) (Keywords, error) {
	return e.unmarshalInputKeywords(v)
}
func (e *executableSchema) unmarshalOuterInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐOuterInput(v interface{}) (OuterInput, error) {
	return e.unmarshalInputOuterInput(v)
}
func (e *executableSchema) unmarshalRecursiveInputSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚐRecursiveInputSlice(v interface{}) (RecursiveInputSlice, error) {
	return e.unmarshalInputRecursiveInputSlice(v)
}
func (e *executableSchema) unmarshalInt2int(v interface{}) (int, error) {
	return graphql.UnmarshalInt(v)
}
func (e *executableSchema) unmarshalChanges2map(v interface{}) (map[string]interface{}, error) {
	return v.(map[string]interface{}), nil
}
func (e *executableSchema) unmarshalString2string(v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

// endregion ***************************** type.gotpl *****************************
