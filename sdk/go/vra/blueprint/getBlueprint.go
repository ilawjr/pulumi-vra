// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides information about a cloud template (blueprint) in vRA.
//
// ## Example Usage
// ### S
//
// This is an example of how to get a vRA cloud template by its name.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/blueprint"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/blueprint"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blueprint.LookupBlueprint(ctx, &blueprint.LookupBlueprintArgs{
//				Name: pulumi.StringRef(vra_blueprint.This.Name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// This is an example of how to get a vRA cloud template by its id.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/blueprint"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/blueprint"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blueprint.LookupBlueprint(ctx, &blueprint.LookupBlueprintArgs{
//				Id: pulumi.StringRef(vra_blueprint.This.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBlueprint(ctx *pulumi.Context, args *LookupBlueprintArgs, opts ...pulumi.InvokeOption) (*LookupBlueprintResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupBlueprintResult
	err := ctx.Invoke("vra:blueprint/getBlueprint:getBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBlueprint.
type LookupBlueprintArgs struct {
	// The id of this cloud template. One of `id` or `name` must be provided.
	Id *string `pulumi:"id"`
	// Name of the cloud template. One of `id` or `name` must be provided.
	Name *string `pulumi:"name"`
	// The id of the project to narrow the search while looking for cloud templates.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getBlueprint.
type LookupBlueprintResult struct {
	// Blueprint YAML content.
	Content string `pulumi:"content"`
	// The id of the content source.
	ContentSourceId string `pulumi:"contentSourceId"`
	// Content source path.
	ContentSourcePath string `pulumi:"contentSourcePath"`
	// Content source last sync at.
	ContentSourceSyncAt string `pulumi:"contentSourceSyncAt"`
	// Content source last sync messages.
	ContentSourceSyncMessages []string `pulumi:"contentSourceSyncMessages"`
	// Content source last sync status. Supported values: `SUCCESSFUL`, `FAILED`.
	ContentSourceSyncStatus string `pulumi:"contentSourceSyncStatus"`
	// Content source type.
	ContentSourceType string `pulumi:"contentSourceType"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// The user the entity was created by.
	CreatedBy string `pulumi:"createdBy"`
	// A human-friendly description.
	Description string  `pulumi:"description"`
	Id          string  `pulumi:"id"`
	Name        *string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId     string `pulumi:"orgId"`
	ProjectId string `pulumi:"projectId"`
	// The name of the project the entity belongs to.
	ProjectName string `pulumi:"projectName"`
	// Flag to indicate whether this blueprint can be requested from any project in the organization this entity belongs to.
	RequestScopeOrg bool `pulumi:"requestScopeOrg"`
	// HATEOAS of the entity.
	SelfLink string `pulumi:"selfLink"`
	// Status of the cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
	Status string `pulumi:"status"`
	// Total number of released versions.
	TotalReleasedVersions int `pulumi:"totalReleasedVersions"`
	// Total number of versions.
	TotalVersions int `pulumi:"totalVersions"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
	// The user the entity was last updated by.
	UpdatedBy string `pulumi:"updatedBy"`
	// Flag to indicate if the current content of the cloud template is valid.
	Valid bool `pulumi:"valid"`
	// List of validations messages.
	// * message - Validation message.
	ValidationMessages []GetBlueprintValidationMessage `pulumi:"validationMessages"`
}

func LookupBlueprintOutput(ctx *pulumi.Context, args LookupBlueprintOutputArgs, opts ...pulumi.InvokeOption) LookupBlueprintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlueprintResult, error) {
			args := v.(LookupBlueprintArgs)
			r, err := LookupBlueprint(ctx, &args, opts...)
			var s LookupBlueprintResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlueprintResultOutput)
}

// A collection of arguments for invoking getBlueprint.
type LookupBlueprintOutputArgs struct {
	// The id of this cloud template. One of `id` or `name` must be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the cloud template. One of `id` or `name` must be provided.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The id of the project to narrow the search while looking for cloud templates.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupBlueprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlueprintArgs)(nil)).Elem()
}

// A collection of values returned by getBlueprint.
type LookupBlueprintResultOutput struct{ *pulumi.OutputState }

func (LookupBlueprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlueprintResult)(nil)).Elem()
}

func (o LookupBlueprintResultOutput) ToLookupBlueprintResultOutput() LookupBlueprintResultOutput {
	return o
}

func (o LookupBlueprintResultOutput) ToLookupBlueprintResultOutputWithContext(ctx context.Context) LookupBlueprintResultOutput {
	return o
}

// Blueprint YAML content.
func (o LookupBlueprintResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Content }).(pulumi.StringOutput)
}

// The id of the content source.
func (o LookupBlueprintResultOutput) ContentSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ContentSourceId }).(pulumi.StringOutput)
}

// Content source path.
func (o LookupBlueprintResultOutput) ContentSourcePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ContentSourcePath }).(pulumi.StringOutput)
}

// Content source last sync at.
func (o LookupBlueprintResultOutput) ContentSourceSyncAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ContentSourceSyncAt }).(pulumi.StringOutput)
}

// Content source last sync messages.
func (o LookupBlueprintResultOutput) ContentSourceSyncMessages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBlueprintResult) []string { return v.ContentSourceSyncMessages }).(pulumi.StringArrayOutput)
}

// Content source last sync status. Supported values: `SUCCESSFUL`, `FAILED`.
func (o LookupBlueprintResultOutput) ContentSourceSyncStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ContentSourceSyncStatus }).(pulumi.StringOutput)
}

// Content source type.
func (o LookupBlueprintResultOutput) ContentSourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ContentSourceType }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupBlueprintResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The user the entity was created by.
func (o LookupBlueprintResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupBlueprintResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlueprintResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The id of the organization this entity belongs to.
func (o LookupBlueprintResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o LookupBlueprintResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The name of the project the entity belongs to.
func (o LookupBlueprintResultOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.ProjectName }).(pulumi.StringOutput)
}

// Flag to indicate whether this blueprint can be requested from any project in the organization this entity belongs to.
func (o LookupBlueprintResultOutput) RequestScopeOrg() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlueprintResult) bool { return v.RequestScopeOrg }).(pulumi.BoolOutput)
}

// HATEOAS of the entity.
func (o LookupBlueprintResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Status of the cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
func (o LookupBlueprintResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.Status }).(pulumi.StringOutput)
}

// Total number of released versions.
func (o LookupBlueprintResultOutput) TotalReleasedVersions() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBlueprintResult) int { return v.TotalReleasedVersions }).(pulumi.IntOutput)
}

// Total number of versions.
func (o LookupBlueprintResultOutput) TotalVersions() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBlueprintResult) int { return v.TotalVersions }).(pulumi.IntOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupBlueprintResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The user the entity was last updated by.
func (o LookupBlueprintResultOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlueprintResult) string { return v.UpdatedBy }).(pulumi.StringOutput)
}

// Flag to indicate if the current content of the cloud template is valid.
func (o LookupBlueprintResultOutput) Valid() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBlueprintResult) bool { return v.Valid }).(pulumi.BoolOutput)
}

// List of validations messages.
// * message - Validation message.
func (o LookupBlueprintResultOutput) ValidationMessages() GetBlueprintValidationMessageArrayOutput {
	return o.ApplyT(func(v LookupBlueprintResult) []GetBlueprintValidationMessage { return v.ValidationMessages }).(GetBlueprintValidationMessageArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlueprintResultOutput{})
}