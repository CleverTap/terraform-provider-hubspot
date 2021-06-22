
This terraform provider enables to perform Create, Read, 
Update, Delete and Import operations on Hubspot users.


## Requirements 

* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* [Hubspot](https://www.hubspot.com/) Account (APIs are supported in all plans.)
* [Hubspot API Documentation](https://developers.hubspot.com/docs/api/settings/user-provisioning)


## Setup Hubspot Account

### Setup<a id="setup"></a>
1. Create a Hubspot account at https://www.hubspot.com/<br>
2. Go to your developer account.
3. Click on `Manage Apps` or `Create Apps`.
4. Click on `Create app`. Create app with required information. This app will provide us with Client Id, Client Secret and Scopes which will be needed to configure our provider and to make request.<br>

### API Authentication
1. Hubspot uses OAuth for authentication which provides Access Token to authenticate to the API. <br>
2. Provider need Client Id, Client Secret and Refresh Token to generate Access Token. <br>
3. Get the Client Id and Client Secret from your app. <br>  
4. For generating Refresh Token, follow this page <br> (https://developers.hubspot.com/docs/api/oauth-quickstart-guide) <br>


## Building The Provider
Clone the repository, add all the dependencies and create a vendor directory that contains all dependencies. For this, run the following commands: <br>
```
cd terraform-provider-hubspot
go mod init terraform-provider-hubspot
go mod tidy
go mod vendor
```

## Managing Terraform Plugins
*For Windows:*
1. Run the following command to create a sub-directory (`%APPDATA%/terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${OS_ARCH}`) which will consist of all terraform plugins. <br> 
Command: 
```bash
mkdir -p %APPDATA%/terraform.d/plugins/hashicorp.com/user/hubspot/1.0.0/windows_amd64
```
2. Run `go build -o terraform-provider-hubspot.exe` to generate the binary in present working directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-hubspot.exe %APPDATA%\terraform.d\plugins\hashicorp.com/user/hubspot\1.0.0\windows_amd64
 ``` 
<p align="center">[OR]</p>
 
3. Manually move the file from current directory to destination directory (`%APPDATA%\terraform.d\plugins\hashicorp.com/user/hubspot\1.0.0\windows_amd64`).<br>


## Working with Terraform

### Application Credential Integration in terraform
1. Add `terraform` block and `provider` block as shown in [example usage](#example-usage).
2. Get the Client Id, Client Secret and Refresh Token.
3. Assign the above credentials to the respective field in the `provider` block.

### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the changes.

### Create User
1. Add the `email` and  `role_id` in the respective field in `resource` block as shown in [example usage](#example-usage).
2. Run the basic terraform commands.<br>
3. On successful execution, sends an account setup mail to user.<br>

### Update the User
1. Update the data of the user in the `resource` block as show in [example usage](#example-usage) and run the basic terraform commands to update user. 
   User is not allowed to update `email`.

### Read the User Data
Add `data` and `output` blocks as shown in the [example usage](#example-usage) and run the basic terraform commands.

### Delete the user
Delete the `resource` block of the user and run `terraform apply`.
 
### Import a User Data
1. Write manually a `resource` configuration block for the user as shown in [example usage](#example-usage). Imported user will be mapped to this block.
2. Run the command `terraform import hubspot_user.user1 [EMAIL_ID]` to import user.
3. Run `terraform plan`, if output shows `0 to addd, 0 to change and 0 to destroy` user import is successful.
4. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.


## Example Usage 
```terraform
terraform {
    required_providers {
        hubspot = {
            version = "1.0.0"
            source  = "hashicorp.com/user/hubspot"
        }
    }
}

provider "hubspot" {
    client_id     = "_REPLACE_CLIENT_ID_"
    client_secret = "_REPLACE_CLIENT_SECRET_"
    refresh_token = "_REPLACE_REFRESH_TOKEN"
}

resource "hubspot_user" "user1" {
    email  = "user@domain.com"
    role_id = "12345"
}

data "hubspot_user" "user2" {
    id = "user@domain.com"
}

output "user" {
    value = data.hubspot_user.user2
}
```


## Argument Reference

* `client_id`     (Required, String)  - The Hubspot App's Client Id. This may also be set via the `"HUBSPOT_CLIENT_ID"` environment variable.
* `client_secret` (Required, String)  - The Hubspot App's Client Secert. This may also be set via the `"HUBSPOT_CLIENT_SECRET"` environment variable.
* `refresh_token` (Required, String)  - The Refresh Token. This may also be set via the `"HUBSPOT_REFRESH_TOKEN"` environment variable.
* `email`         (Required, String)  - The email id associated with the user account.
* `role_id`        (Optional, String)  - The role id assigned to the user.
* `id`            (Required, string)  - Email of particular user that has to be read.

## Exceptions

1. You have to generate Refresh Token, it can not be automated.
2. Role id can be taken by two ways
* User Interface 
  1. Go to `Settings -> Users & Teams -> Roles -> click on any Role`.<br>
  2. Then in the URL of that page, the last entry is the Id of that Role. <br>
  For exmple in the below URL, `76891` is Id of that Role. 
  `
  https://app.hubspot.com/settings/20060307/users/permissions/76891
  `
* The API (https://developers.hubspot.com/docs/api/settings/user-provisioning).
3. `Super Admin` role can not be assigned to a user through API. It should be done through UI. But it can be changed to another role through API.<br>
4. A user's Role can not be updated to `No Role`.

