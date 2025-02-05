$OSLIST = "linux", "windows", "darwin"
$ARCHLIST = "amd64", "arm", "arm64"

foreach ($os in $OSLIST) {
    foreach ($arch in $ARCHLIST) {
        if ("$os/$arch" -match "windows/arm64|darwin/arm") { continue }

        Write-Host "Building binary for $os $arch"

        $env:CGO_ENABLED = 0
        $env:GOOS = $os
        $env:GOARCH = $arch

#        for inmemory repo version
#        go build -tags=inmemory -o "releases\$os\$arch\"

#        for SQLite repo version
        go build -o "releases\$os\$arch\"
    }
}
