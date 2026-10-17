param(
    [Parameter(Position=0)]
    [ValidateSet("run", "build", "test", "clean", "help")]
    [string]$Task = "help"
)

$ErrorActionPreference = "Stop"
$BinDir = Join-Path $PSScriptRoot "bin"
$ServerExe = Join-Path $BinDir "server.exe"

function Show-Help {
    Write-Host "Usage: .\tasks.ps1 <target>" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "Targets:" -ForegroundColor Cyan
    Write-Host "  run     - go run ./cmd/server"
    Write-Host "  build   - go build -o .\bin\server.exe .\cmd\server"
    Write-Host "  test    - go test -count=1 ./..."
    Write-Host "  clean   - remove .\bin"
}

switch ($Task) {
    "run" {
        Write-Host ">>> go run ./cmd/server" -ForegroundColor Green
        go run ./cmd/server
    }
    "build" {
        if (-not (Test-Path $BinDir)) {
            New-Item -ItemType Directory -Path $BinDir | Out-Null
        }
        Write-Host ">>> go build -o $ServerExe ./cmd/server" -ForegroundColor Green
        go build -o $ServerExe ./cmd/server
        Write-Host "Built: $ServerExe" -ForegroundColor Green
    }
    "test" {
        Write-Host ">>> go test -count=1 ./..." -ForegroundColor Green
        go test -count=1 ./...
    }
    "clean" {
        if (Test-Path $BinDir) {
            Remove-Item -Recurse -Force $BinDir
            Write-Host "Removed $BinDir" -ForegroundColor Yellow
        } else {
            Write-Host "Nothing to clean" -ForegroundColor Yellow
        }
    }
    default { Show-Help }
}