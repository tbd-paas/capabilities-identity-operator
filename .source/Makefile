VERSION ?= latest

#
# code generation
#
INIT_OPTS=init \
	--workload-config .source/workload.yaml \
   	--repo github.com/tbd-paas/capabilities-identity-operator \
	--controller-image quay.io/tbd-paas/identity-operator:$(VERSION) \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .source/workload.yaml \
	--controller \
	--resource \
	--force

MANAGED_ASSETS=../config/install \
	../README.md \
	../Makefile \
	../Dockerfile \
	../go.mod \
	../controllers/capabilities/certificatescapability_phases.go \
	../phases \
	../.gitignore

# preserve manually mananged assets
preserve:
	@mkdir -p assets; \
	for ASSET in $(MANAGED_ASSETS); do \
		DEST=`echo $$ASSET | sed 's/\.\.\///g'`; \
		if [ -f $$ASSET ]; then \
			echo "preserving file $$ASSET"; \
			DIR=`dirname $$DEST`; \
			mkdir -p assets/$$DIR; \
			cp -R $$ASSET assets/$$DEST; \
		elif [ -d $$ASSET ]; then \
			echo "preserving directory $$ASSET"; \
			mkdir -p assets/$$DEST; \
			cp -R $$ASSET/* assets/$$DEST; \
		else \
			echo "not preserving non-existing asset $$ASSET"; \
		fi \
	done

# restore manually managed assets
restore:
	@cp -R assets/* ..

operator-clean: preserve
	rm -rf ../*

operator-init:
	cd .. && operator-builder $(INIT_OPTS)

operator-create:
	cd .. && operator-builder $(CREATE_OPTS)

#
# project management
#

# create scaffold
PROJECT ?=
project:
	@mkdir -p $(PROJECT)/config $(PROJECT)/vendor $(PROJECT)/manifests
	@touch $(PROJECT)/config/values.yaml $(PROJECT)/config/vendor.yaml $(PROJECT)/config/overlays.yaml

# download from upstream
download = (vendir sync --file $1/config/vendor.yaml --lock-file $1/config/vendor.yaml.lock)
download:
	$(call download, $(PROJECT))

download-all:
	@for project in $$(find . -maxdepth 1 -type d ! -name "assets" ! -name "." -exec basename {} \;); do \
		$(call download, $$project); \
	done

# overlay changes
overlay = (yot -I 2 -i $1/config/overlays.yaml -o $1 -f $1/config/values.yaml)
overlay:
	$(call overlay, $(PROJECT))

overlay-all:
	@for project in $$(find . -maxdepth 1 -type d ! -name "assets" ! -name "." -exec basename {} \;); do \
		$(call overlay, $$project); \
	done

# controller manifests
CONTROLLER_IMG ?= quay.io/tbd-paas/identity-operator:latest
controller-manifests:
	cd .. && \
	make kustomize && \
	mkdir -p config/deploy && \
	cd config/manager && ../../bin/kustomize edit set image controller=$(CONTROLLER_IMG) && \
	cd ../../ && bin/kustomize build config/default > config/deploy/manifests.yaml
