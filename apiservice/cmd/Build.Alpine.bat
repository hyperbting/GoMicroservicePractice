set message=`go env GOOS`
echo %message%
%rem GOOS=windows GOARCH=amd64 go install

set GOOS=linux
set GOARCH=amd64
go build -o api ../main.go

%rem 

set CUR_YYYY=%date:~10,4%
set CUR_MM=%date:~4,2%
set CUR_DD=%date:~7,2%
set CUR_HH=%time:~0,2%
if %CUR_HH% lss 10 (set CUR_HH=0%time:~1,1%)

set CUR_NN=%time:~3,2%
set CUR_SS=%time:~6,2%
set CUR_MS=%time:~9,2%

set SUBFILENAME=%CUR_YYYY%%CUR_MM%%CUR_DD%-%CUR_HH%%CUR_NN%%CUR_SS%


docker build -t api_service:%SUBFILENAME% -f ../Dockerfile .