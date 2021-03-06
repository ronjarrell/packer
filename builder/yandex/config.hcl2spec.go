// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package yandex

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	Endpoint                  *string           `mapstructure:"endpoint" required:"false" cty:"endpoint"`
	FolderID                  *string           `mapstructure:"folder_id" required:"true" cty:"folder_id"`
	ServiceAccountKeyFile     *string           `mapstructure:"service_account_key_file" required:"false" cty:"service_account_key_file"`
	Token                     *string           `mapstructure:"token" required:"true" cty:"token"`
	DiskName                  *string           `mapstructure:"disk_name" required:"false" cty:"disk_name"`
	DiskSizeGb                *int              `mapstructure:"disk_size_gb" required:"false" cty:"disk_size_gb"`
	DiskType                  *string           `mapstructure:"disk_type" required:"false" cty:"disk_type"`
	ImageDescription          *string           `mapstructure:"image_description" required:"false" cty:"image_description"`
	ImageFamily               *string           `mapstructure:"image_family" required:"false" cty:"image_family"`
	ImageLabels               map[string]string `mapstructure:"image_labels" required:"false" cty:"image_labels"`
	ImageName                 *string           `mapstructure:"image_name" required:"false" cty:"image_name"`
	ImageProductIDs           []string          `mapstructure:"image_product_ids" required:"false" cty:"image_product_ids"`
	InstanceCores             *int              `mapstructure:"instance_cores" required:"false" cty:"instance_cores"`
	InstanceGpus              *int              `mapstructure:"instance_gpus" cty:"instance_gpus"`
	InstanceMemory            *int              `mapstructure:"instance_mem_gb" required:"false" cty:"instance_mem_gb"`
	InstanceName              *string           `mapstructure:"instance_name" required:"false" cty:"instance_name"`
	Labels                    map[string]string `mapstructure:"labels" required:"false" cty:"labels"`
	PlatformID                *string           `mapstructure:"platform_id" required:"false" cty:"platform_id"`
	MaxRetries                *int              `mapstructure:"max_retries" cty:"max_retries"`
	Metadata                  map[string]string `mapstructure:"metadata" required:"false" cty:"metadata"`
	MetadataFromFile          map[string]string `mapstructure:"metadata_from_file" cty:"metadata_from_file"`
	Preemptible               *bool             `mapstructure:"preemptible" cty:"preemptible"`
	SerialLogFile             *string           `mapstructure:"serial_log_file" required:"false" cty:"serial_log_file"`
	SourceImageFamily         *string           `mapstructure:"source_image_family" required:"true" cty:"source_image_family"`
	SourceImageFolderID       *string           `mapstructure:"source_image_folder_id" required:"false" cty:"source_image_folder_id"`
	SourceImageID             *string           `mapstructure:"source_image_id" required:"false" cty:"source_image_id"`
	SourceImageName           *string           `mapstructure:"source_image_name" cty:"source_image_name"`
	SubnetID                  *string           `mapstructure:"subnet_id" required:"false" cty:"subnet_id"`
	UseIPv4Nat                *bool             `mapstructure:"use_ipv4_nat" required:"false" cty:"use_ipv4_nat"`
	UseIPv6                   *bool             `mapstructure:"use_ipv6" required:"false" cty:"use_ipv6"`
	UseInternalIP             *bool             `mapstructure:"use_internal_ip" required:"false" cty:"use_internal_ip"`
	Zone                      *string           `mapstructure:"zone" required:"false" cty:"zone"`
	StateTimeout              *string           `mapstructure:"state_timeout" required:"false" cty:"state_timeout"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"endpoint":                     &hcldec.AttrSpec{Name: "endpoint", Type: cty.String, Required: false},
		"folder_id":                    &hcldec.AttrSpec{Name: "folder_id", Type: cty.String, Required: false},
		"service_account_key_file":     &hcldec.AttrSpec{Name: "service_account_key_file", Type: cty.String, Required: false},
		"token":                        &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"disk_name":                    &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"disk_size_gb":                 &hcldec.AttrSpec{Name: "disk_size_gb", Type: cty.Number, Required: false},
		"disk_type":                    &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"image_description":            &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_family":                 &hcldec.AttrSpec{Name: "image_family", Type: cty.String, Required: false},
		"image_labels":                 &hcldec.BlockAttrsSpec{TypeName: "image_labels", ElementType: cty.String, Required: false},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_product_ids":            &hcldec.AttrSpec{Name: "image_product_ids", Type: cty.List(cty.String), Required: false},
		"instance_cores":               &hcldec.AttrSpec{Name: "instance_cores", Type: cty.Number, Required: false},
		"instance_gpus":                &hcldec.AttrSpec{Name: "instance_gpus", Type: cty.Number, Required: false},
		"instance_mem_gb":              &hcldec.AttrSpec{Name: "instance_mem_gb", Type: cty.Number, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"labels":                       &hcldec.BlockAttrsSpec{TypeName: "labels", ElementType: cty.String, Required: false},
		"platform_id":                  &hcldec.AttrSpec{Name: "platform_id", Type: cty.String, Required: false},
		"max_retries":                  &hcldec.AttrSpec{Name: "max_retries", Type: cty.Number, Required: false},
		"metadata":                     &hcldec.BlockAttrsSpec{TypeName: "metadata", ElementType: cty.String, Required: false},
		"metadata_from_file":           &hcldec.BlockAttrsSpec{TypeName: "metadata_from_file", ElementType: cty.String, Required: false},
		"preemptible":                  &hcldec.AttrSpec{Name: "preemptible", Type: cty.Bool, Required: false},
		"serial_log_file":              &hcldec.AttrSpec{Name: "serial_log_file", Type: cty.String, Required: false},
		"source_image_family":          &hcldec.AttrSpec{Name: "source_image_family", Type: cty.String, Required: false},
		"source_image_folder_id":       &hcldec.AttrSpec{Name: "source_image_folder_id", Type: cty.String, Required: false},
		"source_image_id":              &hcldec.AttrSpec{Name: "source_image_id", Type: cty.String, Required: false},
		"source_image_name":            &hcldec.AttrSpec{Name: "source_image_name", Type: cty.String, Required: false},
		"subnet_id":                    &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"use_ipv4_nat":                 &hcldec.AttrSpec{Name: "use_ipv4_nat", Type: cty.Bool, Required: false},
		"use_ipv6":                     &hcldec.AttrSpec{Name: "use_ipv6", Type: cty.Bool, Required: false},
		"use_internal_ip":              &hcldec.AttrSpec{Name: "use_internal_ip", Type: cty.Bool, Required: false},
		"zone":                         &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"state_timeout":                &hcldec.AttrSpec{Name: "state_timeout", Type: cty.String, Required: false},
	}
	return s
}
