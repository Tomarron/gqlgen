// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testserver

import (
	"context"

	introspection1 "github.com/99designs/gqlgen/codegen/testserver/introspection"
	"github.com/99designs/gqlgen/codegen/testserver/invalid-packagename"
)

type Stub struct {
	ForcedResolverResolver struct {
		Field func(ctx context.Context, obj *ForcedResolver) (*Circle, error)
	}
	ModelMethodsResolver struct {
		ResolverField func(ctx context.Context, obj *ModelMethods) (bool, error)
	}
	QueryResolver struct {
		InvalidIdentifier      func(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error)
		Collision              func(ctx context.Context) (*introspection1.It, error)
		MapInput               func(ctx context.Context, input map[string]interface{}) (*bool, error)
		Recursive              func(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
		NestedInputs           func(ctx context.Context, input [][]*OuterInput) (*bool, error)
		NestedOutputs          func(ctx context.Context) ([][]*OuterObject, error)
		Keywords               func(ctx context.Context, input *Keywords) (bool, error)
		Shapes                 func(ctx context.Context) ([]Shape, error)
		ErrorBubble            func(ctx context.Context) (*Error, error)
		ModelMethods           func(ctx context.Context) (*ModelMethods, error)
		Valid                  func(ctx context.Context) (string, error)
		User                   func(ctx context.Context, id int) (User, error)
		NullableArg            func(ctx context.Context, arg *int) (*string, error)
		DirectiveArg           func(ctx context.Context, arg string) (*string, error)
		DirectiveNullableArg   func(ctx context.Context, arg *int, arg2 *int) (*string, error)
		DirectiveInputNullable func(ctx context.Context, arg *InputDirectives) (*string, error)
		DirectiveInput         func(ctx context.Context, arg InputDirectives) (*string, error)
		KeywordArgs            func(ctx context.Context, breakArg string, defaultArg string, funcArg string, interfaceArg string, selectArg string, caseArg string, deferArg string, goArg string, mapArg string, structArg string, chanArg string, elseArg string, gotoArg string, packageArg string, switchArg string, constArg string, fallthroughArg string, ifArg string, rangeArg string, typeArg string, continueArg string, forArg string, importArg string, returnArg string, varArg string) (bool, error)
	}
	SubscriptionResolver struct {
		Updated     func(ctx context.Context) (<-chan string, error)
		InitPayload func(ctx context.Context) (<-chan string, error)
	}
	UserResolver struct {
		Friends func(ctx context.Context, obj *User) ([]User, error)
	}
}

func (r *Stub) ForcedResolver() ForcedResolverResolver {
	return &stubForcedResolver{r}
}
func (r *Stub) ModelMethods() ModelMethodsResolver {
	return &stubModelMethods{r}
}
func (r *Stub) Query() QueryResolver {
	return &stubQuery{r}
}
func (r *Stub) Subscription() SubscriptionResolver {
	return &stubSubscription{r}
}
func (r *Stub) User() UserResolver {
	return &stubUser{r}
}

type stubForcedResolver struct{ *Stub }

func (r *stubForcedResolver) Field(ctx context.Context, obj *ForcedResolver) (*Circle, error) {
	return r.ForcedResolverResolver.Field(ctx, obj)
}

type stubModelMethods struct{ *Stub }

func (r *stubModelMethods) ResolverField(ctx context.Context, obj *ModelMethods) (bool, error) {
	return r.ModelMethodsResolver.ResolverField(ctx, obj)
}

type stubQuery struct{ *Stub }

func (r *stubQuery) InvalidIdentifier(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error) {
	return r.QueryResolver.InvalidIdentifier(ctx)
}
func (r *stubQuery) Collision(ctx context.Context) (*introspection1.It, error) {
	return r.QueryResolver.Collision(ctx)
}
func (r *stubQuery) MapInput(ctx context.Context, input map[string]interface{}) (*bool, error) {
	return r.QueryResolver.MapInput(ctx, input)
}
func (r *stubQuery) Recursive(ctx context.Context, input *RecursiveInputSlice) (*bool, error) {
	return r.QueryResolver.Recursive(ctx, input)
}
func (r *stubQuery) NestedInputs(ctx context.Context, input [][]*OuterInput) (*bool, error) {
	return r.QueryResolver.NestedInputs(ctx, input)
}
func (r *stubQuery) NestedOutputs(ctx context.Context) ([][]*OuterObject, error) {
	return r.QueryResolver.NestedOutputs(ctx)
}
func (r *stubQuery) Keywords(ctx context.Context, input *Keywords) (bool, error) {
	return r.QueryResolver.Keywords(ctx, input)
}
func (r *stubQuery) Shapes(ctx context.Context) ([]Shape, error) {
	return r.QueryResolver.Shapes(ctx)
}
func (r *stubQuery) ErrorBubble(ctx context.Context) (*Error, error) {
	return r.QueryResolver.ErrorBubble(ctx)
}
func (r *stubQuery) ModelMethods(ctx context.Context) (*ModelMethods, error) {
	return r.QueryResolver.ModelMethods(ctx)
}
func (r *stubQuery) Valid(ctx context.Context) (string, error) {
	return r.QueryResolver.Valid(ctx)
}
func (r *stubQuery) User(ctx context.Context, id int) (User, error) {
	return r.QueryResolver.User(ctx, id)
}
func (r *stubQuery) NullableArg(ctx context.Context, arg *int) (*string, error) {
	return r.QueryResolver.NullableArg(ctx, arg)
}
func (r *stubQuery) DirectiveArg(ctx context.Context, arg string) (*string, error) {
	return r.QueryResolver.DirectiveArg(ctx, arg)
}
func (r *stubQuery) DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int) (*string, error) {
	return r.QueryResolver.DirectiveNullableArg(ctx, arg, arg2)
}
func (r *stubQuery) DirectiveInputNullable(ctx context.Context, arg *InputDirectives) (*string, error) {
	return r.QueryResolver.DirectiveInputNullable(ctx, arg)
}
func (r *stubQuery) DirectiveInput(ctx context.Context, arg InputDirectives) (*string, error) {
	return r.QueryResolver.DirectiveInput(ctx, arg)
}
func (r *stubQuery) KeywordArgs(ctx context.Context, breakArg string, defaultArg string, funcArg string, interfaceArg string, selectArg string, caseArg string, deferArg string, goArg string, mapArg string, structArg string, chanArg string, elseArg string, gotoArg string, packageArg string, switchArg string, constArg string, fallthroughArg string, ifArg string, rangeArg string, typeArg string, continueArg string, forArg string, importArg string, returnArg string, varArg string) (bool, error) {
	return r.QueryResolver.KeywordArgs(ctx, breakArg, defaultArg, funcArg, interfaceArg, selectArg, caseArg, deferArg, goArg, mapArg, structArg, chanArg, elseArg, gotoArg, packageArg, switchArg, constArg, fallthroughArg, ifArg, rangeArg, typeArg, continueArg, forArg, importArg, returnArg, varArg)
}

type stubSubscription struct{ *Stub }

func (r *stubSubscription) Updated(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.Updated(ctx)
}
func (r *stubSubscription) InitPayload(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.InitPayload(ctx)
}

type stubUser struct{ *Stub }

func (r *stubUser) Friends(ctx context.Context, obj *User) ([]User, error) {
	return r.UserResolver.Friends(ctx, obj)
}
