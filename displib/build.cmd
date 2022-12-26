@REM go build -buildmode=c-archive -o ../
go build -buildmode=c-shared --ldflags=-aslr=false -o ../displib.dll
