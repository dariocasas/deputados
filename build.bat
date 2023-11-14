if not exist build (
    mkdir build
)
cd build

go build ../src/cmd/server/main.go 

if exist server.exe (
    del server.exe
)
ren main.exe server.exe

copy "..\src\cmd\server\config.json" "." 
if exist "..\src\cmd\server\.env" (
    copy "..\src\cmd\server\.env" "." 
)

cd ..\frontend\deputados

call flutter build windows --release 

copy "build\windows\runner\Release\deputados.exe" "..\..\build" 
copy "build\windows\runner\Release\*.dll" "..\..\build" 
xcopy "build\windows\runner\Release\data\" "..\..\build\data\" /E /Y

cd ../..
