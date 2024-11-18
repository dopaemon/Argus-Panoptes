GO = go
SERVER_MAIN = ./cmd/argus
AGENT_MAIN = ./cmd/argus-agent
SERVER_BIN = argus
AGENT_BIN = argus-agent
BUILD_DIR = build

.PHONY: complete
complete: agent server

.PHONY: server
server:
	@echo "Building argus-server..."
	$(GO) build -o $(BUILD_DIR)/$(SERVER_BIN) $(SERVER_MAIN)

.PHONY: agent
agent:
	@echo "Building argus-agent..."
	$(GO) build -o $(BUILD_DIR)/$(AGENT_BIN) $(AGENT_MAIN)

.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	rm -f $(BUILD_DIR)
