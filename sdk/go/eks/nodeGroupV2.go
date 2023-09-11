// Code generated by pulumi-gen-eks DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-eks/sdk/v2/go/eks/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// NodeGroup is a component that wraps the AWS EC2 instances that provide compute capacity for an EKS cluster.
type NodeGroupV2 struct {
	pulumi.ResourceState

	// The AutoScalingGroup for the Node group.
	AutoScalingGroup autoscaling.GroupOutput `pulumi:"autoScalingGroup"`
	// The additional security groups for the node group that captures user-specific rules.
	ExtraNodeSecurityGroups ec2.SecurityGroupArrayOutput `pulumi:"extraNodeSecurityGroups"`
	// The security group for the node group to communicate with the cluster.
	NodeSecurityGroup ec2.SecurityGroupOutput `pulumi:"nodeSecurityGroup"`
}

// NewNodeGroupV2 registers a new resource with the given unique name, arguments, and options.
func NewNodeGroupV2(ctx *pulumi.Context,
	name string, args *NodeGroupV2Args, opts ...pulumi.ResourceOption) (*NodeGroupV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NodeGroupV2
	err := ctx.RegisterRemoteComponentResource("eks:index:NodeGroupV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type nodeGroupV2Args struct {
	// The AMI ID to use for the worker nodes.
	//
	// Defaults to the latest recommended EKS Optimized Linux AMI from the AWS Systems Manager Parameter Store.
	//
	// Note: `amiId` and `gpu` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html.
	AmiId *string `pulumi:"amiId"`
	// The AMI Type to use for the worker nodes.
	//
	// Only applicable when setting an AMI ID that is of type `arm64`.
	//
	// Note: `amiType` and `gpu` are mutually exclusive.
	AmiType *string `pulumi:"amiType"`
	// The tags to apply to the NodeGroup's AutoScalingGroup in the CloudFormation Stack.
	//
	// Per AWS, all stack-level tags, including automatically created tags, and the `cloudFormationTags` option are propagated to resources that AWS CloudFormation supports, including the AutoScalingGroup. See https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html
	//
	// Note: Given the inheritance of auto-generated CF tags and `cloudFormationTags`, you should either supply the tag in `autoScalingGroupTags` or `cloudFormationTags`, but not both.
	AutoScalingGroupTags map[string]string `pulumi:"autoScalingGroupTags"`
	// Additional args to pass directly to `/etc/eks/bootstrap.sh`. For details on available options, see: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh. Note that the `--apiserver-endpoint`, `--b64-cluster-ca` and `--kubelet-extra-args` flags are included automatically based on other configuration parameters.
	BootstrapExtraArgs *string `pulumi:"bootstrapExtraArgs"`
	// The tags to apply to the CloudFormation Stack of the Worker NodeGroup.
	//
	// Note: Given the inheritance of auto-generated CF tags and `cloudFormationTags`, you should either supply the tag in `autoScalingGroupTags` or `cloudFormationTags`, but not both.
	CloudFormationTags map[string]string `pulumi:"cloudFormationTags"`
	// The target EKS cluster.
	Cluster interface{} `pulumi:"cluster"`
	// The ingress rule that gives node group access.
	ClusterIngressRule *ec2.SecurityGroupRule `pulumi:"clusterIngressRule"`
	// The number of worker nodes that should be running in the cluster. Defaults to 2.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// Encrypt the root block device of the nodes in the node group.
	EncryptRootBlockDevice *bool `pulumi:"encryptRootBlockDevice"`
	// Extra security groups to attach on all nodes in this worker node group.
	//
	// This additional set of security groups captures any user application rules that will be needed for the nodes.
	ExtraNodeSecurityGroups []*ec2.SecurityGroup `pulumi:"extraNodeSecurityGroups"`
	// Use the latest recommended EKS Optimized Linux AMI with GPU support for the worker nodes from the AWS Systems Manager Parameter Store.
	//
	// Defaults to false.
	//
	// Note: `gpu` and `amiId` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id.html
	Gpu *bool `pulumi:"gpu"`
	// The ingress rule that gives node group access.
	InstanceProfile *iam.InstanceProfile `pulumi:"instanceProfile"`
	// The instance type to use for the cluster's nodes. Defaults to "t2.medium".
	InstanceType *string `pulumi:"instanceType"`
	// Name of the key pair to use for SSH access to worker nodes.
	KeyName *string `pulumi:"keyName"`
	// Extra args to pass to the Kubelet. Corresponds to the options passed in the `--kubeletExtraArgs` flag to `/etc/eks/bootstrap.sh`. For example, '--port=10251 --address=0.0.0.0'. Note that the `labels` and `taints` properties will be applied to this list (using `--node-labels` and `--register-with-taints` respectively) after to the explicit `kubeletExtraArgs`.
	KubeletExtraArgs *string `pulumi:"kubeletExtraArgs"`
	// Custom k8s node labels to be attached to each worker node. Adds the given key/value pairs to the `--node-labels` kubelet argument.
	Labels map[string]string `pulumi:"labels"`
	// The tag specifications to apply to the launch template.
	LaunchTemplateTagSpecifications []ec2.LaunchTemplateTagSpecification `pulumi:"launchTemplateTagSpecifications"`
	// The maximum number of worker nodes running in the cluster. Defaults to 2.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum amount of instances that should remain available during an instance refresh, expressed as a percentage. Defaults to 50.
	MinRefreshPercentage *int `pulumi:"minRefreshPercentage"`
	// The minimum number of worker nodes running in the cluster. Defaults to 1.
	MinSize *int `pulumi:"minSize"`
	// Whether or not to auto-assign public IP addresses on the EKS worker nodes. If this toggle is set to true, the EKS workers will be auto-assigned public IPs. If false, they will not be auto-assigned public IPs.
	NodeAssociatePublicIpAddress *bool `pulumi:"nodeAssociatePublicIpAddress"`
	// Public key material for SSH access to worker nodes. See allowed formats at:
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	// If not provided, no SSH access is enabled on VMs.
	NodePublicKey *string `pulumi:"nodePublicKey"`
	// The size in GiB of a cluster node's root volume. Defaults to 20.
	NodeRootVolumeSize *int `pulumi:"nodeRootVolumeSize"`
	// The security group for the worker node group to communicate with the cluster.
	//
	// This security group requires specific inbound and outbound rules.
	//
	// See for more details:
	// https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html
	//
	// Note: The `nodeSecurityGroup` option and the cluster option`nodeSecurityGroupTags` are mutually exclusive.
	NodeSecurityGroup *ec2.SecurityGroup `pulumi:"nodeSecurityGroup"`
	// The set of subnets to override and use for the worker node group.
	//
	// Setting this option overrides which subnets to use for the worker node group, regardless if the cluster's `subnetIds` is set, or if `publicSubnetIds` and/or `privateSubnetIds` were set.
	NodeSubnetIds []string `pulumi:"nodeSubnetIds"`
	// Extra code to run on node startup. This code will run after the AWS EKS bootstrapping code and before the node signals its readiness to the managing CloudFormation stack. This code must be a typical user data script: critically it must begin with an interpreter directive (i.e. a `#!`).
	NodeUserData *string `pulumi:"nodeUserData"`
	// User specified code to run on node startup. This code is expected to handle the full AWS EKS bootstrapping code and signal node readiness to the managing CloudFormation stack. This code must be a complete and executable user data script in bash (Linux) or powershell (Windows).
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/worker.html
	NodeUserDataOverride *string `pulumi:"nodeUserDataOverride"`
	// Bidding price for spot instance. If set, only spot instances will be added as worker node.
	SpotPrice *string `pulumi:"spotPrice"`
	// Custom k8s node taints to be attached to each worker node. Adds the given taints to the `--register-with-taints` kubelet argument
	Taints map[string]Taint `pulumi:"taints"`
	// Desired Kubernetes master / control plane version. If you do not specify a value, the latest available version is used.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a NodeGroupV2 resource.
type NodeGroupV2Args struct {
	// The AMI ID to use for the worker nodes.
	//
	// Defaults to the latest recommended EKS Optimized Linux AMI from the AWS Systems Manager Parameter Store.
	//
	// Note: `amiId` and `gpu` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html.
	AmiId pulumi.StringPtrInput
	// The AMI Type to use for the worker nodes.
	//
	// Only applicable when setting an AMI ID that is of type `arm64`.
	//
	// Note: `amiType` and `gpu` are mutually exclusive.
	AmiType pulumi.StringPtrInput
	// The tags to apply to the NodeGroup's AutoScalingGroup in the CloudFormation Stack.
	//
	// Per AWS, all stack-level tags, including automatically created tags, and the `cloudFormationTags` option are propagated to resources that AWS CloudFormation supports, including the AutoScalingGroup. See https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html
	//
	// Note: Given the inheritance of auto-generated CF tags and `cloudFormationTags`, you should either supply the tag in `autoScalingGroupTags` or `cloudFormationTags`, but not both.
	AutoScalingGroupTags pulumi.StringMapInput
	// Additional args to pass directly to `/etc/eks/bootstrap.sh`. For details on available options, see: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh. Note that the `--apiserver-endpoint`, `--b64-cluster-ca` and `--kubelet-extra-args` flags are included automatically based on other configuration parameters.
	BootstrapExtraArgs *string
	// The tags to apply to the CloudFormation Stack of the Worker NodeGroup.
	//
	// Note: Given the inheritance of auto-generated CF tags and `cloudFormationTags`, you should either supply the tag in `autoScalingGroupTags` or `cloudFormationTags`, but not both.
	CloudFormationTags pulumi.StringMapInput
	// The target EKS cluster.
	Cluster pulumi.Input
	// The ingress rule that gives node group access.
	ClusterIngressRule *ec2.SecurityGroupRule
	// The number of worker nodes that should be running in the cluster. Defaults to 2.
	DesiredCapacity pulumi.IntPtrInput
	// Encrypt the root block device of the nodes in the node group.
	EncryptRootBlockDevice pulumi.BoolPtrInput
	// Extra security groups to attach on all nodes in this worker node group.
	//
	// This additional set of security groups captures any user application rules that will be needed for the nodes.
	ExtraNodeSecurityGroups []*ec2.SecurityGroup
	// Use the latest recommended EKS Optimized Linux AMI with GPU support for the worker nodes from the AWS Systems Manager Parameter Store.
	//
	// Defaults to false.
	//
	// Note: `gpu` and `amiId` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id.html
	Gpu pulumi.BoolPtrInput
	// The ingress rule that gives node group access.
	InstanceProfile *iam.InstanceProfile
	// The instance type to use for the cluster's nodes. Defaults to "t2.medium".
	InstanceType pulumi.StringPtrInput
	// Name of the key pair to use for SSH access to worker nodes.
	KeyName pulumi.StringPtrInput
	// Extra args to pass to the Kubelet. Corresponds to the options passed in the `--kubeletExtraArgs` flag to `/etc/eks/bootstrap.sh`. For example, '--port=10251 --address=0.0.0.0'. Note that the `labels` and `taints` properties will be applied to this list (using `--node-labels` and `--register-with-taints` respectively) after to the explicit `kubeletExtraArgs`.
	KubeletExtraArgs *string
	// Custom k8s node labels to be attached to each worker node. Adds the given key/value pairs to the `--node-labels` kubelet argument.
	Labels map[string]string
	// The tag specifications to apply to the launch template.
	LaunchTemplateTagSpecifications ec2.LaunchTemplateTagSpecificationArrayInput
	// The maximum number of worker nodes running in the cluster. Defaults to 2.
	MaxSize pulumi.IntPtrInput
	// The minimum amount of instances that should remain available during an instance refresh, expressed as a percentage. Defaults to 50.
	MinRefreshPercentage pulumi.IntPtrInput
	// The minimum number of worker nodes running in the cluster. Defaults to 1.
	MinSize pulumi.IntPtrInput
	// Whether or not to auto-assign public IP addresses on the EKS worker nodes. If this toggle is set to true, the EKS workers will be auto-assigned public IPs. If false, they will not be auto-assigned public IPs.
	NodeAssociatePublicIpAddress *bool
	// Public key material for SSH access to worker nodes. See allowed formats at:
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	// If not provided, no SSH access is enabled on VMs.
	NodePublicKey pulumi.StringPtrInput
	// The size in GiB of a cluster node's root volume. Defaults to 20.
	NodeRootVolumeSize pulumi.IntPtrInput
	// The security group for the worker node group to communicate with the cluster.
	//
	// This security group requires specific inbound and outbound rules.
	//
	// See for more details:
	// https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html
	//
	// Note: The `nodeSecurityGroup` option and the cluster option`nodeSecurityGroupTags` are mutually exclusive.
	NodeSecurityGroup *ec2.SecurityGroup
	// The set of subnets to override and use for the worker node group.
	//
	// Setting this option overrides which subnets to use for the worker node group, regardless if the cluster's `subnetIds` is set, or if `publicSubnetIds` and/or `privateSubnetIds` were set.
	NodeSubnetIds pulumi.StringArrayInput
	// Extra code to run on node startup. This code will run after the AWS EKS bootstrapping code and before the node signals its readiness to the managing CloudFormation stack. This code must be a typical user data script: critically it must begin with an interpreter directive (i.e. a `#!`).
	NodeUserData pulumi.StringPtrInput
	// User specified code to run on node startup. This code is expected to handle the full AWS EKS bootstrapping code and signal node readiness to the managing CloudFormation stack. This code must be a complete and executable user data script in bash (Linux) or powershell (Windows).
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/worker.html
	NodeUserDataOverride pulumi.StringPtrInput
	// Bidding price for spot instance. If set, only spot instances will be added as worker node.
	SpotPrice pulumi.StringPtrInput
	// Custom k8s node taints to be attached to each worker node. Adds the given taints to the `--register-with-taints` kubelet argument
	Taints map[string]TaintArgs
	// Desired Kubernetes master / control plane version. If you do not specify a value, the latest available version is used.
	Version pulumi.StringPtrInput
}

func (NodeGroupV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupV2Args)(nil)).Elem()
}

type NodeGroupV2Input interface {
	pulumi.Input

	ToNodeGroupV2Output() NodeGroupV2Output
	ToNodeGroupV2OutputWithContext(ctx context.Context) NodeGroupV2Output
}

func (*NodeGroupV2) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroupV2)(nil)).Elem()
}

func (i *NodeGroupV2) ToNodeGroupV2Output() NodeGroupV2Output {
	return i.ToNodeGroupV2OutputWithContext(context.Background())
}

func (i *NodeGroupV2) ToNodeGroupV2OutputWithContext(ctx context.Context) NodeGroupV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupV2Output)
}

func (i *NodeGroupV2) ToOutput(ctx context.Context) pulumix.Output[*NodeGroupV2] {
	return pulumix.Output[*NodeGroupV2]{
		OutputState: i.ToNodeGroupV2OutputWithContext(ctx).OutputState,
	}
}

// NodeGroupV2ArrayInput is an input type that accepts NodeGroupV2Array and NodeGroupV2ArrayOutput values.
// You can construct a concrete instance of `NodeGroupV2ArrayInput` via:
//
//	NodeGroupV2Array{ NodeGroupV2Args{...} }
type NodeGroupV2ArrayInput interface {
	pulumi.Input

	ToNodeGroupV2ArrayOutput() NodeGroupV2ArrayOutput
	ToNodeGroupV2ArrayOutputWithContext(context.Context) NodeGroupV2ArrayOutput
}

type NodeGroupV2Array []NodeGroupV2Input

func (NodeGroupV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeGroupV2)(nil)).Elem()
}

func (i NodeGroupV2Array) ToNodeGroupV2ArrayOutput() NodeGroupV2ArrayOutput {
	return i.ToNodeGroupV2ArrayOutputWithContext(context.Background())
}

func (i NodeGroupV2Array) ToNodeGroupV2ArrayOutputWithContext(ctx context.Context) NodeGroupV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupV2ArrayOutput)
}

func (i NodeGroupV2Array) ToOutput(ctx context.Context) pulumix.Output[[]*NodeGroupV2] {
	return pulumix.Output[[]*NodeGroupV2]{
		OutputState: i.ToNodeGroupV2ArrayOutputWithContext(ctx).OutputState,
	}
}

// NodeGroupV2MapInput is an input type that accepts NodeGroupV2Map and NodeGroupV2MapOutput values.
// You can construct a concrete instance of `NodeGroupV2MapInput` via:
//
//	NodeGroupV2Map{ "key": NodeGroupV2Args{...} }
type NodeGroupV2MapInput interface {
	pulumi.Input

	ToNodeGroupV2MapOutput() NodeGroupV2MapOutput
	ToNodeGroupV2MapOutputWithContext(context.Context) NodeGroupV2MapOutput
}

type NodeGroupV2Map map[string]NodeGroupV2Input

func (NodeGroupV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeGroupV2)(nil)).Elem()
}

func (i NodeGroupV2Map) ToNodeGroupV2MapOutput() NodeGroupV2MapOutput {
	return i.ToNodeGroupV2MapOutputWithContext(context.Background())
}

func (i NodeGroupV2Map) ToNodeGroupV2MapOutputWithContext(ctx context.Context) NodeGroupV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupV2MapOutput)
}

func (i NodeGroupV2Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*NodeGroupV2] {
	return pulumix.Output[map[string]*NodeGroupV2]{
		OutputState: i.ToNodeGroupV2MapOutputWithContext(ctx).OutputState,
	}
}

type NodeGroupV2Output struct{ *pulumi.OutputState }

func (NodeGroupV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroupV2)(nil)).Elem()
}

func (o NodeGroupV2Output) ToNodeGroupV2Output() NodeGroupV2Output {
	return o
}

func (o NodeGroupV2Output) ToNodeGroupV2OutputWithContext(ctx context.Context) NodeGroupV2Output {
	return o
}

func (o NodeGroupV2Output) ToOutput(ctx context.Context) pulumix.Output[*NodeGroupV2] {
	return pulumix.Output[*NodeGroupV2]{
		OutputState: o.OutputState,
	}
}

// The AutoScalingGroup for the Node group.
func (o NodeGroupV2Output) AutoScalingGroup() autoscaling.GroupOutput {
	return o.ApplyT(func(v *NodeGroupV2) autoscaling.GroupOutput { return v.AutoScalingGroup }).(autoscaling.GroupOutput)
}

// The additional security groups for the node group that captures user-specific rules.
func (o NodeGroupV2Output) ExtraNodeSecurityGroups() ec2.SecurityGroupArrayOutput {
	return o.ApplyT(func(v *NodeGroupV2) ec2.SecurityGroupArrayOutput { return v.ExtraNodeSecurityGroups }).(ec2.SecurityGroupArrayOutput)
}

// The security group for the node group to communicate with the cluster.
func (o NodeGroupV2Output) NodeSecurityGroup() ec2.SecurityGroupOutput {
	return o.ApplyT(func(v *NodeGroupV2) ec2.SecurityGroupOutput { return v.NodeSecurityGroup }).(ec2.SecurityGroupOutput)
}

type NodeGroupV2ArrayOutput struct{ *pulumi.OutputState }

func (NodeGroupV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeGroupV2)(nil)).Elem()
}

func (o NodeGroupV2ArrayOutput) ToNodeGroupV2ArrayOutput() NodeGroupV2ArrayOutput {
	return o
}

func (o NodeGroupV2ArrayOutput) ToNodeGroupV2ArrayOutputWithContext(ctx context.Context) NodeGroupV2ArrayOutput {
	return o
}

func (o NodeGroupV2ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NodeGroupV2] {
	return pulumix.Output[[]*NodeGroupV2]{
		OutputState: o.OutputState,
	}
}

func (o NodeGroupV2ArrayOutput) Index(i pulumi.IntInput) NodeGroupV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodeGroupV2 {
		return vs[0].([]*NodeGroupV2)[vs[1].(int)]
	}).(NodeGroupV2Output)
}

type NodeGroupV2MapOutput struct{ *pulumi.OutputState }

func (NodeGroupV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeGroupV2)(nil)).Elem()
}

func (o NodeGroupV2MapOutput) ToNodeGroupV2MapOutput() NodeGroupV2MapOutput {
	return o
}

func (o NodeGroupV2MapOutput) ToNodeGroupV2MapOutputWithContext(ctx context.Context) NodeGroupV2MapOutput {
	return o
}

func (o NodeGroupV2MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NodeGroupV2] {
	return pulumix.Output[map[string]*NodeGroupV2]{
		OutputState: o.OutputState,
	}
}

func (o NodeGroupV2MapOutput) MapIndex(k pulumi.StringInput) NodeGroupV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodeGroupV2 {
		return vs[0].(map[string]*NodeGroupV2)[vs[1].(string)]
	}).(NodeGroupV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupV2Input)(nil)).Elem(), &NodeGroupV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupV2ArrayInput)(nil)).Elem(), NodeGroupV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupV2MapInput)(nil)).Elem(), NodeGroupV2Map{})
	pulumi.RegisterOutputType(NodeGroupV2Output{})
	pulumi.RegisterOutputType(NodeGroupV2ArrayOutput{})
	pulumi.RegisterOutputType(NodeGroupV2MapOutput{})
}
