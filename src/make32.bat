setlocal
  SET PATH=C:\Go\32\go.120\bin;C:\go\gcc\mingw32\bin;%PATH%
  SET GOARCH=386
  SET GOOS=windows
  SET CGO_ENABLED=1

  copy /y go.mod.32 go.mod
  go mod tidy

  DEL /Q ..\dist\stats32.exe ..\dist\stats_32.exe

  go build -ldflags="-s -w -X 'firstwails/domain.Mode=production'" -o ../dist/stats_32.exe ./cmd/alcogo
  upx --force-overwrite -o ../dist/stats32.exe ../dist/stats_32.exe

endlocal
