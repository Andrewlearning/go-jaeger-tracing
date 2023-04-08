.PHONY: run
run:
	go run ./sc/sc.go &\
	sleep 1;\
	go run ./sb/sb.go &\
	sleep 1;\
	go run ./sa/sa.go

	echo "[remember to kill sb and sc process, kill PID]"
	lsof -i:8081
	lsof -i:8082


.PHONY: sa
sa:
	 go run ./sa/sa.go
