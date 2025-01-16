setlocal
  SET PATH=C:\go\gcc\mingw64\bin;%PATH%
  SET GOARCH=amd64
  SET GOOS=windows
  SET CGO_ENABLED=1

  rem   copy /y app64\resource.syso_64 app64\resource.syso
  DEL /Q ..\dist\stats.exe ..\dist\stats_64.exe

  go mod tidy

  go build -ldflags="-s -w -X 'firstwails/domain.Mode=production'" -o ../dist/stats_64.exe ./cmd/alcogo
  upx --force-overwrite -o ../dist/stats.exe ../dist/stats_64.exe

endlocal
