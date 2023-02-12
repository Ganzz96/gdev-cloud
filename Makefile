.PHONY: build

openapi-codegen:
	@oapi-codegen -config=config/openapi/conn_mgr.yaml api/conn_mgr_api.yaml
