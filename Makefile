GO = go
SERVER_MAIN = ./cmd/argus
AGENT_MAIN = ./cmd/argus-agent
SERVER_BIN = argus
AGENT_BIN = argus-agent
BUILD_DIR = ./build
BUILD_SERVER_DIR = $(BUILD_DIR)/server
BUILD_AGENT_DIR = $(BUILD_DIR)/agent
CONFIGS_DIR = configs

.PHONY: complete
complete: agent server

.PHONY: server
server:
	@echo "Building argus-server..."
	$(GO) build -v -o $(BUILD_SERVER_DIR)/$(SERVER_BIN) $(SERVER_MAIN)
	mkdir -p $(BUILD_SERVER_DIR)/$(CONFIGS_DIR)/
	cp -r ./configs/server.toml $(BUILD_SERVER_DIR)/$(CONFIGS_DIR)/

.PHONY: agent
agent:
	@echo "Building argus-agent..."
	$(GO) build -v -o $(BUILD_AGENT_DIR)/$(AGENT_BIN) $(AGENT_MAIN)
	mkdir -p $(BUILD_AGENT_DIR)/$(CONFIGS_DIR)/
	cp -r ./configs/agent.toml $(BUILD_AGENT_DIR)/$(CONFIGS_DIR)/

.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	rm -rf $(BUILD_DIR)
