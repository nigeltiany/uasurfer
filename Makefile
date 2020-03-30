.PHONY: types delete-fixed

GITHOST := github.com
GITPROJECT := nigeltiany/uasurfer

types:
	protoc \
	--proto_path=. \
	--go_out=. \
	types/browser/*.proto \
	types/device/*.proto \
	types/os/*.proto \
	types/platform/*.proto \
	types/versioning/*.proto

	$(call fix_golang_output,types/browser)
	$(call fix_golang_output,types/device)
	$(call fix_golang_output,types/os)
	$(call fix_golang_output,types/platform)
	$(call fix_golang_output,types/versioning)

	make delete-fixed

define fix_golang_output
	mv ./$(GITHOST)/$(GITPROJECT)/${1}/* ./${1}
	rm -R ./$(GITHOST)/$(GITPROJECT)/${1}
endef

delete-fixed:
	rm -R ./$(GITHOST)