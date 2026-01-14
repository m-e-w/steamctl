# Installation

## Linux

### Linux One-liner (64-bit)

[View install script in browser](https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.sh)

Below is a one-line install command you can use to download & install steamctl on Linux platforms.

```bash
curl -fsSL https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.sh | bash
```

If you would like to inspect the script first (always advisable), you can download it first with:
```bash
curl -fsSLO https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.sh
```

And then open using your text editor of choice e.g. vi, nano, etc
```bash
vi install.sh
```

You can then install after using:
```bash
cat install.sh | bash
```

Verify installation
```bash
steamctl --help
```

### Linux Manual Install (64-bit)

1. Download the binary and checksums.txt file
    ```bash
    curl -LO "https://github.com/m-e-w/steamctl/releases/latest/download/steamctl-linux-amd64"
    curl -LO "https://github.com/m-e-w/steamctl/releases/latest/download/checksums.txt"
    ```
2. Verify checksum
    ```bash
    grep " steamctl-linux-amd64" checksums.txt | sha256sum -c -
    ```
    The output should match `steamctl-linux-amd64: OK`

3. Rename binary to steamctl
    ```bash
    mv steamctl-linux-amd64 steamctl
    ```
4. Make the binary executable
    ```bash
    chmod +x steamctl
    ```
5. Ensure your user bin directory exists
    ```bash
    mkdir -p ~/.local/bin
    ```
6. Install to your user bin
    ```bash
    mv steamctl ~/.local/bin
    ```
7. Verify installation
    ```bash
    steamctl --help
    ```

## Windows

### Windows One-liner (64-bit)

[View install script in browser](https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.ps1)

Below is a one-line install command you can use to download & install steamctl on Windows platforms.

```powershell
curl.exe -fsSLO https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.ps1; ./install.ps1; Remove-Item "install.ps1"
```

If you would like to inspect the script first (always advisable), you can download it first with:
```powershell
curl.exe -fsSLO https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.ps1
```

And then open using your text editor of choice e.g. notepad, vscode, etc
```powershell
notepad.exe install.ps1
```

You can then install after using:
```powershell
./install.ps1
```

Verify installation
```powershell
steamctl --help
```

**Important**

By default, running PowerShell scripts may be disabled on Windows systems. If you try running the script below you may see the following error:
```powershell
PS C:\Users\devtest\tmp> .\install.ps1
.\install.ps1 : File C:\Users\devtest\tmp\install.ps1 cannot be loaded because running scripts is disabled on this system. For more information, see about_Execution_Policies at https:/go.microsoft.com/fwlink/?LinkID=135170.
At line:1 char:1
+ .\install.ps1
+ ~~~~~~~~~~~~~
    + CategoryInfo          : SecurityError: (:) [], PSSecurityException
    + FullyQualifiedErrorId : UnauthorizedAccess
```

This is normal and expected.

To get around this, you must temporarily adjust the execution policy for your user.

To get the effective execution policy for the current PowerShell session, use the `Get-ExecutionPolicy` cmdlet.
```powershell
Get-ExecutionPolicy
```

If you see: `Restricted` you need to change your execution policy as such:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope Process
```

Note: This will only adjust the execution policy for the current PowerShell session (meaning there is no need to try to revert this command after as it only lasts/affects the PowerShell session you ran it under)

To confirm it took effect, run `Get-ExecutionPolicy` again and you should now see: `RemoteSigned`. You should now be able to run the install script using `./install.ps1`

For more details on working with Execution Policies in PowerShell, see: [about_Execution_Policies](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-7.5)