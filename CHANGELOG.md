## 1.4.1 (Unreleased)
## 1.4.0 (September 30, 2019)

NOTES:

* The provider has switched to the standalone TF SDK, there should be no noticeable impact on compatibility. ([#32](https://github.com/terraform-providers/terraform-provider-local/issues/32))

NEW FEATURES:

* r/local_file: allow for configurable permissions ([#30](https://github.com/terraform-providers/terraform-provider-local/issues/30))
* r/local_directory: allow for the creation of directories ([#14](https://github.com/hashicorp/terraform-provider-local/pull/14))

## 1.3.0 (June 26, 2019)

* Add support for base64 encoded content ([#29](https://github.com/terraform-providers/terraform-provider-local/issues/29))

## 1.2.2 (May 01, 2019)

* This releases includes another Terraform SDK upgrade intended to align with that being used for other providers as we prepare for the Core v0.12.0 release. It should have no significant changes in behavior for this provider.

## 1.2.1 (April 11, 2019)

* This releases includes only a Terraform SDK upgrade intended to align with that being used for other providers as we prepare for the Core v0.12.0 release. It should have no significant changes in behavior for this provider.

## 1.2.0 (March 20, 2019)

NEW FEATURES:

* The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions.
* `local_file` resource has optional `sensitive_content` attribute, which can be used instead of `content` in situations where the content contains sensitive information that should not be displayed in a rendered diff. ([#9](https://github.com/terraform-providers/terraform-provider-local/issues/9))

## 1.1.0 (January 04, 2018)

NEW FEATURES:

* `local_file` data source, for reading files in a way that participates in Terraform's dependency graph, which allows reading of files that are created dynamically during `terraform apply`. ([#6](https://github.com/terraform-providers/terraform-provider-local/issues/6))

## 1.0.0 (September 15, 2017)

* No changes from 0.1.0; just adjusting to [the new version numbering scheme](https://www.hashicorp.com/blog/hashicorp-terraform-provider-versioning/).

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
