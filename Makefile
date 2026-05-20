GENERATORS := $(wildcard generators/*)
TELEMETRY := telemetry/cmd/telemetry
LOG_DIR  := $(shell pwd)/.logs
PID_DIR  := $(shell pwd)/.pids

export APP_MODE   := development
export BROKER_PWD := mypassword
export BROKER_USR := myuser

.PHONY: run-docker-compose run-compose-clean stop-docker clean-docker run-all run-telemetry stop clean

run-docker-compose:
	docker compose up --build

run-compose-clean: stop-docker clean-docker

stop-docker:
	docker compose down

clean-docker:
	docker compose down --rmi all

run-all: clean run-telemetry
	@mkdir -p $(LOG_DIR) $(PID_DIR)
	@$(foreach proj, $(GENERATORS), \
		echo "Building $(proj)..."; \
		(cd $(proj) && go build -o app . && APP_MODE=$(APP_MODE) BROKER_PWD=$(BROKER_PWD) BROKER_USR=$(BROKER_USR) ./app > $(LOG_DIR)/$(notdir $(proj)).log 2>&1) & \
		echo $$! > $(PID_DIR)/$(notdir $(proj)).pid; \
	)
	@echo "All projects running. Use 'make stop' to stop them."

run-telemetry:
	@mkdir -p $(LOG_DIR) $(PID_DIR)
	@echo "Building $(TELEMETRY)..."
	@(cd $(TELEMETRY) && go build -o app . && APP_MODE=$(APP_MODE) BROKER_PWD=$(BROKER_PWD) BROKER_USR=$(BROKER_USR) ./app > $(LOG_DIR)/telemetry.log 2>&1) & \
	echo $$! > $(PID_DIR)/telemetry.pid
	@echo "Telemetry running"

stop:
	@echo "Stopping all projects..."
	@$(foreach proj, $(GENERATORS), \
		if [ -f $(PID_DIR)/$(notdir $(proj)).pid ]; then \
			kill $$(cat $(PID_DIR)/$(notdir $(proj)).pid) 2>/dev/null || true; \
			rm -f $(PID_DIR)/$(notdir $(proj)).pid; \
		fi; \
	)
	@if [ -f $(PID_DIR)/telemetry.pid ]; then \
		kill $$(cat $(PID_DIR)/telemetry.pid) 2>/dev/null || true; \
		rm -f $(PID_DIR)/telemetry.pid; \
	fi
	@echo "All stopped."

clean:
	@rm -rf $(LOG_DIR) $(PID_DIR)
	@$(foreach proj, $(GENERATORS), rm -f $(proj)/app;)
	@rm -f $(TELEMETRY)/app
