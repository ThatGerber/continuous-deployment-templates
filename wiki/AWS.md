## 0. Setup

### Getting Started on AWS

#### Account

1. Sign Up
    * Go to [sign-in][aws_portal_signin] page.
    * Add Email and choose "I am a new user." radial.
    * Confirm email & add password.
    * Confirm Contact/Company information
    * Add payment information
    * Confirm Contact Info via email.

2. Securing Account
    1. Create IAM Users/Groups.
        1. Highly-secure Setup
            * Create three groups to cover 3 roles that would need to be
              fulfilled to securely split up management and eliminate root
              credentials.
                * Administrator - Can provision users and create groups in IAM.
                * Manager - Can assign users to groups.
                * FinanceManager - Can access billing, usage, payment methods, and budgets.
            * This makes it impossible for one user to be able to wreak havoc on an
              account. Any attempt to provide extended privileges requires one
              person to create an elevated privilege group and a second person to
              assign users to that group.
        2. Simple (common) Setup
            * Create one group for account administrators that is assigned the
              "job function" policies "AdministratorAccess" and "Billing".
        * **NOTE** There are some tasks that still require the root account.
          Please see [Notes][additional_notes.md] for a link to the list.
    2. Create IAM users for each group created above.
        * Create users with dashboard access (username/password) that can be
          placed within each of the groups created above.
    3. Add MFA token to root account.
        * For highly secure systems, associate the root account's MFA with a
          physical token and store the token in a physically secure location,
          like a safe, locker or safe deposit box.
        * Remember that the battery in the device can wear out, which would
          require a long verify/reset/removal process, so plan to rotate the
          key yearly.
        * [Gemalta Hardware MFA Token from Amazon.com][aws_hardware_mfa_link]
    4. [Create a strong password policy][aws_password_policy].
    5. Create a Cloudwatch Alarm to monitor root logins.
        * [Receive Notifications When Your AWS Account’s Root Access Keys Are Used][root_cloudwatch_alarm]

3. Configuring CLI

    1. Install AWS CLI
    2. Generate access keys and add them to `$HOME/.aws/config` or
       `$HOME/.aws/credentials`.

[aws_hardware_mfa_link]: http://a.co/0omvLT7
[aws_password_policy]: http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_enable-user-change.html
[aws_portal_signin]: https://portal.aws.amazon.com/gp/aws/developer/registration/index.html?nc2=h_ct
[root_cloudwatch_alarm]: https://aws.amazon.com/blogs/security/how-to-receive-notifications-when-your-aws-accounts-root-access-keys-are-used/
