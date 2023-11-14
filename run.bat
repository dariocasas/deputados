
@echo off
cd build
copy ".env" "./data/flutter_assets" >nul
start "" "server.exe" 
start "" "deputados.exe"

cd ..