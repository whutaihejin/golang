.PHONY : ALL
ALL :  goconfig gofile struct t4.1 t4.3

goconfig:
	go build goconfig.go
gofile:
	go build gofile.go
struct:
	go build struct.go
t4.1:
	go build t4.1.go
t4.3:
	go build t4.3.go

.PHONY : clean
clean :
	 rm -rf  goconfig gofile struct t4.1 t4.3
