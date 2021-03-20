help:               ## list all commands available for make
				@IFS=$$'\n' ; \
				help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
				for help_line in $${help_lines[@]}; do \
						IFS=$$'#' ; \
						help_split=($$help_line) ; \
						help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
						help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
						printf "%-30s %s\n" $$help_command $$help_info ; \
				done
testAll:            ## run test
				go test ./...
startTest:          ## run go applicaiton
				go run main.go
format:             ## run go format
				gofmt -s -w .
startBuild:         ## run build container for local testing
				docker image build -t recovery-local-test .