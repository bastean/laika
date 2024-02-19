#* ---------- VARS ----------

git-reset-hard = git reset --hard HEAD

go-tidy = go mod tidy -e

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

compose = docker compose
compose-env = ${compose} --env-file

#* ---------- RULES ----------

genesis:
	@git clean -e .env* -fdx
	@${git-reset-hard}
	@make init

from-zero:
	@git init
	@make init
	@${npx} husky install

upgrade-manager:
	@npm upgrade -g

upgrade-node:
	@${npx} ncu -u
	@rm -f package-lock.json
	@npm i --legacy-peer-deps

upgrade-reset:
	@${git-reset-hard}
	@${npm-ci}

upgrade:
	@go run scripts/upgrade.go

init: upgrade-manager
	@${npm-ci}
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11

lint:
	@gofmt -l -s -w . tests/
	@${npx} prettier --ignore-unknown --write .
	@templ generate
	@templ fmt .
	@${go-tidy}
	@cd tests/server && ${go-tidy}

lint-check:
	@staticcheck ./... ./tests/...
	@${npx} prettier --check .

commit:
	@${npx} cz

release:
	@${release-it}

release-alpha:
	@${release-it} --preRelease=alpha
	
release-beta:
	@${release-it} --preRelease=beta

release-ci:
	@${release-it} --ci $(OPTIONS)

release-dry:
	@${release-it-dry}
 
release-dry-version:
	@${release-it-dry} --release-version

release-dry-changelog:
	@${release-it-dry} --changelog

test-run-server:
	@PORT=8080 ALLOWED_HOSTS=localhost:8080 air

test:
	@go clean -testcache
	@cd tests/ && mkdir -p reports && go test -v -cover ./... > reports/report.txt
