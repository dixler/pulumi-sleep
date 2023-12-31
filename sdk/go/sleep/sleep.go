// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sleep

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type Sleep struct {
	pulumi.CustomResourceState

	Result  pulumi.StringOutput `pulumi:"result"`
	Seconds pulumi.IntOutput    `pulumi:"seconds"`
}

// NewSleep registers a new resource with the given unique name, arguments, and options.
func NewSleep(ctx *pulumi.Context,
	name string, args *SleepArgs, opts ...pulumi.ResourceOption) (*Sleep, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Seconds == nil {
		return nil, errors.New("invalid value for required argument 'Seconds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sleep
	err := ctx.RegisterResource("sleep:index:Sleep", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSleep gets an existing Sleep resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSleep(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SleepState, opts ...pulumi.ResourceOption) (*Sleep, error) {
	var resource Sleep
	err := ctx.ReadResource("sleep:index:Sleep", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sleep resources.
type sleepState struct {
}

type SleepState struct {
}

func (SleepState) ElementType() reflect.Type {
	return reflect.TypeOf((*sleepState)(nil)).Elem()
}

type sleepArgs struct {
	Seconds int `pulumi:"seconds"`
}

// The set of arguments for constructing a Sleep resource.
type SleepArgs struct {
	Seconds pulumi.IntInput
}

func (SleepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sleepArgs)(nil)).Elem()
}

type SleepInput interface {
	pulumi.Input

	ToSleepOutput() SleepOutput
	ToSleepOutputWithContext(ctx context.Context) SleepOutput
}

func (*Sleep) ElementType() reflect.Type {
	return reflect.TypeOf((**Sleep)(nil)).Elem()
}

func (i *Sleep) ToSleepOutput() SleepOutput {
	return i.ToSleepOutputWithContext(context.Background())
}

func (i *Sleep) ToSleepOutputWithContext(ctx context.Context) SleepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SleepOutput)
}

type SleepOutput struct{ *pulumi.OutputState }

func (SleepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sleep)(nil)).Elem()
}

func (o SleepOutput) ToSleepOutput() SleepOutput {
	return o
}

func (o SleepOutput) ToSleepOutputWithContext(ctx context.Context) SleepOutput {
	return o
}

func (o SleepOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *Sleep) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

func (o SleepOutput) Seconds() pulumi.IntOutput {
	return o.ApplyT(func(v *Sleep) pulumi.IntOutput { return v.Seconds }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SleepInput)(nil)).Elem(), &Sleep{})
	pulumi.RegisterOutputType(SleepOutput{})
}
