# Terraform Plugin SDK

This is a fork of Hashicorp's [terraform-plugin-sdk](https://github.com/hashicorp/terraform-plugin-sdk).

Key Differences:

- Exported functions using helper/schema/exports.go for use by the TF bridge
- Enhanced `PlanResourceChange` functionality:
  - Added `PlanResourceChangeExtra` function with instance diff transformation
  - Added `TransformInstanceDiff` callback support
- State upgrade functions made public:
  - `UpgradeFlatmapState` and `UpgradeJSONState` exported for external use
- Provider validation control:
  - Added `RunProviderInternalValidation` flag to control when internal validation runs

These changes maintain backward compatibility while extending the SDK's capabilities for Pulumi's specific use cases.

## Updating the fork

When a new version of the terraform-plugin-sdk comes out, we need to update our fork. This is necessary to avoid build errors in newer providers.

To update our fork:
1. Checkout the latest upstream branch: `git checkout upstream-${LATEST_VERSION}`.
2. Create a new branch: `git checkout -b upstream-${NEW_VERSION}`.
3. Ensure that original repository is a remote and up to date:

   ``` shell
   𝛌 git remote -v
   origin  https://github.com/pulumi/terraform-plugin-sdk.git (fetch)
   origin  https://github.com/pulumi/terraform-plugin-sdk.git (push)
   tf      https://github.com/hashicorp/terraform-plugin-sdk.git (fetch)
   tf      https://github.com/hashicorp/terraform-plugin-sdk.git (push)
   𝛌 git fetch tf
   ```
4. Merge the new version into your branch: `git merge ${NEW_VERSION}`.
5. Resolve any conflicts and push the new version.

We do not use the `master` branch for this.

