$ErrorActionPreference = "Stop"

# Config
$Repo    = "m-e-w/steamctl"
$Version = "latest"
$OS      = "windows"
$Arch    = "amd64"
$BinName = "steamctl-$OS-$Arch.exe"
$InstallDir = Join-Path $env:LOCALAPPDATA "steamctl"

$BaseUrl = "https://github.com/$Repo/releases/$Version/download"

# Temp files
$TempDir = New-Item -ItemType Directory -Path ([System.IO.Path]::GetTempPath() + [System.Guid]::NewGuid()) -Force
$BinPath = Join-Path $TempDir $BinName
$ChecksumPath = Join-Path $TempDir "checksums.txt"

Write-Host "Downloading steamctl..."

Invoke-WebRequest "$BaseUrl/$BinName" -OutFile $BinPath -UseBasicParsing
Invoke-WebRequest "$BaseUrl/checksums.txt" -OutFile $ChecksumPath -UseBasicParsing

# Verify checksum
$Expected = Select-String -Path $ChecksumPath -Pattern " $BinName$" |
    ForEach-Object { $_.Line.Split(' ')[0] }

if (-not $Expected) {
    Write-Error "Checksum entry not found for $BinName"
}

$Actual = (Get-FileHash $BinPath -Algorithm SHA256).Hash.ToLower()

Write-Host "Verifying checksum..."
if ($Actual -ne $Expected.ToLower()) {
    Write-Error "Checksum verification failed"
}
Write-Host "Checksum verification passed."

# Create install dir if missing
if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}

# Install binary
$FinalPath = Join-Path $InstallDir "steamctl.exe"
Copy-Item $BinPath $FinalPath -Force

# Add to user PATH (idempotent)
$UserPath = [Environment]::GetEnvironmentVariable("PATH", "User")

if ($UserPath.EndsWith(";")) {
    $NewPath = "$UserPath$InstallDir"
} else {
    $NewPath = "$UserPath;$InstallDir"
}

if ($UserPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable(
        "PATH",
        "$NewPath",
        "User"
    )
}


# Cleanup
Remove-Item $TempDir -Recurse -Force

Write-Host ""
Write-Host "steamctl was installed to $InstallDir"
Write-Host ""
Write-Host "Next step:"
Write-Host "  steamctl configure"
Write-Host ""
Write-Host "Restart your terminal to use steamctl"

