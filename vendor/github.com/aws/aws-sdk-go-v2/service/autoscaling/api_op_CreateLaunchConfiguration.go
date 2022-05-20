// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a launch configuration. If you exceed your maximum limit of launch
// configurations, the call fails. To query this limit, call the
// DescribeAccountLimits API. For information about updating this limit, see Amazon
// EC2 Auto Scaling service quotas
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Launch
// configurations
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchConfiguration.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) CreateLaunchConfiguration(ctx context.Context, params *CreateLaunchConfigurationInput, optFns ...func(*Options)) (*CreateLaunchConfigurationOutput, error) {
	if params == nil {
		params = &CreateLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchConfiguration", params, optFns, c.addOperationCreateLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationInput struct {

	// The name of the launch configuration. This name must be unique per Region per
	// account.
	//
	// This member is required.
	LaunchConfigurationName *string

	// For Auto Scaling groups that are running in a virtual private cloud (VPC),
	// specifies whether to assign a public IP address to the group's instances. If you
	// specify true, each instance in the Auto Scaling group receives a unique public
	// IP address. For more information, see Launching Auto Scaling instances in a VPC
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the
	// Amazon EC2 Auto Scaling User Guide. If you specify this parameter, you must
	// specify at least one subnet for VPCZoneIdentifier when you create your group. If
	// the instance is launched into a default subnet, the default is to assign a
	// public IP address, unless you disabled the option to assign a public IP address
	// on the subnet. If the instance is launched into a nondefault subnet, the default
	// is not to assign a public IP address, unless you enabled the option to assign a
	// public IP address on the subnet.
	AssociatePublicIpAddress *bool

	// A block device mapping, which specifies the block devices for the instance. You
	// can specify virtual devices and EBS volumes. For more information, see Block
	// Device Mapping
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	BlockDeviceMappings []types.BlockDeviceMapping

	// EC2-Classic retires on August 15, 2022. This parameter is not supported after
	// that date. The ID of a ClassicLink-enabled VPC to link your EC2-Classic
	// instances to. For more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	ClassicLinkVPCId *string

	// EC2-Classic retires on August 15, 2022. This parameter is not supported after
	// that date. The IDs of one or more security groups for the specified
	// ClassicLink-enabled VPC. For more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon EC2 User Guide for Linux Instances. If you specify the
	// ClassicLinkVPCId parameter, you must specify this parameter.
	ClassicLinkVPCSecurityGroups []string

	// Specifies whether the launch configuration is optimized for EBS I/O (true) or
	// not (false). The optimization provides dedicated throughput to Amazon EBS and an
	// optimized configuration stack to provide optimal I/O performance. This
	// optimization is not available with all instance types. Additional fees are
	// incurred when you enable EBS optimization for an instance type that is not
	// EBS-optimized by default. For more information, see Amazon EBS-optimized
	// instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the
	// Amazon EC2 User Guide for Linux Instances. The default value is false.
	EbsOptimized *bool

	// The name or the Amazon Resource Name (ARN) of the instance profile associated
	// with the IAM role for the instance. The instance profile contains the IAM role.
	// For more information, see IAM role for applications that run on Amazon EC2
	// instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	IamInstanceProfile *string

	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	// For more information, see Finding an AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the
	// Amazon EC2 User Guide for Linux Instances. If you do not specify InstanceId, you
	// must specify ImageId.
	ImageId *string

	// The ID of the instance to use to create the launch configuration. The new launch
	// configuration derives attributes from the instance, except for the block device
	// mapping. To create a launch configuration with a block device mapping or
	// override any other instance attributes, specify them as part of the same
	// request. For more information, see Creating a launch configuration using an EC2
	// instance
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-lc-with-instanceID.html)
	// in the Amazon EC2 Auto Scaling User Guide. If you do not specify InstanceId, you
	// must specify both ImageId and InstanceType.
	InstanceId *string

	// Controls whether instances in this group are launched with detailed (true) or
	// basic (false) monitoring. The default value is true (enabled). When detailed
	// monitoring is enabled, Amazon CloudWatch generates metrics every minute and your
	// account is charged a fee. When you disable detailed monitoring, CloudWatch
	// generates metrics every 5 minutes. For more information, see Configure
	// Monitoring for Auto Scaling Instances
	// (https://docs.aws.amazon.com/autoscaling/latest/userguide/enable-as-instance-metrics.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceMonitoring *types.InstanceMonitoring

	// Specifies the instance type of the EC2 instance. For information about available
	// instance types, see Available Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes)
	// in the Amazon EC2 User Guide for Linux Instances. If you do not specify
	// InstanceId, you must specify InstanceType.
	InstanceType *string

	// The ID of the kernel associated with the AMI.
	KernelId *string

	// The name of the key pair. For more information, see Amazon EC2 Key Pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon EC2 User Guide for Linux Instances.
	KeyName *string

	// The metadata options for the instances. For more information, see Configuring
	// the Instance Metadata Options
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds)
	// in the Amazon EC2 Auto Scaling User Guide.
	MetadataOptions *types.InstanceMetadataOptions

	// The tenancy of the instance. An instance with dedicated tenancy runs on
	// isolated, single-tenant hardware and can only be launched into a VPC. To launch
	// dedicated instances into a shared tenancy VPC (a VPC with the instance placement
	// tenancy attribute set to default), you must set the value of this parameter to
	// dedicated. If you specify PlacementTenancy, you must specify at least one subnet
	// for VPCZoneIdentifier when you create your group. For more information, see
	// Configuring instance tenancy with Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html)
	// in the Amazon EC2 Auto Scaling User Guide. Valid Values: default | dedicated
	PlacementTenancy *string

	// The ID of the RAM disk to select.
	RamdiskId *string

	// A list that contains the security groups to assign to the instances in the Auto
	// Scaling group. [EC2-VPC] Specify the security group IDs. For more information,
	// see Security Groups for Your VPC
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
	// in the Amazon Virtual Private Cloud User Guide. [EC2-Classic] Specify either the
	// security group names or the security group IDs. For more information, see Amazon
	// EC2 Security Groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	SecurityGroups []string

	// The maximum hourly price to be paid for any Spot Instance launched to fulfill
	// the request. Spot Instances are launched when the price you specify exceeds the
	// current Spot price. For more information, see Requesting Spot Instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-launch-spot-instances.html)
	// in the Amazon EC2 Auto Scaling User Guide. When you change your maximum price by
	// creating a new launch configuration, running instances will continue to run as
	// long as the maximum price for those running instances is higher than the current
	// Spot price.
	SpotPrice *string

	// The user data to make available to the launched EC2 instances. For more
	// information, see Instance metadata and user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// (Linux) and Instance metadata and user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html)
	// (Windows). If you are using a command line tool, base64-encoding is performed
	// for you, and you can load the text from a file. Otherwise, you must provide
	// base64-encoded text. User data is limited to 16 KB.
	UserData *string

	noSmithyDocumentSerde
}

type CreateLaunchConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateLaunchConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CreateLaunchConfiguration",
	}
}
